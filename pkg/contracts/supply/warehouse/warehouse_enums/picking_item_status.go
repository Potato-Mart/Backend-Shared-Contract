package warehouse_enums

type PickingItemStatus string

const (
	PickingItemStatusPending  PickingItemStatus = "pending"
	PickingItemStatusPartial  PickingItemStatus = "partial"
	PickingItemStatusComplete PickingItemStatus = "complete"
	PickingItemStatusSkipped  PickingItemStatus = "skipped"
)

// IsValid reports whether s is a known InboundReceiptStatus.
func (s PickingItemStatus) IsValid() bool {
	switch s {
	case PickingItemStatusPending, PickingItemStatusPartial,
		PickingItemStatusComplete, PickingItemStatusSkipped:
		return true
	}
	return false
}

func (s PickingItemStatus) String() string { return string(s) }
