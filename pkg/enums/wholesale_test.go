package enums_test

import (
	"testing"

	wholesaleenum "github.com/Potato-Mart/Backend-Shared-Contract/v16/pkg/enums/wholesale"
)

func TestWholesaleEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "wholesaleenum.OrganisationAccessStatus", valid: []stringEnum{wholesaleenum.OrganisationAccessStatusPending, wholesaleenum.OrganisationAccessStatusActive, wholesaleenum.OrganisationAccessStatusSuspended, wholesaleenum.OrganisationAccessStatusRevoked}, invalid: wholesaleenum.OrganisationAccessStatus("__invalid__")},
		{name: "wholesaleenum.WholesaleBuyerRole", valid: []stringEnum{wholesaleenum.WholesaleBuyerRoleOwner, wholesaleenum.WholesaleBuyerRoleBuyer, wholesaleenum.WholesaleBuyerRoleFinance, wholesaleenum.WholesaleBuyerRoleReadOnly, wholesaleenum.WholesaleBuyerRoleGroupOrderManager}, invalid: wholesaleenum.WholesaleBuyerRole("__invalid__")},
		{name: "wholesaleenum.WholesalePermission", valid: []stringEnum{wholesaleenum.WholesalePermissionProductsView, wholesaleenum.WholesalePermissionCartWrite, wholesaleenum.WholesalePermissionCheckoutSubmit, wholesaleenum.WholesalePermissionOrdersViewOwn, wholesaleenum.WholesalePermissionOrdersViewOrg, wholesaleenum.WholesalePermissionOrdersReorder, wholesaleenum.WholesalePermissionInvoicesViewOwn, wholesaleenum.WholesalePermissionInvoicesViewOrg, wholesaleenum.WholesalePermissionInvoicesPay, wholesaleenum.WholesalePermissionAccountView, wholesaleenum.WholesalePermissionTeamView, wholesaleenum.WholesalePermissionFavouritesWrite, wholesaleenum.WholesalePermissionOrderListsViewOwn, wholesaleenum.WholesalePermissionOrderListsWriteOwn, wholesaleenum.WholesalePermissionOrderListsViewOrg, wholesaleenum.WholesalePermissionOrderListsWriteOrg, wholesaleenum.WholesalePermissionGroupOrdersViewOrg, wholesaleenum.WholesalePermissionGroupOrdersManage, wholesaleenum.WholesalePermissionGroupOrdersInvite, wholesaleenum.WholesalePermissionGroupOrdersSubmit, wholesaleenum.WholesalePermissionGroupOrderDiscountApply}, invalid: wholesaleenum.WholesalePermission("__invalid__")},
		{name: "wholesaleenum.WholesaleApplicationState", valid: []stringEnum{wholesaleenum.WholesaleApplicationStateMissing, wholesaleenum.WholesaleApplicationStatePending, wholesaleenum.WholesaleApplicationStateApproved, wholesaleenum.WholesaleApplicationStateRejected, wholesaleenum.WholesaleApplicationStateSuspended}, invalid: wholesaleenum.WholesaleApplicationState("__invalid__")},
		{name: "wholesaleenum.WholesaleOrganisationStatus", valid: []stringEnum{wholesaleenum.WholesaleOrganisationStatusPending, wholesaleenum.WholesaleOrganisationStatusApproved, wholesaleenum.WholesaleOrganisationStatusSuspended, wholesaleenum.WholesaleOrganisationStatusRejected, wholesaleenum.WholesaleOrganisationStatusClosed}, invalid: wholesaleenum.WholesaleOrganisationStatus("__invalid__")},
	})
}
