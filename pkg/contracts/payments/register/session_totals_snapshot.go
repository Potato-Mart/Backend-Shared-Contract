package register

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/temporal"
	"time"
)

// SessionTotalsSnapshot is the X/Z read model for one register session:
// per-method sale totals, refunds, and cash movements at generation time.
type SessionTotalsSnapshot struct {
	SessionID         string        `json:"session_id"`
	RegisterID        string        `json:"register_id,omitempty"`
	DepotCode         string        `json:"depot_code,omitempty"`
	BusinessDate      temporal.Date `json:"business_date,omitempty"`
	MethodTotals      []MethodTotal `json:"method_totals,omitempty"`
	RefundTotal       money.Money   `json:"refund_total"`
	CashMovementTotal money.Money   `json:"cash_movement_total"`
	// ManualDiscountTotal and ManualDiscountCount summarize the cart-level
	// manual discounts on the session's sales. Both are omitted when the
	// session had none or the report predates them.
	ManualDiscountTotal *money.Money `json:"manual_discount_total,omitempty"`
	ManualDiscountCount int64        `json:"manual_discount_count,omitempty"`
	// SplitSaleCount is the number of the session's sales paid with a POS
	// split tender. Because MethodTotals counts payments, such a sale adds
	// one to the Count of every method it used. Omitted when the session
	// had none or the report predates it.
	SplitSaleCount int64     `json:"split_sale_count,omitempty"`
	GeneratedAt    time.Time `json:"generated_at"`
}
