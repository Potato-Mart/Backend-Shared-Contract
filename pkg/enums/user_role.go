package enums

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
//
// UserRoleClient is retained for backward compatibility with shared
// contract v3.0.x – v3.2.x where the only two roles were "admin" and
// "user". New code should prefer UserRoleCustomer.
type UserRole string

const (
	UserRoleSuperAdmin        UserRole = "superAdmin"
	UserRoleAdmin             UserRole = "admin"
	UserRoleSales             UserRole = "sales"
	UserRoleWarehouse         UserRole = "warehouse"
	UserRoleWarehouseOperator UserRole = "warehouseOperator"
	UserRoleMarketing         UserRole = "marketing"
	UserRoleCustomer          UserRole = "customer"

	// Deprecated: use UserRoleCustomer. Retained for backward
	// compatibility with shared contract <= v3.2.x.
	UserRoleClient UserRole = "user"
)

// IsValid reports whether r is a known UserRole value.
func (r UserRole) IsValid() bool {
	switch r {
	case UserRoleSuperAdmin, UserRoleAdmin, UserRoleSales,
		UserRoleWarehouse, UserRoleWarehouseOperator, UserRoleMarketing,
		UserRoleCustomer, UserRoleClient:
		return true
	}
	return false
}

// IsStaff reports whether the legacy role represents an internal staff member.
// Deprecated for portal admission: retained for backward compatibility only.
// New code should use AccountType and PortalAccess as the platform gate.
func (r UserRole) IsStaff() bool {
	switch r {
	case UserRoleSuperAdmin, UserRoleAdmin, UserRoleSales,
		UserRoleWarehouse, UserRoleWarehouseOperator, UserRoleMarketing:
		return true
	}
	return false
}

// IsAdmin reports whether the legacy role has unrestricted administrative
// privileges (superAdmin or admin). Deprecated for portal admission: retained
// for backward compatibility only.
func (r UserRole) IsAdmin() bool {
	return r == UserRoleSuperAdmin || r == UserRoleAdmin
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
