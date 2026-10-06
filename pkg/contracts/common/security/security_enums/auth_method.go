package security_enums

// AuthMethod records the primary authentication mechanism used for a login
// or privileged action.
type AuthMethod string

const (
	AuthMethodPassword     AuthMethod = "password"
	AuthMethodMFA          AuthMethod = "mfa"
	AuthMethodPasskey      AuthMethod = "passkey"
	AuthMethodSSO          AuthMethod = "sso"
	AuthMethodRefreshToken AuthMethod = "refresh_token"
	AuthMethodAPIKey       AuthMethod = "api_key"
	// AuthMethodPIN records reusable staff PIN authentication through
	// Identity's authorized POS flow; it does not imply MFA or higher assurance.
	AuthMethodPIN AuthMethod = "pin"
)

func (m AuthMethod) IsValid() bool {
	switch m {
	case AuthMethodPassword, AuthMethodMFA, AuthMethodPasskey,
		AuthMethodSSO, AuthMethodRefreshToken, AuthMethodAPIKey, AuthMethodPIN:
		return true
	}
	return false
}

func (m AuthMethod) String() string { return string(m) }
