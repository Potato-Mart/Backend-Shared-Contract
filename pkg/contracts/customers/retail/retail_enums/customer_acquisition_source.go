package retail_enums

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
