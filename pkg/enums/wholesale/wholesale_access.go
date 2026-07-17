package wholesaleenum

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

// WholesalePermission is a persisted/signed permission value. Role resolution
// and forbidden-permission policy are owned by Management.
type WholesalePermission string

const (
	WholesalePermissionProductsView            WholesalePermission = "products.view"
	WholesalePermissionCartWrite               WholesalePermission = "cart.write"
	WholesalePermissionCheckoutSubmit          WholesalePermission = "checkout.submit"
	WholesalePermissionOrdersViewOwn           WholesalePermission = "orders.view_own"
	WholesalePermissionOrdersViewOrg           WholesalePermission = "orders.view_org"
	WholesalePermissionOrdersReorder           WholesalePermission = "orders.reorder"
	WholesalePermissionInvoicesViewOwn         WholesalePermission = "invoices.view_own"
	WholesalePermissionInvoicesViewOrg         WholesalePermission = "invoices.view_org"
	WholesalePermissionInvoicesPay             WholesalePermission = "invoices.pay"
	WholesalePermissionAccountView             WholesalePermission = "account.view"
	WholesalePermissionTeamView                WholesalePermission = "team.view"
	WholesalePermissionFavouriteListsViewOrg   WholesalePermission = "favourite_lists.view_org"
	WholesalePermissionFavouriteListsWriteOrg  WholesalePermission = "favourite_lists.write_org"
	WholesalePermissionGroupOrdersViewOrg      WholesalePermission = "group_orders.view_org"
	WholesalePermissionGroupOrdersManage       WholesalePermission = "group_orders.manage_org"
	WholesalePermissionGroupOrdersInvite       WholesalePermission = "group_orders.invite"
	WholesalePermissionGroupOrdersSubmit       WholesalePermission = "group_orders.submit"
	WholesalePermissionGroupOrderDiscountApply WholesalePermission = "group_order_discount.apply"
)

func (p WholesalePermission) String() string { return string(p) }

func (p WholesalePermission) IsValid() bool {
	switch p {
	case WholesalePermissionProductsView, WholesalePermissionCartWrite,
		WholesalePermissionCheckoutSubmit, WholesalePermissionOrdersViewOwn,
		WholesalePermissionOrdersViewOrg, WholesalePermissionOrdersReorder,
		WholesalePermissionInvoicesViewOwn, WholesalePermissionInvoicesViewOrg,
		WholesalePermissionInvoicesPay, WholesalePermissionAccountView,
		WholesalePermissionTeamView, WholesalePermissionFavouriteListsViewOrg,
		WholesalePermissionFavouriteListsWriteOrg,
		WholesalePermissionGroupOrdersViewOrg, WholesalePermissionGroupOrdersManage,
		WholesalePermissionGroupOrdersInvite, WholesalePermissionGroupOrdersSubmit,
		WholesalePermissionGroupOrderDiscountApply:
		return true
	default:
		return false
	}
}
