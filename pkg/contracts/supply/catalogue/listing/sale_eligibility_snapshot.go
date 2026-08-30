package listing

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/warehouse"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/warehouse/warehouse_enums"
)

// SaleEligibilitySnapshot is the Supply-owned inventory and listing evidence
// captured for one SKU in one market at quote time. It carries no commercial
// price: Supply returns evidence and Pricing decides the money.
//
// ValidityToken is the opaque handle a caller returns when reserving stock or
// re-quoting. A stale listing, inventory, or eligibility revision requires a
// fresh quote rather than a silent recalculation.
type SaleEligibilitySnapshot struct {
	MarketCode      string `json:"market_code"`
	SKUCode         string `json:"sku_code"`
	ListingRevision int64  `json:"listing_revision"`
	TaxCategoryCode string `json:"tax_category_code"`

	DepotCode     string                     `json:"depot_code"`
	StockLocation warehouse.StockLocationRef `json:"stock_location"`
	BucketID      string                     `json:"bucket_id,omitempty"`
	StockUnitID   string                     `json:"stock_unit_id,omitempty"`
	LotID         string                     `json:"lot_id,omitempty"`

	Condition   warehouse_enums.InventoryCondition   `json:"condition"`
	Disposition warehouse_enums.InventoryDisposition `json:"disposition"`
	DateMark    *warehouse.InventoryDateMark         `json:"date_mark,omitempty"`
	// ExpiryLeadDays is the soon-expiry lead time that was in force, after
	// any listing override.
	ExpiryLeadDays int32               `json:"expiry_lead_days"`
	DamageApproval *DamageSaleApproval `json:"damage_approval,omitempty"`

	AvailableBaseUnits int64 `json:"available_base_units"`
	InventoryRevision  int64 `json:"inventory_revision"`

	ValidityToken string    `json:"validity_token"`
	ValidUntil    time.Time `json:"valid_until"`
	CapturedAt    time.Time `json:"captured_at"`
}
