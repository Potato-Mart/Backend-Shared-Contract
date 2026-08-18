// Package cost holds the Supply-owned private operational cost models. Base
// acquisition cost and carrying cost are internal values: they never appear in
// a storefront, order, invoice, or receipt response.
package cost

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/money"
)

// BaseAcquisitionCost is the current tax-exclusive acquisition cost for one
// SKU in one currency. It is overwritten in place under an optimistic revision
// rather than versioned, and a change regenerates Pricing draft suggestions
// without disturbing approved sale prices.
type BaseAcquisitionCost struct {
	ID       string             `json:"id"`
	SKUCode  string             `json:"sku_code"`
	Currency money.CurrencyCode `json:"currency"`
	// Amount is tax exclusive. Recoverable input tax is never included.
	Amount        money.Money `json:"amount"`
	SourceType    string      `json:"source_type,omitempty"`
	SourceID      string      `json:"source_id,omitempty"`
	Revision      int64       `json:"revision"`
	EffectiveFrom time.Time   `json:"effective_from"`

	audit.AuditFields
}

// DepotCarryingCost is the exact weighted-average carrying cost for one SKU at
// one depot in one currency.
//
// TotalCarryingCostMinor and BaseUnitQuantity are the stored authority; the
// average is derived from them so no float ever enters the calculation.
// Depots aggregate only when their currency matches. Freight and duty are
// recorded separately on the supplier invoice and are never blended in.
type DepotCarryingCost struct {
	ID        string `json:"id"`
	SKUCode   string `json:"sku_code"`
	DepotCode string `json:"depot_code"`

	Currency         money.CurrencyCode     `json:"currency"`
	CurrencyExponent money.CurrencyExponent `json:"currency_exponent"`

	TotalCarryingCostMinor int64 `json:"total_carrying_cost_minor"`
	BaseUnitQuantity       int64 `json:"base_unit_quantity"`
	// CurrentAverageUnitCost is the derived per-base-unit average. It is
	// absent when BaseUnitQuantity is zero, which is also when the total
	// carrying cost is zero.
	CurrentAverageUnitCost *money.Money `json:"current_average_unit_cost,omitempty"`

	Revision int64     `json:"revision"`
	AsOf     time.Time `json:"as_of"`

	audit.AuditFields
}

// CarryingCostMovement is one auditable change to a depot carrying-cost
// balance. A supplier return reverses the cost of its originating receipt
// rather than the current average, so ReversesMovementID is required on a
// reversal and makes retries idempotent.
type CarryingCostMovement struct {
	ID        string `json:"id"`
	SKUCode   string `json:"sku_code"`
	DepotCode string `json:"depot_code"`

	ReferenceType string `json:"reference_type"`
	ReferenceID   string `json:"reference_id"`
	// IdempotencyKey makes a concurrent retry of the same receipt, return,
	// or transfer a no-op.
	IdempotencyKey string `json:"idempotency_key"`
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
