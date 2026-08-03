package enums_test

import (
	"testing"

	geographyenum "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/enums/geography"
)

func TestGeographyEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "geographyenum.AdministrativeAreaType", valid: []stringEnum{geographyenum.AdministrativeAreaState, geographyenum.AdministrativeAreaTerritory, geographyenum.AdministrativeAreaProvince, geographyenum.AdministrativeAreaPrefecture, geographyenum.AdministrativeAreaRegion, geographyenum.AdministrativeAreaDistrict}, invalid: geographyenum.AdministrativeAreaType("__invalid__")},
		{name: "geographyenum.GeographicScopeMode", valid: []stringEnum{geographyenum.GeographicScopeModeGlobal, geographyenum.GeographicScopeModeTargeted}, invalid: geographyenum.GeographicScopeMode("__invalid__")},
		{name: "geographyenum.GeographicTargetKind", valid: []stringEnum{geographyenum.GeographicTargetCountry, geographyenum.GeographicTargetSubdivision, geographyenum.GeographicTargetDepotRegion, geographyenum.GeographicTargetDepot}, invalid: geographyenum.GeographicTargetKind("__invalid__")},
		{name: "geographyenum.GeographicContextSource", valid: []stringEnum{geographyenum.GeographicContextSourceRetailCustomerProfile, geographyenum.GeographicContextSourceWholesaleOrganisationProfile, geographyenum.GeographicContextSourceGlobalFallback}, invalid: geographyenum.GeographicContextSource("__invalid__")},
	})
}
