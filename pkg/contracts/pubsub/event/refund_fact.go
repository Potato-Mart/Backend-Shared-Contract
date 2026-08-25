package event

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/money"
	analytics "github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/insights/analytics"
)

// RefundFact is the immutable analytical projection of a refund event.
type RefundFact struct {
	EventID     string                     `json:"event_id"`
	RefundID    string                     `json:"refund_id"`
	OrderNumber string                     `json:"order_number"`
	Status      string                     `json:"status"`
	Amount      money.Money                `json:"amount"`
	Items       []analytics.RefundItemFact `json:"items,omitempty"`
	// MarketCode and CountryCode are the denormalized geography the event
	// belongs to. They are absent on every event published before v28.0.0;
	// a consumer that persists a geographically scoped record treats an
	// absent value as "no evidence" and fails closed rather than defaulting.
	MarketCode  string                `json:"market_code,omitempty"`
	CountryCode geography.CountryCode `json:"country_code,omitempty"`
	OccurredAt  time.Time             `json:"occurred_at"`
}
