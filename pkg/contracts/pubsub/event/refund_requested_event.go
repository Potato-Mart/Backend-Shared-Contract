package event

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/money"
)

// RefundRequestedEvent is emitted on the refund-events topic when a refund is
// raised against an order's captured payments.
type RefundRequestedEvent struct {
	RefundID    string      `json:"refund_id"`
	OrderID     string      `json:"order_id,omitempty"`
	OrderNumber string      `json:"order_number"`
	Amount      money.Money `json:"amount"`
	RequestedBy string      `json:"requested_by,omitempty"`
	// MarketCode and CountryCode are the denormalized geography the event
	// belongs to. Empty values provide no geographic evidence; a consumer
	// that persists a geographically scoped record must fail closed rather
	// than defaulting them.
	MarketCode  string                `json:"market_code,omitempty"`
	CountryCode geography.CountryCode `json:"country_code,omitempty"`
	RequestedAt time.Time             `json:"requested_at"`
	Reason      string                `json:"reason,omitempty"`
	RequestID   string                `json:"request_id,omitempty"`
}
