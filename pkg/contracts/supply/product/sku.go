package product

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/measurement"

	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/supply/product/product_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/supply/warehouse/warehouse_enums"
)

// SKU is the globally identifiable sellable and stockable inventory identity
// beneath one Product. One Product has many SKUs, for example a 375 mL can and
// a 600 mL bottle of the same drink.
//
// A SKU carries no market-specific price, tax treatment, or availability.
// Supplier cases and multipacks are package composition on this SKU, not
// separate SKUs, and they expand to base units for inventory and pricing.
// Damaged, expiring, or clearance units are inventory conditions on stock, not
// separate SKUs.
type SKU struct {
	ID        string `json:"id"`
	ProductID string `json:"product_id"`
	Code      string `json:"code"`

	PackageOptions     []ProductPackageOption      `json:"package_options"`
	BarcodeAssignments []ProductBarcodeAssignment  `json:"barcode_assignments,omitempty"`
	NetContent         *measurement.NetContent     `json:"net_content,omitempty"`
	StorageType        warehouse_enums.StorageType `json:"storage_type,omitempty"`
	Status             product_enums.SKUStatus     `json:"status"`

	audit.AuditFields
}
