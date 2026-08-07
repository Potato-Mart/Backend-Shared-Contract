package enums_test

import (
	"testing"

	securityenum "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/common/security"
	identityenum "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/identity/account"
	roleenum "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/identity/role"
)

func TestIdentityEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "identityenum.AuthIdentityProvider", valid: []stringEnum{identityenum.AuthIdentityProviderPassword, identityenum.AuthIdentityProviderGoogle, identityenum.AuthIdentityProviderApple, identityenum.AuthIdentityProviderAzureAD, identityenum.AuthIdentityProviderOkta, identityenum.AuthIdentityProviderPasskey, identityenum.AuthIdentityProviderServiceToken, identityenum.AuthIdentityProviderLine, identityenum.AuthIdentityProviderDiscord, identityenum.AuthIdentityProviderMicrosoft, identityenum.AuthIdentityProviderOIDC}, invalid: identityenum.AuthIdentityProvider("__invalid__")},
		{name: "identityenum.AuthIdentityStatus", valid: []stringEnum{identityenum.AuthIdentityStatusActive, identityenum.AuthIdentityStatusDisabled, identityenum.AuthIdentityStatusRevoked}, invalid: identityenum.AuthIdentityStatus("__invalid__")},
		{name: "identityenum.DeviceType", valid: []stringEnum{identityenum.DeviceTypeDesktop, identityenum.DeviceTypeMobile, identityenum.DeviceTypeTablet, identityenum.DeviceTypeAPI}, invalid: identityenum.DeviceType("__invalid__")},
		{name: "securityenum.IdentityDomain", valid: []stringEnum{securityenum.IdentityDomainCustomer, securityenum.IdentityDomainWorkforce, securityenum.IdentityDomainService}, invalid: securityenum.IdentityDomain("__invalid__")},
		{name: "identityenum.UserPreferredLanguage", valid: []stringEnum{identityenum.PreferredLanguageEnglish, identityenum.PreferredLanguageTraditionalChinese, identityenum.PreferredLanguageSimplifiedChinese}, invalid: identityenum.UserPreferredLanguage("__invalid__")},
		{name: "roleenum.UserRole", valid: []stringEnum{roleenum.UserRoleSuperAdmin, roleenum.UserRoleAdmin, roleenum.UserRoleSales, roleenum.UserRoleWarehouse, roleenum.UserRoleWarehouseOperator, roleenum.UserRoleMarketing}, invalid: roleenum.UserRole("__invalid__")},
	})
}
