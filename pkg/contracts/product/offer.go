package product

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/common"
	warehouseenum "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/enums/warehouse"
)

// SellableOfferDateMarkSnapshot freezes the lot date mark exposed by an offer.
type SellableOfferDateMarkSnapshot struct {
	Kind       warehouseenum.InventoryDateMarkKind `json:"kind"`
	DateMarkAt time.Time                           `json:"date_mark_at"`
	Timezone   string                              `json:"timezone"`
}

// SellableOfferDiscountSnapshot records one discount included in an offer.
type SellableOfferDiscountSnapshot struct {
	ID     string       `json:"id"`
	Code   string       `json:"code,omitempty"`
	Amount common.Money `json:"amount"`
}

// SellableOffer is a revisioned package-aware offer for one product at one
// depot. Source IDs are populated when the offer binds exact inventory.
type SellableOffer struct {
	ID                    string                             `json:"id"`
	ProductSKUCode        string                             `json:"product_sku_code"`
	DepotCode             string                             `json:"depot_code"`
	SourceBucketID        string                             `json:"source_bucket_id,omitempty"`
	SourceStockUnitID     string                             `json:"source_stock_unit_id,omitempty"`
	PackageOption         ProductPackageOptionSnapshot       `json:"package_option"`
	AvailablePackageCount int64                              `json:"available_package_count"`
	AvailableBaseUnits    int64                              `json:"available_base_units"`
	Condition             warehouseenum.InventoryCondition   `json:"condition"`
	Disposition           warehouseenum.InventoryDisposition `json:"disposition"`
	DateMark              *SellableOfferDateMarkSnapshot     `json:"date_mark,omitempty"`
	Revision              int64                              `json:"revision"`
	InventoryRevision     int64                              `json:"inventory_revision"`
	PackagePrice          common.Money                       `json:"package_price"`
	TaxAmount             common.Money                       `json:"tax_amount"`
	Discounts             []SellableOfferDiscountSnapshot    `json:"discounts,omitempty"`
	ValidFrom             time.Time                          `json:"valid_from"`
	ValidTo               *time.Time                         `json:"valid_to,omitempty"`
	Timezone              string                             `json:"timezone"`
	GeographicContext     common.GeographicContext           `json:"geographic_context"`
}

// SellableOfferSnapshot freezes an accepted sellable offer and its inventory,
// pricing, package, time, and geographic revisions.
type SellableOfferSnapshot struct {
	ID                    string                             `json:"id"`
	ProductSKUCode        string                             `json:"product_sku_code"`
	DepotCode             string                             `json:"depot_code"`
	SourceBucketID        string                             `json:"source_bucket_id,omitempty"`
	SourceStockUnitID     string                             `json:"source_stock_unit_id,omitempty"`
	PackageOption         ProductPackageOptionSnapshot       `json:"package_option"`
	AvailablePackageCount int64                              `json:"available_package_count"`
	AvailableBaseUnits    int64                              `json:"available_base_units"`
	Condition             warehouseenum.InventoryCondition   `json:"condition"`
	Disposition           warehouseenum.InventoryDisposition `json:"disposition"`
	DateMark              *SellableOfferDateMarkSnapshot     `json:"date_mark,omitempty"`
	Revision              int64                              `json:"revision"`
	InventoryRevision     int64                              `json:"inventory_revision"`
	PackagePrice          common.Money                       `json:"package_price"`
	TaxAmount             common.Money                       `json:"tax_amount"`
	Discounts             []SellableOfferDiscountSnapshot    `json:"discounts,omitempty"`
	ValidFrom             time.Time                          `json:"valid_from"`
	ValidTo               *time.Time                         `json:"valid_to,omitempty"`
	Timezone              string                             `json:"timezone"`
	GeographicContext     common.GeographicContext           `json:"geographic_context"`
	CapturedAt            time.Time                          `json:"captured_at"`
}
