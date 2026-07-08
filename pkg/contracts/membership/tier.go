package membership

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v14/pkg/common"
	membershipenum "github.com/Potato-Mart/Backend-Shared-Contract/v14/pkg/enums/membership"
)

// MembershipTier defines the qualification and benefit rules for the global
// membership programme. Spending points does not affect tier qualification.
type MembershipTier struct {
	TierKey               string                              `json:"tier_key"`
	Label                 string                              `json:"label"`
	QualificationMetric   membershipenum.MembershipTierMetric `json:"qualification_metric"`
	MinQualifyingSpend    common.Money                        `json:"min_qualifying_spend"`
	PointMultiplier       float64                             `json:"point_multiplier"`
	DiscountPercent       float64                             `json:"discount_percent"`
	FreeShippingThreshold *common.Money                       `json:"free_shipping_threshold,omitempty"`
	BirthdayBonusPoints   int                                 `json:"birthday_bonus_points"`
	Perks                 common.Metadata                     `json:"perks,omitempty"`
	SortOrder             int                                 `json:"sort_order"`
	IsActive              bool                                `json:"is_active"`
	IsSystem              bool                                `json:"is_system"`

	common.AuditFields
}
