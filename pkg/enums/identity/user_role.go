package identityenum

// UserRole identifies an RBAC role key for permissions inside a portal,
// account, or organisation context. It is not a portal admission or
// account-type discriminator.
//
// Wire values are the camelCase strings used by the frontend role picker
// (e.g. "superAdmin", "warehouseOperator") so a JWT payload's role claim
// can be compared directly against the frontend constant without
// translation.
//
// New code must use AccountType plus PortalAccess to decide whether an
// account/persona may enter a portal. Use UserRole, Role, Permission, and
// RoleAssignment only after portal admission has selected the account context.
type UserRole string

const (
	UserRoleSuperAdmin        UserRole = "superAdmin"
	UserRoleAdmin             UserRole = "admin"
	UserRoleSales             UserRole = "sales"
	UserRoleWarehouse         UserRole = "warehouse"
	UserRoleWarehouseOperator UserRole = "warehouseOperator"
	UserRoleMarketing         UserRole = "marketing"
	UserRoleCustomer          UserRole = "customer"
)

// IsValid reports whether r is a known UserRole value.
func (r UserRole) IsValid() bool {
	switch r {
	case UserRoleSuperAdmin, UserRoleAdmin, UserRoleSales,
		UserRoleWarehouse, UserRoleWarehouseOperator, UserRoleMarketing,
		UserRoleCustomer:
		return true
	}
	return false
}

func (r UserRole) String() string { return string(r) }

// AllStaffRoles returns every non-customer role. Useful for RBAC defaults
// that should grant access to "any staff member".
func AllStaffRoles() []UserRole {
	return []UserRole{
		UserRoleSuperAdmin,
		UserRoleAdmin,
		UserRoleSales,
		UserRoleWarehouse,
		UserRoleWarehouseOperator,
		UserRoleMarketing,
	}
}
