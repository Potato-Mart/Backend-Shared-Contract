package enums

// CustomerType distinguishes retail individuals from wholesale clients.
// Only users with role CLIENT carry a meaningful CustomerType; for
// admins the value is ignored by business logic.
type CustomerType string

const (
	CustomerTypeIndividual CustomerType = "INDIVIDUAL"
	CustomerTypeWholesaler CustomerType = "COMPANY"
)

// IsValid reports whether c is a known CustomerType.
func (c CustomerType) IsValid() bool {
	switch c {
	case CustomerTypeIndividual, CustomerTypeWholesaler:
		return true
	}
	return false
}

func (c CustomerType) String() string { return string(c) }
