package product

import (
	"time"
)

// ProductSalesRollup is the product-stats topic payload: per-SKU sales
// statistics published on an interval by the sales-order owner and
// materialised by the catalog owner. AggregateID is the SKU code.
type ProductSalesRollup struct {
	ProductSKUCode string           `json:"product_sku_code"`
	Last7Days      SalesWindowStats `json:"last_7_days"`
	Last30Days     SalesWindowStats `json:"last_30_days"`
	Last90Days     SalesWindowStats `json:"last_90_days"`
	Lifetime       SalesTotals      `json:"lifetime"`
	LastSoldAt     *time.Time       `json:"last_sold_at,omitempty"`
	AsOf           time.Time        `json:"as_of"`
	Timezone       string           `json:"timezone,omitempty"`
}
