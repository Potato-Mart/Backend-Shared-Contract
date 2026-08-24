package membership_enums

// MembershipAccountStatus describes the lifecycle of a programme account.
type MembershipAccountStatus string

const (
	MembershipAccountStatusActive    MembershipAccountStatus = "active"
	MembershipAccountStatusSuspended MembershipAccountStatus = "suspended"
	MembershipAccountStatusClosed    MembershipAccountStatus = "closed"
)

func (s MembershipAccountStatus) IsValid() bool {
	return s == MembershipAccountStatusActive || s == MembershipAccountStatusSuspended || s == MembershipAccountStatusClosed
}
func (s MembershipAccountStatus) String() string { return string(s) }
