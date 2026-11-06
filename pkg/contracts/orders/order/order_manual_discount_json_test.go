package order_test

import (
	"encoding/json"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
	sales "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/orders/order"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/pricing/quote"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/pricing/quote/quote_enums"
)

func TestOrderJSONFreezesCartManualDiscountEvidence(t *testing.T) {
	raw, err := os.ReadFile("testdata/order_manual_discount.json")
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}
	var decoded sales.Order
	if err := json.Unmarshal(raw, &decoded); err != nil {
		t.Fatalf("unmarshal fixture: %v", err)
	}
	want := &quote.CartManualDiscountEvidence{
		Mode:             quote_enums.ManualDiscountModeSetTotal,
		TargetTotal:      &money.Money{AmountMinor: 4000, Currency: "AUD"},
		PreDiscountTotal: money.Money{AmountMinor: 4500, Currency: "AUD"},
		AppliedAmount:    money.Money{AmountMinor: 500, Currency: "AUD"},
		AppliedTax:       money.Money{AmountMinor: 45, Currency: "AUD"},
		CouponCode:       "SPRING5",
		CouponAmount:     &money.Money{AmountMinor: 500, Currency: "AUD"},
		ReasonCode:       "price_match",
		ActorUserID:      "cashier_1",
		AppliedAt:        time.Date(2026, 11, 6, 2, 3, 4, 0, time.UTC),
	}
	if !reflect.DeepEqual(decoded.ManualDiscount, want) {
		t.Fatalf("manual discount decoded to %+v, want %+v", decoded.ManualDiscount, want)
	}
	encoded, err := json.Marshal(decoded)
	if err != nil {
		t.Fatal(err)
	}
	wantJSON := `"discount_amount":{"amount_minor":500,"currency":"AUD"},"manual_discount":{"mode":"set_total","target_total":{"amount_minor":4000,"currency":"AUD"},"pre_discount_total":{"amount_minor":4500,"currency":"AUD"},"applied_amount":{"amount_minor":500,"currency":"AUD"},"applied_tax":{"amount_minor":45,"currency":"AUD"},"coupon_code":"SPRING5","coupon_amount":{"amount_minor":500,"currency":"AUD"},"reason_code":"price_match","actor_user_id":"cashier_1","applied_at":"2026-11-06T02:03:04Z"},"shipping_amount"`
	if !strings.Contains(string(encoded), wantJSON) {
		t.Fatalf("order JSON = %s, want %s", encoded, wantJSON)
	}
}

func TestOrderJSONOmitsAbsentManualDiscount(t *testing.T) {
	raw, err := json.Marshal(sales.Order{OrderNumber: "SO-1"})
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(raw), `"manual_discount"`) {
		t.Fatalf("absent manual discount must be omitted: %s", raw)
	}
}
