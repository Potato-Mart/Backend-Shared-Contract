package accountenum

// Portal identifies which front-door application a user, session, or
// device interaction belongs to.
type Portal string

const (
	PortalControl Portal = "control"
	PortalStore   Portal = "store"
	PortalPartner Portal = "partner"
)

// IsValid reports whether p is a known Portal.
func (p Portal) IsValid() bool {
	switch p {
	case PortalControl, PortalStore, PortalPartner:
		return true
	}
	return false
}

// AllowedAccountTypes returns the account types accepted by this portal.
func (p Portal) AllowedAccountTypes() []AccountType {
	accountType, ok := p.RequiredAccountType()
	if !ok {
		return nil
	}
	return []AccountType{accountType}
}

// RequiresAccountType reports whether accountType is the account type required
// by this portal.
func (p Portal) RequiresAccountType(accountType AccountType) bool {
	required, ok := p.RequiredAccountType()
	return ok && required == accountType
}

// RequiredAccountType returns the single account type required by this portal.
func (p Portal) RequiredAccountType() (AccountType, bool) {
	switch p {
	case PortalControl:
		return AccountTypeAdminUser, true
	case PortalStore:
		return AccountTypeGeneralCustomer, true
	case PortalPartner:
		return AccountTypeWholesaleCustomer, true
	}
	return "", false
}

func (p Portal) String() string { return string(p) }
