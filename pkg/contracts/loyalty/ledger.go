package loyalty

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v6/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v6/pkg/enums"
)

// LoyaltyLedgerEntry is a single points transaction for a customer.
// Positive delta = earned; negative = redeemed or expired.
// The remaining field tracks how many earned points from this row
// are still available for FIFO redemption.
type LoyaltyLedgerEntry struct {
	ID                 string                    `json:"id"`
	CustomerID         string                    `json:"customer_id"`
	Delta              int                       `json:"delta"`
	Reason             enums.LoyaltyLedgerReason `json:"reason"`
	RelatedOrderID     string                    `json:"related_order_id,omitempty"`
	RelatedOrderNumber string                    `json:"related_order_number,omitempty"`
	BalanceAfter       int                       `json:"balance_after"`
	Remaining          int                       `json:"remaining"` // unspent from this earn row (FIFO)
	ExpiresAt          *time.Time                `json:"expires_at,omitempty"`
	Allocations        []LoyaltyPointAllocation  `json:"allocations,omitempty"`
	Note               string                    `json:"note,omitempty"`
	CreatedBy          string                    `json:"created_by,omitempty"`
	CreatedAt          time.Time                 `json:"created_at"`
}

// LoyaltyPointAllocation shows which earned points row was consumed by a
// redemption or expiry ledger entry.
type LoyaltyPointAllocation struct {
	LedgerEntryID string     `json:"ledger_entry_id"`
	Points        int        `json:"points"`
	ExpiresAt     *time.Time `json:"expires_at,omitempty"`
}

// LoyaltyPointBucket is one currently available points batch for display in
// customer-facing and admin expiry views.
type LoyaltyPointBucket struct {
	Points              int                       `json:"points"`
	ExpiresAt           *time.Time                `json:"expires_at,omitempty"`
	SourceLedgerEntryID string                    `json:"source_ledger_entry_id"`
	Reason              enums.LoyaltyLedgerReason `json:"reason"`
	RelatedOrderID      string                    `json:"related_order_id,omitempty"`
	RelatedOrderNumber  string                    `json:"related_order_number,omitempty"`
}

// LoyaltyBalanceBreakdown is a projected view of active point buckets.
type LoyaltyBalanceBreakdown struct {
	CustomerID     string               `json:"customer_id"`
	TotalPoints    int                  `json:"total_points"`
	ExpiringPoints int                  `json:"expiring_points"`
	Buckets        []LoyaltyPointBucket `json:"buckets,omitempty"`
	CalculatedAt   time.Time            `json:"calculated_at"`
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

	common.AuditFields `bson:",inline"`
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
