package packaging

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v25/pkg/contracts/common/packaging/packaging_enums"
)

func TestPackageCompositionJSONPreservesKnownZeroQuantities(t *testing.T) {
	composition := PackageCompositionSnapshot{
		TotalBaseUnits: 0,
		Components: []PackageComponentSnapshot{{
			PackageOptionID: "pkg_case_12",
			HandlingUnit:    packaging_enums.PackageHandlingUnitCase,
			PackageCount:    0,
			UnitsPerPackage: 12,
			BaseUnits:       0,
		}},
	}
	body, err := json.Marshal(composition)
	if err != nil {
		t.Fatalf("marshal package composition: %v", err)
	}
	for _, want := range []string{`"total_base_units":0`, `"package_count":0`, `"units_per_package":12`, `"base_units":0`, `"handling_unit":"CASE"`} {
		if !strings.Contains(string(body), want) {
			t.Fatalf("package composition JSON missing %s: %s", want, body)
		}
	}
	if !packaging_enums.PackageHandlingUnitEach.IsValid() || !packaging_enums.PackageHandlingUnitCase.IsValid() || packaging_enums.PackageHandlingUnit("BOX").IsValid() {
		t.Fatal("package handling unit validity mismatch")
	}
}
