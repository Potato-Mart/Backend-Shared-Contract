package analytics_enums

// ChurnRisk is the computed churn risk level derived from RFM recency analysis.
type ChurnRisk string

const (
	ChurnRiskLow    ChurnRisk = "LOW"
	ChurnRiskMedium ChurnRisk = "MEDIUM"
	ChurnRiskHigh   ChurnRisk = "HIGH"
)

// IsValid reports whether c is a known ChurnRisk value.
func (c ChurnRisk) IsValid() bool {
	switch c {
	case ChurnRiskLow, ChurnRiskMedium, ChurnRiskHigh:
		return true
	}
	return false
}

// String returns the wire value for c.
func (c ChurnRisk) String() string { return string(c) }
