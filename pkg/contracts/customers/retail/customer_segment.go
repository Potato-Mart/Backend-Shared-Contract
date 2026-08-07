package retail

// CustomerStatus is the lifecycle status of a retail or wholesale customer
// business profile. It is not used to decide portal admission.
type CustomerStatus string

const (
	CustomerStatusActive   CustomerStatus = "ACTIVE"
	CustomerStatusInactive CustomerStatus = "INACTIVE"
	CustomerStatusBlocked  CustomerStatus = "BLOCKED"
	CustomerStatusClosed   CustomerStatus = "CLOSED"
)

// IsValid reports whether c is a known CustomerStatus value.
func (c CustomerStatus) IsValid() bool {
	switch c {
	case CustomerStatusActive, CustomerStatusInactive, CustomerStatusBlocked,
		CustomerStatusClosed:
		return true
	}
	return false
}

// String returns the wire value for c.
func (c CustomerStatus) String() string { return string(c) }

// CustomerAcquisitionSource identifies how a customer profile was acquired.
type CustomerAcquisitionSource string

const (
	CustomerAcquisitionSourceOnline CustomerAcquisitionSource = "online"
	CustomerAcquisitionSourcePOS    CustomerAcquisitionSource = "pos"
	CustomerAcquisitionSourceImport CustomerAcquisitionSource = "import"
	CustomerAcquisitionSourceManual CustomerAcquisitionSource = "manual"
	CustomerAcquisitionSourcePhone  CustomerAcquisitionSource = "phone"
)

// IsValid reports whether s is a known CustomerAcquisitionSource value.
func (s CustomerAcquisitionSource) IsValid() bool {
	switch s {
	case CustomerAcquisitionSourceOnline, CustomerAcquisitionSourcePOS,
		CustomerAcquisitionSourceImport, CustomerAcquisitionSourceManual,
		CustomerAcquisitionSourcePhone:
		return true
	}
	return false
}

// String returns the wire value for s.
func (s CustomerAcquisitionSource) String() string { return string(s) }

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
