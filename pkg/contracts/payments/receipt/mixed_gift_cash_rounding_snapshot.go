package receipt

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/pricing/quote/quote_enums"
)

// MixedGiftCashRoundingSnapshot records cash rounding when the only tender
// sources are committed gift-card applications and a positive cash remainder.
// It does not cover point redemptions, electronic tenders, or other splits.
// Order consideration and tax remain unchanged.
//
// Consumers must verify OrderTotal equals the containing receipt's total, all
// amounts share its currency, and the positive gift total matches owner
// evidence. ExactCashAmountDue must be positive and equal OrderTotal less
// GiftCardAmount; CashAmountDue must be positive and equal that remainder plus
// RoundingAdjustment. With overflow-safe arithmetic, consumers must also
// verify the adjustment and rounded amount against the authoritative market
// policy, including its valid increment, mode and application time. The owning
// services validate this evidence; the record does not authorize payment or
// refund.
type MixedGiftCashRoundingSnapshot struct {
	OrderTotal         money.Money              `json:"order_total"`
	GiftCardAmount     money.Money              `json:"gift_card_amount"`
	ExactCashAmountDue money.Money              `json:"exact_cash_amount_due"`
	RoundingAdjustment money.Money              `json:"rounding_adjustment"`
	CashAmountDue      money.Money              `json:"cash_amount_due"`
	IncrementMinor     int64                    `json:"increment_minor"`
	Mode               quote_enums.RoundingMode `json:"mode"`
	AppliedAt          time.Time                `json:"applied_at"`
}
