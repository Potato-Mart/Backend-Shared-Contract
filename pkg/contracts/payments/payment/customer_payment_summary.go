package payment

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/payments/payment/payment_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/pricing/wallet/wallet_enums"
)

// CustomerPaymentSummary is the customer-safe projection of how an order was
// paid: totals, per-source allocations, points earned, and an explicit
// completeness marker when parts could not be assembled.
type CustomerPaymentSummary struct {
	OrderNumber   string                      `json:"order_number"`
	Status        payment_enums.PaymentStatus `json:"status"`
	Provider      string                      `json:"provider,omitempty"`
	ProcessorFee  *money.Money                `json:"processor_fee,omitempty"`
	CapturedTotal *money.Money                `json:"captured_total,omitempty"`
	RefundedTotal *money.Money                `json:"refunded_total,omitempty"`
	TotalPaid     *money.Money                `json:"total_paid,omitempty"`
	// AmountRemaining is the server-side amount the order still owes, at
	// exact minor units before any cash rounding: zero once the order is
	// fully paid. Payments sets it for POS split-tender orders; it is
	// omitted when Payments did not compute it.
	AmountRemaining   *money.Money                  `json:"amount_remaining,omitempty"`
	Allocations       []CustomerPaymentAllocation   `json:"allocations,omitempty"`
	PointsAwardStatus wallet_enums.PointAwardStatus `json:"points_award_status,omitempty"`
	// PointsHeld is prepared but excluded from the spendable wallet balance.
	PointsHeld          *int                                    `json:"points_held,omitempty"`
	PointsEarned        *int                                    `json:"points_earned,omitempty"`
	PointsAppliedToDebt *int                                    `json:"points_applied_to_debt,omitempty"`
	PointsNetCredited   *int                                    `json:"points_net_credited,omitempty"`
	PointDebtRemaining  *int                                    `json:"point_debt_remaining,omitempty"`
	PointsLedgerEntryID string                                  `json:"points_ledger_entry_id,omitempty"`
	PointsEarnedAt      *time.Time                              `json:"points_earned_at,omitempty"`
	Completeness        payment_enums.PaymentCompleteness       `json:"completeness"`
	MissingComponents   []payment_enums.PaymentSummaryComponent `json:"missing_components,omitempty"`
	PaidAt              *time.Time                              `json:"paid_at,omitempty"`
	UpdatedAt           *time.Time                              `json:"updated_at,omitempty"`
}
