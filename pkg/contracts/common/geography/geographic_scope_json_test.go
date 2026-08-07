package geography

import (
	"encoding/json"
	"testing"
)

func TestGeographicScopeAndContextJSON(t *testing.T) {
	scope := GeographicScope{
		Mode: GeographicScopeModeTargeted,
		Targets: []GeographicTarget{
			{Kind: GeographicTargetCountry, Code: "AU"},
			{Kind: GeographicTargetDepot, Code: "AU-VIC-MEL-DC-01"},
		},
	}
	payload, err := json.Marshal(scope)
	if err != nil {
		t.Fatalf("marshal geographic scope: %v", err)
	}
	if string(payload) != `{"mode":"TARGETED","targets":[{"kind":"COUNTRY","code":"AU"},{"kind":"DEPOT","code":"AU-VIC-MEL-DC-01"}]}` {
		t.Fatalf("GeographicScope JSON = %s", payload)
	}

	context := GeographicContext{
		Source:             GeographicContextSourceRetailCustomerProfile,
		CountryCode:        "AU",
		SubdivisionCode:    "AU-VIC",
		DepotRegionCode:    "AU-VIC-MEL",
		DepotCode:          "AU-VIC-MEL-DC-01",
		MatchedTargetKind:  GeographicTargetDepot,
		MatchedTargetCode:  "AU-VIC-MEL-DC-01",
		ScopeRevision:      7,
		RuleRevision:       11,
		EvaluationTimezone: "Australia/Melbourne",
	}
	payload, err = json.Marshal(context)
	if err != nil {
		t.Fatalf("marshal geographic context: %v", err)
	}
	var got map[string]any
	if err := json.Unmarshal(payload, &got); err != nil {
		t.Fatalf("unmarshal geographic context: %v", err)
	}
	for _, key := range []string{"source", "country_code", "subdivision_code", "depot_region_code", "depot_code", "matched_target_kind", "matched_target_code", "scope_revision", "rule_revision", "evaluation_timezone"} {
		if _, ok := got[key]; !ok {
			t.Fatalf("GeographicContext JSON missing %q: %s", key, payload)
		}
	}
}

func TestGlobalFallbackGeographicContextOmitsUnresolvedProfileGeography(t *testing.T) {
	payload, err := json.Marshal(GeographicContext{
		Source:             GeographicContextSourceGlobalFallback,
		ScopeRevision:      2,
		RuleRevision:       4,
		EvaluationTimezone: "Etc/UTC",
	})
	if err != nil {
		t.Fatalf("marshal global fallback context: %v", err)
	}
	var got map[string]any
	if err := json.Unmarshal(payload, &got); err != nil {
		t.Fatalf("unmarshal global fallback context: %v", err)
	}
	if got["source"] != "GLOBAL_FALLBACK" || got["evaluation_timezone"] != "Etc/UTC" {
		t.Fatalf("global fallback context mismatch: %s", payload)
	}
	for _, key := range []string{"country_code", "subdivision_code", "depot_region_code", "depot_code", "matched_target_kind", "matched_target_code"} {
		if _, ok := got[key]; ok {
			t.Fatalf("global fallback context should omit unresolved %q: %s", key, payload)
		}
	}
}
