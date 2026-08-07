package warehouse

// WMSDraftType classifies a WMS draft as an inbound or outbound operation.
type WMSDraftType string

const (
	WMSDraftTypeInbound  WMSDraftType = "INBOUND"
	WMSDraftTypeOutbound WMSDraftType = "OUTBOUND"
)

func (w WMSDraftType) IsValid() bool {
	switch w {
	case WMSDraftTypeInbound, WMSDraftTypeOutbound:
		return true
	}
	return false
}

func (w WMSDraftType) String() string { return string(w) }

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
