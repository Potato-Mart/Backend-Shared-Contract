package inventory

import (
	"time"
)

// ProductStockSummary is a derived, revisioned availability projection. It is
// operational data and is not part of immutable transaction product records.
type ProductStockSummary struct {
	SKUCode        string                              `json:"sku_code"`
	AllDepots      ProductStockQuantitySnapshot        `json:"all_depots"`
	Depots         []DepotProductStockSnapshot         `json:"depots,omitempty"`
	Locations      []LocationProductStockSnapshot      `json:"locations,omitempty"`
	PackageOptions []PackageOptionProductStockSnapshot `json:"package_options,omitempty"`
	Channels       []ChannelProductStockSnapshot       `json:"channels,omitempty"`
	Revision       int64                               `json:"revision"`
	Timezone       string                              `json:"timezone"`
	IsOutOfStock   bool                                `json:"is_out_of_stock"`
	AsOf           time.Time                           `json:"as_of"`
}
