package wallet_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	geography "github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/geography/geography_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/pricing/promotion"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/pricing/promotion/promotion_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/pricing/wallet"
)

func TestCouponReusesPromotionScopePeriodTermsAndControls(t *testing.T) {
	startsAt := time.Date(2026, 8, 9, 0, 0, 0, 0, time.UTC)
	minimumAmount := money.Money{AmountMinor: 2_000, Currency: "AUD"}
	payload, err := json.Marshal(wallet.Coupon{
		ID: "coupon_1", Code: "NSW10", Description: "NSW potato coupon",
		Scope: promotion.PromotionScope{
			MatchMode: promotion_enums.PromotionMatchModeAll,
			Groups: []promotion.PromotionScopeGroup{{
				MatchMode:      promotion_enums.PromotionMatchModeAny,
				CategoryTagIDs: []string{"tag_potato"},
			}},
		},
		Period: promotion.PromotionPeriod{StartsAt: &startsAt, Timezone: "Australia/Sydney"},
		Terms:  []promotion.PromotionTerm{{Key: "minimum_order_amount", MoneyValue: &minimumAmount}},
		Controls: promotion.PromotionControls{
			GeographicScope: geography.GeographicScope{
				Mode:    geography_enums.GeographicScopeModeTargeted,
				Targets: []geography.GeographicTarget{{Kind: geography_enums.GeographicTargetSubdivision, Code: "AU-NSW"}},
			},
		},
	})
	if err != nil {
		t.Fatalf("marshal coupon: %v", err)
	}
	for _, want := range []string{`"scope":{"unrestricted":false`, `"period":{"starts_at":"2026-08-09T00:00:00Z","timezone":"Australia/Sydney"}`, `"terms":[{`, `"controls":{"geographic_scope"`} {
		if !strings.Contains(string(payload), want) {
			t.Fatalf("coupon JSON = %s, want %s", payload, want)
		}
	}
	for _, retired := range []string{`"applies_to"`, `"product_sku_codes"`, `"category_tags"`, `"discount_type"`, `"discount_value"`, `"usage_limit"`, `"used_count"`, `"per_customer_limit"`, `"schedule_timezone"`} {
		if strings.Contains(string(payload), retired) {
			t.Fatalf("coupon JSON retained %s: %s", retired, payload)
		}
	}

	var got wallet.Coupon
	if err := json.Unmarshal(payload, &got); err != nil {
		t.Fatalf("unmarshal coupon: %v", err)
	}
	if got.Scope.Groups[0].CategoryTagIDs[0] != "tag_potato" || got.Period.StartsAt == nil || got.Terms[0].MoneyValue == nil || got.Controls.GeographicScope.Mode != geography_enums.GeographicScopeModeTargeted {
		t.Fatalf("coupon reusable promotion fields did not round-trip: %+v", got)
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
