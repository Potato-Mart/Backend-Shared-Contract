package shipping

// DeliveryDateGroup groups the selectable windows for one store-local date.
// Date uses the YYYY-MM-DD calendar representation in the schedule timezone.
type DeliveryDateGroup struct {
	Date  string         `json:"date"`
	Label string         `json:"label,omitempty"`
	Slots []DeliverySlot `json:"slots"`
}
