package account

// AccountType identifies the persona/access-domain record attached to a
// canonical user. It is used for portal admission and is not an RBAC role.
type AccountType string

const (
	// AccountTypeAdminUser is the internal/admin-platform account persona.
	AccountTypeAdminUser AccountType = "adminUser"
	// AccountTypeRetailCustomer is the retail storefront customer account persona.
	AccountTypeRetailCustomer AccountType = "retailCustomer"
	// AccountTypeWholesaleCustomer is the wholesale organisation-principal account persona.
	AccountTypeWholesaleCustomer AccountType = "wholesaleCustomer"
)

// IsValid reports whether t is a known AccountType value.
func (t AccountType) IsValid() bool {
	switch t {
	case AccountTypeAdminUser, AccountTypeRetailCustomer, AccountTypeWholesaleCustomer:
		return true
	}
	return false
}

// String returns the wire value for t.
func (t AccountType) String() string { return string(t) }
