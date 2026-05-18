package enums

// UserRole identifies the access role of a backend user.
//
// Wire values are the camelCase strings used by the frontend role picker
// (e.g. "superAdmin", "warehouseOperator") so a JWT payload's role claim
// can be compared directly against the frontend constant without
// translation.
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

// IsStaff reports whether the role represents an internal staff member
// (anything except customer). Staff roles are allowed into potato-control;
// customers are restricted to potato-store.
func (r UserRole) IsStaff() bool {
	switch r {
	case UserRoleSuperAdmin, UserRoleAdmin, UserRoleSales,
		UserRoleWarehouse, UserRoleWarehouseOperator, UserRoleMarketing:
		return true
	}
	return false
}

// IsAdmin reports whether the role has unrestricted administrative
// privileges (superAdmin or admin).
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
