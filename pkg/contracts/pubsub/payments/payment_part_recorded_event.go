package payments

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/payments/payment/payment_enums"
)

// PaymentPartRecordedEvent is payment.part_recorded on the payment-events
// topic. AggregateID is the order number. Payments emits one for every
// captured part of a POS split-tender order (order.TenderPlanSnapshot),
// including the last one; Orders projects the order's partially paid state
// from it. payment.captured keeps its meaning: it is still emitted once, when
// the order is fully paid, with Amount equal to the order total.
//
// EventID is the stable identity of this part record, the same on every
// redelivery. PaymentID, Method, Provider and Amount describe the part's
// captured Payment; Amount is the part's exact Amount, before any cash
// rounding. TenderKey and TenderSequence name the plan part
// (Payment.TenderKey and Payment.TenderSequence). AmountRemaining is what the
// order still owes after this part, at exact minor units, in Amount's
// currency; OrderSettled is true exactly when it is zero, that is when this
// part completed the order.
type PaymentPartRecordedEvent struct {
	EventID         string                      `json:"event_id"`
	OrderNumber     string                      `json:"order_number"`
	PaymentID       string                      `json:"payment_id"`
	Method          payment_enums.PaymentMethod `json:"method"`
	Provider        string                      `json:"provider,omitempty"`
	Amount          money.Money                 `json:"amount"`
	TenderKey       string                      `json:"tender_key"`
	TenderSequence  int                         `json:"tender_sequence"`
	OrderSettled    bool                        `json:"order_settled"`
	AmountRemaining money.Money                 `json:"amount_remaining"`
	OccurredAt      time.Time                   `json:"occurred_at"`
}
