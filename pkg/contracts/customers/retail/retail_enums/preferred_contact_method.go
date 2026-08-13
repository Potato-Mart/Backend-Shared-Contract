package retail_enums

// PreferredContactMethod captures which contact channel a retail customer
// prefers for person-to-person contact.
type PreferredContactMethod string

const (
	PreferredContactMethodEmail PreferredContactMethod = "email"
	PreferredContactMethodPhone PreferredContactMethod = "phone"
)

// IsValid reports whether m is a known PreferredContactMethod value.
func (m PreferredContactMethod) IsValid() bool {
	switch m {
	case PreferredContactMethodEmail, PreferredContactMethodPhone:
		return true
	}
	return false
}

func (m PreferredContactMethod) String() string { return string(m) }
