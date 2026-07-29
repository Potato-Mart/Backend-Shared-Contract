package enums_test

import (
	"testing"

	accountenum "github.com/Potato-Mart/Backend-Shared-Contract/v21/pkg/enums/account"
)

func TestAccountEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "accountenum.AccountStatus", valid: []stringEnum{accountenum.AccountStatusPending, accountenum.AccountStatusActive, accountenum.AccountStatusSuspended, accountenum.AccountStatusClosed, accountenum.AccountStatusDeleted}, invalid: accountenum.AccountStatus("__invalid__")},
		{name: "accountenum.AccountType", valid: []stringEnum{accountenum.AccountTypeAdminUser, accountenum.AccountTypeRetailCustomer, accountenum.AccountTypeWholesaleCustomer}, invalid: accountenum.AccountType("__invalid__")},
		{name: "accountenum.Portal", valid: []stringEnum{accountenum.PortalControl, accountenum.PortalRetail, accountenum.PortalWholesale}, invalid: accountenum.Portal("__invalid__")},
		{name: "accountenum.PortalAccessStatus", valid: []stringEnum{accountenum.PortalAccessStatusPending, accountenum.PortalAccessStatusActive, accountenum.PortalAccessStatusSuspended, accountenum.PortalAccessStatusRevoked}, invalid: accountenum.PortalAccessStatus("__invalid__")},
	})
}
