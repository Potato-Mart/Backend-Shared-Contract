package order_test

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/orders/order"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/orders/order/order_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/payments/payment/payment_enums"
)

func aud(minor int64) money.Money { return money.Money{AmountMinor: minor, Currency: "AUD"} }

// A sale of 100.03 paid 60.00 by EFTPOS, then 40.03 in cash rounded to 40.05.
func splitTenderOrder() order.Order {
	frozenAt := time.Date(2026, 11, 15, 1, 2, 3, 0, time.UTC)
	adjustment := aud(2)
	return order.Order{
		OrderNumber:   "POS-1",
		Total:         aud(10003),
		PaymentMethod: payment_enums.PaymentMethodEFTPOS,
		TenderPlan: &order.TenderPlanSnapshot{
			PlanID:   "plan-1",
			Revision: 2,
			FrozenAt: &frozenAt,
			Parts: []order.TenderPlanPart{{
				PartID: "part-1", TenderKey: "tender-key-1", Sequence: 1,
				Method: payment_enums.PaymentMethodEFTPOS,
				Amount: aud(6000), AmountDue: aud(6000),
				Status: order_enums.TenderPartStatusCaptured, PaymentID: "pay-1",
			}, {
				PartID: "part-2", TenderKey: "tender-key-2", Sequence: 2,
				Method: payment_enums.PaymentMethodCash,
				Amount: aud(4003), AmountDue: aud(4005), RoundingAdjustment: &adjustment,
				Status: order_enums.TenderPartStatusCaptured, PaymentID: "pay-2",
				CashTender: &order.OfflineCashTenderEvidence{AmountDue: aud(4005), AmountReceived: aud(5000), Change: aud(995)},
			}},
		},
	}
}

func TestOrderTenderPlanRoundTrip(t *testing.T) {
	want := splitTenderOrder()
	encoded, err := json.Marshal(want)
	if err != nil {
		t.Fatal(err)
	}
	for _, fragment := range []string{
		`"tender_plan":{"plan_id":"plan-1","revision":2,"parts":[`,
		`{"part_id":"part-1","tender_key":"tender-key-1","sequence":1,"method":"eftpos","amount":{"amount_minor":6000,"currency":"AUD"},"amount_due":{"amount_minor":6000,"currency":"AUD"},"status":"captured","payment_id":"pay-1"}`,
		`{"part_id":"part-2","tender_key":"tender-key-2","sequence":2,"method":"cash","amount":{"amount_minor":4003,"currency":"AUD"},"amount_due":{"amount_minor":4005,"currency":"AUD"},"rounding_adjustment":{"amount_minor":2,"currency":"AUD"},"status":"captured","payment_id":"pay-2","cash_tender":{"amount_due":{"amount_minor":4005,"currency":"AUD"},"amount_received":{"amount_minor":5000,"currency":"AUD"},"change":{"amount_minor":995,"currency":"AUD"}}}`,
		`],"frozen_at":"2026-11-15T01:02:03Z"}`,
		`"payment_method":"eftpos"`,
	} {
		if !strings.Contains(string(encoded), fragment) {
			t.Fatalf("tender plan wire shape changed: missing %s in %s", fragment, encoded)
		}
	}
	var got order.Order
	if err := json.Unmarshal(encoded, &got); err != nil || !reflect.DeepEqual(got.TenderPlan, want.TenderPlan) {
		t.Fatalf("tender plan round trip: %+v %v", got.TenderPlan, err)
	}
	assertTenderPlanInvariants(t, got)
}

// A plan still being revised omits frozen_at, and a pending part omits its
// payment, rounding and cash evidence.
func TestOrderTenderPlanOptionalParts(t *testing.T) {
	encoded, err := json.Marshal(order.TenderPlanSnapshot{PlanID: "plan-1", Revision: 1, Parts: []order.TenderPlanPart{{
		PartID: "part-1", TenderKey: "tender-key-1", Sequence: 1, Method: payment_enums.PaymentMethodEFTPOS,
		Amount: aud(500), AmountDue: aud(500), Status: order_enums.TenderPartStatusPending,
	}}})
	if err != nil {
		t.Fatal(err)
	}
	want := `{"plan_id":"plan-1","revision":1,"parts":[{"part_id":"part-1","tender_key":"tender-key-1","sequence":1,"method":"eftpos","amount":{"amount_minor":500,"currency":"AUD"},"amount_due":{"amount_minor":500,"currency":"AUD"},"status":"pending"}]}`
	if string(encoded) != want {
		t.Fatalf("pending tender plan JSON = %s, want %s", encoded, want)
	}
}

func TestOrderTenderPlanIsOptional(t *testing.T) {
	encoded, err := json.Marshal(order.Order{OrderNumber: "SO-1"})
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(encoded), `"tender_plan"`) {
		t.Fatalf("an order without a split tender must omit tender_plan: %s", encoded)
	}
	var decoded order.Order
	if err := json.Unmarshal([]byte(`{"order_number":"SO-1","payment_method":"card"}`), &decoded); err != nil || decoded.TenderPlan != nil {
		t.Fatalf("legacy order gained a tender plan: %+v %v", decoded.TenderPlan, err)
	}
}

// The documented plan invariants hold for the reference split: one currency,
// sequences 1..n used once, parts summing to the total less gift-card and
// points redemptions, and rounding only on the one cash part.
func assertTenderPlanInvariants(t *testing.T, o order.Order) {
	t.Helper()
	plan := o.TenderPlan
	remainder := o.Total.AmountMinor
	for _, redemption := range o.GiftCardRedemptions {
		remainder -= redemption.AppliedAmount.AmountMinor
	}
	if o.PointRedemption != nil {
		remainder -= o.PointRedemption.DiscountAmount.AmountMinor
	}
	seen := make(map[int]bool, len(plan.Parts))
	var sum int64
	rounded := 0
	for _, part := range plan.Parts {
		if part.Sequence < 1 || part.Sequence > len(plan.Parts) || seen[part.Sequence] {
			t.Fatalf("sequence %d is not a unique position in 1..%d", part.Sequence, len(plan.Parts))
		}
		seen[part.Sequence] = true
		for _, amount := range []money.Money{part.Amount, part.AmountDue} {
			if amount.Currency != o.Total.Currency {
				t.Fatalf("part %s mixes currency %s with %s", part.PartID, amount.Currency, o.Total.Currency)
			}
		}
		if !part.Status.IsValid() || !part.Method.IsValid() {
			t.Fatalf("part %s has an unknown status or method", part.PartID)
		}
		due := part.Amount.AmountMinor
		if part.RoundingAdjustment != nil {
			rounded++
			if part.Method != payment_enums.PaymentMethodCash || part.RoundingAdjustment.Currency != o.Total.Currency {
				t.Fatalf("only a cash part in the order currency may round: %+v", part)
			}
			due += part.RoundingAdjustment.AmountMinor
		}
		if part.AmountDue.AmountMinor != due {
			t.Fatalf("part %s amount due %d, want %d", part.PartID, part.AmountDue.AmountMinor, due)
		}
		if part.CashTender != nil && part.CashTender.AmountDue != part.AmountDue {
			t.Fatalf("part %s cash evidence disagrees with its amount due", part.PartID)
		}
		sum += part.Amount.AmountMinor
	}
	if sum != remainder {
		t.Fatalf("parts sum to %d, want %d", sum, remainder)
	}
	if rounded > 1 {
		t.Fatalf("%d parts carry a rounding adjustment, want at most one", rounded)
	}
}
