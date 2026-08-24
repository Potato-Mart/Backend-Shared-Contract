package retail_enums

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
