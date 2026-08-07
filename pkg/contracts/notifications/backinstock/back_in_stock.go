package backinstock

import (
	"time"

	identity "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/identity/account"
)

// BackInStockConsentSnapshot records the account and customer-level consent
// evidence captured when the customer asked to be notified.
type BackInStockConsentSnapshot struct {
	AccountPreferences *identity.UserNotificationPreferences `json:"account_preferences,omitempty"`
	EmailConsent       bool                                  `json:"email_consent"`
	SMSConsent         bool                                  `json:"sms_consent"`
	PhonePresent       bool                                  `json:"phone_present"`
	CapturedAt         time.Time                             `json:"captured_at"`
}

// BackInStockDeliveryError keeps the last delivery failure visible without
// coupling the contract to a specific email or SMS provider.
type BackInStockDeliveryError struct {
	Channel     BackInStockChannel `json:"channel"`
	Code        string             `json:"code,omitempty"`
	Message     string             `json:"message,omitempty"`
	AttemptedAt time.Time          `json:"attempted_at"`
}

// BackInStockSubscription is a one-shot request to notify an authenticated
// account when a SKU becomes storefront-visible and sellable again.
type BackInStockSubscription struct {
	ID                  string                     `json:"id"`
	ProductSKUCode      string                     `json:"product_sku_code"`
	UserID              string                     `json:"user_id"`
	CustomerType        BackInStockCustomerType    `json:"customer_type"`
	Channel             BackInStockChannel         `json:"channel"`
	Locale              string                     `json:"locale,omitempty"`
	Status              BackInStockStatus          `json:"status"`
	ConsentSnapshot     BackInStockConsentSnapshot `json:"consent_snapshot"`
	RequestedAt         time.Time                  `json:"requested_at"`
	NotifiedAt          *time.Time                 `json:"notified_at,omitempty"`
	CancelledAt         *time.Time                 `json:"cancelled_at,omitempty"`
	LastDeliveryError   *BackInStockDeliveryError  `json:"last_delivery_error,omitempty"`
	NotificationEventID string                     `json:"notification_event_id,omitempty"`
}
