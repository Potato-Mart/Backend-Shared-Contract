package geography

import (
	"encoding/json"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/geography/geography_enums"
)

func TestAddressUsesTypedGeographyJSON(t *testing.T) {
	value := Address{
		Label:      "Main depot",
		Line1:      "1 Example Drive",
		Locality:   "Dandenong South",
		PostalCode: "3175",
		AdministrativeArea: &AdministrativeAreaRef{
			Code: "AU-VIC",
			Name: "Victoria",
			Type: geography_enums.AdministrativeAreaState,
		},
		Country: CountryRef{Code: "AU", Name: "Australia"},
	}

	payload, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("marshal address: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(payload, &got); err != nil {
		t.Fatalf("unmarshal address JSON: %v", err)
	}
	for _, key := range []string{"line1", "locality", "administrative_area", "postal_code", "country"} {
		if _, ok := got[key]; !ok {
			t.Fatalf("Address JSON missing %q: %s", key, payload)
		}
	}
	for _, key := range []string{"city", "state", "postcode"} {
		if _, ok := got[key]; ok {
			t.Fatalf("Address JSON retained removed key %q: %s", key, payload)
		}
	}
	area, ok := got["administrative_area"].(map[string]any)
	if !ok || area["code"] != "AU-VIC" || area["name"] != "Victoria" || area["type"] != "STATE" {
		t.Fatalf("administrative_area = %#v, want typed Victoria reference", got["administrative_area"])
	}
	country, ok := got["country"].(map[string]any)
	if !ok || country["code"] != "AU" || country["name"] != "Australia" {
		t.Fatalf("country = %#v, want typed Australia reference", got["country"])
	}
}
