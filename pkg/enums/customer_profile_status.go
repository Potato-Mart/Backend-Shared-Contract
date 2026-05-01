package enums

// CustomerSegment distinguishes wholesale B2B clients from retail individuals
// in the CRM customer_profiles table.
type CustomerSegment string

const (
	CustomerSegmentWholesale CustomerSegment = "WHOLESALE"
	CustomerSegmentRetail    CustomerSegment = "RETAIL"
)

func (c CustomerSegment) IsValid() bool {
	switch c {
	case CustomerSegmentWholesale, CustomerSegmentRetail:
		return true
	}
	return false
}

func (c CustomerSegment) String() string { return string(c) }

// CustomerProfileStatus is the lifecycle status of a CRM customer profile.
type CustomerProfileStatus string

const (
	CustomerProfileStatusActive   CustomerProfileStatus = "ACTIVE"
	CustomerProfileStatusInactive CustomerProfileStatus = "INACTIVE"
	CustomerProfileStatusBlocked  CustomerProfileStatus = "BLOCKED"
)

func (c CustomerProfileStatus) IsValid() bool {
	switch c {
	case CustomerProfileStatusActive, CustomerProfileStatusInactive, CustomerProfileStatusBlocked:
		return true
	}
	return false
}

func (c CustomerProfileStatus) String() string { return string(c) }

// ChurnRisk is the computed churn risk level derived from RFM recency analysis.
type ChurnRisk string

const (
	ChurnRiskLow    ChurnRisk = "LOW"
	ChurnRiskMedium ChurnRisk = "MEDIUM"
	ChurnRiskHigh   ChurnRisk = "HIGH"
)

func (c ChurnRisk) IsValid() bool {
	switch c {
	case ChurnRiskLow, ChurnRiskMedium, ChurnRiskHigh:
		return true
	}
	return false
}

func (c ChurnRisk) String() string { return string(c) }
