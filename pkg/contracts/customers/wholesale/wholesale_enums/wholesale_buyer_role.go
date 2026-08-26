package wholesale_enums

// WholesaleBuyerRole is the organisation-scoped role a person holds inside one
// wholesale organisation. It is deliberately not role_enums.UserRole: that
// enum is the workforce ladder, and these two vocabularies never mix.
//
// A buyer role carries no rank, because organisation membership is not a
// hierarchy of platform authority. Its scope is the organisation the access
// record names, not a geography, so it holds no country, market, or depot.
// The role-to-permission matrix and the forbidden-permission policy are owned
// and seeded by Backend-Customers; Identity only transports the resolved keys.
type WholesaleBuyerRole string

const (
	WholesaleBuyerRoleOwner             WholesaleBuyerRole = "owner"
	WholesaleBuyerRoleBuyer             WholesaleBuyerRole = "buyer"
	WholesaleBuyerRoleFinance           WholesaleBuyerRole = "finance"
	WholesaleBuyerRoleReadOnly          WholesaleBuyerRole = "read_only"
	WholesaleBuyerRoleGroupOrderManager WholesaleBuyerRole = "group_order_manager"
)

func (r WholesaleBuyerRole) String() string { return string(r) }

func (r WholesaleBuyerRole) IsValid() bool {
	switch r {
	case WholesaleBuyerRoleOwner, WholesaleBuyerRoleBuyer, WholesaleBuyerRoleFinance,
		WholesaleBuyerRoleReadOnly, WholesaleBuyerRoleGroupOrderManager:
		return true
	default:
		return false
	}
}
