package enums_test

import (
	"reflect"
	"testing"

	accountenum "github.com/Potato-Mart/Backend-Shared-Contract/v15/pkg/enums/account"
)

func TestAccountEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "accountenum.AccountStatus", valid: []stringEnum{accountenum.AccountStatusPending, accountenum.AccountStatusActive, accountenum.AccountStatusSuspended, accountenum.AccountStatusClosed, accountenum.AccountStatusDeleted}, invalid: accountenum.AccountStatus("__invalid__")},
		{name: "accountenum.AccountType", valid: []stringEnum{accountenum.AccountTypeAdminUser, accountenum.AccountTypeRetailCustomer, accountenum.AccountTypeWholesaleCustomer}, invalid: accountenum.AccountType("__invalid__")},
		{name: "accountenum.Portal", valid: []stringEnum{accountenum.PortalControl, accountenum.PortalRetail, accountenum.PortalWholesale}, invalid: accountenum.Portal("__invalid__")},
		{name: "accountenum.PortalAccessStatus", valid: []stringEnum{accountenum.PortalAccessStatusPending, accountenum.PortalAccessStatusActive, accountenum.PortalAccessStatusSuspended, accountenum.PortalAccessStatusRevoked}, invalid: accountenum.PortalAccessStatus("__invalid__")},
	})
}

func TestAccountTypePortalAdmission(t *testing.T) {
	allowed := map[accountenum.AccountType]accountenum.Portal{
		accountenum.AccountTypeAdminUser:         accountenum.PortalControl,
		accountenum.AccountTypeRetailCustomer:    accountenum.PortalRetail,
		accountenum.AccountTypeWholesaleCustomer: accountenum.PortalWholesale,
	}

	for accountType, portal := range allowed {
		if !accountType.IsAllowedInPortal(portal) {
			t.Fatalf("%s should be allowed in %s", accountType, portal)
		}
	}

	rejected := []struct {
		accountType accountenum.AccountType
		portal      accountenum.Portal
	}{
		{accountenum.AccountTypeAdminUser, accountenum.PortalRetail},
		{accountenum.AccountTypeAdminUser, accountenum.PortalWholesale},
		{accountenum.AccountTypeRetailCustomer, accountenum.PortalControl},
		{accountenum.AccountTypeRetailCustomer, accountenum.PortalWholesale},
		{accountenum.AccountTypeWholesaleCustomer, accountenum.PortalControl},
		{accountenum.AccountTypeWholesaleCustomer, accountenum.PortalRetail},
	}

	for _, tt := range rejected {
		if tt.accountType.IsAllowedInPortal(tt.portal) {
			t.Fatalf("%s should not be allowed in %s", tt.accountType, tt.portal)
		}
	}
}

func TestPortalAccountTypeHelpers(t *testing.T) {
	tests := []struct {
		portal      accountenum.Portal
		accountType accountenum.AccountType
	}{
		{accountenum.PortalControl, accountenum.AccountTypeAdminUser},
		{accountenum.PortalRetail, accountenum.AccountTypeRetailCustomer},
		{accountenum.PortalWholesale, accountenum.AccountTypeWholesaleCustomer},
	}

	for _, tt := range tests {
		if !tt.portal.RequiresAccountType(tt.accountType) {
			t.Fatalf("%s should require %s", tt.portal, tt.accountType)
		}

		required, ok := tt.portal.RequiredAccountType()
		if !ok {
			t.Fatalf("%s should have a required account type", tt.portal)
		}
		if required != tt.accountType {
			t.Fatalf("%s required account type = %s, want %s", tt.portal, required, tt.accountType)
		}

		want := []accountenum.AccountType{tt.accountType}
		if got := tt.portal.AllowedAccountTypes(); !reflect.DeepEqual(got, want) {
			t.Fatalf("%s allowed account types = %#v, want %#v", tt.portal, got, want)
		}
		if got := accountenum.AccountTypesForPortal(tt.portal); !reflect.DeepEqual(got, want) {
			t.Fatalf("accountenum.AccountTypesForPortal(%s) = %#v, want %#v", tt.portal, got, want)
		}
	}

	if _, ok := accountenum.Portal("__invalid__").RequiredAccountType(); ok {
		t.Fatal("invalid portal should not have a required account type")
	}
	if got := accountenum.Portal("__invalid__").AllowedAccountTypes(); got != nil {
		t.Fatalf("invalid portal allowed account types = %#v, want nil", got)
	}
}

func TestPortalAccessStatusCanAccess(t *testing.T) {
	if !accountenum.PortalAccessStatusActive.CanAccess() {
		t.Fatal("active portal access should allow access")
	}

	for _, status := range []accountenum.PortalAccessStatus{
		accountenum.PortalAccessStatusPending,
		accountenum.PortalAccessStatusSuspended,
		accountenum.PortalAccessStatusRevoked,
	} {
		if status.CanAccess() {
			t.Fatalf("%s should not allow access", status)
		}
	}
}

func TestAccountStatusTerminalState(t *testing.T) {
	for _, status := range []accountenum.AccountStatus{accountenum.AccountStatusClosed, accountenum.AccountStatusDeleted} {
		if !status.IsTerminal() {
			t.Fatalf("%s should be terminal", status)
		}
	}

	for _, status := range []accountenum.AccountStatus{
		accountenum.AccountStatusPending,
		accountenum.AccountStatusActive,
		accountenum.AccountStatusSuspended,
	} {
		if status.IsTerminal() {
			t.Fatalf("%s should not be terminal", status)
		}
	}
}
