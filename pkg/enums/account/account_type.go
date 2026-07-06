package accountenum

// AccountType identifies the persona/access-domain record attached to a
// canonical user. It is used for portal admission and is not an RBAC role.
type AccountType string

const (
	// AccountTypeAdminUser is the internal/admin-platform account persona.
	AccountTypeAdminUser AccountType = "adminUser"
	// AccountTypeGeneralCustomer is the storefront customer account persona.
	AccountTypeGeneralCustomer AccountType = "generalCustomer"
	// AccountTypeWholesaleCustomer is the wholesale/partner account persona.
	AccountTypeWholesaleCustomer AccountType = "wholesaleCustomer"
)

// IsValid reports whether t is a known AccountType value.
func (t AccountType) IsValid() bool {
	switch t {
	case AccountTypeAdminUser, AccountTypeGeneralCustomer, AccountTypeWholesaleCustomer:
		return true
	}
	return false
}

// String returns the wire value for t.
func (t AccountType) String() string { return string(t) }

// IsAllowedInPortal reports whether t is accepted by the given front-door
// portal.
func (t AccountType) IsAllowedInPortal(portal Portal) bool {
	return portal.RequiresAccountType(t)
}

// AllAccountTypes returns every account type known to this contract.
func AllAccountTypes() []AccountType {
	return []AccountType{
		AccountTypeAdminUser,
		AccountTypeGeneralCustomer,
		AccountTypeWholesaleCustomer,
	}
}

// AccountTypesForPortal returns the account types accepted by portal.
func AccountTypesForPortal(portal Portal) []AccountType {
	return portal.AllowedAccountTypes()
}
