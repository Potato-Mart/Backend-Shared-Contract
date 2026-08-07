package enums_test

import (
	"testing"

	commonenum "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/common/shared"
	accessenum "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/identity/access"
	accountenum "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/identity/account"
)

func TestAccountEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "accountenum.AccountStatus", valid: []stringEnum{accountenum.AccountStatusPending, accountenum.AccountStatusActive, accountenum.AccountStatusSuspended, accountenum.AccountStatusClosed, accountenum.AccountStatusDeleted}, invalid: accountenum.AccountStatus("__invalid__")},
		{name: "accountenum.AccountType", valid: []stringEnum{accountenum.AccountTypeAdminUser, accountenum.AccountTypeRetailCustomer, accountenum.AccountTypeWholesaleCustomer}, invalid: accountenum.AccountType("__invalid__")},
		{name: "commonenum.Portal", valid: []stringEnum{commonenum.PortalControl, commonenum.PortalRetail, commonenum.PortalWholesale}, invalid: commonenum.Portal("__invalid__")},
		{name: "accessenum.PortalAccessStatus", valid: []stringEnum{accessenum.PortalAccessStatusPending, accessenum.PortalAccessStatusActive, accessenum.PortalAccessStatusSuspended, accessenum.PortalAccessStatusRevoked}, invalid: accessenum.PortalAccessStatus("__invalid__")},
	})
}
