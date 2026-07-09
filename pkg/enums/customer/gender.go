package customerenum

// CustomerGender captures the supported self-declared retail profile genders.
type CustomerGender string

const (
	CustomerGenderFemale    CustomerGender = "female"
	CustomerGenderMale      CustomerGender = "male"
	CustomerGenderNonBinary CustomerGender = "non_binary"
)

// IsValid reports whether g is a known CustomerGender value.
func (g CustomerGender) IsValid() bool {
	switch g {
	case CustomerGenderFemale, CustomerGenderMale, CustomerGenderNonBinary:
		return true
	}
	return false
}

func (g CustomerGender) String() string { return string(g) }
