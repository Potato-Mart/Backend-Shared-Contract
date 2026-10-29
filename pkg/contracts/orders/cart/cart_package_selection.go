package cart

import "github.com/Potato-Mart/Backend-Shared-Contract/v49/pkg/contracts/common/packaging"

// CartPackageSelection identifies the quantity selected before pricing.
// It does not assert a price, stock availability, or sale eligibility.
type CartPackageSelection struct {
	PackageOption packaging.PackageOptionRef `json:"package_option"`
	PackageCount  int64                      `json:"package_count"`
	BaseUnits     int64                      `json:"base_units"`
}
