package product

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/measurement"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/packaging/packaging_enums"
)

// SellingProductPackageOption is an active package option that a customer may
// select. Lifecycle and effective-window metadata stay in the Product master.
type SellingProductPackageOption struct {
	Code            string                              `json:"code"`
	HandlingUnit    packaging_enums.PackageHandlingUnit `json:"handling_unit"`
	UnitsPerPackage int64                               `json:"units_per_package"`
	Dimensions      *measurement.Dimensions             `json:"dimensions,omitempty"`
	Weight          *measurement.Weight                 `json:"weight,omitempty"`
	IsCanonical     bool                                `json:"is_canonical"`
}
