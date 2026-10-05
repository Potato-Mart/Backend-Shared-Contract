package product

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v43/pkg/contracts/common/packaging"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v43/pkg/contracts/supply/catalogue/product/product_enums"
)

// ProductBarcodeAssignment binds one manufacturer barcode to one package
// option. SKUCode must match its enclosing Product when embedded there; it
// remains for independent catalogued-component and snapshot identity. Barcode
// values are not product identities and need not be unique across product SKU
// codes.
type ProductBarcodeAssignment struct {
	Code             string                      `json:"code"`
	SKUCode          string                      `json:"sku_code"`
	PackageOption    packaging.PackageOptionRef  `json:"package_option"`
	Value            string                      `json:"value"`
	Format           product_enums.BarcodeFormat `json:"format"`
	ManufacturerCode string                      `json:"manufacturer_code,omitempty"`
	IsPrimary        bool                        `json:"is_primary"`
	EffectiveFrom    time.Time                   `json:"effective_from"`
	EffectiveTo      *time.Time                  `json:"effective_to,omitempty"`
}
