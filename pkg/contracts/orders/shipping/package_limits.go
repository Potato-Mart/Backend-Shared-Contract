package shipping

import "github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/measurement"

// PackageLimits describes the physical package constraints for a shipping
// option. All units use the shared measurement primitives to avoid carrier-
// or service-specific unit assumptions.
type PackageLimits struct {
	MinPackageWeight        *measurement.Weight     `json:"min_package_weight,omitempty"`
	MaxPackageWeight        *measurement.Weight     `json:"max_package_weight,omitempty"`
	MaxShipmentWeight       *measurement.Weight     `json:"max_shipment_weight,omitempty"`
	MaxPackageDimensions    *measurement.Dimensions `json:"max_package_dimensions,omitempty"`
	MaxPackageLinearMM      int64                   `json:"max_package_linear_mm,omitempty"`
	MaxPackageVolumeCubicMM int64                   `json:"max_package_volume_cubic_mm,omitempty"`
	MaxPackageCount         int                     `json:"max_package_count,omitempty"`
}
