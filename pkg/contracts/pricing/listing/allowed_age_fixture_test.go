package listing

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v44/pkg/contracts/pricing/listing/listing_enums"
)

// This locks the shared service acceptance fixture and canonical wire. Services
// run its expected outcomes against their evaluators; Contract contains none.
func TestAllowedAgeAcceptanceFixture(t *testing.T) {
	payload, err := os.ReadFile("testdata/allowed-age-groups.json")
	if err != nil {
		t.Fatal(err)
	}
	var fixture struct {
		ThresholdCases []struct {
			Policy  SaleRestriction `json:"policy"`
			Allowed []int32         `json:"allowed_completed_ages"`
			Refused []int32         `json:"refused_completed_ages"`
		} `json:"threshold_cases"`
		CalendarCases []struct {
			DOB  string `json:"dob"`
			Date string `json:"market_date"`
			Age  int32  `json:"completed_age"`
		} `json:"calendar_cases"`
		ApplicabilityCases []struct {
			Buyer    string `json:"buyer"`
			DOB      bool   `json:"usable_saved_dob"`
			View     bool   `json:"can_view"`
			Purchase bool   `json:"age_rule_allows_purchase"`
		} `json:"applicability_cases"`
	}
	if err := json.Unmarshal(payload, &fixture); err != nil {
		t.Fatal(err)
	}
	if len(fixture.ThresholdCases) != 2 {
		t.Fatal("both allowed groups required")
	}
	for i, tc := range fixture.ThresholdCases {
		comparison := []listing_enums.AgeComparison{listing_enums.AgeComparisonBelow, listing_enums.AgeComparisonAbove}[i]
		if tc.Policy.Kind != listing_enums.SaleRestrictionKindAge || tc.Policy.AgeYears == nil || *tc.Policy.AgeYears != 18 || tc.Policy.AgeComparison != comparison {
			t.Fatalf("invalid canonical policy: %+v", tc.Policy)
		}
		allowed := [][]int32{{17}, {18, 19}}[i]
		refused := [][]int32{{18, 19}, {17}}[i]
		if !reflect.DeepEqual(tc.Allowed, allowed) || !reflect.DeepEqual(tc.Refused, refused) {
			t.Fatalf("allowed-group boundary changed: %+v", tc)
		}
		wire, err := json.Marshal(tc.Policy)
		if err != nil {
			t.Fatal(err)
		}
		want := `{"kind":"age","age_years":18,"age_comparison":"` + string(comparison) + `"}`
		if string(wire) != want {
			t.Fatalf("policy wire = %s", wire)
		}
	}
	if len(fixture.CalendarCases) != 4 {
		t.Fatal("leap and non-leap birthday cases required")
	}
	for i, tc := range fixture.CalendarCases {
		if tc.DOB != "2008-02-29" || tc.Date != []string{"2026-02-28", "2026-03-01", "2028-02-28", "2028-02-29"}[i] || tc.Age != []int32{17, 18, 19, 20}[i] {
			t.Fatalf("birthday fixture changed: %+v", tc)
		}
	}
	if len(fixture.ApplicabilityCases) != 3 {
		t.Fatal("retail, guest and wholesale cases required")
	}
	for i, tc := range fixture.ApplicabilityCases {
		if tc.Buyer != []string{"retail", "guest", "wholesale"}[i] || tc.DOB || !tc.View || tc.Purchase != (i == 2) {
			t.Fatalf("applicability fixture changed: %+v", tc)
		}
	}
}
