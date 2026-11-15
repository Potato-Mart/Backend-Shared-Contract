package order

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
)

// GiftCardRedemptionSnapshot records one ordered gift-card allocation applied
// to an order. WalletTransactionID links the snapshot and completed gift-card
// payment to the committed wallet ledger entry.
//
// GiftCardID is Pricing's opaque wallet gift-card id (1-64 characters, never
// the card code). It is set only when a till tendered a card bound to the
// order's customer by tapping it in that customer's wallet, so the till never
// held the full code. Any copy of such a redemption returned to a till carries
// the masked code (for example GC********1234) in GiftCardCode, never the full
// code. Redemptions tendered by code omit GiftCardID and are unchanged.
type GiftCardRedemptionSnapshot struct {
	GiftCardCode        string      `json:"gift_card_code"`
	GiftCardID          string      `json:"gift_card_id,omitempty"`
	AppliedAmount       money.Money `json:"applied_amount"`
	ReservationID       string      `json:"reservation_id,omitempty"`
	WalletTransactionID string      `json:"wallet_transaction_id,omitempty"`
	OccurredAt          *time.Time  `json:"occurred_at,omitempty"`
}
