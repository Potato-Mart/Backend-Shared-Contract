package event

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/supply/listing/listing_enums"
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

// CatalogListingChangedEvent is emitted on the catalog-events topic for every
// MarketListing lifecycle or configuration change. AggregateID is the listing
// ID, which keeps a listing's changes ordered.
//
// It carries the listing identity plus the revision that produced it so
// Pricing can invalidate drafts and cached quotes idempotently: a consumer
// applies the event only when Revision is newer than the revision it already
// holds. The event never carries a commercial price.
type CatalogListingChangedEvent struct {
	ListingID  string `json:"listing_id"`
	MarketCode string `json:"market_code"`
	SKUCode    string `json:"sku_code"`

	PreviousStatus listing_enums.MarketListingStatus `json:"previous_status,omitempty"`
	Status         listing_enums.MarketListingStatus `json:"status"`

	TaxCategoryCode        string `json:"tax_category_code"`
	UnitPricingRequired    bool   `json:"unit_pricing_required"`
	ExpiryLeadDaysOverride *int32 `json:"expiry_lead_days_override,omitempty"`

	PreviousRevision int64 `json:"previous_revision"`
	Revision         int64 `json:"revision"`

	AvailableFrom  time.Time  `json:"available_from"`
	AvailableUntil *time.Time `json:"available_until,omitempty"`
	ChangedBy      string     `json:"changed_by,omitempty"`
	OccurredAt     time.Time  `json:"occurred_at"`
}
