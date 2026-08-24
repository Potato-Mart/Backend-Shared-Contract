package account_enums

// AuthIdentityStatus describes whether a login identity can be used.
type AuthIdentityStatus string

const (
	AuthIdentityStatusActive   AuthIdentityStatus = "active"
	AuthIdentityStatusDisabled AuthIdentityStatus = "disabled"
	AuthIdentityStatusRevoked  AuthIdentityStatus = "revoked"
)

// IsValid reports whether s is a known AuthIdentityStatus value.
func (s AuthIdentityStatus) IsValid() bool {
	switch s {
	case AuthIdentityStatusActive, AuthIdentityStatusDisabled, AuthIdentityStatusRevoked:
		return true
	}
	return false
}

// String returns the wire value for s.
func (s AuthIdentityStatus) String() string { return string(s) }
