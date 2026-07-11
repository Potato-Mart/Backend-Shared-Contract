package notification_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v15/pkg/contracts/notification"
	notificationenum "github.com/Potato-Mart/Backend-Shared-Contract/v15/pkg/enums/notification"
)

func TestPreorderAvailabilityCommandExcludesCallerControlledDeliveryCopy(t *testing.T) {
	now := time.Date(2026, 7, 12, 2, 3, 4, 0, time.UTC)
	command := notification.PreorderAvailabilityCommand{
		EventID:              "preorder:PO-1:SKU-1:available",
		RetailCustomerNumber: "RET-1",
		OrderNumber:          "PO-1",
		ProductSKUCode:       "SKU-1",
		ProductName:          "Potato",
		Quantity:             2,
		AvailableAt:          now,
		Locale:               "en-AU",
	}
	payload, err := json.Marshal(command)
	if err != nil {
		t.Fatalf("marshal command: %v", err)
	}
	for _, forbidden := range []string{"recipient", "email_address", "subject", "body"} {
		if strings.Contains(string(payload), forbidden) {
			t.Fatalf("command leaked caller-controlled %q: %s", forbidden, payload)
		}
	}

	var decoded notification.PreorderAvailabilityCommand
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal command: %v", err)
	}
	if decoded.EventID != command.EventID || decoded.RetailCustomerNumber != command.RetailCustomerNumber {
		t.Fatalf("decoded command = %#v", decoded)
	}
}

func TestCustomerNotificationAndReceiptJSONRoundTrip(t *testing.T) {
	now := time.Date(2026, 7, 12, 2, 3, 4, 0, time.UTC)
	delivery := notification.CustomerNotificationDelivery{
		Channel:       notificationenum.CustomerNotificationChannelEmail,
		Status:        notificationenum.CustomerNotificationDeliveryStatusDelivered,
		AttemptCount:  1,
		LastAttemptAt: &now,
		DeliveredAt:   &now,
	}
	n := notification.CustomerNotification{
		ID:             "ntf_1",
		EventID:        "event_1",
		Topic:          notificationenum.CustomerNotificationTopicPreorderAvailable,
		Title:          "Your preorder is ready",
		Message:        "Your order can now be processed.",
		ActionURL:      "/en/account/orders/PO-1",
		OrderNumber:    "PO-1",
		ProductSKUCode: "SKU-1",
		ProductName:    "Potato",
		Deliveries:     []notification.CustomerNotificationDelivery{delivery},
		CreatedAt:      now,
	}
	receipt := notification.NotificationDeliveryReceipt{
		EventID:        n.EventID,
		NotificationID: n.ID,
		Topic:          n.Topic,
		Deliveries:     n.Deliveries,
		CreatedAt:      now,
	}

	for name, value := range map[string]any{"notification": n, "receipt": receipt} {
		payload, err := json.Marshal(value)
		if err != nil {
			t.Fatalf("marshal %s: %v", name, err)
		}
		if !strings.Contains(string(payload), `"event_id":"event_1"`) ||
			!strings.Contains(string(payload), `"status":"delivered"`) {
			t.Fatalf("%s JSON = %s", name, payload)
		}
	}
}

func TestCustomerNotificationPaths(t *testing.T) {
	if notification.PathCustomerNotifications != "/v1/notifications" ||
		notification.PathCustomerNotificationRead != "/v1/notifications/:id/read" ||
		notification.PathInternalPreorderAvailability != "/v1/internal/notifications/preorder-availability" {
		t.Fatal("notification path constants drifted")
	}
}
