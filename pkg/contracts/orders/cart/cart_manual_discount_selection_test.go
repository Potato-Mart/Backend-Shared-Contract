package cart

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

func TestCartManualDiscountSelectionFixtureRoundTrip(t *testing.T) {
	raw, err := os.ReadFile("testdata/cart_manual_discount_selection.json")
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}
	var decoded CartManualDiscountSelection
	if err := json.Unmarshal(raw, &decoded); err != nil {
		t.Fatalf("unmarshal fixture: %v", err)
	}
	want := CartManualDiscountSelection{
		Mode:        quote_enums.ManualDiscountModeAmount,
		Amount:      &money.Money{AmountMinor: 300, Currency: "AUD"},
		ReasonCode:  "staff_discretion",
		Reason:      "Dented tin on shelf",
		ActorUserID: "cashier_1",
		SetAt:       time.Date(2026, 11, 6, 2, 0, 0, 0, time.UTC),
	}
	if !reflect.DeepEqual(decoded, want) {
		t.Fatalf("fixture decoded to %+v, want %+v", decoded, want)
	}
	encoded, err := json.Marshal(Cart{ID: "cart_1", ManualDiscount: &decoded})
	if err != nil {
		t.Fatal(err)
	}
	wantJSON := `"manual_discount":{"mode":"amount","amount":{"amount_minor":300,"currency":"AUD"},"reason_code":"staff_discretion","reason":"Dented tin on shelf","actor_user_id":"cashier_1","set_at":"2026-11-06T02:00:00Z"}`
	if !strings.Contains(string(encoded), wantJSON) {
		t.Fatalf("cart JSON = %s, want %s", encoded, wantJSON)
	}
	var cart Cart
	if err := json.Unmarshal(encoded, &cart); err != nil || !reflect.DeepEqual(cart.ManualDiscount, &want) {
		t.Fatalf("cart manual discount round trip: %+v %v", cart.ManualDiscount, err)
	}
}

func TestCartManualDiscountSelectionModesAndAbsence(t *testing.T) {
	legacy, err := json.Marshal(Cart{ID: "cart_1"})
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(legacy), `"manual_discount"`) {
		t.Fatalf("absent manual discount must be omitted: %s", legacy)
	}
	basisPoints := int64(1500)
	for _, tc := range []struct {
		selection CartManualDiscountSelection
		want      string
	}{
		{CartManualDiscountSelection{Mode: quote_enums.ManualDiscountModePercent, PercentBasisPoints: &basisPoints}, `{"mode":"percent","percent_basis_points":1500,"reason_code":"","actor_user_id":"","set_at":"0001-01-01T00:00:00Z"}`},
		{CartManualDiscountSelection{Mode: quote_enums.ManualDiscountModeSetTotal, TargetTotal: &money.Money{AmountMinor: 2000, Currency: "AUD"}}, `{"mode":"set_total","target_total":{"amount_minor":2000,"currency":"AUD"},"reason_code":"","actor_user_id":"","set_at":"0001-01-01T00:00:00Z"}`},
	} {
		payload, err := json.Marshal(tc.selection)
		if err != nil {
			t.Fatal(err)
		}
		if string(payload) != tc.want {
			t.Fatalf("selection JSON = %s, want %s", payload, tc.want)
		}
	}
}
