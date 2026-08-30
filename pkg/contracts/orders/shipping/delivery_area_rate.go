package shipping

import (
	geography "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/money"
)

// DeliveryAreaRate is the customer-safe geographic rate and depot projection.
type DeliveryAreaRate struct {
	CountryCode            geography.CountryCode     `json:"country_code"`
	AdministrativeAreaCode geography.SubdivisionCode `json:"administrative_area_code,omitempty"`
	PostalCode             string                    `json:"postal_code"`
	Locality               string                    `json:"locality,omitempty"`
	ZoneID                 string                    `json:"zone_id"`
	DepotRegionCode        string                    `json:"depot_region_code,omitempty"`
	DepotCode              string                    `json:"depot_code"`
	DepotName              string                    `json:"depot_name"`
	ShippingFee            money.Money               `json:"shipping_fee"`
	FreeShippingThreshold  money.Money               `json:"free_shipping_threshold"`
}
