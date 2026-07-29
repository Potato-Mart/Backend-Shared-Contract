package enums_test

import (
	"testing"

	customerenum "github.com/Potato-Mart/Backend-Shared-Contract/v20/pkg/enums/customer"
)

func TestCustomerEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "customerenum.ChurnRisk", valid: []stringEnum{customerenum.ChurnRiskLow, customerenum.ChurnRiskMedium, customerenum.ChurnRiskHigh}, invalid: customerenum.ChurnRisk("__invalid__")},
		{name: "customerenum.CustomerAcquisitionSource", valid: []stringEnum{customerenum.CustomerAcquisitionSourceOnline, customerenum.CustomerAcquisitionSourcePOS, customerenum.CustomerAcquisitionSourceImport, customerenum.CustomerAcquisitionSourceManual, customerenum.CustomerAcquisitionSourcePhone}, invalid: customerenum.CustomerAcquisitionSource("__invalid__")},
		{name: "customerenum.CustomerGender", valid: []stringEnum{customerenum.CustomerGenderFemale, customerenum.CustomerGenderMale, customerenum.CustomerGenderNonBinary}, invalid: customerenum.CustomerGender("__invalid__")},
		{name: "customerenum.CustomerIdentityKind", valid: []stringEnum{customerenum.CustomerIdentityKindPhone, customerenum.CustomerIdentityKindEmail, customerenum.CustomerIdentityKindLine, customerenum.CustomerIdentityKindMemberCard, customerenum.CustomerIdentityKindPOSID, customerenum.CustomerIdentityKindExternal}, invalid: customerenum.CustomerIdentityKind("__invalid__")},
		{name: "customerenum.CustomerStatus", valid: []stringEnum{customerenum.CustomerStatusActive, customerenum.CustomerStatusInactive, customerenum.CustomerStatusBlocked, customerenum.CustomerStatusClosed}, invalid: customerenum.CustomerStatus("__invalid__")},
		{name: "customerenum.BuyerType", valid: []stringEnum{customerenum.BuyerTypeGuestRetail, customerenum.BuyerTypeRetailCustomer, customerenum.BuyerTypeWholesaleOrganisation}, invalid: customerenum.BuyerType("__invalid__")},
	})
}
