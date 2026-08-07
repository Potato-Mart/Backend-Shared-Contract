package wholesale_enums

// WholesaleApplicationStatus is the persisted review decision for a wholesale
// application.
type WholesaleApplicationStatus string

const (
	WholesaleApplicationStatusPending  WholesaleApplicationStatus = "pending"
	WholesaleApplicationStatusApproved WholesaleApplicationStatus = "approved"
	WholesaleApplicationStatusRejected WholesaleApplicationStatus = "rejected"
)

func (s WholesaleApplicationStatus) IsValid() bool {
	switch s {
	case WholesaleApplicationStatusPending, WholesaleApplicationStatusApproved,
		WholesaleApplicationStatusRejected:
		return true
	default:
		return false
	}
}

func (s WholesaleApplicationStatus) String() string { return string(s) }
