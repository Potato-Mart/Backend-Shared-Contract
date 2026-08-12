package listing

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/supply/warehouse"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/supply/warehouse/warehouse_enums"
)

// DamageSaleApproval is the explicit quality decision that allows damaged
// inventory to be sold at a reduced price. Unsafe dispositions are never
// approved.
type DamageSaleApproval struct {
	QualityAssessmentID string                         `json:"quality_assessment_id"`
	Tier                warehouse_enums.DamageSaleTier `json:"tier"`
	ReasonCode          string                         `json:"reason_code"`
	ApprovedBy          string                         `json:"approved_by"`
	ApprovedAt          time.Time                      `json:"approved_at"`
}

// SaleEligibilitySnapshot is the Supply-owned inventory and listing evidence
// captured for one SKU in one market at quote time. It carries no commercial
// price: Supply returns evidence and Pricing decides the money.
//
// ValidityToken is the opaque handle a caller returns when reserving stock or
// re-quoting. A stale listing, inventory, or eligibility revision requires a
// fresh quote rather than a silent recalculation.
type SaleEligibilitySnapshot struct {
	MarketID        string `json:"market_id"`
	SKUID           string `json:"sku_id"`
	ListingID       string `json:"listing_id"`
	ListingRevision int64  `json:"listing_revision"`
	TaxCategoryID   string `json:"tax_category_id"`

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
