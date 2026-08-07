package packaging_enums

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
