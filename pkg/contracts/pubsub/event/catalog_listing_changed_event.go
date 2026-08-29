package event

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/listing/listing_enums"
)

// CatalogListingChangedEvent is emitted on the catalog-events topic for every
// MarketListing lifecycle or configuration change. AggregateID is the
// immutable market-code/SKU-code composite, which keeps a listing's changes
// ordered without persisting a catalog-semantic database identifier.
//
// It carries the listing identity plus the revision that produced it so
// Pricing can invalidate drafts and cached quotes idempotently: a consumer
// applies the event only when Revision is newer than the revision it already
// holds. The event never carries a commercial price.
type CatalogListingChangedEvent struct {
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
