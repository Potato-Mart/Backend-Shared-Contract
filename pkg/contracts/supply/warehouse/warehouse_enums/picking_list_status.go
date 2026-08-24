package warehouse_enums

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
