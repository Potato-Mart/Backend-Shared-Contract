package promotionenum

// GroupOrderDiscountState is the lifecycle of a per-group-order discount
// application: a wholesale group-order manager submits it and a staff approver
// either issues (approves) or rejects the benefit. Added v13.1.0.
type GroupOrderDiscountState string

const (
	GroupOrderDiscountStatePending  GroupOrderDiscountState = "pending"
	GroupOrderDiscountStateApproved GroupOrderDiscountState = "approved"
	GroupOrderDiscountStateRejected GroupOrderDiscountState = "rejected"
)

func (s GroupOrderDiscountState) IsValid() bool {
	switch s {
	case GroupOrderDiscountStatePending, GroupOrderDiscountStateApproved, GroupOrderDiscountStateRejected:
		return true
	}
	return false
}

func (s GroupOrderDiscountState) String() string { return string(s) }
