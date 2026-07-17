package notification_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v18/pkg/contracts/notification"
	notificationenum "github.com/Potato-Mart/Backend-Shared-Contract/v18/pkg/enums/notification"
)

func TestCustomerNotificationJSONRoundTrip(t *testing.T) {
	now := time.Date(2026, 7, 12, 2, 3, 4, 0, time.UTC)
	n := notification.CustomerNotification{
		ID: "ntf_1", EventID: "event_1",
		Topic: notificationenum.CustomerNotificationTopicPreorderAvailable,
		Title: "Your preorder is ready", Message: "Your order can now be processed.",
		Status:    notificationenum.CustomerNotificationStatusUnread,
		ExpiresAt: now.Add(14 * 24 * time.Hour),
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
		!strings.Contains(string(payload), `"status":"delivered"`) ||
		!strings.Contains(string(payload), `"status":"unread"`) ||
		!strings.Contains(string(payload), `"expires_at":`) {
		t.Fatalf("notification JSON = %s", payload)
	}
}
