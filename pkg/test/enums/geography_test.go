package enums_test

import (
	geography "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/common/geography"
	"testing"
)

func TestGeographyEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "geography.AdministrativeAreaType", valid: []stringEnum{geography.AdministrativeAreaState, geography.AdministrativeAreaTerritory, geography.AdministrativeAreaProvince, geography.AdministrativeAreaPrefecture, geography.AdministrativeAreaRegion, geography.AdministrativeAreaDistrict}, invalid: geography.AdministrativeAreaType("__invalid__")},
		{name: "geography.GeographicScopeMode", valid: []stringEnum{geography.GeographicScopeModeGlobal, geography.GeographicScopeModeTargeted}, invalid: geography.GeographicScopeMode("__invalid__")},
		{name: "geography.GeographicTargetKind", valid: []stringEnum{geography.GeographicTargetCountry, geography.GeographicTargetSubdivision, geography.GeographicTargetDepotRegion, geography.GeographicTargetDepot}, invalid: geography.GeographicTargetKind("__invalid__")},
		{name: "geography.GeographicContextSource", valid: []stringEnum{geography.GeographicContextSourceRetailCustomerProfile, geography.GeographicContextSourceWholesaleOrganisationProfile, geography.GeographicContextSourceGlobalFallback}, invalid: geography.GeographicContextSource("__invalid__")},
	})
}
