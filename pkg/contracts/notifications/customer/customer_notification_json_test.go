package customer_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	notification "github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/notifications/customer"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/notifications/customer/customer_enums"
)

func TestCustomerNotificationJSONRoundTrip(t *testing.T) {
	now := time.Date(2026, 7, 12, 2, 3, 4, 0, time.UTC)
	n := notification.CustomerNotification{
		ID: "ntf_1", EventID: "event_1",
		Topic: customer_enums.CustomerNotificationTopicPreorderAvailable,
		Title: "Your preorder is ready", Message: "Your order can now be processed.",
		Status:    customer_enums.CustomerNotificationStatusUnread,
		ExpiresAt: now.Add(14 * 24 * time.Hour),
		Deliveries: []notification.CustomerNotificationDelivery{{
			Channel:      customer_enums.CustomerNotificationChannelEmail,
			Status:       customer_enums.CustomerNotificationDeliveryStatusDelivered,
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

func TestReadCustomerNotificationRemainsVisible(t *testing.T) {
	readAt := time.Date(2026, 7, 30, 4, 5, 6, 0, time.UTC)
	n := notification.CustomerNotification{
		ID: "ntf_read", EventID: "event_read",
		Topic: customer_enums.CustomerNotificationTopicOrderConfirmed,
		Title: "Order confirmed", Message: "Your order is confirmed.",
		Status:    customer_enums.CustomerNotificationStatusRead,
		ReadAt:    &readAt,
		CreatedAt: readAt.Add(-time.Hour),
		ExpiresAt: readAt.Add(30 * 24 * time.Hour),
	}

	payload, err := json.Marshal(n)
	if err != nil {
		t.Fatalf("marshal read notification: %v", err)
	}
	if !strings.Contains(string(payload), `"status":"read"`) ||
		!strings.Contains(string(payload), `"read_at":"2026-07-30T04:05:06Z"`) {
		t.Fatalf("read notification JSON = %s", payload)
	}

	var decoded notification.CustomerNotification
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal read notification: %v", err)
	}
	if decoded.Status != customer_enums.CustomerNotificationStatusRead || decoded.ReadAt == nil || !decoded.ReadAt.Equal(readAt) {
		t.Fatalf("read notification did not round-trip: %+v", decoded)
	}
}

func TestCustomerLifecycleNotificationTopicJSONValues(t *testing.T) {
	topics := []customer_enums.CustomerNotificationTopic{
		customer_enums.CustomerNotificationTopicOrderConfirmed,
		customer_enums.CustomerNotificationTopicOrderPlaced,
		customer_enums.CustomerNotificationTopicOrderCancelled,
		customer_enums.CustomerNotificationTopicPaymentReceived,
		customer_enums.CustomerNotificationTopicPaymentFailed,
		customer_enums.CustomerNotificationTopicPaymentRefunded,
		customer_enums.CustomerNotificationTopicPackingStarted,
		customer_enums.CustomerNotificationTopicOrderPacked,
		customer_enums.CustomerNotificationTopicOrderDispatched,
		customer_enums.CustomerNotificationTopicOrderDelivered,
		customer_enums.CustomerNotificationTopicInvoiceAvailable,
		customer_enums.CustomerNotificationTopicPromotionAvailable,
		customer_enums.CustomerNotificationTopicAnnouncement,
	}

	payload, err := json.Marshal(topics)
	if err != nil {
		t.Fatalf("marshal lifecycle topics: %v", err)
	}
	wantValues := []string{
		"order_confirmed",
		"order_placed",
		"order_cancelled",
		"payment_received",
		"payment_failed",
		"payment_refunded",
		"packing_started",
		"order_packed",
		"order_dispatched",
		"order_delivered",
		"invoice_available",
		"promotion_available",
		"announcement",
	}
	for _, want := range wantValues {
		if !strings.Contains(string(payload), `"`+want+`"`) {
			t.Fatalf("lifecycle topic JSON = %s, missing %q", payload, want)
		}
	}
}

func TestCampaignNotificationReferenceAndPushChannelJSON(t *testing.T) {
	now := time.Date(2026, 7, 29, 1, 2, 3, 0, time.UTC)
	n := notification.CustomerNotification{
		ID: "notification_1", EventID: "event_1",
		Topic:   customer_enums.CustomerNotificationTopicPromotionAvailable,
		Title:   "Promotion available",
		Message: "Open the app to see the current promotion.",
		Campaign: &notification.CampaignReference{
			CampaignID: "campaign_1", CampaignKey: "winter-sale", PromotionID: "promotion_1",
			ActivationRevision: 2, ContentRevision: 5,
		},
		Deliveries: []notification.CustomerNotificationDelivery{{
			Channel: customer_enums.CustomerNotificationChannelPush,
			Status:  customer_enums.CustomerNotificationDeliveryStatusPending,
		}},
		CreatedAt: now, ExpiresAt: now.Add(30 * 24 * time.Hour),
		Status: customer_enums.CustomerNotificationStatusUnread,
	}
	payload, err := json.Marshal(n)
	if err != nil {
		t.Fatalf("marshal campaign push notification: %v", err)
	}
	for _, field := range []string{`"channel":"push"`, `"campaign_key":"winter-sale"`, `"promotion_id":"promotion_1"`, `"activation_revision":2`, `"content_revision":5`} {
		if !strings.Contains(string(payload), field) {
			t.Fatalf("campaign push notification missing %s: %s", field, payload)
		}
	}
}
