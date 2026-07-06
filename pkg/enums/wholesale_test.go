package enums_test

import (
	"testing"

	wholesaleenum "github.com/Potato-Mart/Backend-Shared-Contract/v13/pkg/enums/wholesale"
)

func TestWholesaleEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "wholesaleenum.OrganisationAccessStatus", valid: []stringEnum{wholesaleenum.OrganisationAccessStatusPending, wholesaleenum.OrganisationAccessStatusActive, wholesaleenum.OrganisationAccessStatusSuspended, wholesaleenum.OrganisationAccessStatusRevoked}, invalid: wholesaleenum.OrganisationAccessStatus("__invalid__")},
		{name: "wholesaleenum.WholesaleBuyerRole", valid: []stringEnum{wholesaleenum.WholesaleBuyerRoleOwner, wholesaleenum.WholesaleBuyerRoleBuyer, wholesaleenum.WholesaleBuyerRoleFinance, wholesaleenum.WholesaleBuyerRoleReadOnly}, invalid: wholesaleenum.WholesaleBuyerRole("__invalid__")},
		{name: "wholesaleenum.WholesalePermission", valid: []stringEnum{wholesaleenum.WholesalePermissionProductsView, wholesaleenum.WholesalePermissionCartWrite, wholesaleenum.WholesalePermissionCheckoutSubmit, wholesaleenum.WholesalePermissionOrdersViewOwn, wholesaleenum.WholesalePermissionOrdersViewOrg, wholesaleenum.WholesalePermissionOrdersReorder, wholesaleenum.WholesalePermissionInvoicesViewOwn, wholesaleenum.WholesalePermissionInvoicesViewOrg, wholesaleenum.WholesalePermissionInvoicesPay, wholesaleenum.WholesalePermissionAccountView, wholesaleenum.WholesalePermissionTeamView, wholesaleenum.WholesalePermissionFavouritesWrite, wholesaleenum.WholesalePermissionOrderListsViewOwn, wholesaleenum.WholesalePermissionOrderListsWriteOwn, wholesaleenum.WholesalePermissionOrderListsViewOrg, wholesaleenum.WholesalePermissionOrderListsWriteOrg}, invalid: wholesaleenum.WholesalePermission("__invalid__")},
		{name: "wholesaleenum.WholesaleApplicationState", valid: []stringEnum{wholesaleenum.WholesaleApplicationStateMissing, wholesaleenum.WholesaleApplicationStatePending, wholesaleenum.WholesaleApplicationStateApproved, wholesaleenum.WholesaleApplicationStateRejected, wholesaleenum.WholesaleApplicationStateSuspended}, invalid: wholesaleenum.WholesaleApplicationState("__invalid__")},
		{name: "wholesaleenum.WholesaleOrganisationStatus", valid: []stringEnum{wholesaleenum.WholesaleOrganisationStatusPending, wholesaleenum.WholesaleOrganisationStatusApproved, wholesaleenum.WholesaleOrganisationStatusSuspended, wholesaleenum.WholesaleOrganisationStatusRejected, wholesaleenum.WholesaleOrganisationStatusClosed}, invalid: wholesaleenum.WholesaleOrganisationStatus("__invalid__")},
	})
}

