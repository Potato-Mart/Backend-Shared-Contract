package payment_test

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/payments/payment"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/payments/payment/payment_enums"
)

func TestPaymentCarriesSplitTenderPart(t *testing.T) {
	paidAt := time.Date(2026, 11, 15, 1, 2, 3, 0, time.UTC)
	want := payment.Payment{
		ID: "pay-2", OrderNumber: "POS-1",
		Amount: money.Money{AmountMinor: 4003, Currency: "AUD"},
		Method: payment_enums.PaymentMethodCash, Status: payment_enums.PaymentRecordStatusCompleted,
		TenderPlanID: "plan-1", TenderPartID: "part-2", TenderKey: "tender-key-2", TenderSequence: 2,
		PaidAt: &paidAt,
	}
	encoded, err := json.Marshal(want)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(encoded), `"tender_plan_id":"plan-1","tender_part_id":"part-2","tender_key":"tender-key-2","tender_sequence":2`) {
		t.Fatalf("payment tender wire shape changed: %s", encoded)
	}
	var got payment.Payment
	if err := json.Unmarshal(encoded, &got); err != nil || !reflect.DeepEqual(got, want) {
		t.Fatalf("payment tender round trip: %+v %v", got, err)
	}
}

func TestPaymentOutsideSplitTenderOmitsTenderFields(t *testing.T) {
	encoded, err := json.Marshal(payment.Payment{ID: "pay-1", OrderNumber: "SO-1", Method: payment_enums.PaymentMethodCard})
	if err != nil {
		t.Fatal(err)
	}
	for _, key := range []string{"tender_plan_id", "tender_part_id", "tender_key", "tender_sequence"} {
		if strings.Contains(string(encoded), `"`+key+`"`) {
			t.Fatalf("a payment outside a split tender must omit %s: %s", key, encoded)
		}
	}
}

func TestCustomerPaymentAllocationNamesTenderPart(t *testing.T) {
	occurredAt := time.Date(2026, 11, 15, 1, 2, 3, 0, time.UTC)
	want := payment.CustomerPaymentAllocation{
		AllocationID: "pay-1", Kind: payment_enums.CustomerPaymentAllocationKindExternalPayment,
		Method: payment_enums.PaymentMethodEFTPOS, Provider: "TYRO",
		Amount: money.Money{AmountMinor: 6000, Currency: "AUD"}, Status: "completed",
		OccurredAt: occurredAt, TenderPartID: "part-1",
	}
	encoded, err := json.Marshal(want)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(encoded), `"occurred_at":"2026-11-15T01:02:03Z","tender_part_id":"part-1"}`) {
		t.Fatalf("allocation tender part wire shape changed: %s", encoded)
	}
	var got payment.CustomerPaymentAllocation
	if err := json.Unmarshal(encoded, &got); err != nil || !reflect.DeepEqual(got, want) {
		t.Fatalf("allocation round trip: %+v %v", got, err)
	}

	want.TenderPartID = ""
	encoded, err = json.Marshal(want)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(encoded), "tender_part_id") {
		t.Fatalf("an allocation outside a split tender must omit tender_part_id: %s", encoded)
	}
}

// A known zero remaining amount (the order is paid) stays distinct from an
// absent one (Payments did not compute it).
func TestCustomerPaymentSummaryAmountRemaining(t *testing.T) {
	for name, remaining := range map[string]*money.Money{
		"outstanding": {AmountMinor: 4003, Currency: "AUD"},
		"paid":        {AmountMinor: 0, Currency: "AUD"},
		"absent":      nil,
	} {
		t.Run(name, func(t *testing.T) {
			want := payment.CustomerPaymentSummary{
				OrderNumber: "POS-1", Status: payment_enums.PaymentStatusPartiallyPaid,
				AmountRemaining: remaining, Completeness: payment_enums.PaymentCompletenessComplete,
			}
			encoded, err := json.Marshal(want)
			if err != nil {
				t.Fatal(err)
			}
			var fields map[string]json.RawMessage
			if err := json.Unmarshal(encoded, &fields); err != nil {
				t.Fatal(err)
			}
			raw, present := fields["amount_remaining"]
			if present != (remaining != nil) {
				t.Fatalf("amount_remaining presence = %t: %s", present, encoded)
			}
			if remaining != nil {
				var got money.Money
				if err := json.Unmarshal(raw, &got); err != nil || got != *remaining {
					t.Fatalf("amount_remaining = %s, want %+v", raw, *remaining)
				}
			}
			var decoded payment.CustomerPaymentSummary
			if err := json.Unmarshal(encoded, &decoded); err != nil || !reflect.DeepEqual(decoded, want) {
				t.Fatalf("summary round trip: %+v %v", decoded, err)
			}
		})
	}
}
