package receipt

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/pricing/quote/quote_enums"
)

// ReceiptManualDiscount is the customer-facing line for one cart-level manual
// discount, copied from the order's Pricing evidence. Amount is the gross
// amount Pricing applied to merchandise; the frozen receipt lines already
// carry its per-line allocation. It never names the staff member who set it.
type ReceiptManualDiscount struct {
	Mode       quote_enums.ManualDiscountMode `json:"mode"`
	ReasonCode string                         `json:"reason_code"`
	Reason     string                         `json:"reason,omitempty"`
	Amount     money.Money                    `json:"amount"`
}
