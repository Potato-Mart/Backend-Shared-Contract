package shipping

import "github.com/Potato-Mart/Backend-Shared-Contract/v6/pkg/common"

// PackageLimits describes the physical package constraints for a shipping
// option. All units use the shared measurement primitives to avoid carrier,
// database, or service-specific unit assumptions.
type PackageLimits struct {
	MinPackageWeight        *common.Weight     `json:"min_package_weight,omitempty"`
	MaxPackageWeight        *common.Weight     `json:"max_package_weight,omitempty"`
	MaxShipmentWeight       *common.Weight     `json:"max_shipment_weight,omitempty"`
	MaxPackageDimensions    *common.Dimensions `json:"max_package_dimensions,omitempty"`
	MaxPackageLinearMM      int64              `json:"max_package_linear_mm,omitempty"`
	MaxPackageVolumeCubicMM int64              `json:"max_package_volume_cubic_mm,omitempty"`
	MaxPackageCount         int                `json:"max_package_count,omitempty"`
}
