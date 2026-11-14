package order

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/orders/order/order_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/payments/payment/payment_enums"
)

// TenderPlanPart is one tender of a TenderPlanSnapshot.
//
// PartID identifies the part within the plan and TenderKey is the
// idempotency key under which the till collects it; Payments copies both,
// with the plan's PlanID and the part's Sequence, onto the Payment it records
// for the part. Sequence is the part's 1-based collection position.
//
// Amount is the exact share of the order this part settles, at exact minor
// units, and is the Amount of the part's Payment. AmountDue is what the till
// actually collects for the part: it equals Amount, except for a cash part
// rounded to the market's cash increment, where it is Amount plus
// RoundingAdjustment. RoundingAdjustment (which may be negative) is present
// only on that rounded cash part, and at most one part of a plan carries it;
// the receipt's SplitTenderCashRoundingSnapshot names the same part.
//
// PaymentID is the captured Payment for the part, absent until there is one.
// CashTender records the cash received and the change given for a cash part;
// its AmountDue equals the part's AmountDue.
type TenderPlanPart struct {
	PartID             string                       `json:"part_id"`
	TenderKey          string                       `json:"tender_key"`
	Sequence           int                          `json:"sequence"`
	Method             payment_enums.PaymentMethod  `json:"method"`
	Amount             money.Money                  `json:"amount"`
	AmountDue          money.Money                  `json:"amount_due"`
	RoundingAdjustment *money.Money                 `json:"rounding_adjustment,omitempty"`
	Status             order_enums.TenderPartStatus `json:"status"`
	PaymentID          string                       `json:"payment_id,omitempty"`
	CashTender         *OfflineCashTenderEvidence   `json:"cash_tender,omitempty"`
}
