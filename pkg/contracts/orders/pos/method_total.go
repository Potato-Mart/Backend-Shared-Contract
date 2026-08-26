package pos

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/payments/payment/payment_enums"
)

// MethodTotal is one payment-method line of a session totals snapshot.
type MethodTotal struct {
	Method   payment_enums.PaymentMethod `json:"method"`
	Provider string                      `json:"provider,omitempty"`
	Amount   money.Money                 `json:"amount"`
	Count    int                         `json:"count"`
}
