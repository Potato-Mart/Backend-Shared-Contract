package membership

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v24/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v24/pkg/contracts/common/money"

	"github.com/Potato-Mart/Backend-Shared-Contract/v24/pkg/contracts/pricing/membership/membership_enums"
)

// MembershipTier defines the qualification and benefit rules for the global
// membership programme. Spending points does not affect tier qualification.
type MembershipTier struct {
	TierKey               string                                `json:"tier_key"`
	Label                 string                                `json:"label"`
	QualificationMetric   membership_enums.MembershipTierMetric `json:"qualification_metric"`
	MinQualifyingSpend    money.Money                           `json:"min_qualifying_spend"`
	PointMultiplier       float64                               `json:"point_multiplier"`
	DiscountPercent       float64                               `json:"discount_percent"`
	FreeShippingThreshold *money.Money                          `json:"free_shipping_threshold,omitempty"`
	BirthdayBonusPoints   int                                   `json:"birthday_bonus_points"`
	// Benefits is the typed, localized benefit list.
	Benefits []TierBenefit `json:"benefits,omitempty"`
	IsActive bool          `json:"is_active"`
	IsSystem bool          `json:"is_system"`

	audit.AuditFields
}
