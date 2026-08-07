package membership

import (
	common "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/common/shared"
)

// MembershipTier defines the qualification and benefit rules for the global
// membership programme. Spending points does not affect tier qualification.
type MembershipTier struct {
	TierKey               string               `json:"tier_key"`
	Label                 string               `json:"label"`
	QualificationMetric   MembershipTierMetric `json:"qualification_metric"`
	MinQualifyingSpend    common.Money         `json:"min_qualifying_spend"`
	PointMultiplier       float64              `json:"point_multiplier"`
	DiscountPercent       float64              `json:"discount_percent"`
	FreeShippingThreshold *common.Money        `json:"free_shipping_threshold,omitempty"`
	BirthdayBonusPoints   int                  `json:"birthday_bonus_points"`
	// Benefits is the typed, localized benefit list.
	Benefits []TierBenefit `json:"benefits,omitempty"`
	IsActive bool          `json:"is_active"`
	IsSystem bool          `json:"is_system"`

	common.AuditFields
}
