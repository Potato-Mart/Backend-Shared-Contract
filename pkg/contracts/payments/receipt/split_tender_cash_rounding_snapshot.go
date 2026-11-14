package receipt

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/pricing/quote/quote_enums"
)

// SplitTenderCashRoundingSnapshot records cash rounding on a POS split-tender
// sale: a tender plan (order.TenderPlanSnapshot) whose cash part is rounded to
// the market's cash increment, alongside non-cash tenders (card or EFTPOS
// parts, gift cards, points). TenderPartID names that cash part; its
// RoundingAdjustment and AmountDue equal the ones here. Order consideration
// and tax remain unchanged, and the non-cash tenders stay exact.
//
// OrderTotal equals the containing receipt's total, and every amount shares
// its currency. NonCashAmount is everything not paid in cash: the plan's
// non-cash parts, gift-card redemptions and the points redemption.
// ExactCashAmountDue is OrderTotal less NonCashAmount and must be positive;
// CashAmountDue is ExactCashAmountDue plus RoundingAdjustment and must be
// positive. With overflow-safe arithmetic, consumers must verify the
// adjustment and rounded amount against the authoritative market policy,
// including its increment, mode and application time.
//
// A receipt carries at most one of CashRounding, MixedGiftCashRounding and
// SplitTenderCashRounding. The owning services validate this evidence; the
// record does not authorize payment or refund.
type SplitTenderCashRoundingSnapshot struct {
	OrderTotal         money.Money              `json:"order_total"`
	NonCashAmount      money.Money              `json:"non_cash_amount"`
	ExactCashAmountDue money.Money              `json:"exact_cash_amount_due"`
	RoundingAdjustment money.Money              `json:"rounding_adjustment"`
	CashAmountDue      money.Money              `json:"cash_amount_due"`
	IncrementMinor     int64                    `json:"increment_minor"`
	Mode               quote_enums.RoundingMode `json:"mode"`
	AppliedAt          time.Time                `json:"applied_at"`
	TenderPartID       string                   `json:"tender_part_id"`
}
