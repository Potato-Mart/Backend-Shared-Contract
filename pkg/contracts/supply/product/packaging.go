package product

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/measurement"
	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/packaging/packaging_enums"

	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/supply/product/product_enums"
)

// ProductPackageOption identifies one stable physical package configuration
// for a product SKU. A changed case size or measurement receives a new ID.
type ProductPackageOption struct {
	Code            string                              `json:"code"`
	SKUCode         string                              `json:"sku_code"`
	HandlingUnit    packaging_enums.PackageHandlingUnit `json:"handling_unit"`
	UnitsPerPackage int64                               `json:"units_per_package"`
	Dimensions      *measurement.Dimensions             `json:"dimensions,omitempty"`
	Weight          *measurement.Weight                 `json:"weight,omitempty"`
	IsCanonical     bool                                `json:"is_canonical"`
	IsActive        bool                                `json:"is_active"`
	EffectiveFrom   time.Time                           `json:"effective_from"`
	EffectiveTo     *time.Time                          `json:"effective_to,omitempty"`
}

// ProductBarcodeAssignment binds one manufacturer barcode to one package
// option. Barcode values are not product identities and need not be unique
// across product SKUs.
type ProductBarcodeAssignment struct {
	Code              string                      `json:"code"`
	SKUCode           string                      `json:"sku_code"`
	PackageOptionCode string                      `json:"package_option_code"`
	Value             string                      `json:"value"`
	Format            product_enums.BarcodeFormat `json:"format"`
	ManufacturerCode  string                      `json:"manufacturer_code,omitempty"`
	IsPrimary         bool                        `json:"is_primary"`
	EffectiveFrom     time.Time                   `json:"effective_from"`
	EffectiveTo       *time.Time                  `json:"effective_to,omitempty"`
}
