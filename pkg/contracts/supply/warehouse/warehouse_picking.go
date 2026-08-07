package warehouse

type PickingListStatus string

const (
	PickingListStatusPending    PickingListStatus = "pending"
	PickingListStatusInProgress PickingListStatus = "in_progress"
	PickingListStatusComplete   PickingListStatus = "complete"
	PickingListStatusCancelled  PickingListStatus = "cancelled"
)

// IsValid reports whether s is a known PickingListStatus.
func (s PickingListStatus) IsValid() bool {
	switch s {
	case PickingListStatusPending, PickingListStatusInProgress,
		PickingListStatusComplete, PickingListStatusCancelled:
		return true
	}
	return false
}

func (s PickingListStatus) String() string { return string(s) }

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
