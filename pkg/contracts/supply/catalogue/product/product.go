package product

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v51/pkg/contracts/common/measurement"

	"github.com/Potato-Mart/Backend-Shared-Contract/v51/pkg/contracts/supply/catalogue/classification"
	"github.com/Potato-Mart/Backend-Shared-Contract/v51/pkg/contracts/supply/catalogue/classification/classification_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v51/pkg/contracts/supply/catalogue/product/product_enums"
)

// Product is the canonical catalogue record for one sellable product identity.
// SKUCode is the product identifier; SKU series defines only its prefix. A
// Product carries its physical package and barcode facts but no commercial
// price, market availability, tax treatment, or customer-specific benefit.
// Pricing resolves those concerns into SellingProduct.
type Product struct {
	ID                 string                           `json:"id"`
	SKUCode            string                           `json:"sku_code"`
	StorageType        classification_enums.StorageType `json:"storage_type"`
	Status             product_enums.ProductStatus      `json:"status"`
	Content            ProductContent                   `json:"content"`
	Classification     ProductClassification            `json:"classification"`
	PackageOptions     []ProductPackageOption           `json:"package_options"`
	BarcodeAssignments []ProductBarcodeAssignment       `json:"barcode_assignments,omitempty"`
	NetContent         *measurement.NetContent          `json:"net_content,omitempty"`
	Supply             *classification.ProductSupply    `json:"supply,omitempty"`
	// Revision is advanced by every successful Supply write to the product
	// document and is the token admin clients echo back as expected_revision.
	// Zero only appears on documents written before revisions were backfilled.
	Revision       int64                  `json:"revision"`
	Administration *ProductAdministration `json:"administration,omitempty"`
}
