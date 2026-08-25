package cost

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/money"
)

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
