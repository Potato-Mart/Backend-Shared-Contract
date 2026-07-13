package identityenum

// UserRole is a persisted workforce RBAC role key. Role-to-permission policy
// is owned by Management and enforced independently by each backend.
type UserRole string

const (
	UserRoleSuperAdmin        UserRole = "superAdmin"
	UserRoleAdmin             UserRole = "admin"
	UserRoleSales             UserRole = "sales"
	UserRoleWarehouse         UserRole = "warehouse"
	UserRoleWarehouseOperator UserRole = "warehouseOperator"
	UserRoleMarketing         UserRole = "marketing"
)

func (r UserRole) IsValid() bool {
	switch r {
	case UserRoleSuperAdmin, UserRoleAdmin, UserRoleSales,
		UserRoleWarehouse, UserRoleWarehouseOperator, UserRoleMarketing:
		return true
	default:
		return false
	}
}

func (r UserRole) String() string { return string(r) }
