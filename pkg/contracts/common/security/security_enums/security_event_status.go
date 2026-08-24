package security_enums

type SecurityEventStatus string

const (
	SecurityEventStatusDetected      SecurityEventStatus = "detected"
	SecurityEventStatusTriaged       SecurityEventStatus = "triaged"
	SecurityEventStatusInvestigating SecurityEventStatus = "investigating"
	SecurityEventStatusContained     SecurityEventStatus = "contained"
	SecurityEventStatusResolved      SecurityEventStatus = "resolved"
	SecurityEventStatusFalsePositive SecurityEventStatus = "false_positive"
)

func (s SecurityEventStatus) IsValid() bool {
	switch s {
	case SecurityEventStatusDetected, SecurityEventStatusTriaged,
		SecurityEventStatusInvestigating, SecurityEventStatusContained,
		SecurityEventStatusResolved, SecurityEventStatusFalsePositive:
		return true
	}
	return false
}

func (s SecurityEventStatus) String() string { return string(s) }
