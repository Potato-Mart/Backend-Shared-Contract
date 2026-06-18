package enums

// AuthIdentityProvider identifies a non-secret login provider attached to a
// canonical user.
type AuthIdentityProvider string

const (
	AuthIdentityProviderPassword     AuthIdentityProvider = "password"
	AuthIdentityProviderGoogle       AuthIdentityProvider = "google"
	AuthIdentityProviderApple        AuthIdentityProvider = "apple"
	AuthIdentityProviderAzureAD      AuthIdentityProvider = "azureAD"
	AuthIdentityProviderOkta         AuthIdentityProvider = "okta"
	AuthIdentityProviderPasskey      AuthIdentityProvider = "passkey"
	AuthIdentityProviderServiceToken AuthIdentityProvider = "serviceToken"
)

// IsValid reports whether p is a known AuthIdentityProvider value.
func (p AuthIdentityProvider) IsValid() bool {
	switch p {
	case AuthIdentityProviderPassword, AuthIdentityProviderGoogle,
		AuthIdentityProviderApple, AuthIdentityProviderAzureAD,
		AuthIdentityProviderOkta, AuthIdentityProviderPasskey,
		AuthIdentityProviderServiceToken:
		return true
	}
	return false
}

// String returns the wire value for p.
func (p AuthIdentityProvider) String() string { return string(p) }

// IdentityDomain identifies the broad trust domain for an auth identity.
type IdentityDomain string

const (
	IdentityDomainCustomer  IdentityDomain = "customer"
	IdentityDomainWorkforce IdentityDomain = "workforce"
	IdentityDomainPartner   IdentityDomain = "partner"
	IdentityDomainService   IdentityDomain = "service"
)

// IsValid reports whether d is a known IdentityDomain value.
func (d IdentityDomain) IsValid() bool {
	switch d {
	case IdentityDomainCustomer, IdentityDomainWorkforce, IdentityDomainPartner,
		IdentityDomainService:
		return true
	}
	return false
}

// String returns the wire value for d.
func (d IdentityDomain) String() string { return string(d) }

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
