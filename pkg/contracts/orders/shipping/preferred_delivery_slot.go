package shipping

// PreferredDeliverySlot is the customer's cart-independent display snapshot.
type PreferredDeliverySlot struct {
	Date             string `json:"date"`
	SlotID           string `json:"slot_id,omitempty"`
	StartAt          string `json:"start_at,omitempty"`
	EndAt            string `json:"end_at,omitempty"`
	ScheduleRevision int64  `json:"schedule_revision,omitempty"`
}
