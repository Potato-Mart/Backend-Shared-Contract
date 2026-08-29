package payments

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/money"
	analytics "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/insights/sales"
)

// RefundFact is the immutable analytical projection of a refund event.
type RefundFact struct {
	FactID      string                     `json:"fact_id"`
	RefundID    string                     `json:"refund_id"`
	OrderNumber string                     `json:"order_number"`
	Status      string                     `json:"status"`
	Amount      money.Money                `json:"amount"`
	Items       []analytics.RefundItemFact `json:"items,omitempty"`
	// MarketCode and CountryCode are the denormalized geography the event
	// belongs to. Empty values provide no geographic evidence; a consumer
	// that persists a geographically scoped record must fail closed rather
	// than defaulting them.
	MarketCode     string                `json:"market_code,omitempty"`
	CountryCode    geography.CountryCode `json:"country_code,omitempty"`
	FactOccurredAt time.Time             `json:"fact_occurred_at"`
}
