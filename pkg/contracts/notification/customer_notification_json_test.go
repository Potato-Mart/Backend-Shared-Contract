package notification_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v17/pkg/contracts/notification"
	notificationenum "github.com/Potato-Mart/Backend-Shared-Contract/v17/pkg/enums/notification"
)

func TestCustomerNotificationJSONRoundTrip(t *testing.T) {
	now := time.Date(2026, 7, 12, 2, 3, 4, 0, time.UTC)
	n := notification.CustomerNotification{
		ID: "ntf_1", EventID: "event_1",
		Topic: notificationenum.CustomerNotificationTopicPreorderAvailable,
		Title: "Your preorder is ready", Message: "Your order can now be processed.",
		Deliveries: []notification.CustomerNotificationDelivery{{
			Channel:      notificationenum.CustomerNotificationChannelEmail,
			Status:       notificationenum.CustomerNotificationDeliveryStatusDelivered,
			AttemptCount: 1, DeliveredAt: &now,
		}},
		CreatedAt: now,
	}
	payload, err := json.Marshal(n)
	if err != nil {
		t.Fatalf("marshal notification: %v", err)
	}
	if !strings.Contains(string(payload), `"event_id":"event_1"`) ||
		!strings.Contains(string(payload), `"status":"delivered"`) {
		t.Fatalf("notification JSON = %s", payload)
	}
}
