package orders

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/payments/payment/payment_enums"
)

// OrderPaidTender is one captured part of a POS split-tender order, listed on
// order.paid: the part's method and its exact Amount (before any cash
// rounding).
type OrderPaidTender struct {
	Method payment_enums.PaymentMethod `json:"method"`
	Amount money.Money                 `json:"amount"`
}
