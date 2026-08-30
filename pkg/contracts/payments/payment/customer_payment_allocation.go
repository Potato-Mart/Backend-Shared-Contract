package payment

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/payments/payment/payment_enums"
)

// CustomerPaymentAllocation is one customer-safe row of an order's payment
// composition. It is also the reusable invoice payment-row shape.
type CustomerPaymentAllocation struct {
	AllocationID   string                                      `json:"allocation_id,omitempty"`
	Kind           payment_enums.CustomerPaymentAllocationKind `json:"kind"`
	Method         payment_enums.PaymentMethod                 `json:"method,omitempty"`
	Provider       string                                      `json:"provider,omitempty"`
	Amount         money.Money                                 `json:"amount"`
	Points         *int                                        `json:"points,omitempty"`
	Status         string                                      `json:"status,omitempty"`
	OccurredAt     time.Time                                   `json:"occurred_at"`
	RefundedAmount *money.Money                                `json:"refunded_amount,omitempty"`
}
