package pos

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/temporal"
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
	GeneratedAt       time.Time     `json:"generated_at"`
}
