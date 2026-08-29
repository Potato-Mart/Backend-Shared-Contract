package product

import "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/product/product_enums"

// SellingProductBarcode is a safe barcode projection for an active package
// option. Manufacturer administration and effective-window metadata remain in
// the Product master.
type SellingProductBarcode struct {
	PackageOptionCode string                      `json:"package_option_code"`
	Value             string                      `json:"value"`
	Format            product_enums.BarcodeFormat `json:"format"`
	IsPrimary         bool                        `json:"is_primary"`
}
