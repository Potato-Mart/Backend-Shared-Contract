package enums_test

import (
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/customers/wholesale/wholesale_enums"
)

func TestWholesaleEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "wholesaleenum.OrganisationAccessStatus", valid: []stringEnum{wholesale_enums.OrganisationAccessStatusPending, wholesale_enums.OrganisationAccessStatusActive, wholesale_enums.OrganisationAccessStatusSuspended, wholesale_enums.OrganisationAccessStatusRevoked}, invalid: wholesale_enums.OrganisationAccessStatus("__invalid__")},
		{name: "wholesaleenum.WholesaleBuyerRole", valid: []stringEnum{wholesale_enums.WholesaleBuyerRoleOwner, wholesale_enums.WholesaleBuyerRoleBuyer, wholesale_enums.WholesaleBuyerRoleFinance, wholesale_enums.WholesaleBuyerRoleReadOnly, wholesale_enums.WholesaleBuyerRoleGroupOrderManager}, invalid: wholesale_enums.WholesaleBuyerRole("__invalid__")},
		{name: "wholesaleenum.WholesalePermission", valid: []stringEnum{wholesale_enums.WholesalePermissionProductsView, wholesale_enums.WholesalePermissionCartWrite, wholesale_enums.WholesalePermissionCheckoutSubmit, wholesale_enums.WholesalePermissionOrdersViewOwn, wholesale_enums.WholesalePermissionOrdersViewOrg, wholesale_enums.WholesalePermissionOrdersReorder, wholesale_enums.WholesalePermissionInvoicesViewOwn, wholesale_enums.WholesalePermissionInvoicesViewOrg, wholesale_enums.WholesalePermissionInvoicesPay, wholesale_enums.WholesalePermissionAccountView, wholesale_enums.WholesalePermissionTeamView, wholesale_enums.WholesalePermissionFavouriteListsViewOrg, wholesale_enums.WholesalePermissionFavouriteListsWriteOrg, wholesale_enums.WholesalePermissionGroupOrdersViewOrg, wholesale_enums.WholesalePermissionGroupOrdersManage, wholesale_enums.WholesalePermissionGroupOrdersInvite, wholesale_enums.WholesalePermissionGroupOrdersSubmit, wholesale_enums.WholesalePermissionGroupOrderDiscountApply}, invalid: wholesale_enums.WholesalePermission("__invalid__")},
		{name: "wholesaleenum.WholesaleApplicationState", valid: []stringEnum{wholesale_enums.WholesaleApplicationStateMissing, wholesale_enums.WholesaleApplicationStatePending, wholesale_enums.WholesaleApplicationStateApproved, wholesale_enums.WholesaleApplicationStateRejected, wholesale_enums.WholesaleApplicationStateSuspended}, invalid: wholesale_enums.WholesaleApplicationState("__invalid__")},
		{name: "wholesaleenum.WholesaleOrganisationStatus", valid: []stringEnum{wholesale_enums.WholesaleOrganisationStatusPending, wholesale_enums.WholesaleOrganisationStatusApproved, wholesale_enums.WholesaleOrganisationStatusSuspended, wholesale_enums.WholesaleOrganisationStatusRejected, wholesale_enums.WholesaleOrganisationStatusClosed}, invalid: wholesale_enums.WholesaleOrganisationStatus("__invalid__")},
	})
}
