package procurement

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/money"
)

// CarryingCostMovement is one auditable change to a depot carrying-cost
// balance. A supplier return reverses the cost of its originating receipt
// rather than the current average, so ReversesMovementID is required on a
// reversal.
type CarryingCostMovement struct {
	ID        string `json:"id"`
	SKUCode   string `json:"sku_code"`
	DepotCode string `json:"depot_code"`

	ReferenceType string `json:"reference_type"`
	ReferenceID   string `json:"reference_id"`
	// ReversesMovementID links a reversal to the movement whose cost it
	// unwinds. Unlinked, duplicate, excessive, and negative-stock reversals
	// are rejected by Supply.
	ReversesMovementID string `json:"reverses_movement_id,omitempty"`

	BaseUnitDelta                 int64 `json:"base_unit_delta"`
	CarryingCostDelta             int64 `json:"carrying_cost_minor_delta"`
	BalanceBaseUnitsAfter         int64 `json:"balance_base_units_after"`
	BalanceCarryingCostMinorAfter int64 `json:"balance_carrying_cost_minor_after"`

	Currency   money.CurrencyCode `json:"currency"`
	Revision   int64              `json:"revision"`
	OccurredAt time.Time          `json:"occurred_at"`
}
