package enums

type PickingListStatus string

const (
	PickingListStatusPending    PickingListStatus = "PENDING"
	PickingListStatusInProgress PickingListStatus = "IN_PROGRESS"
	PickingListStatusComplete   PickingListStatus = "COMPLETE"
	PickingListStatusCancelled  PickingListStatus = "CANCELLED"
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
	PickingItemStatusPending  PickingItemStatus = "PENDING"
	PickingItemStatusPartial  PickingItemStatus = "PARTIAL"
	PickingItemStatusComplete PickingItemStatus = "COMPLETE"
	PickingItemStatusSkipped  PickingItemStatus = "SKIPPED"
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
