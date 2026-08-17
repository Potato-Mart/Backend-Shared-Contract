package shipping

import (
	geography "github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/geography"

	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/money"
)

// DeliverySlot is one customer-selectable delivery window. ID is opaque and
// stable for the area, date, window, and schedule revision that produced it.
// Availability is a service-owned wire value such as available, limited, or
// full.
type DeliverySlot struct {
	ID           string       `json:"id"`
	StartAt      time.Time    `json:"start_at"`
	EndAt        time.Time    `json:"end_at"`
	Label        string       `json:"label,omitempty"`
	Availability string       `json:"availability"`
	Fee          *money.Money `json:"fee,omitempty"`
}

// DeliveryDateGroup groups the selectable windows for one store-local date.
// Date uses the YYYY-MM-DD calendar representation in the schedule timezone.
type DeliveryDateGroup struct {
	Date  string         `json:"date"`
	Label string         `json:"label,omitempty"`
	Slots []DeliverySlot `json:"slots"`
}

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

// DeliverySchedule is a cart-free, revisioned view of delivery windows for an
// area.
type DeliverySchedule struct {
	Availability      string              `json:"availability"`
	UnavailableReason string              `json:"unavailable_reason,omitempty"`
	Revision          int64               `json:"revision"`
	Timezone          string              `json:"timezone"`
	Carrier           string              `json:"carrier,omitempty"`
	AreaRate          *DeliveryAreaRate   `json:"area_rate,omitempty"`
	DateGroups        []DeliveryDateGroup `json:"date_groups"`
}

// PreferredDeliverySlot is the customer's cart-independent display snapshot.
type PreferredDeliverySlot struct {
	Date             string `json:"date"`
	SlotID           string `json:"slot_id,omitempty"`
	StartAt          string `json:"start_at,omitempty"`
	EndAt            string `json:"end_at,omitempty"`
	ScheduleRevision int64  `json:"schedule_revision,omitempty"`
}
