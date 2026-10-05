package product

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v44/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v44/pkg/contracts/supply/catalogue/product/product_enums"
)

// SellingProductBarcode is a safe barcode projection for an active package
// option. Manufacturer administration and effective-window metadata remain in
// the Product master.
type SellingProductBarcode struct {
	PackageOption packaging.PackageOptionRef  `json:"package_option"`
	Value         string                      `json:"value"`
	Format        product_enums.BarcodeFormat `json:"format"`
	IsPrimary     bool                        `json:"is_primary"`
}
