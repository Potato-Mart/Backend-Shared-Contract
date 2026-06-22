package membership

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v8/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v8/pkg/enums"
)

// PermissionMembershipPointsSpend is the permission a wholesale organisation
// access grant needs before it may spend organisation-owned points.
const PermissionMembershipPointsSpend = "membership.points.spend"

// PointLedgerEntry is a single points transaction for a membership account.
// Positive delta = earned; negative = redeemed or expired. Remaining tracks how
// many points from an earn row are still available for FIFO redemption.
type PointLedgerEntry struct {
	ID                        string                      `json:"id"`
	MembershipAccountID       string                      `json:"membership_account_id"`
	Owner                     MembershipOwnerRef          `json:"owner"`
	Delta                     int                         `json:"delta"`
	Reason                    enums.MembershipPointReason `json:"reason"`
	RelatedOrderID            string                      `json:"related_order_id,omitempty"`
	RelatedOrderNumber        string                      `json:"related_order_number,omitempty"`
	RelatedReservationID      string                      `json:"related_reservation_id,omitempty"`
	RelatedRewardRedemptionID string                      `json:"related_reward_redemption_id,omitempty"`
	BalanceAfter              int                         `json:"balance_after"`
	Remaining                 int                         `json:"remaining"`
	ExpiresAt                 *time.Time                  `json:"expires_at,omitempty"`
	Allocations               []PointAllocation           `json:"allocations,omitempty"`
	Note                      string                      `json:"note,omitempty"`
	CreatedBy                 string                      `json:"created_by,omitempty"`
	CreatedAt                 time.Time                   `json:"created_at"`
}

// PointAllocation shows which earned points row was consumed by a redemption
// or expiry ledger entry.
type PointAllocation struct {
	LedgerEntryID string     `json:"ledger_entry_id"`
	Points        int        `json:"points"`
	ExpiresAt     *time.Time `json:"expires_at,omitempty"`
}

// PointBucket is one currently available points batch for expiry-aware
// customer-facing and admin wallet views.
type PointBucket struct {
	Points              int                         `json:"points"`
	ExpiresAt           *time.Time                  `json:"expires_at,omitempty"`
	SourceLedgerEntryID string                      `json:"source_ledger_entry_id"`
	Reason              enums.MembershipPointReason `json:"reason"`
	RelatedOrderID      string                      `json:"related_order_id,omitempty"`
	RelatedOrderNumber  string                      `json:"related_order_number,omitempty"`
}

// PointBalanceBreakdown is a projected view of active point buckets.
type PointBalanceBreakdown struct {
	MembershipAccountID string             `json:"membership_account_id"`
	Owner               MembershipOwnerRef `json:"owner"`
	TotalPoints         int                `json:"total_points"`
	ReservedPoints      int                `json:"reserved_points"`
	AvailablePoints     int                `json:"available_points"`
	ExpiringPoints      int                `json:"expiring_points"`
	Buckets             []PointBucket      `json:"buckets,omitempty"`
	CalculatedAt        time.Time          `json:"calculated_at"`
}

// CanReserve reports whether the projected balance can reserve points without
// going negative.
func (b PointBalanceBreakdown) CanReserve(points int) bool {
	return points > 0 && b.AvailablePoints >= points
}

// PointReservation holds points during checkout or reward redemption. A
// reservation must be committed before a negative ledger entry is created.
type PointReservation struct {
	ID                        string                         `json:"id"`
	MembershipAccountID       string                         `json:"membership_account_id"`
	Owner                     MembershipOwnerRef             `json:"owner"`
	OrganisationAccessID      string                         `json:"organisation_access_id,omitempty"`
	Points                    int                            `json:"points"`
	Status                    enums.PointReservationStatus   `json:"status"`
	Reason                    enums.MembershipPointReason    `json:"reason"`
	RedemptionType            enums.MembershipRedemptionType `json:"redemption_type"`
	RelatedOrderID            string                         `json:"related_order_id,omitempty"`
	RelatedRewardID           string                         `json:"related_reward_id,omitempty"`
	RelatedRewardRedemptionID string                         `json:"related_reward_redemption_id,omitempty"`
	Allocations               []PointAllocation              `json:"allocations,omitempty"`
	ExpiresAt                 time.Time                      `json:"expires_at"`
	CommittedAt               *time.Time                     `json:"committed_at,omitempty"`
	CancelledAt               *time.Time                     `json:"cancelled_at,omitempty"`
	CancelReason              string                         `json:"cancel_reason,omitempty"`
	CreatedBy                 string                         `json:"created_by,omitempty"`
	CreatedAt                 time.Time                      `json:"created_at"`
}

// CanCommit reports whether a reservation is still eligible to become a
// committed ledger spend at the supplied time.
func (r PointReservation) CanCommit(now time.Time) bool {
	return r.Status == enums.PointReservationStatusReserved && !now.After(r.ExpiresAt)
}

// HasRequiredSpendAccess reports whether this reservation has the required
// access context for the owner. Retail owners do not need organisation access;
// wholesale organisation owners need both an access ID and spend permission.
func (r PointReservation) HasRequiredSpendAccess(hasSpendPermission bool) bool {
	if !r.Owner.RequiresOrganisationAccess() {
		return true
	}
	return r.OrganisationAccessID != "" && hasSpendPermission
}

// PointPromotion is a time-limited points multiplier event. The effective earn
// rate is tier multiplier times promotion multiplier.
type PointPromotion struct {
	ID             string                          `json:"id"`
	Name           string                          `json:"name"`
	Description    string                          `json:"description,omitempty"`
	Multiplier     float64                         `json:"multiplier"`
	StartAt        time.Time                       `json:"start_at"`
	EndAt          time.Time                       `json:"end_at"`
	AppliesTo      enums.MembershipPromotionTarget `json:"applies_to"`
	TargetTierKeys []string                        `json:"target_tier_keys,omitempty"`
	MinOrderAmount *common.Money                   `json:"min_order_amount,omitempty"`
	IsActive       bool                            `json:"is_active"`

	common.AuditFields `bson:",inline"`
}

// MemberCheckIn records a daily check-in for streak-based point awards.
type MemberCheckIn struct {
	ID                  string    `json:"id"`
	MembershipAccountID string    `json:"membership_account_id"`
	CheckInDate         time.Time `json:"check_in_date"`
	StreakCount         int       `json:"streak_count"`
	PointsAwarded       int       `json:"points_awarded"`
	CreatedAt           time.Time `json:"created_at"`
}
