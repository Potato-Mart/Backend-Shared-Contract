package event

import (
	"time"

	product "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/product"
)

// ProductSalesRollup is the product-stats topic payload: per-SKU, per-market
// sales statistics published on an interval by the sales-order owner and
// materialised by the catalog owner. AggregateID is the SKU ID. Rollups are
// separated by market and are never merged across markets.
type ProductSalesRollup struct {
	SKUCode    string                   `json:"sku_code"`
	MarketCode string                   `json:"market_code"`
	Last7Days  product.SalesWindowStats `json:"last_7_days"`
	Last30Days product.SalesWindowStats `json:"last_30_days"`
	Last90Days product.SalesWindowStats `json:"last_90_days"`
	Lifetime   product.SalesTotals      `json:"lifetime"`
	LastSoldAt *time.Time               `json:"last_sold_at,omitempty"`
	AsOf       time.Time                `json:"as_of"`
	Timezone   string                   `json:"timezone,omitempty"`
}
