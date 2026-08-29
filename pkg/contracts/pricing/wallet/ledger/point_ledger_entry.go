package ledger

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/pricing/wallet/points"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/pricing/wallet/wallet_enums"
)

// PointLedgerEntry is an immutable customer-wallet points transaction.
type PointLedgerEntry struct {
	ID                        string                         `json:"id"`
	CustomerNumber            string                         `json:"customer_number"`
	Delta                     int                            `json:"delta"`
	Reason                    wallet_enums.PointLedgerReason `json:"reason"`
	RelatedOrderNumber        string                         `json:"related_order_number,omitempty"`
	RelatedReservationID      string                         `json:"related_reservation_id,omitempty"`
	RelatedRewardRedemptionID string                         `json:"related_reward_redemption_id,omitempty"`
	BalanceAfter              int                            `json:"balance_after"`
	DebtDelta                 *int                           `json:"debt_delta,omitempty"`
	DebtAfter                 *int                           `json:"debt_after,omitempty"`
	Remaining                 int                            `json:"remaining"`
	ExpiresAt                 *time.Time                     `json:"expires_at,omitempty"`
	Allocations               []points.PointAllocation       `json:"allocations,omitempty"`
	Note                      string                         `json:"note,omitempty"`
	CreatedBy                 string                         `json:"created_by,omitempty"`
	CreatedAt                 time.Time                      `json:"created_at"`
}
