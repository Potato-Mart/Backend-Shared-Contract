package role_enums

// UserRole is a persisted workforce RBAC role key. Role-to-permission policy
// is owned by Identity and enforced independently by each backend.
//
// The six built-in roles form a strict rank order from 1 (widest) to 6
// (narrowest); Role.Rank carries that number. Ranks 2 and below are
// geographically scoped: countryAdmin holds one country, marketing holds
// markets, and depotManager, warehouseManager, and warehouseOperator hold
// depots. Only superAdmin is global.
//
// There is no selling rank. Working a register is a permission every rank may
// hold, bounded by the geographic scope the principal is granted, so neither
// a `sales` nor a `cashier` role exists.
type UserRole string

const (
	// UserRoleSuperAdmin is rank 1 and the only globally scoped role.
	UserRoleSuperAdmin UserRole = "superAdmin"
	// UserRoleCountryAdmin is rank 2 and administers one country.
	UserRoleCountryAdmin UserRole = "countryAdmin"
	// UserRoleDepotManager is rank 3 and manages the depots it is granted.
	UserRoleDepotManager UserRole = "depotManager"
	// UserRoleMarketing is rank 4 and works the markets it is granted.
	UserRoleMarketing UserRole = "marketing"
	// UserRoleWarehouseManager is rank 5 and runs depot warehouse operations.
	UserRoleWarehouseManager UserRole = "warehouseManager"
	// UserRoleWarehouseOperator is rank 6 and executes depot warehouse tasks.
	UserRoleWarehouseOperator UserRole = "warehouseOperator"
)

func (r UserRole) IsValid() bool {
	switch r {
	case UserRoleSuperAdmin, UserRoleCountryAdmin, UserRoleDepotManager,
		UserRoleMarketing, UserRoleWarehouseManager, UserRoleWarehouseOperator:
		return true
	default:
		return false
	}
}

func (r UserRole) String() string { return string(r) }
