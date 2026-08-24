package security_enums

type SecurityEventSeverity string

const (
	SecurityEventSeverityInfo     SecurityEventSeverity = "info"
	SecurityEventSeverityLow      SecurityEventSeverity = "low"
	SecurityEventSeverityMedium   SecurityEventSeverity = "medium"
	SecurityEventSeverityHigh     SecurityEventSeverity = "high"
	SecurityEventSeverityCritical SecurityEventSeverity = "critical"
)

func (s SecurityEventSeverity) IsValid() bool {
	switch s {
	case SecurityEventSeverityInfo, SecurityEventSeverityLow,
		SecurityEventSeverityMedium, SecurityEventSeverityHigh,
		SecurityEventSeverityCritical:
		return true
	}
	return false
}

func (s SecurityEventSeverity) String() string { return string(s) }
