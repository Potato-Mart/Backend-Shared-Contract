package enums_test

import (
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v24/pkg/contracts/customers/retail/retail_enums"
)

func TestCustomerEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "customerenum.ChurnRisk", valid: []stringEnum{retail_enums.ChurnRiskLow, retail_enums.ChurnRiskMedium, retail_enums.ChurnRiskHigh}, invalid: retail_enums.ChurnRisk("__invalid__")},
		{name: "customerenum.CustomerAcquisitionSource", valid: []stringEnum{retail_enums.CustomerAcquisitionSourceOnline, retail_enums.CustomerAcquisitionSourcePOS, retail_enums.CustomerAcquisitionSourceImport, retail_enums.CustomerAcquisitionSourceManual, retail_enums.CustomerAcquisitionSourcePhone}, invalid: retail_enums.CustomerAcquisitionSource("__invalid__")},
		{name: "customerenum.CustomerGender", valid: []stringEnum{retail_enums.CustomerGenderFemale, retail_enums.CustomerGenderMale, retail_enums.CustomerGenderNonBinary}, invalid: retail_enums.CustomerGender("__invalid__")},
		{name: "customerenum.CustomerIdentityKind", valid: []stringEnum{retail_enums.CustomerIdentityKindPhone, retail_enums.CustomerIdentityKindEmail, retail_enums.CustomerIdentityKindLine, retail_enums.CustomerIdentityKindMemberCard, retail_enums.CustomerIdentityKindPOSID, retail_enums.CustomerIdentityKindExternal}, invalid: retail_enums.CustomerIdentityKind("__invalid__")},
		{name: "customerenum.CustomerStatus", valid: []stringEnum{retail_enums.CustomerStatusActive, retail_enums.CustomerStatusInactive, retail_enums.CustomerStatusBlocked, retail_enums.CustomerStatusClosed}, invalid: retail_enums.CustomerStatus("__invalid__")},
		{name: "customerenum.BuyerType", valid: []stringEnum{retail_enums.BuyerTypeGuestRetail, retail_enums.BuyerTypeRetailCustomer, retail_enums.BuyerTypeWholesaleOrganisation}, invalid: retail_enums.BuyerType("__invalid__")},
	})
}
