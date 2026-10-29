package quote

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v48/pkg/contracts/common/money"
)

func TestOfflineUnsupportedPromotionDoesNotInventDifference(t *testing.T) {
	for _, promotionID := range []string{"promo-1", ""} {
		want := OfflinePriceReviewFinding{ClientSaleID: "sale-1", LineNumber: 1, SKUCode: "SKU", PromotionID: promotionID, Reason: "unsupported_promotion"}
		raw, err := json.Marshal(want)
		if err != nil {
			t.Fatal(err)
		}
		for _, forbidden := range []string{`"charged_amount"`, `"expected_amount"`, `"difference"`} {
			if strings.Contains(string(raw), forbidden) {
				t.Fatalf("unknown amount was serialized: %s", raw)
			}
		}
		var got OfflinePriceReviewFinding
		if err := json.Unmarshal(raw, &got); err != nil {
			t.Fatal(err)
		}
		if got.PromotionID != want.PromotionID || got.Difference != nil {
			t.Fatalf("finding changed: %#v", got)
		}
	}
}

func TestCustomPriceOverrideCarriesIncreasesAndDecreases(t *testing.T) {
	for _, amount := range []int64{500, 1500} {
		want := PriceSnapshot{CustomOverride: &CustomPriceOverrideEvidence{
			ActorUserID: "staff-1", Reason: "customer adjustment", SourceApprovedPrice: money.Money{AmountMinor: 1000, Currency: "AUD"},
			OverrideGrossAmount: money.Money{AmountMinor: amount, Currency: "AUD"}, OverriddenAt: time.Date(2026, 10, 1, 0, 0, 0, 0, time.UTC),
		}}
		raw, err := json.Marshal(want)
		if err != nil {
			t.Fatal(err)
		}
		var got PriceSnapshot
		if err := json.Unmarshal(raw, &got); err != nil {
			t.Fatal(err)
		}
		if got.CustomOverride == nil || got.CustomOverride.ActorUserID != "staff-1" || got.CustomOverride.Reason == "" || got.CustomOverride.OverrideGrossAmount.AmountMinor != amount {
			t.Fatalf("override evidence changed: %#v", got.CustomOverride)
		}
	}
}
