package security_enums

// SecurityRiskLevel is a coarse risk label used by sessions, devices,
// requests, and security events.
type SecurityRiskLevel string

const (
	SecurityRiskLevelLow      SecurityRiskLevel = "low"
	SecurityRiskLevelMedium   SecurityRiskLevel = "medium"
	SecurityRiskLevelHigh     SecurityRiskLevel = "high"
	SecurityRiskLevelCritical SecurityRiskLevel = "critical"
)

func (r SecurityRiskLevel) IsValid() bool {
	switch r {
	case SecurityRiskLevelLow, SecurityRiskLevelMedium,
		SecurityRiskLevelHigh, SecurityRiskLevelCritical:
		return true
	}
	return false
}

func (r SecurityRiskLevel) String() string { return string(r) }
