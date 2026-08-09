package shipping

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	geography "github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/geography"
)

func TestZoneUsesTypedGeographicCoverage(t *testing.T) {
	payload, err := json.Marshal(Zone{
		ID:                      "zone_1",
		Name:                    "Sydney metro",
		CountryCode:             "AU",
		AdministrativeAreaCodes: []geography.SubdivisionCode{"AU-NSW"},
		PostalCodes:             []string{"2000"},
		IsActive:                true,
		CreatedAt:               time.Date(2026, 8, 1, 0, 0, 0, 0, time.UTC),
	})
	if err != nil {
		t.Fatalf("marshal shipping zone: %v", err)
	}
	for _, field := range []string{`"country_code":"AU"`, `"administrative_area_codes":["AU-NSW"]`, `"postal_codes":["2000"]`} {
		if !strings.Contains(string(payload), field) {
			t.Fatalf("shipping zone missing %s: %s", field, payload)
		}
	}
	for _, removed := range []string{`"states"`, `"postcodes"`, `"is_local"`} {
		if strings.Contains(string(payload), removed) {
			t.Fatalf("shipping zone retained removed field %s: %s", removed, payload)
		}
	}
}
