package event

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/money"
)

// CatalogBaseCostChangedEvent is emitted on the catalog-events topic when a
// SKU's base acquisition cost is overwritten. AggregateID is the SKU ID.
//
// The event is idempotent: consumers dedupe on the envelope event ID together
// with Revision, so a redelivered or out-of-order message can never regenerate
// a stale suggestion. Receiving it regenerates or replaces Pricing draft
// suggestions and never activates a draft or disturbs an approved price.
type CatalogBaseCostChangedEvent struct {
	SKUCode  string             `json:"sku_code"`
	Currency money.CurrencyCode `json:"currency"`
	// PreviousAmount is absent the first time a cost is recorded.
	PreviousAmount *money.Money `json:"previous_amount,omitempty"`
	// Amount is tax exclusive and is never exposed to customers.
	Amount           money.Money `json:"amount"`
	PreviousRevision int64       `json:"previous_revision"`
	Revision         int64       `json:"revision"`
	SourceType       string      `json:"source_type,omitempty"`
	SourceID         string      `json:"source_id,omitempty"`
	EffectiveFrom    time.Time   `json:"effective_from"`
	OccurredAt       time.Time   `json:"occurred_at"`
}
