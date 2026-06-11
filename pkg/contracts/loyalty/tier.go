package loyalty

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v4/pkg/common"
)

// LoyaltyTier defines the rules for a membership tier in the loyalty programme.
// The tier_key is a stable text PK (e.g. "standard", "silver", "gold", "platinum").
// Admin-defined tiers can be added beyond the four system defaults.
type LoyaltyTier struct {
	TierKey               string          `json:"tier_key"`
	Label                 string          `json:"label"`
	MinAnnualSpend        common.Money    `json:"min_annual_spend"`
	PointMultiplier       float64         `json:"point_multiplier"`                  // $1 spend = N points
	DiscountPercent       float64         `json:"discount_percent"`                  // 0-100
	FreeShippingThreshold *common.Money   `json:"free_shipping_threshold,omitempty"` // nil = no change; zero amount = always free
	BirthdayBonusPoints   int             `json:"birthday_bonus_points"`
	Perks                 common.Metadata `json:"perks,omitempty"`
	SortOrder             int             `json:"sort_order"`
	IsActive              bool            `json:"is_active"`
	IsSystem              bool            `json:"is_system"` // system tiers cannot be deleted

	common.AuditFields
}
