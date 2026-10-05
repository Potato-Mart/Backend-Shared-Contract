package geography

import (
	"encoding/json"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v44/pkg/contracts/common/geography/geography_enums"
)

func TestGeographicScopeAndContextJSON(t *testing.T) {
	scope := GeographicScope{
		Mode: geography_enums.GeographicScopeModeTargeted,
		Targets: []GeographicPath{
			{CountryCode: "AU"},
			{CountryCode: "AU", MarketCode: "mkt_au_vic", StateCode: "AU-VIC", DepotCode: "AU-VIC-MEL-DC-01"},
		},
	}
	payload, err := json.Marshal(scope)
	if err != nil {
		t.Fatalf("marshal geographic scope: %v", err)
	}
	if string(payload) != `{"mode":"TARGETED","targets":[{"country_code":"AU"},{"country_code":"AU","market_code":"mkt_au_vic","state_code":"AU-VIC","depot_code":"AU-VIC-MEL-DC-01"}]}` {
		t.Fatalf("GeographicScope JSON = %s", payload)
	}

	context := GeographicContext{
		Source:             geography_enums.GeographicContextSourceDeliveryAddress,
		MarketCode:         "mkt_au_vic",
		Path:               &GeographicPath{CountryCode: "AU", MarketCode: "mkt_au_vic", StateCode: "AU-VIC", DepotCode: "AU-VIC-MEL-DC-01"},
		MatchedPath:        &GeographicPath{CountryCode: "AU", MarketCode: "mkt_au_vic", StateCode: "AU-VIC", DepotCode: "AU-VIC-MEL-DC-01"},
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
	for _, key := range []string{"source", "market_code", "path", "matched_path", "scope_revision", "rule_revision", "evaluation_timezone"} {
		if _, ok := got[key]; !ok {
			t.Fatalf("GeographicContext JSON missing %q: %s", key, payload)
		}
	}
}

func TestGlobalFallbackGeographicContextOmitsUnresolvedProfileGeography(t *testing.T) {
	payload, err := json.Marshal(GeographicContext{
		Source:             geography_enums.GeographicContextSourceGlobalFallback,
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
	for _, key := range []string{"path", "matched_path"} {
		if _, ok := got[key]; ok {
			t.Fatalf("global fallback context should omit unresolved %q: %s", key, payload)
		}
	}
}
