package wallet_test

import (
	"encoding/json"
	"testing"
	"time"

	geography "github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/geography/geography_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/pricing/promotion"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/pricing/wallet"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/pricing/wallet/wallet_enums"
)

func TestCouponCarriesRequiredGeographicScope(t *testing.T) {
	payload, err := json.Marshal(wallet.Coupon{
		ID: "coupon_1", Code: "NSW10",
		AppliesTo:    wallet_enums.CouponAppliesToAll,
		ActiveWindow: promotion.ActiveWindow{ScheduleTimezone: "Australia/Sydney", IsActive: true},
		GeographicScope: geography.GeographicScope{
			Mode:    geography_enums.GeographicScopeModeTargeted,
			Targets: []geography.GeographicTarget{{Kind: geography_enums.GeographicTargetSubdivision, Code: "AU-NSW"}},
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

func TestCouponUsageRecordFreezesGeographicContext(t *testing.T) {
	now := time.Date(2026, 8, 1, 0, 0, 0, 0, time.UTC)
	context := geography.GeographicContext{
		Source:      geography_enums.GeographicContextSourceWholesaleOrganisationProfile,
		CountryCode: "AU", SubdivisionCode: "AU-VIC", DepotRegionCode: "AU-VIC-MEL",
		MatchedTargetKind: geography_enums.GeographicTargetSubdivision, MatchedTargetCode: "AU-VIC",
		ScopeRevision: 5, RuleRevision: 8, EvaluationTimezone: "Australia/Melbourne",
	}
	payload, err := json.Marshal(wallet.CouponUsageRecord{
		ID: "usage_1", CouponCode: "VIC10", RedeemedOrderNumber: "order_1",
		DiscountAmount:    money.Money{AmountMinor: 1000, Currency: "AUD"},
		GeographicContext: context, RedeemedAt: now, CreatedAt: now,
	})
	if err != nil {
		t.Fatalf("marshal coupon usage: %v", err)
	}
	var got map[string]any
	if err := json.Unmarshal(payload, &got); err != nil {
		t.Fatalf("unmarshal coupon usage: %v", err)
	}
	resolved, ok := got["geographic_context"].(map[string]any)
	if !ok || resolved["source"] != "WHOLESALE_ORGANISATION_PROFILE" || resolved["matched_target_code"] != "AU-VIC" {
		t.Fatalf("coupon usage geographic context mismatch: %s", payload)
	}
}
