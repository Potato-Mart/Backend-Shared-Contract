package product

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v43/pkg/contracts/common/measurement"
	"github.com/Potato-Mart/Backend-Shared-Contract/v43/pkg/contracts/common/packaging/packaging_enums"
)

// ProductPackageOption identifies one stable physical package configuration.
// SKUCode must match its enclosing Product when embedded there; it remains on
// the component so immutable order, POS, and supplier snapshots can identify
// the product when the component is used independently. A changed case size
// receives its canonical EACH/CASE{units} code. Changed physical specifications
// create a new immutable Version; references retain SKUCode+Code+Version.
type ProductPackageOption struct {
	Version         int64                               `json:"version"`
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
