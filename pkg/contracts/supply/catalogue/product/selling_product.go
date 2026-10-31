package product

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/measurement"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/pricing/pricebook"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/catalogue/classification/classification_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/warehouse"
)

// SellingProduct is the customer-safe, market-scoped product projection
// published to storefront, app, and POS consumers. It is emitted only after
// the owning backend has resolved an authorized, approved, effective, and
// non-hidden SellingPrice for the requested market, channel, and audience.
//
// SellingProduct intentionally excludes catalogue identifiers, lifecycle and
// audit data, supply administration, stock, price-book identity/revision,
// source cost, wallet holdings, and promotion calculations. Checkout captures
// its own transaction price snapshot and never treats this projection as an
// immutable order price.
//
// ExpiryDate is an optional, display-only date mark for the product as
// offered. Kind reuses warehouse_enums.InventoryDateMarkKind: BEST_BEFORE is a
// quality date and EXPIRY is a safety (use-by) date. The owning backend
// decides which lot's date mark is shown; it is not stock or sale-eligibility
// evidence, and Supply's SaleEligibilitySnapshot remains authoritative at
// quote time. It is omitted when no date mark is known.
type SellingProduct struct {
	SKUCode            string                           `json:"sku_code"`
	SKUSeriesCode      string                           `json:"sku_series_code"`
	StorageType        classification_enums.StorageType `json:"storage_type"`
	Content            SellingProductContent            `json:"content"`
	Classification     SellingProductClassification     `json:"classification"`
	PackageOptions     []SellingProductPackageOption    `json:"package_options"`
	BarcodeAssignments []SellingProductBarcode          `json:"barcode_assignments,omitempty"`
	NetContent         *measurement.NetContent          `json:"net_content,omitempty"`
	ExpiryDate         *warehouse.InventoryDateMark     `json:"expiry_date,omitempty"`
	Price              pricebook.SellingPrice           `json:"price"`
}
