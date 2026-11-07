package membership_enums

// MembershipTierChangeReason explains why Pricing recomputed a customer's
// effective membership tier on membership.tier_changed.
type MembershipTierChangeReason string

const (
	// MembershipTierChangeReasonSpend is qualifying spend from a paid order.
	MembershipTierChangeReasonSpend MembershipTierChangeReason = "spend"
	// MembershipTierChangeReasonRefund is qualifying spend reversed by a refund.
	MembershipTierChangeReasonRefund MembershipTierChangeReason = "refund"
	// MembershipTierChangeReasonAssignment is a staff tier assignment or removal.
	MembershipTierChangeReasonAssignment MembershipTierChangeReason = "assignment"
	// MembershipTierChangeReasonEnrolment is the customer joining the programme.
	MembershipTierChangeReasonEnrolment MembershipTierChangeReason = "enrolment"
	// MembershipTierChangeReasonReconcile is a scheduled or operator
	// recomputation, such as a qualification window rolling over.
	MembershipTierChangeReasonReconcile MembershipTierChangeReason = "reconcile"
)

func (r MembershipTierChangeReason) IsValid() bool {
	switch r {
	case MembershipTierChangeReasonSpend, MembershipTierChangeReasonRefund,
		MembershipTierChangeReasonAssignment, MembershipTierChangeReasonEnrolment,
		MembershipTierChangeReasonReconcile:
		return true
	default:
		return false
	}
}

func (r MembershipTierChangeReason) String() string { return string(r) }
