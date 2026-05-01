package loyalty

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v3/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v3/pkg/enums"
)

// LoyaltyLedgerEntry is a single points transaction for a customer.
// Positive delta = earned; negative = redeemed or expired.
// The remaining field tracks how many earned points from this row
// are still available for FIFO redemption.
type LoyaltyLedgerEntry struct {
	ID                 string                    `json:"id"`
	CustomerProfileID  string                    `json:"customer_profile_id"`
	Delta              int                       `json:"delta"`
	Reason             enums.LoyaltyLedgerReason `json:"reason"`
	RelatedOrderID     string                    `json:"related_order_id,omitempty"`
	RelatedOrderNumber string                    `json:"related_order_number,omitempty"`
	BalanceAfter       int                       `json:"balance_after"`
	Remaining          int                       `json:"remaining"` // unspent from this earn row (FIFO)
	ExpiresAt          *time.Time                `json:"expires_at,omitempty"`
	Note               string                    `json:"note,omitempty"`
	CreatedBy          string                    `json:"created_by,omitempty"`
	CreatedAt          time.Time                 `json:"created_at"`
}

// LoyaltyPromotion is a time-limited points multiplier event.
// The effective earn rate is tier_multiplier × promo_multiplier.
type LoyaltyPromotion struct {
	ID             string                       `json:"id"`
	Name           string                       `json:"name"`
	Description    string                       `json:"description,omitempty"`
	Multiplier     float64                      `json:"multiplier"`
	StartAt        time.Time                    `json:"start_at"`
	EndAt          time.Time                    `json:"end_at"`
	AppliesTo      enums.LoyaltyPromotionTarget `json:"applies_to"`
	TargetTierKeys []string                     `json:"target_tier_keys,omitempty"` // for TIER_SPECIFIC
	MinOrderAmount *common.Money                `json:"min_order_amount,omitempty"`
	IsActive       bool                         `json:"is_active"`
	CreatedAt      time.Time                    `json:"created_at"`
	UpdatedAt      time.Time                    `json:"updated_at"`
}

// CustomerCheckIn records a daily check-in for streak-based point awards.
type CustomerCheckIn struct {
	ID                string    `json:"id"`
	CustomerProfileID string    `json:"customer_profile_id"`
	CheckInDate       time.Time `json:"check_in_date"`
	StreakCount       int       `json:"streak_count"`
	PointsAwarded     int       `json:"points_awarded"`
	CreatedAt         time.Time `json:"created_at"`
}
