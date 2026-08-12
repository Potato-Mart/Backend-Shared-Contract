package enums_test

import (
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/security/security_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/identity/account/account_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/identity/role/role_enums"
)

func TestIdentityEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "identityenum.AuthIdentityProvider", valid: []stringEnum{account_enums.AuthIdentityProviderPassword, account_enums.AuthIdentityProviderGoogle, account_enums.AuthIdentityProviderApple, account_enums.AuthIdentityProviderAzureAD, account_enums.AuthIdentityProviderOkta, account_enums.AuthIdentityProviderPasskey, account_enums.AuthIdentityProviderServiceToken, account_enums.AuthIdentityProviderLine, account_enums.AuthIdentityProviderDiscord, account_enums.AuthIdentityProviderMicrosoft, account_enums.AuthIdentityProviderOIDC}, invalid: account_enums.AuthIdentityProvider("__invalid__")},
		{name: "identityenum.AuthIdentityStatus", valid: []stringEnum{account_enums.AuthIdentityStatusActive, account_enums.AuthIdentityStatusDisabled, account_enums.AuthIdentityStatusRevoked}, invalid: account_enums.AuthIdentityStatus("__invalid__")},
		{name: "identityenum.DeviceType", valid: []stringEnum{account_enums.DeviceTypeDesktop, account_enums.DeviceTypeMobile, account_enums.DeviceTypeTablet, account_enums.DeviceTypeAPI}, invalid: account_enums.DeviceType("__invalid__")},
		{name: "securityenum.IdentityDomain", valid: []stringEnum{security_enums.IdentityDomainCustomer, security_enums.IdentityDomainWorkforce, security_enums.IdentityDomainService}, invalid: security_enums.IdentityDomain("__invalid__")},
		{name: "identityenum.UserPreferredLanguage", valid: []stringEnum{account_enums.PreferredLanguageEnglish, account_enums.PreferredLanguageTraditionalChinese, account_enums.PreferredLanguageSimplifiedChinese}, invalid: account_enums.UserPreferredLanguage("__invalid__")},
		{name: "roleenum.UserRole", valid: []stringEnum{role_enums.UserRoleSuperAdmin, role_enums.UserRoleAdmin, role_enums.UserRoleSales, role_enums.UserRoleWarehouse, role_enums.UserRoleWarehouseOperator, role_enums.UserRoleMarketing, role_enums.UserRoleCashier}, invalid: role_enums.UserRole("__invalid__")},
	})
}
