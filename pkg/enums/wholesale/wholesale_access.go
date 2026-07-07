package wholesaleenum

// WholesaleBuyerRole is the organisation-scoped buyer role assigned to an
// OrganisationAccess record.
type WholesaleBuyerRole string

const (
	WholesaleBuyerRoleOwner    WholesaleBuyerRole = "owner"
	WholesaleBuyerRoleBuyer    WholesaleBuyerRole = "buyer"
	WholesaleBuyerRoleFinance  WholesaleBuyerRole = "finance"
	WholesaleBuyerRoleReadOnly WholesaleBuyerRole = "read_only"
)

func (r WholesaleBuyerRole) String() string { return string(r) }

func (r WholesaleBuyerRole) IsValid() bool {
	switch r {
	case WholesaleBuyerRoleOwner,
		WholesaleBuyerRoleBuyer,
		WholesaleBuyerRoleFinance,
		WholesaleBuyerRoleReadOnly:
		return true
	default:
		return false
	}
}

// WholesalePermission is a signed storefront permission derived from a buyer
// role and stamped into store JWTs by Management.
type WholesalePermission string

const (
	WholesalePermissionProductsView       WholesalePermission = "products.view"
	WholesalePermissionCartWrite          WholesalePermission = "cart.write"
	WholesalePermissionCheckoutSubmit     WholesalePermission = "checkout.submit"
	WholesalePermissionOrdersViewOwn      WholesalePermission = "orders.view_own"
	WholesalePermissionOrdersViewOrg      WholesalePermission = "orders.view_org"
	WholesalePermissionOrdersReorder      WholesalePermission = "orders.reorder"
	WholesalePermissionInvoicesViewOwn    WholesalePermission = "invoices.view_own"
	WholesalePermissionInvoicesViewOrg    WholesalePermission = "invoices.view_org"
	WholesalePermissionInvoicesPay        WholesalePermission = "invoices.pay"
	WholesalePermissionAccountView        WholesalePermission = "account.view"
	WholesalePermissionTeamView           WholesalePermission = "team.view"
	WholesalePermissionFavouritesWrite    WholesalePermission = "favourites.write"
	WholesalePermissionOrderListsViewOwn  WholesalePermission = "order_lists.view_own"
	WholesalePermissionOrderListsWriteOwn WholesalePermission = "order_lists.write_own"
	WholesalePermissionOrderListsViewOrg  WholesalePermission = "order_lists.view_org"
	WholesalePermissionOrderListsWriteOrg WholesalePermission = "order_lists.write_org"

	// Group-order manager permissions (added v13.1.0). The wholesale
	// group-order manager creates/manages the group order, invites retail
	// participants by email, submits it, and applies for a per-group discount.
	WholesalePermissionGroupOrdersManage       WholesalePermission = "group_orders.manage_org"
	WholesalePermissionGroupOrdersInvite       WholesalePermission = "group_orders.invite"
	WholesalePermissionGroupOrdersSubmit       WholesalePermission = "group_orders.submit"
	WholesalePermissionGroupOrderDiscountApply WholesalePermission = "group_order_discount.apply"
)

func (p WholesalePermission) String() string { return string(p) }

func (p WholesalePermission) IsValid() bool {
	switch p {
	case WholesalePermissionProductsView,
		WholesalePermissionCartWrite,
		WholesalePermissionCheckoutSubmit,
		WholesalePermissionOrdersViewOwn,
		WholesalePermissionOrdersViewOrg,
		WholesalePermissionOrdersReorder,
		WholesalePermissionInvoicesViewOwn,
		WholesalePermissionInvoicesViewOrg,
		WholesalePermissionInvoicesPay,
		WholesalePermissionAccountView,
		WholesalePermissionTeamView,
		WholesalePermissionFavouritesWrite,
		WholesalePermissionOrderListsViewOwn,
		WholesalePermissionOrderListsWriteOwn,
		WholesalePermissionOrderListsViewOrg,
		WholesalePermissionOrderListsWriteOrg,
		WholesalePermissionGroupOrdersManage,
		WholesalePermissionGroupOrdersInvite,
		WholesalePermissionGroupOrdersSubmit,
		WholesalePermissionGroupOrderDiscountApply:
		return true
	default:
		return false
	}
}

// PermissionsForWholesaleBuyerRole returns the default permission set for a
// role. Backends may add compatibility handling before choosing the role.
func PermissionsForWholesaleBuyerRole(role WholesaleBuyerRole) []WholesalePermission {
	switch role {
	case WholesaleBuyerRoleOwner:
		return []WholesalePermission{
			WholesalePermissionProductsView,
			WholesalePermissionCartWrite,
			WholesalePermissionCheckoutSubmit,
			WholesalePermissionOrdersViewOwn,
			WholesalePermissionOrdersViewOrg,
			WholesalePermissionOrdersReorder,
			WholesalePermissionInvoicesViewOwn,
			WholesalePermissionInvoicesViewOrg,
			WholesalePermissionInvoicesPay,
			WholesalePermissionAccountView,
			WholesalePermissionTeamView,
			WholesalePermissionFavouritesWrite,
			WholesalePermissionOrderListsViewOwn,
			WholesalePermissionOrderListsWriteOwn,
			WholesalePermissionOrderListsViewOrg,
			WholesalePermissionOrderListsWriteOrg,
			WholesalePermissionGroupOrdersManage,
			WholesalePermissionGroupOrdersInvite,
			WholesalePermissionGroupOrdersSubmit,
			WholesalePermissionGroupOrderDiscountApply,
		}
	case WholesaleBuyerRoleBuyer:
		return []WholesalePermission{
			WholesalePermissionProductsView,
			WholesalePermissionCartWrite,
			WholesalePermissionCheckoutSubmit,
			WholesalePermissionOrdersViewOwn,
			WholesalePermissionOrdersReorder,
			WholesalePermissionInvoicesViewOwn,
			WholesalePermissionAccountView,
			WholesalePermissionFavouritesWrite,
			WholesalePermissionOrderListsViewOwn,
			WholesalePermissionOrderListsWriteOwn,
			WholesalePermissionOrderListsViewOrg,
		}
	case WholesaleBuyerRoleFinance:
		return []WholesalePermission{
			WholesalePermissionProductsView,
			WholesalePermissionOrdersViewOwn,
			WholesalePermissionInvoicesViewOwn,
			WholesalePermissionInvoicesViewOrg,
			WholesalePermissionInvoicesPay,
			WholesalePermissionAccountView,
		}
	case WholesaleBuyerRoleReadOnly:
		return []WholesalePermission{
			WholesalePermissionProductsView,
			WholesalePermissionOrdersViewOwn,
			WholesalePermissionInvoicesViewOwn,
			WholesalePermissionAccountView,
			WholesalePermissionTeamView,
			WholesalePermissionOrderListsViewOrg,
		}
	default:
		return nil
	}
}

func WholesalePermissionStrings(perms []WholesalePermission) []string {
	if len(perms) == 0 {
		return nil
	}
	out := make([]string, 0, len(perms))
	for _, perm := range perms {
		out = append(out, perm.String())
	}
	return out
}

func HasWholesalePermission(perms []string, required WholesalePermission) bool {
	if !required.IsValid() {
		return false
	}
	needle := required.String()
	for _, perm := range perms {
		if perm == needle {
			return true
		}
	}
	return false
}
