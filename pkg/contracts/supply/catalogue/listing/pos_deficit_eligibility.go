package listing

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/packaging/packaging_enums"
)

// POSDeficitEligibility is Supply's evidence that a POS sale line is eligible
// even though the depot cannot physically cover all of it. It is only issued
// for a depot that has a configured POS deficit location; the short quantity
// is booked against DeficitLocationCode when the sale settles. It asserts
// neither a reservation nor future stock, and the enclosing snapshot's
// validity token still governs revalidation.
type POSDeficitEligibility struct {
	DeficitLocationCode string                              `json:"deficit_location_code"`
	PackageOption       packaging.PackageOptionRef          `json:"package_option"`
	HandlingUnit        packaging_enums.PackageHandlingUnit `json:"handling_unit"`
	UnitsPerPackage     int64                               `json:"units_per_package"`
	ShortPackageCount   int64                               `json:"short_package_count"`
	ShortBaseUnits      int64                               `json:"short_base_units"`
}
