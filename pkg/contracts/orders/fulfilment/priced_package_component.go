package fulfilment

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v48/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v48/pkg/contracts/pricing/quote"
)

// PricedPackageComponent freezes evidence, quantity, and commercial totals
// for one CASE or EACH component. Exactly one snapshot is supplied: the
// Pricing-owned online quote or the journaled offline POS price. Transaction
// owners enforce this choice and reconcile the common frozen totals.
type PricedPackageComponent struct {
	PriceSnapshot         *quote.PriceSnapshot        `json:"price_snapshot,omitempty"`
	OfflinePriceSnapshot  *quote.OfflinePriceSnapshot `json:"offline_price_snapshot,omitempty"`
	RequestedPackageCount int64                       `json:"requested_package_count"`
	RequestedBaseUnits    int64                       `json:"requested_base_units"`
	PackagePrice          money.Money                 `json:"package_price"`
	TaxAmount             money.Money                 `json:"tax_amount"`
	DiscountAmount        money.Money                 `json:"discount_amount"`
	ComponentTotal        money.Money                 `json:"component_total"`
}
