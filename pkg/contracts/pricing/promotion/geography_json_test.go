package promotion

import (
	"encoding/json"

	geography "github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/geography"

	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/geography/geography_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/pricing/promotion/promotion_enums"
)

func TestEffectivePromotionFreezesGeographicContext(t *testing.T) {
	context := geography.GeographicContext{
		Source:      geography_enums.GeographicContextSourceWholesaleOrganisationProfile,
		CountryCode: "AU", SubdivisionCode: "AU-VIC", DepotRegionCode: "AU-VIC-MEL",
		MatchedTargetKind: geography_enums.GeographicTargetSubdivision, MatchedTargetCode: "AU-VIC",
		ScopeRevision: 5, RuleRevision: 8, EvaluationTimezone: "Australia/Melbourne",
	}
	value := EffectivePromotion{
		PromotionID: "prm_1", SeriesKey: "series_1", Class: promotion_enums.PromotionClassNormal,
		TargetScope: promotion_enums.DiscountScopeProduct, ScheduleTimezone: "Australia/Melbourne",
		GeographicContext: context,
	}
	payload, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("marshal effective promotion: %v", err)
	}
	var got map[string]any
	if err := json.Unmarshal(payload, &got); err != nil {
		t.Fatalf("unmarshal effective promotion: %v", err)
	}
	resolved, ok := got["geographic_context"].(map[string]any)
	if !ok || resolved["source"] != "WHOLESALE_ORGANISATION_PROFILE" || resolved["matched_target_code"] != "AU-VIC" {
		t.Fatalf("effective promotion geographic context mismatch: %s", payload)
	}
}
