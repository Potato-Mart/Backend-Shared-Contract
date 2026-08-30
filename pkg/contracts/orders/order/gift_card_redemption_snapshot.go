package order

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/money"
)

// GiftCardRedemptionSnapshot records one ordered gift-card allocation applied
// to an order. WalletTransactionID links the snapshot and completed gift-card
// payment to the committed wallet ledger entry.
type GiftCardRedemptionSnapshot struct {
	GiftCardCode        string      `json:"gift_card_code"`
	AppliedAmount       money.Money `json:"applied_amount"`
	ReservationID       string      `json:"reservation_id,omitempty"`
	WalletTransactionID string      `json:"wallet_transaction_id,omitempty"`
	OccurredAt          *time.Time  `json:"occurred_at,omitempty"`
}
