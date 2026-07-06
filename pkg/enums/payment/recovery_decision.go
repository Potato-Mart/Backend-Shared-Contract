package paymentenum

// RecoveryDecision captures the merchant's manual answer to the
// override dialog ("Was the transaction successful?") shown when a
// terminal transaction times out or returns an Unknown outcome. It is
// the source of truth for closing the sale in those cases.
type RecoveryDecision string

const (
	RecoveryDecisionPending  RecoveryDecision = "pending"
	RecoveryDecisionApproved RecoveryDecision = "approved"
	RecoveryDecisionDeclined RecoveryDecision = "declined"
)

// IsValid reports whether d is a known RecoveryDecision.
func (d RecoveryDecision) IsValid() bool {
	switch d {
	case RecoveryDecisionPending, RecoveryDecisionApproved, RecoveryDecisionDeclined:
		return true
	}
	return false
}

func (d RecoveryDecision) String() string { return string(d) }
