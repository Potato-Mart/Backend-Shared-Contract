package payments

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/payments/payment/payment_enums"
	"time"
)

// PaymentCapturedEvent is emitted on the payment-events topic when a payment
// is captured (online provider webhook, terminal settlement or manual
// record). AggregateID is the order number.
type PaymentCapturedEvent struct {
	PaymentID         string                      `json:"payment_id"`
	OrderID           string                      `json:"order_id,omitempty"`
	OrderNumber       string                      `json:"order_number"`
	Method            payment_enums.PaymentMethod `json:"method,omitempty"`
	Amount            money.Money                 `json:"amount"`
	ProviderSessionID string                      `json:"provider_session_id,omitempty"`
	// MarketCode, DepotCode, and CountryCode are the denormalized geography
	// the event belongs to. Empty values provide no geographic evidence; a
	// consumer that persists a geographically scoped record must fail closed
	// rather than defaulting them.
	MarketCode  string                `json:"market_code,omitempty"`
	CountryCode geography.CountryCode `json:"country_code,omitempty"`
	DepotCode   string                `json:"depot_code,omitempty"`
	CapturedAt  time.Time             `json:"captured_at"`
	RequestID   string                `json:"request_id,omitempty"`
}
