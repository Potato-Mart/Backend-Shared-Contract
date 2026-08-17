package backinstock

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/geography"
	identity "github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/identity/account"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/notifications/backinstock/backinstock_enums"
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
	Channel     backinstock_enums.BackInStockChannel `json:"channel"`
	Code        string                               `json:"code,omitempty"`
	Message     string                               `json:"message,omitempty"`
	AttemptedAt time.Time                            `json:"attempted_at"`
}

// BackInStockSubscription is a one-shot request to notify an authenticated
// account when a SKU becomes storefront-visible and sellable again.
type BackInStockSubscription struct {
	ID    string `json:"id"`
	SKUID string `json:"sku_id"`
	// MarketID is the market the subscriber expects to buy the SKU in.
	MarketID string `json:"market_id"`
	// CountryCode is the denormalized country of MarketID, carried so a
	// country-scoped staff query is a plain indexed match.
	CountryCode         geography.CountryCode                     `json:"country_code,omitempty"`
	UserID              string                                    `json:"user_id"`
	CustomerType        backinstock_enums.BackInStockCustomerType `json:"customer_type"`
	Channel             backinstock_enums.BackInStockChannel      `json:"channel"`
	Locale              string                                    `json:"locale,omitempty"`
	Status              backinstock_enums.BackInStockStatus       `json:"status"`
	ConsentSnapshot     BackInStockConsentSnapshot                `json:"consent_snapshot"`
	RequestedAt         time.Time                                 `json:"requested_at"`
	NotifiedAt          *time.Time                                `json:"notified_at,omitempty"`
	CancelledAt         *time.Time                                `json:"cancelled_at,omitempty"`
	LastDeliveryError   *BackInStockDeliveryError                 `json:"last_delivery_error,omitempty"`
	NotificationEventID string                                    `json:"notification_event_id,omitempty"`
}
