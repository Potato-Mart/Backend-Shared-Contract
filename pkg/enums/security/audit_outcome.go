package securityenum

// AuditOutcome is a coarse success/failure flag for audit reporting.
type AuditOutcome string

const (
	AuditOutcomeSuccess AuditOutcome = "success"
	AuditOutcomeFailure AuditOutcome = "failure"
	AuditOutcomeDenied  AuditOutcome = "denied"
)

// IsValid reports whether o is a known AuditOutcome.
func (o AuditOutcome) IsValid() bool {
	switch o {
	case AuditOutcomeSuccess, AuditOutcomeFailure, AuditOutcomeDenied:
		return true
	}
	return false
}

func (o AuditOutcome) String() string { return string(o) }
