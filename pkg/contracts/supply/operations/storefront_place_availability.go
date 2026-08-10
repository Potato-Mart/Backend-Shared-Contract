package operations

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/supply/product/product_enums"
)

// StorefrontPlaceAvailability is the customer-safe stock availability of one
// product at one publicly visible depot, carrying real-world place identity
// (depot name, region, administrative area, locality) for storefront display.
// It is a derived projection: it never exposes locations, lots, reservations,
// staging, or quality-hold quantities.
type StorefrontPlaceAvailability struct {
	DepotCode              string                             `json:"depot_code"`
	DepotName              string                             `json:"depot_name"`
	RegionCode             string                             `json:"region_code,omitempty"`
	RegionName             string                             `json:"region_name,omitempty"`
	CountryCode            geography.CountryCode              `json:"country_code,omitempty"`
	AdministrativeAreaCode geography.SubdivisionCode          `json:"administrative_area_code,omitempty"`
	Locality               string                             `json:"locality,omitempty"`
	StockState             product_enums.StorefrontStockState `json:"stock_state"`
	AvailableBaseUnits     int64                              `json:"available_base_units"`
	AsOf                   time.Time                          `json:"as_of"`
}
