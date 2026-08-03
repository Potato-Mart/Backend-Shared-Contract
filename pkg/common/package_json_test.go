package common

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestPackageCompositionJSONPreservesKnownZeroQuantities(t *testing.T) {
	composition := PackageCompositionSnapshot{
		TotalBaseUnits: 0,
		Components: []PackageComponentSnapshot{{
			PackageOptionID: "pkg_case_12",
			HandlingUnit:    PackageHandlingUnitCase,
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
	if !PackageHandlingUnitEach.IsValid() || !PackageHandlingUnitCase.IsValid() || PackageHandlingUnit("BOX").IsValid() {
		t.Fatal("package handling unit validity mismatch")
	}
}
