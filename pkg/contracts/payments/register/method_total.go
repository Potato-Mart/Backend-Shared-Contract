package register

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/payments/payment/payment_enums"
)

// MethodTotal is one payment-method line of a session totals snapshot.
//
// Amount is the total taken through the method and Count is the number of
// payments (tender lines) through it, not the number of sales: a POS
// split-tender sale paid by EFTPOS and cash counts once under each method.
// SessionTotalsSnapshot.SplitSaleCount counts such sales.
type MethodTotal struct {
	Method   payment_enums.PaymentMethod `json:"method"`
	Provider string                      `json:"provider,omitempty"`
	Amount   money.Money                 `json:"amount"`
	Count    int                         `json:"count"`
}
