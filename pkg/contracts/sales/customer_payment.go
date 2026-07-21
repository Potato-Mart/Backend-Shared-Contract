package sales

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v18/pkg/common"
	paymentenum "github.com/Potato-Mart/Backend-Shared-Contract/v18/pkg/enums/payment"
)

// CustomerPaymentSummary is the customer-safe projection of how an order was
// paid: totals, per-source allocations, points earned, and an explicit
// completeness marker when parts could not be assembled.
type CustomerPaymentSummary struct {
	OrderNumber         string                                `json:"order_number"`
	Status              paymentenum.PaymentStatus             `json:"status"`
	Provider            string                                `json:"provider,omitempty"`
	ProcessorFee        *common.Money                         `json:"processor_fee,omitempty"`
	CapturedTotal       *common.Money                         `json:"captured_total,omitempty"`
	RefundedTotal       *common.Money                         `json:"refunded_total,omitempty"`
	TotalPaid           *common.Money                         `json:"total_paid,omitempty"`
	Allocations         []CustomerPaymentAllocation           `json:"allocations,omitempty"`
	PointsEarned        *int                                  `json:"points_earned,omitempty"`
	PointsLedgerEntryID string                                `json:"points_ledger_entry_id,omitempty"`
	PointsEarnedAt      *time.Time                            `json:"points_earned_at,omitempty"`
	Completeness        paymentenum.PaymentCompleteness       `json:"completeness"`
	MissingComponents   []paymentenum.PaymentSummaryComponent `json:"missing_components,omitempty"`
	PaidAt              *time.Time                            `json:"paid_at,omitempty"`
	UpdatedAt           *time.Time                            `json:"updated_at,omitempty"`
}

// CustomerPaymentAllocation is one customer-safe row of an order's payment
// composition. It is also the reusable invoice payment-row shape.
type CustomerPaymentAllocation struct {
	AllocationID   string                                    `json:"allocation_id,omitempty"`
	Kind           paymentenum.CustomerPaymentAllocationKind `json:"kind"`
	Method         paymentenum.PaymentMethod                 `json:"method,omitempty"`
	Provider       string                                    `json:"provider,omitempty"`
	Amount         common.Money                              `json:"amount"`
	Points         *int                                      `json:"points,omitempty"`
	Status         string                                    `json:"status,omitempty"`
	OccurredAt     time.Time                                 `json:"occurred_at"`
	RefundedAmount *common.Money                             `json:"refunded_amount,omitempty"`
}

// InvoiceEmailDelivery records one customer-requested invoice email
// redelivery and its lifecycle state.
type InvoiceEmailDelivery struct {
	InvoiceNumber string                          `json:"invoice_number"`
	DeliveryID    string                          `json:"delivery_id"`
	Status        paymentenum.InvoiceResendStatus `json:"status"`
	RecipientHint string                          `json:"recipient_hint,omitempty"`
	RequestedAt   time.Time                       `json:"requested_at"`
}
