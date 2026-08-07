package role

// UserRole is a persisted workforce RBAC role key. Role-to-permission policy
// is owned by Identity and enforced independently by each backend.
type UserRole string

const (
	UserRoleSuperAdmin        UserRole = "superAdmin"
	UserRoleAdmin             UserRole = "admin"
	UserRoleSales             UserRole = "sales"
	UserRoleWarehouse         UserRole = "warehouse"
	UserRoleWarehouseOperator UserRole = "warehouseOperator"
	UserRoleMarketing         UserRole = "marketing"
	UserRoleCashier           UserRole = "cashier"
)

func (r UserRole) IsValid() bool {
	switch r {
	case UserRoleSuperAdmin, UserRoleAdmin, UserRoleSales,
		UserRoleWarehouse, UserRoleWarehouseOperator, UserRoleMarketing,
		UserRoleCashier:
		return true
	default:
		return false
	}
}

func (r UserRole) String() string { return string(r) }
