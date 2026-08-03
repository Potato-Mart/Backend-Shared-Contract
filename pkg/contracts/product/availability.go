package product

import (
	"time"

	salesenum "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/enums/sales"
)

// ProductStockQuantitySnapshot is a base-unit stock projection for one
// product and one optional stock dimension.
type ProductStockQuantitySnapshot struct {
	OnHandBaseUnits      int64 `json:"on_hand_base_units"`
	AvailableBaseUnits   int64 `json:"available_base_units"`
	SellableBaseUnits    int64 `json:"sellable_base_units"`
	ReducedBaseUnits     int64 `json:"reduced_base_units"`
	ReservedBaseUnits    int64 `json:"reserved_base_units"`
	StagedBaseUnits      int64 `json:"staged_base_units"`
	QualityHoldBaseUnits int64 `json:"quality_hold_base_units"`
}

// DepotProductStockSnapshot qualifies product stock by depot.
type DepotProductStockSnapshot struct {
	DepotCode  string                       `json:"depot_code"`
	Quantities ProductStockQuantitySnapshot `json:"quantities"`
}

// LocationProductStockSnapshot qualifies product stock by depot and location.
type LocationProductStockSnapshot struct {
	DepotCode    string                       `json:"depot_code"`
	LocationCode string                       `json:"location_code"`
	Quantities   ProductStockQuantitySnapshot `json:"quantities"`
}

// PackageOptionProductStockSnapshot qualifies product stock by package option.
type PackageOptionProductStockSnapshot struct {
	PackageOptionID string                       `json:"package_option_id"`
	Quantities      ProductStockQuantitySnapshot `json:"quantities"`
}

// ChannelProductStockSnapshot qualifies product stock by sales channel.
type ChannelProductStockSnapshot struct {
	Channel    salesenum.OrderType          `json:"channel"`
	Quantities ProductStockQuantitySnapshot `json:"quantities"`
}

// ProductStockSummary is a derived, revisioned availability projection. It is
// not part of immutable transaction product snapshots.
type ProductStockSummary struct {
	ProductSKUCode string                              `json:"product_sku_code"`
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
