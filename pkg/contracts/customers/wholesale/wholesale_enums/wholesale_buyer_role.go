package wholesale_enums

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
