package order

import (
	common "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/common/shared"
	promotionenum "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/pricing/promotion"
)

// VolumeDiscountTier defines a CASE-package threshold for a package option.
type VolumeDiscountTier struct {
	ID              string                                `json:"id"`
	Name            string                                `json:"name"`
	PackageOptionID string                                `json:"package_option_id"`
	MinCasePackages int64                                 `json:"min_case_packages"`
	DiscountPercent float64                               `json:"discount_percent"`
	AppliesTo       promotionenum.VolumeDiscountAppliesTo `json:"applies_to"`
	IsActive        bool                                  `json:"is_active"`
	common.AuditFields
}
