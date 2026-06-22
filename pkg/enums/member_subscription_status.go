package enums

// MemberSubscriptionStatus is the lifecycle of a recurring purchase attached
// to a membership account.
type MemberSubscriptionStatus string

const (
	MemberSubscriptionStatusActive    MemberSubscriptionStatus = "ACTIVE"
	MemberSubscriptionStatusPaused    MemberSubscriptionStatus = "PAUSED"
	MemberSubscriptionStatusCancelled MemberSubscriptionStatus = "CANCELLED"
)

func (s MemberSubscriptionStatus) IsValid() bool {
	switch s {
	case MemberSubscriptionStatusActive, MemberSubscriptionStatusPaused, MemberSubscriptionStatusCancelled:
		return true
	}
	return false
}

func (s MemberSubscriptionStatus) String() string { return string(s) }
