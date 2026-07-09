package membership

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v15/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v15/pkg/contracts/product"
	"github.com/Potato-Mart/Backend-Shared-Contract/v15/pkg/contracts/shared"
	membershipenum "github.com/Potato-Mart/Backend-Shared-Contract/v15/pkg/enums/membership"
)

// Reward is a catalog item that can be redeemed with membership points.
type Reward struct {
	ID                     string                               `json:"id"`
	Code                   string                               `json:"code"`
	Name                   string                               `json:"name"`
	Description            string                               `json:"description,omitempty"`
	Type                   membershipenum.MembershipRewardType  `json:"type"`
	PointsCost             int                                  `json:"points_cost"`
	DiscountAmount         *common.Money                        `json:"discount_amount,omitempty"`
	DiscountPercent        float64                              `json:"discount_percent,omitempty"`
	Product                *product.Snapshot                    `json:"product,omitempty"`
	VoucherCodePrefix      string                               `json:"voucher_code_prefix,omitempty"`
	StartsAt               *time.Time                           `json:"starts_at,omitempty"`
	ExpiresAt              *time.Time                           `json:"expires_at,omitempty"`
	IsActive               bool                                 `json:"is_active"`
	UsageLimit             int                                  `json:"usage_limit,omitempty"`
	UsedCount              int                                  `json:"used_count"`
	PerMemberLimit         int                                  `json:"per_member_limit,omitempty"`
	MinimumTierKey         string                               `json:"minimum_tier_key,omitempty"`
	TriggerTierKey         string                               `json:"trigger_tier_key,omitempty"`
	IssueOnTierAchievement bool                                 `json:"issue_on_tier_achievement,omitempty"`
	EligibleOwnerTypes     []membershipenum.MembershipOwnerType `json:"eligible_owner_types,omitempty"`
	Metadata               common.Metadata                      `json:"metadata,omitempty"`
	History                []shared.HistoryEntry                `json:"history,omitempty"`

	common.AuditFields
}

// RewardRedemption records a member's use of a catalog reward.
type RewardRedemption struct {
	ID                   string                                          `json:"id"`
	MembershipAccountID  string                                          `json:"membership_account_id"`
	Owner                MembershipOwnerRef                              `json:"owner"`
	OrganisationAccessID string                                          `json:"organisation_access_id,omitempty"`
	RewardCode           string                                          `json:"reward_code"`
	ReservationID        string                                          `json:"reservation_id,omitempty"`
	PointsSpent          int                                             `json:"points_spent"`
	Status               membershipenum.MembershipRewardRedemptionStatus `json:"status"`
	DiscountAmount       *common.Money                                   `json:"discount_amount,omitempty"`
	RelatedOrderNumber   string                                          `json:"related_order_number,omitempty"`
	VoucherCode          string                                          `json:"voucher_code,omitempty"`
	FulfilledAt          *time.Time                                      `json:"fulfilled_at,omitempty"`
	ExpiresAt            *time.Time                                      `json:"expires_at,omitempty"`
	CreatedBy            string                                          `json:"created_by,omitempty"`
	CreatedAt            time.Time                                       `json:"created_at"`
	History              []shared.HistoryEntry                           `json:"history,omitempty"`
}
