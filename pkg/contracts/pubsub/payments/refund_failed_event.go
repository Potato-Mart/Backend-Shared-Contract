package payments

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
)

// RefundFailedEvent is emitted on the refund-events topic when a refund
// attempt terminally fails. Reason is free text; ReasonCode, when present, is
// the Payments-owned machine-readable code for the failure (for example
// GIFT_CARD_ALREADY_USED when a sold gift card in the refund was spent).
// ReasonCode is an open code: consumers must not reject one they do not know.
type RefundFailedEvent struct {
	RefundID    string      `json:"refund_id"`
	OrderID     string      `json:"order_id,omitempty"`
	OrderNumber string      `json:"order_number"`
	Amount      money.Money `json:"amount,omitempty"`
	Reason      string      `json:"reason,omitempty"`
	ReasonCode  string      `json:"reason_code,omitempty"`
	// MarketCode and CountryCode are the denormalized geography the event
	// belongs to. Empty values provide no geographic evidence; a consumer
	// that persists a geographically scoped record must fail closed rather
	// than defaulting them.
	MarketCode  string                `json:"market_code,omitempty"`
	CountryCode geography.CountryCode `json:"country_code,omitempty"`
	FailedAt    time.Time             `json:"failed_at"`
	RequestID   string                `json:"request_id,omitempty"`
}
