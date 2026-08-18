package backinstock_test

import (
	"encoding/json"
	"testing"
	"time"

	identity "github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/identity/account"
	notification "github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/notifications/backinstock"
	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/notifications/backinstock/backinstock_enums"
)

func TestBackInStockSubscriptionJSONRoundTrip(t *testing.T) {
	now := time.Date(2026, 7, 8, 1, 2, 3, 0, time.UTC)
	sub := notification.BackInStockSubscription{
		ID:           "bis_123",
		SKUCode:      "SKU-001",
		UserID:       "usr_123",
		CustomerType: backinstock_enums.BackInStockCustomerTypeRetail,
		Channel:      backinstock_enums.BackInStockChannelSMS,
		Locale:       "zh-Hant",
		Status:       backinstock_enums.BackInStockStatusPending,
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
	if decoded.SKUCode != sub.SKUCode || decoded.Channel != sub.Channel || decoded.Status != sub.Status {
		t.Fatalf("decoded subscription = %#v", decoded)
	}
	if !decoded.ConsentSnapshot.SMSConsent || decoded.ConsentSnapshot.AccountPreferences == nil {
		t.Fatalf("decoded consent snapshot = %#v", decoded.ConsentSnapshot)
	}
}
