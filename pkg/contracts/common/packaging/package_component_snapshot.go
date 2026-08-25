package packaging

import "github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/packaging/packaging_enums"

// PackageComponentSnapshot records one physical package component in base
// units. PackageCount counts intact packages represented by the component.
type PackageComponentSnapshot struct {
	PackageOptionCode string                              `json:"package_option_code"`
	HandlingUnit      packaging_enums.PackageHandlingUnit `json:"handling_unit"`
	PackageCount      int64                               `json:"package_count"`
	UnitsPerPackage   int64                               `json:"units_per_package"`
	BaseUnits         int64                               `json:"base_units"`
}
