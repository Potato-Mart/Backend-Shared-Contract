package notification_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/contracts/identity"
	"github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/contracts/notification"
	notificationenum "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/enums/notification"
)

func TestBackInStockSubscriptionJSONRoundTrip(t *testing.T) {
	now := time.Date(2026, 7, 8, 1, 2, 3, 0, time.UTC)
	sub := notification.BackInStockSubscription{
		ID:             "bis_123",
		ProductSKUCode: "SKU-001",
		UserID:         "usr_123",
		CustomerType:   notificationenum.BackInStockCustomerTypeRetail,
		Channel:        notificationenum.BackInStockChannelSMS,
		Locale:         "zh-Hant",
		Status:         notificationenum.BackInStockStatusPending,
		ConsentSnapshot: notification.BackInStockConsentSnapshot{
			AccountPreferences: &identity.UserNotificationPreferences{
				Channels: identity.UserNotificationChannels{Email: true, SMS: true},
			},
			EmailConsent: true,
			SMSConsent:   true,
			PhonePresent: true,
			CapturedAt:   now,
		},
		RequestedAt: now,
	}

	body, err := json.Marshal(sub)
	if err != nil {
		t.Fatalf("marshal subscription: %v", err)
	}

	var decoded notification.BackInStockSubscription
	if err := json.Unmarshal(body, &decoded); err != nil {
		t.Fatalf("unmarshal subscription: %v", err)
	}
	if decoded.ProductSKUCode != sub.ProductSKUCode || decoded.Channel != sub.Channel || decoded.Status != sub.Status {
		t.Fatalf("decoded subscription = %#v", decoded)
	}
	if !decoded.ConsentSnapshot.SMSConsent || decoded.ConsentSnapshot.AccountPreferences == nil {
		t.Fatalf("decoded consent snapshot = %#v", decoded.ConsentSnapshot)
	}
}
