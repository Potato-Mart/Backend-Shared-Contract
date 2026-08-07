package product

import (
	"time"

	common "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/common/shared"
)

// ProductPackageOption identifies one stable physical package configuration
// for a product SKU. A changed case size or measurement receives a new ID.
type ProductPackageOption struct {
	ID              string                     `json:"id"`
	Code            string                     `json:"code"`
	ProductSKUCode  string                     `json:"product_sku_code"`
	HandlingUnit    common.PackageHandlingUnit `json:"handling_unit"`
	UnitsPerPackage int64                      `json:"units_per_package"`
	Dimensions      *common.Dimensions         `json:"dimensions,omitempty"`
	Weight          *common.Weight             `json:"weight,omitempty"`
	IsCanonical     bool                       `json:"is_canonical"`
	IsActive        bool                       `json:"is_active"`
	EffectiveFrom   time.Time                  `json:"effective_from"`
	EffectiveTo     *time.Time                 `json:"effective_to,omitempty"`
}

// ProductPackageOptionSnapshot is the immutable package configuration carried
// by offers and transaction lines.
type ProductPackageOptionSnapshot struct {
	ID              string                     `json:"id"`
	Code            string                     `json:"code"`
	ProductSKUCode  string                     `json:"product_sku_code"`
	HandlingUnit    common.PackageHandlingUnit `json:"handling_unit"`
	UnitsPerPackage int64                      `json:"units_per_package"`
	Dimensions      *common.Dimensions         `json:"dimensions,omitempty"`
	Weight          *common.Weight             `json:"weight,omitempty"`
	EffectiveFrom   time.Time                  `json:"effective_from"`
	EffectiveTo     *time.Time                 `json:"effective_to,omitempty"`
	CapturedAt      time.Time                  `json:"captured_at"`
}

// ProductBarcodeAssignment binds one manufacturer barcode to one package
// option. Barcode values are not product identities and need not be unique
// across product SKUs.
type ProductBarcodeAssignment struct {
	ID              string        `json:"id"`
	ProductSKUCode  string        `json:"product_sku_code"`
	PackageOptionID string        `json:"package_option_id"`
	Value           string        `json:"value"`
	Format          BarcodeFormat `json:"format"`
	ManufacturerID  string        `json:"manufacturer_id,omitempty"`
	IsPrimary       bool          `json:"is_primary"`
	EffectiveFrom   time.Time     `json:"effective_from"`
	EffectiveTo     *time.Time    `json:"effective_to,omitempty"`
}

// ProductBarcodeAssignmentSnapshot freezes a barcode assignment used by a
// transaction or customer-safe catalogue projection.
type ProductBarcodeAssignmentSnapshot struct {
	ID              string        `json:"id"`
	ProductSKUCode  string        `json:"product_sku_code"`
	PackageOptionID string        `json:"package_option_id"`
	Value           string        `json:"value"`
	Format          BarcodeFormat `json:"format"`
	ManufacturerID  string        `json:"manufacturer_id,omitempty"`
	IsPrimary       bool          `json:"is_primary"`
	EffectiveFrom   time.Time     `json:"effective_from"`
	EffectiveTo     *time.Time    `json:"effective_to,omitempty"`
	CapturedAt      time.Time     `json:"captured_at"`
}
