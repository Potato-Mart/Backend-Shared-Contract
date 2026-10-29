package fulfilment

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v50/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v50/pkg/contracts/pricing/quote"
)

// PricedPackageComponent freezes evidence, quantity, and commercial totals
// for one CASE or EACH component. Exactly one snapshot is supplied: the
// Pricing-owned online quote or the journaled offline POS price. An Orders-owned
// digital gift face-value instrument uses a synthetic PriceSnapshot, not a
// Pricing quote; Orders verifies its purchase identity and classification.
// Transaction owners enforce this choice and reconcile the common totals.
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
