package procurement

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v35/pkg/contracts/common/money"
)

// NetGoodsValuationSnapshot is Supply's private net-goods valuation advice for
// one SKU, depot and source currency. It is separate from the published
// carrying-cost models: all tax, freight and duty are excluded here. No FX
// conversion or customer-price decision is represented by this model.
type NetGoodsValuationSnapshot struct {
	ID               string                 `json:"id"`
	SKUCode          string                 `json:"sku_code"`
	DepotCode        string                 `json:"depot_code"`
	Currency         money.CurrencyCode     `json:"currency"`
	CurrencyExponent money.CurrencyExponent `json:"currency_exponent"`

	// ValuedBaseUnits includes ProvisionalBaseUnits. UnvaluedBaseUnits is a
	// separate quantity, so total owned units are valued plus unvalued.
	ValuedBaseUnits      int64 `json:"valued_base_units"`
	ProvisionalBaseUnits int64 `json:"provisional_base_units"`
	UnvaluedBaseUnits    int64 `json:"unvalued_base_units"`
	NetGoodsCostMinor    int64 `json:"net_goods_cost_minor"`
	// AverageValuedUnitCost is derived from NetGoodsCostMinor divided by
	// ValuedBaseUnits only. It is absent with no valued units, including when
	// all stock is unvalued. Known zero-cost valued stock may carry zero.
	AverageValuedUnitCost *money.Money `json:"average_valued_unit_cost,omitempty"`
	// Consumed and rounding variances are reconciliation balances separate
	// from the on-hand NetGoodsCostMinor and its weighted average.
	ConsumedCostVarianceMinor int64     `json:"consumed_cost_variance_minor"`
	RoundingVarianceMinor     int64     `json:"rounding_variance_minor"`
	Revision                  int64     `json:"revision"`
	AsOf                      time.Time `json:"as_of"`
}
