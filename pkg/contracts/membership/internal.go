package membership

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v9/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v9/pkg/enums"
)

// Internal membership service endpoints. Provider: Backend-Management.
const (
	PathPointsQuote   = "/v1/internal/membership/points/quote"
	PathPointsReserve = "/v1/internal/membership/points/reserve"
	PathPointsCommit  = "/v1/internal/membership/points/commit"
	PathPointsCancel  = "/v1/internal/membership/points/cancel"
	PathRewardRedeem  = "/v1/internal/membership/rewards/redeem"
)

// PointQuoteCommand asks for the value and eligibility of a points spend before
// checkout or reward reservation.
type PointQuoteCommand struct {
	MembershipAccountID  string                         `json:"membership_account_id,omitempty"`
	Owner                *MembershipOwnerRef            `json:"owner,omitempty"`
	OrganisationAccessID string                         `json:"organisation_access_id,omitempty"`
	RedemptionType       enums.MembershipRedemptionType `json:"redemption_type"`
	Points               int                            `json:"points" binding:"required,gt=0"`
	Currency             string                         `json:"currency" binding:"required"`
	MerchandiseAmount    *common.Money                  `json:"merchandise_amount,omitempty"`
	RelatedOrderID       string                         `json:"related_order_id,omitempty"`
	RewardID             string                         `json:"reward_id,omitempty"`
}

// PointQuoteResult returns the projected discount and available balance.
type PointQuoteResult struct {
	MembershipAccountID string             `json:"membership_account_id"`
	Owner               MembershipOwnerRef `json:"owner"`
	Points              int                `json:"points"`
	DiscountAmount      common.Money       `json:"discount_amount"`
	AvailablePoints     int                `json:"available_points"`
	ExpiresAt           *time.Time         `json:"expires_at,omitempty"`
}

// PointReserveCommand reserves points against checkout or a reward redemption.
type PointReserveCommand struct {
	MembershipAccountID       string                         `json:"membership_account_id"`
	Owner                     MembershipOwnerRef             `json:"owner"`
	OrganisationAccessID      string                         `json:"organisation_access_id,omitempty"`
	Points                    int                            `json:"points" binding:"required,gt=0"`
	Reason                    enums.MembershipPointReason    `json:"reason"`
	RedemptionType            enums.MembershipRedemptionType `json:"redemption_type"`
	RelatedOrderID            string                         `json:"related_order_id,omitempty"`
	RelatedRewardID           string                         `json:"related_reward_id,omitempty"`
	RelatedRewardRedemptionID string                         `json:"related_reward_redemption_id,omitempty"`
	ExpiresAt                 time.Time                      `json:"expires_at"`
	CreatedBy                 string                         `json:"created_by,omitempty"`
}

type PointReserveResult struct {
	Reservation PointReservation      `json:"reservation"`
	Balance     PointBalanceBreakdown `json:"balance"`
}

type PointCommitCommand struct {
	ReservationID             string `json:"reservation_id" binding:"required"`
	RelatedOrderID            string `json:"related_order_id,omitempty"`
	RelatedOrderNumber        string `json:"related_order_number,omitempty"`
	RelatedRewardRedemptionID string `json:"related_reward_redemption_id,omitempty"`
	CreatedBy                 string `json:"created_by,omitempty"`
}

type PointCommitResult struct {
	LedgerEntry PointLedgerEntry      `json:"ledger_entry"`
	Balance     PointBalanceBreakdown `json:"balance"`
}

type PointCancelCommand struct {
	ReservationID string `json:"reservation_id" binding:"required"`
	Reason        string `json:"reason,omitempty"`
	CancelledBy   string `json:"cancelled_by,omitempty"`
}

type PointCancelResult struct {
	Reservation PointReservation      `json:"reservation"`
	Balance     PointBalanceBreakdown `json:"balance"`
}

// RewardRedeemCommand reserves points and creates a reward redemption in one
// Management-owned operation.
type RewardRedeemCommand struct {
	MembershipAccountID  string             `json:"membership_account_id"`
	Owner                MembershipOwnerRef `json:"owner"`
	OrganisationAccessID string             `json:"organisation_access_id,omitempty"`
	RewardID             string             `json:"reward_id" binding:"required"`
	RelatedOrderID       string             `json:"related_order_id,omitempty"`
	CreatedBy            string             `json:"created_by,omitempty"`
}

type RewardRedeemResult struct {
	Redemption  RewardRedemption      `json:"redemption"`
	Reservation PointReservation      `json:"reservation"`
	Balance     PointBalanceBreakdown `json:"balance"`
}
