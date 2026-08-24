package membership_enums

// MembershipTierMetric controls the value that qualifies a membership tier.
type MembershipTierMetric string

const (
	MembershipTierMetricAnnualSpend   MembershipTierMetric = "annual_spend"
	MembershipTierMetricLifetimeSpend MembershipTierMetric = "lifetime_spend"
	MembershipTierMetricManual        MembershipTierMetric = "manual"
)

func (m MembershipTierMetric) IsValid() bool {
	switch m {
	case MembershipTierMetricAnnualSpend, MembershipTierMetricLifetimeSpend, MembershipTierMetricManual:
		return true
	}
	return false
}
func (m MembershipTierMetric) String() string { return string(m) }
