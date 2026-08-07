package enums_test

import (
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v24/pkg/contracts/common/geography/geography_enums"
)

func TestGeographyEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "geography.AdministrativeAreaType", valid: []stringEnum{geography_enums.AdministrativeAreaState, geography_enums.AdministrativeAreaTerritory, geography_enums.AdministrativeAreaProvince, geography_enums.AdministrativeAreaPrefecture, geography_enums.AdministrativeAreaRegion, geography_enums.AdministrativeAreaDistrict}, invalid: geography_enums.AdministrativeAreaType("__invalid__")},
		{name: "geography.GeographicScopeMode", valid: []stringEnum{geography_enums.GeographicScopeModeGlobal, geography_enums.GeographicScopeModeTargeted}, invalid: geography_enums.GeographicScopeMode("__invalid__")},
		{name: "geography.GeographicTargetKind", valid: []stringEnum{geography_enums.GeographicTargetCountry, geography_enums.GeographicTargetSubdivision, geography_enums.GeographicTargetDepotRegion, geography_enums.GeographicTargetDepot}, invalid: geography_enums.GeographicTargetKind("__invalid__")},
		{name: "geography.GeographicContextSource", valid: []stringEnum{geography_enums.GeographicContextSourceRetailCustomerProfile, geography_enums.GeographicContextSourceWholesaleOrganisationProfile, geography_enums.GeographicContextSourceGlobalFallback}, invalid: geography_enums.GeographicContextSource("__invalid__")},
	})
}
