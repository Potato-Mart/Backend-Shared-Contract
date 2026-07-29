package shipping

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v21/pkg/common"
)

// DeliverySlot is one customer-selectable delivery window. ID is opaque and
// stable for the area, date, window, and schedule revision that produced it.
// Availability is a service-owned wire value such as available, limited, or
// full. Checkout remains authoritative for both capacity and final fees.
type DeliverySlot struct {
	ID           string        `json:"id"`
	StartAt      time.Time     `json:"start_at"`
	EndAt        time.Time     `json:"end_at"`
	Label        string        `json:"label,omitempty"`
	Availability string        `json:"availability"`
	Fee          *common.Money `json:"fee,omitempty"`
}

// DeliveryDateGroup groups the selectable windows for one store-local date.
// Date uses the YYYY-MM-DD calendar representation in the schedule timezone.
type DeliveryDateGroup struct {
	Date  string         `json:"date"`
	Label string         `json:"label,omitempty"`
	Slots []DeliverySlot `json:"slots"`
}

// DeliveryAreaRate is the customer-safe rate and depot projection selected by
// postcode and, where supplied, suburb. Monetary values use AUD minor units
// through common.Money.
type DeliveryAreaRate struct {
	Postcode              string       `json:"postcode"`
	Suburb                string       `json:"suburb,omitempty"`
	DeliveryRegion        string       `json:"delivery_region"`
	DepotCode             string       `json:"depot_code"`
	DepotName             string       `json:"depot_name"`
	ShippingFee           common.Money `json:"shipping_fee"`
	FreeShippingThreshold common.Money `json:"free_shipping_threshold"`
}

// DeliverySchedule is a cart-free, revisioned view of delivery windows for an
// area. Checkout revalidates every selected slot and calculates the final fee.
type DeliverySchedule struct {
	Availability      string              `json:"availability"`
	UnavailableReason string              `json:"unavailable_reason,omitempty"`
	Revision          int64               `json:"revision"`
	Timezone          string              `json:"timezone"`
	Carrier           string              `json:"carrier,omitempty"`
	AreaRate          *DeliveryAreaRate   `json:"area_rate,omitempty"`
	DateGroups        []DeliveryDateGroup `json:"date_groups"`
}

// PreferredDeliverySlot is the customer's cart-independent preference. It is
// display metadata only: checkout must revalidate the opaque slot ID and time
// window against the current cart-bound schedule.
type PreferredDeliverySlot struct {
	Date             string `json:"date"`
	SlotID           string `json:"slot_id,omitempty"`
	StartAt          string `json:"start_at,omitempty"`
	EndAt            string `json:"end_at,omitempty"`
	ScheduleRevision int64  `json:"schedule_revision,omitempty"`
}
