package event

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/payments/payment/payment_enums"
	"time"
)

// PaymentFailedEvent is emitted on the payment-events topic when a payment
// attempt terminally fails (provider decline, validation failure, expiry).
type PaymentFailedEvent struct {
	PaymentID            string                      `json:"payment_id,omitempty"`
	OrderID              string                      `json:"order_id,omitempty"`
	OrderNumber          string                      `json:"order_number"`
	Method               payment_enums.PaymentMethod `json:"method,omitempty"`
	Amount               money.Money                 `json:"amount,omitempty"`
	ProviderSessionID    string                      `json:"provider_session_id,omitempty"`
	RetailCustomerNumber string                      `json:"retail_customer_number,omitempty"`
	OrganisationAccessID string                      `json:"organisation_access_id,omitempty"`
	Reason               string                      `json:"reason,omitempty"`
	// MarketCode and CountryCode are the denormalized geography the event
	// belongs to. Empty values provide no geographic evidence; a consumer
	// that persists a geographically scoped record must fail closed rather
	// than defaulting them.
	MarketCode  string                `json:"market_code,omitempty"`
	CountryCode geography.CountryCode `json:"country_code,omitempty"`
	FailedAt    time.Time             `json:"failed_at"`
	RequestID   string                `json:"request_id,omitempty"`
}
