package security_enums

// IdentityDomain identifies the broad trust domain for an auth identity.
type IdentityDomain string

const (
	IdentityDomainCustomer  IdentityDomain = "customer"
	IdentityDomainWorkforce IdentityDomain = "workforce"
	IdentityDomainService   IdentityDomain = "service"
)

// IsValid reports whether d is a known IdentityDomain value.
func (d IdentityDomain) IsValid() bool {
	switch d {
	case IdentityDomainCustomer, IdentityDomainWorkforce, IdentityDomainService:
		return true
	}
	return false
}

// String returns the wire value for d.
func (d IdentityDomain) String() string { return string(d) }
