package quote

import (
	"encoding/json"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/pricing/quote/quote_enums"
)

func TestCartManualDiscountEvidenceFixtureRoundTrip(t *testing.T) {
	raw, err := os.ReadFile("testdata/cart_manual_discount_evidence.json")
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}
	var decoded CartManualDiscountEvidence
	if err := json.Unmarshal(raw, &decoded); err != nil {
		t.Fatalf("unmarshal fixture: %v", err)
	}
	want := cartManualDiscountFixture()
	if !reflect.DeepEqual(decoded, want) {
		t.Fatalf("fixture decoded to %+v, want %+v", decoded, want)
	}
	encoded, err := json.Marshal(decoded)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	wantJSON := `{"mode":"percent","percent_basis_points":1000,"pre_discount_total":{"amount_minor":4500,"currency":"AUD"},"applied_amount":{"amount_minor":450,"currency":"AUD"},"applied_tax":{"amount_minor":41,"currency":"AUD"},"coupon_code":"SPRING5","coupon_amount":{"amount_minor":500,"currency":"AUD"},"reason_code":"staff_discretion","reason":"Regular customer goodwill","actor_user_id":"cashier_1","applied_at":"2026-11-06T02:03:04Z"}`
	if string(encoded) != wantJSON {
		t.Fatalf("manual discount JSON = %s, want %s", encoded, wantJSON)
	}
}

func TestCartManualDiscountEvidenceCarriesOnlyTheRequestedValueForItsMode(t *testing.T) {
	amount := money.Money{AmountMinor: 300, Currency: "AUD"}
	target := money.Money{AmountMinor: 4000, Currency: "AUD"}
	for _, tc := range []struct {
		mode    quote_enums.ManualDiscountMode
		set     func(*CartManualDiscountEvidence)
		present string
		absent  []string
	}{
		{quote_enums.ManualDiscountModeAmount, func(e *CartManualDiscountEvidence) { e.RequestedAmount = &amount }, `"requested_amount":{"amount_minor":300,"currency":"AUD"}`, []string{"percent_basis_points", "target_total"}},
		{quote_enums.ManualDiscountModeSetTotal, func(e *CartManualDiscountEvidence) { e.TargetTotal = &target }, `"target_total":{"amount_minor":4000,"currency":"AUD"}`, []string{"percent_basis_points", "requested_amount"}},
	} {
		t.Run(string(tc.mode), func(t *testing.T) {
			evidence := cartManualDiscountFixture()
			evidence.Mode, evidence.PercentBasisPoints = tc.mode, nil
			evidence.CouponCode, evidence.CouponAmount, evidence.Reason = "", nil, ""
			tc.set(&evidence)
			payload, err := json.Marshal(evidence)
			if err != nil {
				t.Fatal(err)
			}
			if !strings.Contains(string(payload), `"mode":"`+string(tc.mode)+`"`) || !strings.Contains(string(payload), tc.present) {
				t.Fatalf("manual discount JSON = %s, want mode %s and %s", payload, tc.mode, tc.present)
			}
			for _, key := range append(tc.absent, "coupon_code", "coupon_amount", "reason") {
				if strings.Contains(string(payload), `"`+key+`":`) {
					t.Fatalf("manual discount JSON unexpectedly includes %s: %s", key, payload)
				}
			}
			var decoded CartManualDiscountEvidence
			if err := json.Unmarshal(payload, &decoded); err != nil || !reflect.DeepEqual(evidence, decoded) {
				t.Fatalf("manual discount round trip: %+v %v", decoded, err)
			}
		})
	}
}

func TestPriceSnapshotCarriesManualDiscountLineAllocation(t *testing.T) {
	snapshot := PriceSnapshot{QuoteID: "quote_1", LineID: "line_1", ManualDiscount: &ManualDiscountLineAllocation{
		AllocatedGross: money.Money{AmountMinor: 150, Currency: "AUD"},
		AllocatedTax:   money.Money{AmountMinor: 14, Currency: "AUD"},
	}}
	payload, err := json.Marshal(snapshot)
	if err != nil {
		t.Fatal(err)
	}
	want := `"manual_discount":{"allocated_gross":{"amount_minor":150,"currency":"AUD"},"allocated_tax":{"amount_minor":14,"currency":"AUD"}}`
	if !strings.Contains(string(payload), want) {
		t.Fatalf("PriceSnapshot JSON = %s, want %s", payload, want)
	}
	var decoded PriceSnapshot
	if err := json.Unmarshal(payload, &decoded); err != nil || !reflect.DeepEqual(snapshot, decoded) {
		t.Fatalf("manual discount allocation round trip: %+v %v", decoded.ManualDiscount, err)
	}
}

func cartManualDiscountFixture() CartManualDiscountEvidence {
	basisPoints := int64(1000)
	return CartManualDiscountEvidence{
		Mode:               quote_enums.ManualDiscountModePercent,
		PercentBasisPoints: &basisPoints,
		PreDiscountTotal:   money.Money{AmountMinor: 4500, Currency: "AUD"},
		AppliedAmount:      money.Money{AmountMinor: 450, Currency: "AUD"},
		AppliedTax:         money.Money{AmountMinor: 41, Currency: "AUD"},
		CouponCode:         "SPRING5",
		CouponAmount:       &money.Money{AmountMinor: 500, Currency: "AUD"},
		ReasonCode:         "staff_discretion",
		Reason:             "Regular customer goodwill",
		ActorUserID:        "cashier_1",
		AppliedAt:          time.Date(2026, 11, 6, 2, 3, 4, 0, time.UTC),
	}
}
