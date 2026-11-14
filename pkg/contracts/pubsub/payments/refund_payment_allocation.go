package payments

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/payments/payment/payment_enums"
)

// RefundPaymentAllocation is the money a completed refund returned through
// one of the order's payments: PaymentID names the Payment, Method its
// method, and Amount how much of the refund went back through it.
type RefundPaymentAllocation struct {
	PaymentID string                      `json:"payment_id"`
	Method    payment_enums.PaymentMethod `json:"method"`
	Amount    money.Money                 `json:"amount"`
}
