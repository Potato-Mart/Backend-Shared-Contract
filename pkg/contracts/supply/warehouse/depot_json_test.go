package warehouse_test

import (
	"encoding/json"

	geography "github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/geography"

	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/supply/warehouse"
)

func TestDepotHierarchyJSONShapes(t *testing.T) {
	now := time.Date(2026, 8, 4, 4, 30, 0, 0, time.UTC)
	audit := audit.AuditFields{CreatedAt: now, UpdatedAt: now}
	depot := warehouse.Depot{
		ID:         "depot_1",
		Code:       "AU-VIC-MEL-DC-01",
		Name:       "Melbourne DC",
		RegionCode: "AU-VIC-MEL",
		Address: geography.Address{
			Line1:      "1 Example Drive",
			Locality:   "Dandenong South",
			PostalCode: "3175",
			Country:    geography.CountryRef{Code: "AU", Name: "Australia"},
		},
		Timezone:    "Australia/Melbourne",
		IsActive:    true,
		AuditFields: audit,
	}
	shape := marshalDepotObject(t, depot)
	for _, key := range []string{"region_code", "address", "timezone"} {
		if _, ok := shape[key]; !ok {
			t.Fatalf("depot JSON missing %q: %+v", key, shape)
		}
	}
	if _, ok := shape["postcode_rules"]; ok {
		t.Fatalf("depot JSON retained postcode_rules: %+v", shape)
	}
	if _, ok := shape["address"].(map[string]any); !ok {
		t.Fatalf("depot address is not an object: %+v", shape["address"])
	}

	regionShape := marshalDepotObject(t, warehouse.DepotRegion{
		ID: "region_1", Code: "AU-VIC-MEL", Name: "Melbourne Metro",
		CountryCode: "AU", AdministrativeAreaCode: "AU-VIC", IsActive: true,
		AuditFields: audit,
	})
	for _, key := range []string{"country_code", "administrative_area_code"} {
		if _, ok := regionShape[key]; !ok {
			t.Fatalf("depot region JSON missing %q: %+v", key, regionShape)
		}
	}
	if _, ok := regionShape["depots"]; ok {
		t.Fatalf("depot region JSON contains nested depots: %+v", regionShape)
	}

	coverageShape := marshalDepotObject(t, warehouse.DepotCoverageRule{
		ID: "coverage_1", DepotCode: "AU-VIC-MEL-DC-01", CountryCode: "AU",
		Priority: 1, IsActive: true, AuditFields: audit,
	})
	for _, key := range []string{"depot_code", "country_code", "priority", "is_active"} {
		if _, ok := coverageShape[key]; !ok {
			t.Fatalf("depot coverage JSON missing %q: %+v", key, coverageShape)
		}
	}
	if _, ok := coverageShape["administrative_area_code"]; ok {
		t.Fatalf("country-level coverage emitted empty administrative area: %+v", coverageShape)
	}
	if _, ok := coverageShape["postal_code"]; ok {
		t.Fatalf("country-level coverage emitted empty postal code: %+v", coverageShape)
	}
}

func marshalDepotObject(t *testing.T, value any) map[string]any {
	t.Helper()
	raw, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("marshal %T: %v", value, err)
	}
	var shape map[string]any
	if err := json.Unmarshal(raw, &shape); err != nil {
		t.Fatalf("unmarshal %T JSON: %v", value, err)
	}
	return shape
}
