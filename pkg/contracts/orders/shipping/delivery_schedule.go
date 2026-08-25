package shipping

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
