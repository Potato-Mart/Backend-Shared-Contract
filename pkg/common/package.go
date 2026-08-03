package common

// PackageHandlingUnit identifies how a product quantity is physically held.
// EACH represents loose base units and CASE represents intact product cases.
type PackageHandlingUnit string

const (
	PackageHandlingUnitEach PackageHandlingUnit = "EACH"
	PackageHandlingUnitCase PackageHandlingUnit = "CASE"
)

// IsValid reports whether u is a known PackageHandlingUnit value.
func (u PackageHandlingUnit) IsValid() bool {
	switch u {
	case PackageHandlingUnitEach, PackageHandlingUnitCase:
		return true
	}
	return false
}

func (u PackageHandlingUnit) String() string { return string(u) }

// PackageComponentSnapshot records one physical package component in base
// units. PackageCount counts intact packages represented by the component.
type PackageComponentSnapshot struct {
	PackageOptionID string              `json:"package_option_id"`
	HandlingUnit    PackageHandlingUnit `json:"handling_unit"`
	PackageCount    int64               `json:"package_count"`
	UnitsPerPackage int64               `json:"units_per_package"`
	BaseUnits       int64               `json:"base_units"`
}

// PackageCompositionSnapshot is the package-aware representation of a product
// quantity at a point in a sales or fulfilment workflow.
type PackageCompositionSnapshot struct {
	TotalBaseUnits int64                      `json:"total_base_units"`
	Components     []PackageComponentSnapshot `json:"components"`
}
