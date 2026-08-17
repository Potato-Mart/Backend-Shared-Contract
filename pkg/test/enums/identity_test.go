package enums_test

import (
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/security/security_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/identity/access/access_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/identity/account/account_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/identity/role/role_enums"
)

func TestIdentityEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "identityenum.AuthIdentityProvider", valid: []stringEnum{account_enums.AuthIdentityProviderPassword, account_enums.AuthIdentityProviderGoogle, account_enums.AuthIdentityProviderApple, account_enums.AuthIdentityProviderAzureAD, account_enums.AuthIdentityProviderOkta, account_enums.AuthIdentityProviderPasskey, account_enums.AuthIdentityProviderServiceToken, account_enums.AuthIdentityProviderLine, account_enums.AuthIdentityProviderFacebook, account_enums.AuthIdentityProviderDiscord, account_enums.AuthIdentityProviderMicrosoft, account_enums.AuthIdentityProviderOIDC}, invalid: account_enums.AuthIdentityProvider("__invalid__")},
		{name: "identityenum.NotificationTopicGroup", valid: []stringEnum{account_enums.NotificationTopicGroupPayment, account_enums.NotificationTopicGroupOrder, account_enums.NotificationTopicGroupReceipt, account_enums.NotificationTopicGroupProductUpdates, account_enums.NotificationTopicGroupPromotions}, invalid: account_enums.NotificationTopicGroup("__invalid__")},
		{name: "identityenum.AuthIdentityStatus", valid: []stringEnum{account_enums.AuthIdentityStatusActive, account_enums.AuthIdentityStatusDisabled, account_enums.AuthIdentityStatusRevoked}, invalid: account_enums.AuthIdentityStatus("__invalid__")},
		{name: "identityenum.DeviceType", valid: []stringEnum{account_enums.DeviceTypeDesktop, account_enums.DeviceTypeMobile, account_enums.DeviceTypeTablet, account_enums.DeviceTypeAPI}, invalid: account_enums.DeviceType("__invalid__")},
		{name: "securityenum.IdentityDomain", valid: []stringEnum{security_enums.IdentityDomainCustomer, security_enums.IdentityDomainWorkforce, security_enums.IdentityDomainService}, invalid: security_enums.IdentityDomain("__invalid__")},
		{name: "identityenum.UserPreferredLanguage", valid: []stringEnum{account_enums.PreferredLanguageEnglish, account_enums.PreferredLanguageTraditionalChinese, account_enums.PreferredLanguageSimplifiedChinese}, invalid: account_enums.UserPreferredLanguage("__invalid__")},
		{name: "roleenum.UserRole", valid: []stringEnum{role_enums.UserRoleSuperAdmin, role_enums.UserRoleCountryAdmin, role_enums.UserRoleDepotManager, role_enums.UserRoleSales, role_enums.UserRoleMarketing, role_enums.UserRoleWarehouseManager, role_enums.UserRoleWarehouseOperator}, invalid: role_enums.UserRole("__invalid__")},
		{name: "accessenum.ScopeLevel", valid: []stringEnum{access_enums.ScopeLevelGlobal, access_enums.ScopeLevelCountry, access_enums.ScopeLevelMarket, access_enums.ScopeLevelDepot}, invalid: access_enums.ScopeLevel("__invalid__")},
	})
}

// TestWorkforceRoleAndScopeWireValuesAreLocked pins the exact wire string of
// every authorization enum value.
//
// assertStringEnums cannot do this: it compares an enum constant against its
// own reflected string, so editing a constant's value still round-trips and
// still satisfies IsValid. Identity mints these strings into JWT role and
// scope claims, and all seven services plus the admin client authorize on
// them, so a silent edit would re-authorize live principals with no test
// turning red. The literals below are the lock.
func TestWorkforceRoleAndScopeWireValuesAreLocked(t *testing.T) {
	roles := []struct {
		rank int
		role role_enums.UserRole
		wire string
	}{
		{1, role_enums.UserRoleSuperAdmin, "superAdmin"},
		{2, role_enums.UserRoleCountryAdmin, "countryAdmin"},
		{3, role_enums.UserRoleDepotManager, "depotManager"},
		{4, role_enums.UserRoleSales, "sales"},
		{5, role_enums.UserRoleMarketing, "marketing"},
		{6, role_enums.UserRoleWarehouseManager, "warehouseManager"},
		{7, role_enums.UserRoleWarehouseOperator, "warehouseOperator"},
	}
	for index, entry := range roles {
		if entry.rank != index+1 {
			t.Errorf("role %q is listed at rank %d; the built-in hierarchy runs 1..7 in order", entry.wire, entry.rank)
		}
		if got := entry.role.String(); got != entry.wire {
			t.Errorf("rank %d role wire value = %q, want %q", entry.rank, got, entry.wire)
		}
		if !entry.role.IsValid() {
			t.Errorf("rank %d role %q must validate", entry.rank, entry.wire)
		}
	}

	// The v28 cut-over retires these three keys. They must never validate
	// again: a stale token or seeded document carrying one is not a role.
	for _, retired := range []string{"admin", "warehouse", "cashier"} {
		if role_enums.UserRole(retired).IsValid() {
			t.Errorf("retired v27 role %q still validates", retired)
		}
	}

	scopes := map[access_enums.ScopeLevel]string{
		access_enums.ScopeLevelGlobal:  "global",
		access_enums.ScopeLevelCountry: "country",
		access_enums.ScopeLevelMarket:  "market",
		access_enums.ScopeLevelDepot:   "depot",
	}
	for level, wire := range scopes {
		if got := level.String(); got != wire {
			t.Errorf("scope level wire value = %q, want %q", got, wire)
		}
		if !level.IsValid() {
			t.Errorf("scope level %q must validate", wire)
		}
	}

	// Depots are the only site identity in the platform: there is no store
	// scope level, and nothing may reintroduce one.
	if access_enums.ScopeLevel("store").IsValid() {
		t.Error(`scope level "store" validates; depots are the only site identity`)
	}
}
