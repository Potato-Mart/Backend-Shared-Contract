package enums_test

import (
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v25/pkg/contracts/common/identity/identity_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v25/pkg/contracts/identity/access/access_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v25/pkg/contracts/identity/account/account_enums"
)

func TestAccountEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "accountenum.AccountStatus", valid: []stringEnum{account_enums.AccountStatusPending, account_enums.AccountStatusActive, account_enums.AccountStatusSuspended, account_enums.AccountStatusClosed, account_enums.AccountStatusDeleted}, invalid: account_enums.AccountStatus("__invalid__")},
		{name: "accountenum.AccountType", valid: []stringEnum{account_enums.AccountTypeAdminUser, account_enums.AccountTypeRetailCustomer, account_enums.AccountTypeWholesaleCustomer}, invalid: account_enums.AccountType("__invalid__")},
		{name: "commonenum.Portal", valid: []stringEnum{identity_enums.PortalControl, identity_enums.PortalRetail, identity_enums.PortalWholesale}, invalid: identity_enums.Portal("__invalid__")},
		{name: "accessenum.PortalAccessStatus", valid: []stringEnum{access_enums.PortalAccessStatusPending, access_enums.PortalAccessStatusActive, access_enums.PortalAccessStatusSuspended, access_enums.PortalAccessStatusRevoked}, invalid: access_enums.PortalAccessStatus("__invalid__")},
	})
}