func TestWholesaleBuyerRolePermissions(t *testing.T) {
	ownerPerms := wholesaleenum.PermissionsForWholesaleBuyerRole(wholesaleenum.WholesaleBuyerRoleOwner)
	if len(ownerPerms) == 0 {
		t.Fatal("owner should receive permissions")
	}
	if !wholesaleenum.HasWholesalePermission(wholesaleenum.WholesalePermissionStrings(ownerPerms), wholesaleenum.WholesalePermissionTeamView) {
		t.Fatal("owner should receive team.view")
	}
	if !wholesaleenum.HasWholesalePermission(wholesaleenum.WholesalePermissionStrings(ownerPerms), wholesaleenum.WholesalePermissionOrderListsWriteOrg) {
		t.Fatal("owner should receive order_lists.write_org")
	}

	buyerPerms := wholesaleenum.WholesalePermissionStrings(wholesaleenum.PermissionsForWholesaleBuyerRole(wholesaleenum.WholesaleBuyerRoleBuyer))
	if wholesaleenum.HasWholesalePermission(buyerPerms, wholesaleenum.WholesalePermissionInvoicesViewOrg) {
		t.Fatal("buyer should not receive organisation invoice visibility")
	}
	if !wholesaleenum.HasWholesalePermission(buyerPerms, wholesaleenum.WholesalePermissionCheckoutSubmit) {
		t.Fatal("buyer should receive checkout.submit")
	}
	if !wholesaleenum.HasWholesalePermission(buyerPerms, wholesaleenum.WholesalePermissionFavouritesWrite) ||
		!wholesaleenum.HasWholesalePermission(buyerPerms, wholesaleenum.WholesalePermissionOrderListsWriteOwn) ||
		!wholesaleenum.HasWholesalePermission(buyerPerms, wholesaleenum.WholesalePermissionOrderListsViewOrg) {
		t.Fatal("buyer should receive own list mutation, org list view, and favourites.write")
	}
	if wholesaleenum.HasWholesalePermission(buyerPerms, wholesaleenum.WholesalePermissionOrderListsWriteOrg) {
		t.Fatal("buyer should not receive organisation list mutation")
	}

	financePerms := wholesaleenum.WholesalePermissionStrings(wholesaleenum.PermissionsForWholesaleBuyerRole(wholesaleenum.WholesaleBuyerRoleFinance))
	if !wholesaleenum.HasWholesalePermission(financePerms, wholesaleenum.WholesalePermissionInvoicesPay) {
		t.Fatal("finance should receive invoices.pay")
	}
	if wholesaleenum.HasWholesalePermission(financePerms, wholesaleenum.WholesalePermissionCheckoutSubmit) {
		t.Fatal("finance should not receive checkout.submit")
	}
	if wholesaleenum.HasWholesalePermission(financePerms, wholesaleenum.WholesalePermissionFavouritesWrite) ||
		wholesaleenum.HasWholesalePermission(financePerms, wholesaleenum.WholesalePermissionOrderListsViewOwn) ||
		wholesaleenum.HasWholesalePermission(financePerms, wholesaleenum.WholesalePermissionOrderListsViewOrg) {
		t.Fatal("finance should not receive procurement list permissions")
	}

	readOnlyPerms := wholesaleenum.WholesalePermissionStrings(wholesaleenum.PermissionsForWholesaleBuyerRole(wholesaleenum.WholesaleBuyerRoleReadOnly))
	if !wholesaleenum.HasWholesalePermission(readOnlyPerms, wholesaleenum.WholesalePermissionOrderListsViewOrg) {
		t.Fatal("read-only should receive organisation list view")
	}
	if wholesaleenum.HasWholesalePermission(readOnlyPerms, wholesaleenum.WholesalePermissionFavouritesWrite) ||
		wholesaleenum.HasWholesalePermission(readOnlyPerms, wholesaleenum.WholesalePermissionOrderListsWriteOwn) ||
		wholesaleenum.HasWholesalePermission(readOnlyPerms, wholesaleenum.WholesalePermissionCartWrite) {
		t.Fatal("read-only should not receive procurement mutation or cart permissions")
	}

	if got := wholesaleenum.PermissionsForWholesaleBuyerRole(wholesaleenum.WholesaleBuyerRole("__invalid__")); got != nil {
		t.Fatalf("invalid role permissions = %#v, want nil", got)
	}
	if wholesaleenum.HasWholesalePermission([]string{wholesaleenum.WholesalePermissionTeamView.String()}, wholesaleenum.WholesalePermission("__invalid__")) {
		t.Fatal("invalid required permission should never match")
	}
}
