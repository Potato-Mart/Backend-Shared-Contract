package role_enums

// RetiredUserRole is a workforce role key the platform used before the
// six-rank model. A retired key is never a valid role: Identity refuses to
// create an owner-defined role with one, and every backend must reject a
// control-portal token whose role is one.
//
// IsValid reports whether a key IS one of the retired keys, so a consumer
// accepts an owner-defined role key only when
// RetiredUserRole(key).IsValid() is false (alongside its own non-empty,
// whitespace and control-character checks).
type RetiredUserRole string

const (
	RetiredUserRoleAdmin     RetiredUserRole = "admin"
	RetiredUserRoleWarehouse RetiredUserRole = "warehouse"
	RetiredUserRoleCashier   RetiredUserRole = "cashier"
	RetiredUserRoleSales     RetiredUserRole = "sales"
	RetiredUserRoleCustomer  RetiredUserRole = "customer"
)

func (r RetiredUserRole) IsValid() bool {
	switch r {
	case RetiredUserRoleAdmin, RetiredUserRoleWarehouse, RetiredUserRoleCashier,
		RetiredUserRoleSales, RetiredUserRoleCustomer:
		return true
	default:
		return false
	}
}

func (r RetiredUserRole) String() string { return string(r) }
