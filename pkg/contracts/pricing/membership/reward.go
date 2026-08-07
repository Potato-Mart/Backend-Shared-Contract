package membership

import (
	"time"

	security "github.com/Potato-Mart/Backend-Shared-Contract/v24/pkg/contracts/common/security"

	"github.com/Potato-Mart/Backend-Shared-Contract/v24/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v24/pkg/contracts/supply/product"

	"github.com/Potato-Mart/Backend-Shared-Contract/v24/pkg/contracts/common/metadata"
	"github.com/Potato-Mart/Backend-Shared-Contract/v24/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v24/pkg/contracts/pricing/membership/membership_enums"
)

// Reward is a catalog item that can be redeemed with membership points.
type Reward struct {
	ID                     string                                `json:"id"`
	Code                   string                                `json:"code"`
	Name                   string                                `json:"name"`
	Description            string                                `json:"description,omitempty"`
	Type                   membership_enums.MembershipRewardType `json:"type"`
	PointsCost             int                                   `json:"points_cost"`
	DiscountAmount         *money.Money                          `json:"discount_amount,omitempty"`
	DiscountPercent        float64                               `json:"discount_percent,omitempty"`
	Product                *product.Snapshot                     `json:"product,omitempty"`
	VoucherCodePrefix      string                                `json:"voucher_code_prefix,omitempty"`
	StartsAt               *time.Time                            `json:"starts_at,omitempty"`
	ExpiresAt              *time.Time                            `json:"expires_at,omitempty"`
	IsActive               bool                                  `json:"is_active"`
	UsageLimit             int                                   `json:"usage_limit,omitempty"`
	UsedCount              int                                   `json:"used_count"`
	PerMemberLimit         int                                   `json:"per_member_limit,omitempty"`
	MinimumTierKey         string                                `json:"minimum_tier_key,omitempty"`
	TriggerTierKey         string                                `json:"trigger_tier_key,omitempty"`
	IssueOnTierAchievement bool                                  `json:"issue_on_tier_achievement,omitempty"`
	Metadata               metadata.Metadata                     `json:"metadata,omitempty"`
	History                []security.HistoryEntry               `json:"history,omitempty"`

	audit.AuditFields
}

// RewardRedemption records a member's use of a catalog reward.
type RewardRedemption struct {
	ID                 string                                            `json:"id"`
	CustomerNumber     string                                            `json:"customer_number"`
	RewardCode         string                                            `json:"reward_code"`
	ReservationID      string                                            `json:"reservation_id,omitempty"`
	PointsSpent        int                                               `json:"points_spent"`
	Status             membership_enums.MembershipRewardRedemptionStatus `json:"status"`
	DiscountAmount     *money.Money                                      `json:"discount_amount,omitempty"`
	RelatedOrderNumber string                                            `json:"related_order_number,omitempty"`
	VoucherCode        string                                            `json:"voucher_code,omitempty"`
	FulfilledAt        *time.Time                                        `json:"fulfilled_at,omitempty"`
	ExpiresAt          *time.Time                                        `json:"expires_at,omitempty"`
	CreatedBy          string                                            `json:"created_by,omitempty"`
	CreatedAt          time.Time                                         `json:"created_at"`
	History            []security.HistoryEntry                           `json:"history,omitempty"`
}
