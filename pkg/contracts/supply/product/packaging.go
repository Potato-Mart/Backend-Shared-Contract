package product

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/measurement"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/packaging/packaging_enums"

	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/supply/product/product_enums"
)

// ProductPackageOption identifies one stable physical package configuration
// for a product SKU. A changed case size or measurement receives a new ID.
type ProductPackageOption struct {
	ID              string                              `json:"id"`
	Code            string                              `json:"code"`
	SKUID           string                              `json:"sku_id"`
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
	ID              string                      `json:"id"`
	SKUID           string                      `json:"sku_id"`
	PackageOptionID string                      `json:"package_option_id"`
	Value           string                      `json:"value"`
	Format          product_enums.BarcodeFormat `json:"format"`
	ManufacturerID  string                      `json:"manufacturer_id,omitempty"`
	IsPrimary       bool                        `json:"is_primary"`
	EffectiveFrom   time.Time                   `json:"effective_from"`
	EffectiveTo     *time.Time                  `json:"effective_to,omitempty"`
}
