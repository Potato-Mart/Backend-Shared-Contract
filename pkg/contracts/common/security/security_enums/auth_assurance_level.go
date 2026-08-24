package security_enums

// AuthAssuranceLevel follows the common AAL1/AAL2/AAL3 vocabulary for
// how strongly a user was authenticated.
type AuthAssuranceLevel string

const (
	AuthAssuranceLevel1 AuthAssuranceLevel = "aal1"
	AuthAssuranceLevel2 AuthAssuranceLevel = "aal2"
	AuthAssuranceLevel3 AuthAssuranceLevel = "aal3"
)

func (a AuthAssuranceLevel) IsValid() bool {
	switch a {
	case AuthAssuranceLevel1, AuthAssuranceLevel2, AuthAssuranceLevel3:
		return true
	}
	return false
}

func (a AuthAssuranceLevel) String() string { return string(a) }
