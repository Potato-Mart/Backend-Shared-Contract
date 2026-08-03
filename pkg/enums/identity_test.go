package enums_test

import (
	"testing"

	identityenum "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/enums/identity"
)

func TestIdentityEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "identityenum.AuthIdentityProvider", valid: []stringEnum{identityenum.AuthIdentityProviderPassword, identityenum.AuthIdentityProviderGoogle, identityenum.AuthIdentityProviderApple, identityenum.AuthIdentityProviderAzureAD, identityenum.AuthIdentityProviderOkta, identityenum.AuthIdentityProviderPasskey, identityenum.AuthIdentityProviderServiceToken, identityenum.AuthIdentityProviderLine, identityenum.AuthIdentityProviderDiscord, identityenum.AuthIdentityProviderMicrosoft, identityenum.AuthIdentityProviderOIDC}, invalid: identityenum.AuthIdentityProvider("__invalid__")},
		{name: "identityenum.AuthIdentityStatus", valid: []stringEnum{identityenum.AuthIdentityStatusActive, identityenum.AuthIdentityStatusDisabled, identityenum.AuthIdentityStatusRevoked}, invalid: identityenum.AuthIdentityStatus("__invalid__")},
		{name: "identityenum.DeviceType", valid: []stringEnum{identityenum.DeviceTypeDesktop, identityenum.DeviceTypeMobile, identityenum.DeviceTypeTablet, identityenum.DeviceTypeAPI}, invalid: identityenum.DeviceType("__invalid__")},
		{name: "identityenum.IdentityDomain", valid: []stringEnum{identityenum.IdentityDomainCustomer, identityenum.IdentityDomainWorkforce, identityenum.IdentityDomainService}, invalid: identityenum.IdentityDomain("__invalid__")},
		{name: "identityenum.UserPreferredLanguage", valid: []stringEnum{identityenum.PreferredLanguageEnglish, identityenum.PreferredLanguageTraditionalChinese, identityenum.PreferredLanguageSimplifiedChinese}, invalid: identityenum.UserPreferredLanguage("__invalid__")},
		{name: "identityenum.UserRole", valid: []stringEnum{identityenum.UserRoleSuperAdmin, identityenum.UserRoleAdmin, identityenum.UserRoleSales, identityenum.UserRoleWarehouse, identityenum.UserRoleWarehouseOperator, identityenum.UserRoleMarketing}, invalid: identityenum.UserRole("__invalid__")},
	})
}
