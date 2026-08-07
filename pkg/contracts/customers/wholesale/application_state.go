package wholesale

// WholesaleApplicationState is the storefront-facing lifecycle for a trade
// account application. It condenses the organisation, access, and customer
// records into the state a buyer needs to understand.
type WholesaleApplicationState string

const (
	WholesaleApplicationStateMissing   WholesaleApplicationState = "missing"
	WholesaleApplicationStatePending   WholesaleApplicationState = "pending"
	WholesaleApplicationStateApproved  WholesaleApplicationState = "approved"
	WholesaleApplicationStateRejected  WholesaleApplicationState = "rejected"
	WholesaleApplicationStateSuspended WholesaleApplicationState = "suspended"
)

func (s WholesaleApplicationState) IsValid() bool {
	switch s {
	case WholesaleApplicationStateMissing, WholesaleApplicationStatePending,
		WholesaleApplicationStateApproved, WholesaleApplicationStateRejected,
		WholesaleApplicationStateSuspended:
		return true
	}
	return false
}

func (s WholesaleApplicationState) String() string { return string(s) }
