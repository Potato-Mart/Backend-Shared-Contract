package warehouse_enums

// WMSDraftStatus is the lifecycle of a WMS batch draft entry.
type WMSDraftStatus string

const (
	WMSDraftStatusDraft     WMSDraftStatus = "DRAFT"
	WMSDraftStatusSubmitted WMSDraftStatus = "SUBMITTED"
	WMSDraftStatusCancelled WMSDraftStatus = "CANCELLED"
)

func (w WMSDraftStatus) IsValid() bool {
	switch w {
	case WMSDraftStatusDraft, WMSDraftStatusSubmitted, WMSDraftStatusCancelled:
		return true
	}
	return false
}

func (w WMSDraftStatus) String() string { return string(w) }
