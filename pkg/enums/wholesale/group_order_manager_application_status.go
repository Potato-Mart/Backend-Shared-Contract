package wholesaleenum

type GroupOrderManagerApplicationStatus string

const (
	GroupOrderManagerApplicationStatusPending   GroupOrderManagerApplicationStatus = "pending"
	GroupOrderManagerApplicationStatusApproved  GroupOrderManagerApplicationStatus = "approved"
	GroupOrderManagerApplicationStatusRejected  GroupOrderManagerApplicationStatus = "rejected"
	GroupOrderManagerApplicationStatusCancelled GroupOrderManagerApplicationStatus = "cancelled"
)

func (s GroupOrderManagerApplicationStatus) IsValid() bool {
	switch s {
	case GroupOrderManagerApplicationStatusPending,
		GroupOrderManagerApplicationStatusApproved,
		GroupOrderManagerApplicationStatusRejected,
		GroupOrderManagerApplicationStatusCancelled:
		return true
	default:
		return false
	}
}

func (s GroupOrderManagerApplicationStatus) String() string { return string(s) }
