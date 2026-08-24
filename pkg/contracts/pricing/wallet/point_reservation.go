package wallet

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/pricing/wallet/wallet_enums"
)

// PointReservation holds customer wallet points until a redemption commits.
type PointReservation struct {
	ID                        string                              `json:"id"`
	CustomerNumber            string                              `json:"customer_number"`
	Points                    int                                 `json:"points"`
	Status                    wallet_enums.PointReservationStatus `json:"status"`
	Reason                    wallet_enums.PointLedgerReason      `json:"reason"`
	RedemptionType            wallet_enums.PointRedemptionType    `json:"redemption_type"`
	RelatedOrderID            string                              `json:"related_order_id,omitempty"`
	RelatedRewardCode         string                              `json:"related_reward_code,omitempty"`
	RelatedRewardRedemptionID string                              `json:"related_reward_redemption_id,omitempty"`
	Allocations               []PointAllocation                   `json:"allocations,omitempty"`
	ExpiresAt                 time.Time                           `json:"expires_at"`
	CommittedAt               *time.Time                          `json:"committed_at,omitempty"`
	CancelledAt               *time.Time                          `json:"cancelled_at,omitempty"`
	CancelReason              string                              `json:"cancel_reason,omitempty"`
	CreatedBy                 string                              `json:"created_by,omitempty"`
	CreatedAt                 time.Time                           `json:"created_at"`
}
