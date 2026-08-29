package event

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/money"
)

// RefundFailedEvent is emitted on the refund-events topic when a refund
// attempt terminally fails.
type RefundFailedEvent struct {
	RefundID    string      `json:"refund_id"`
	OrderID     string      `json:"order_id,omitempty"`
	OrderNumber string      `json:"order_number"`
	Amount      money.Money `json:"amount,omitempty"`
	Reason      string      `json:"reason,omitempty"`
	// MarketCode and CountryCode are the denormalized geography the event
	// belongs to. Empty values provide no geographic evidence; a consumer
	// that persists a geographically scoped record must fail closed rather
	// than defaulting them.
	MarketCode  string                `json:"market_code,omitempty"`
	CountryCode geography.CountryCode `json:"country_code,omitempty"`
	FailedAt    time.Time             `json:"failed_at"`
	RequestID   string                `json:"request_id,omitempty"`
}
