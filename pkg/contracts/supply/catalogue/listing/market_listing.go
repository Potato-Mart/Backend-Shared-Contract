// Package listing holds the Supply-owned market availability models. A
// MarketListing answers whether one SKU may be sold in one market and carries
// the market-specific facts a quote needs; it never carries an authoritative
// commercial price. Prices live in Pricing price books.
package listing

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/localization"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/catalogue/listing/listing_enums"
)

// MarketListing links one MarketCode and one SKUCode with a lifecycle, the
// Pricing-owned tax category that applies in that market, market restrictions,
// and an optional soon-expiry lead-time override.
//
// TaxCategoryCode is a scalar reference to a Pricing-owned tax category; Supply
// validates that the market and tax category exist before activating a
// listing, and Pricing validates the SKU and an active listing before
// approving a price entry or quoting.
type MarketListing struct {
	ID         string `json:"id"`
	MarketCode string `json:"market_code"`
	// CountryCode is the denormalized country of MarketCode, carried so a
	// country-scoped staff query is a plain indexed match.
	CountryCode geography.CountryCode             `json:"country_code,omitempty"`
	SKUCode     string                            `json:"sku_code"`
	Status      listing_enums.MarketListingStatus `json:"status"`

	TaxCategoryCode string                       `json:"tax_category_code"`
	DisplayName     []localization.LocalizedName `json:"display_name,omitempty"`
	Restrictions    []SaleRestriction            `json:"restrictions,omitempty"`

	// ExpiryLeadDaysOverride replaces the market's default soon-expiry lead
	// time for this listing. Nil keeps the market default.
	ExpiryLeadDaysOverride *int32 `json:"expiry_lead_days_override,omitempty"`
	// UnitPricingRequired marks a listing that may not be activated without
	// the SKU's net quantity and standard measure.
	UnitPricingRequired bool `json:"unit_pricing_required"`

	AvailableFrom  time.Time              `json:"available_from"`
	AvailableUntil *time.Time             `json:"available_until,omitempty"`
	Activation     *audit.LifecycleAction `json:"activation,omitempty"`
	Delisting      *audit.LifecycleAction `json:"delisting,omitempty"`
	Revision       int64                  `json:"revision"`

	audit.AuditFields
}
