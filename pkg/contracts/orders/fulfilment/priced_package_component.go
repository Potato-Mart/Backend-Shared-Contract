package fulfilment

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/pricing/quote"
)

// PricedPackageComponent freezes the immutable Pricing snapshot, requested
// quantity, and commercial totals for one CASE or EACH component. PriceSnapshot
// is the authoritative historical value; the money fields beside it are the
// frozen line projection Orders reconciles against.
type PricedPackageComponent struct {
	PriceSnapshot         quote.PriceSnapshot `json:"price_snapshot"`
	RequestedPackageCount int64               `json:"requested_package_count"`
	RequestedBaseUnits    int64               `json:"requested_base_units"`
	PackagePrice          money.Money         `json:"package_price"`
	TaxAmount             money.Money         `json:"tax_amount"`
	DiscountAmount        money.Money         `json:"discount_amount"`
	ComponentTotal        money.Money         `json:"component_total"`
}
