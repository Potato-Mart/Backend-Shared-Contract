package promotion

import (
	"encoding/json"
	geography "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/common/geography"
	common "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/common/shared"
	"testing"
	"time"
)

func TestCouponCarriesRequiredGeographicScope(t *testing.T) {
	payload, err := json.Marshal(Coupon{
		ID: "coupon_1", Code: "NSW10",
		AppliesTo:    CouponAppliesToAll,
		ActiveWindow: ActiveWindow{ScheduleTimezone: "Australia/Sydney", IsActive: true},
		GeographicScope: geography.GeographicScope{
			Mode:    geography.GeographicScopeModeTargeted,
			Targets: []geography.GeographicTarget{{Kind: geography.GeographicTargetSubdivision, Code: "AU-NSW"}},
		},
	})
	if err != nil {
		t.Fatalf("marshal coupon: %v", err)
	}
	var got map[string]any
	if err := json.Unmarshal(payload, &got); err != nil {
		t.Fatalf("unmarshal coupon JSON: %v", err)
	}
	scope, ok := got["geographic_scope"].(map[string]any)
	if !ok || scope["mode"] != "TARGETED" || got["schedule_timezone"] != "Australia/Sydney" {
		t.Fatalf("coupon geographic scope mismatch: %s", payload)
	}
}

func TestEffectivePromotionAndCouponUsageFreezeGeographicContext(t *testing.T) {
	now := time.Date(2026, 8, 1, 0, 0, 0, 0, time.UTC)
	context := geography.GeographicContext{
		Source:      geography.GeographicContextSourceWholesaleOrganisationProfile,
		CountryCode: "AU", SubdivisionCode: "AU-VIC", DepotRegionCode: "AU-VIC-MEL",
		MatchedTargetKind: geography.GeographicTargetSubdivision, MatchedTargetCode: "AU-VIC",
		ScopeRevision: 5, RuleRevision: 8, EvaluationTimezone: "Australia/Melbourne",
	}
	values := []any{
		EffectivePromotion{
			PromotionID: "prm_1", SeriesKey: "series_1", Class: PromotionClassNormal,
			TargetScope: DiscountScopeProduct, ScheduleTimezone: "Australia/Melbourne",
			GeographicContext: context,
		},
		CouponUsageRecord{
			ID: "usage_1", CouponCode: "VIC10", RedeemedOrderNumber: "order_1",
			DiscountAmount:    common.Money{AmountMinor: 1000, Currency: "AUD"},
			GeographicContext: context, RedeemedAt: now, CreatedAt: now,
		},
	}
	for _, value := range values {
		payload, err := json.Marshal(value)
		if err != nil {
			t.Fatalf("marshal geographic snapshot: %v", err)
		}
		var got map[string]any
		if err := json.Unmarshal(payload, &got); err != nil {
			t.Fatalf("unmarshal geographic snapshot: %v", err)
		}
		resolved, ok := got["geographic_context"].(map[string]any)
		if !ok || resolved["source"] != "WHOLESALE_ORGANISATION_PROFILE" || resolved["matched_target_code"] != "AU-VIC" {
			t.Fatalf("geographic context mismatch: %s", payload)
		}
	}
}
