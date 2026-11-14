package payments

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/payments/payment/payment_enums"
)

// CapturedTender is one captured part of a POS split-tender order, listed on
// payment.captured: the part's Payment, its method and provider, and its
// exact Amount (before any cash rounding).
type CapturedTender struct {
	PaymentID string                      `json:"payment_id"`
	Method    payment_enums.PaymentMethod `json:"method"`
	Provider  string                      `json:"provider,omitempty"`
	Amount    money.Money                 `json:"amount"`
}
