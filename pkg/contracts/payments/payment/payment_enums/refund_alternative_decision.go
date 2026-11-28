package payment_enums

// RefundAlternativeDecision is how an alternative offer on a pending refund
// request was answered. Staff record the answer after speaking to the
// customer; there is no customer-facing acceptance endpoint, so this is always
// a staff-recorded outcome.
type RefundAlternativeDecision string

const (
	// RefundAlternativeDecisionAccepted means the customer took the
	// alternative instead of a refund.
	RefundAlternativeDecisionAccepted RefundAlternativeDecision = "accepted"
	// RefundAlternativeDecisionRejected means the customer turned the
	// alternative down; the refund request stays open.
	RefundAlternativeDecisionRejected RefundAlternativeDecision = "rejected"
	// RefundAlternativeDecisionWithdrawn means staff withdrew the offer
	// before it was answered.
	RefundAlternativeDecisionWithdrawn RefundAlternativeDecision = "withdrawn"
)

// IsValid reports whether d is a known RefundAlternativeDecision.
func (d RefundAlternativeDecision) IsValid() bool {
	switch d {
	case RefundAlternativeDecisionAccepted, RefundAlternativeDecisionRejected, RefundAlternativeDecisionWithdrawn:
		return true
	}
	return false
}

func (d RefundAlternativeDecision) String() string { return string(d) }
