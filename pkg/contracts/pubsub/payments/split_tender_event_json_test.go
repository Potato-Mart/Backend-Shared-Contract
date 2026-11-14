package payments_test

import (
	"encoding/json"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/payments/payment"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/payments/payment/payment_enums"
	event "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/pubsub/payments"
)

func audMoney(minor int64) money.Money { return money.Money{AmountMinor: minor, Currency: "AUD"} }

func TestPaymentPartRecordedEventFixtureRoundTrip(t *testing.T) {
	raw, err := os.ReadFile("testdata/payment_part_recorded.json")
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}
	var decoded event.PaymentPartRecordedEvent
	if err := json.Unmarshal(raw, &decoded); err != nil {
		t.Fatalf("unmarshal fixture: %v", err)
	}
	want := event.PaymentPartRecordedEvent{
		EventID: "payment_part_recorded:pay-1", OrderNumber: "POS-1", PaymentID: "pay-1",
		Method: payment_enums.PaymentMethodEFTPOS, Provider: "TYRO", Amount: audMoney(6000),
		TenderKey: "tender-key-1", TenderSequence: 1,
		OrderSettled: false, AmountRemaining: audMoney(4003),
		OccurredAt: time.Date(2026, 11, 15, 1, 2, 3, 0, time.UTC),
	}
	if !reflect.DeepEqual(decoded, want) {
		t.Fatalf("fixture decoded to %+v, want %+v", decoded, want)
	}
	encoded, err := json.Marshal(decoded)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(encoded, &fields); err != nil {
		t.Fatal(err)
	}
	names := []string{"event_id", "order_number", "payment_id", "method", "provider", "amount", "tender_key", "tender_sequence", "order_settled", "amount_remaining", "occurred_at"}
	for _, name := range names {
		if _, ok := fields[name]; !ok {
			t.Fatalf("payment part recorded JSON missing %q: %s", name, encoded)
		}
	}
	if len(fields) != len(names) {
		t.Fatalf("unexpected payment part recorded fields: %s", encoded)
	}
	var again event.PaymentPartRecordedEvent
	if err := json.Unmarshal(encoded, &again); err != nil || !reflect.DeepEqual(again, want) {
		t.Fatalf("round trip changed the event: %v %+v", err, again)
	}
}

// The settling part keeps order_settled and a known zero remaining amount on
// the wire; only the provider is optional.
func TestPaymentPartRecordedEventSettlingPartKeepsZeroRemaining(t *testing.T) {
	encoded, err := json.Marshal(event.PaymentPartRecordedEvent{
		EventID: "payment_part_recorded:pay-2", OrderNumber: "POS-1", PaymentID: "pay-2",
		Method: payment_enums.PaymentMethodCash, Amount: audMoney(4003),
		TenderKey: "tender-key-2", TenderSequence: 2,
		OrderSettled: true, AmountRemaining: audMoney(0),
		OccurredAt: time.Date(2026, 11, 15, 1, 3, 0, 0, time.UTC),
	})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(encoded), `"order_settled":true,"amount_remaining":{"amount_minor":0,"currency":"AUD"}`) {
		t.Fatalf("settling part lost its zero remaining amount: %s", encoded)
	}
	if strings.Contains(string(encoded), `"provider"`) {
		t.Fatalf("an absent provider must be omitted: %s", encoded)
	}
}

func TestPaymentCapturedEventsListSplitTenders(t *testing.T) {
	tenders := []event.CapturedTender{
		{PaymentID: "pay-1", Method: payment_enums.PaymentMethodEFTPOS, Provider: "TYRO", Amount: audMoney(6000)},
		{PaymentID: "pay-2", Method: payment_enums.PaymentMethodCash, Amount: audMoney(4003)},
	}
	const fragment = `"tenders":[{"payment_id":"pay-1","method":"eftpos","provider":"TYRO","amount":{"amount_minor":6000,"currency":"AUD"}},{"payment_id":"pay-2","method":"cash","amount":{"amount_minor":4003,"currency":"AUD"}}]`
	capturedAt := time.Date(2026, 11, 15, 1, 3, 0, 0, time.UTC)

	v1 := event.PaymentCapturedEvent{PaymentID: "pay-2", OrderNumber: "POS-1", Method: payment_enums.PaymentMethodEFTPOS, Amount: audMoney(10003), CapturedAt: capturedAt, Tenders: tenders}
	v2 := event.PaymentCapturedEventV2{PaymentID: "pay-2", OrderNumber: "POS-1", Method: payment_enums.PaymentMethodEFTPOS, Amount: audMoney(10003), CaptureTiming: payment.CaptureTimingEvidence{}, Tenders: tenders}
	for name, value := range map[string]any{"v1": &v1, "v2": &v2} {
		encoded, err := json.Marshal(value)
		if err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(string(encoded), fragment) {
			t.Fatalf("%s captured tenders wire shape changed: %s", name, encoded)
		}
		decoded := reflect.New(reflect.TypeOf(value).Elem()).Interface()
		if err := json.Unmarshal(encoded, decoded); err != nil || !reflect.DeepEqual(decoded, value) {
			t.Fatalf("%s captured event round trip: %+v %v", name, decoded, err)
		}
	}
	var sum int64
	for _, tender := range tenders {
		sum += tender.Amount.AmountMinor
	}
	if sum != v2.Amount.AmountMinor {
		t.Fatalf("tenders of a split without gift cards or points must sum to the order total: %d", sum)
	}

	for name, value := range map[string]any{
		"v1": event.PaymentCapturedEvent{PaymentID: "pay-1", OrderNumber: "SO-1", Amount: audMoney(100), CapturedAt: capturedAt},
		"v2": event.PaymentCapturedEventV2{PaymentID: "pay-1", OrderNumber: "SO-1", Amount: audMoney(100)},
	} {
		encoded, err := json.Marshal(value)
		if err != nil {
			t.Fatal(err)
		}
		if strings.Contains(string(encoded), `"tenders"`) {
			t.Fatalf("%s capture of an order that was not split must omit tenders: %s", name, encoded)
		}
	}
}

func TestRefundCompletedEventPaymentAllocations(t *testing.T) {
	completedAt := time.Date(2026, 11, 15, 2, 0, 0, 0, time.UTC)
	giftCardRefund := audMoney(1501)
	want := event.RefundCompletedEvent{
		RefundID: "refund-1", OrderNumber: "POS-2", Amount: audMoney(7501),
		GiftCardRefundAmount: &giftCardRefund, CompletedAt: completedAt,
		PaymentAllocations: []event.RefundPaymentAllocation{
			{PaymentID: "pay-1", Method: payment_enums.PaymentMethodEFTPOS, Amount: audMoney(4000)},
			{PaymentID: "pay-2", Method: payment_enums.PaymentMethodEFTPOS, Amount: audMoney(2000)},
		},
	}
	encoded, err := json.Marshal(want)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(encoded), `"payment_allocations":[{"payment_id":"pay-1","method":"eftpos","amount":{"amount_minor":4000,"currency":"AUD"}},{"payment_id":"pay-2","method":"eftpos","amount":{"amount_minor":2000,"currency":"AUD"}}]`) {
		t.Fatalf("refund payment allocation wire shape changed: %s", encoded)
	}
	var got event.RefundCompletedEvent
	if err := json.Unmarshal(encoded, &got); err != nil || !reflect.DeepEqual(got, want) {
		t.Fatalf("round trip changed the event: %v %+v", err, got)
	}
	var sum int64
	for _, allocation := range got.PaymentAllocations {
		sum += allocation.Amount.AmountMinor
	}
	if sum != got.Amount.AmountMinor-got.GiftCardRefundAmount.AmountMinor {
		t.Fatalf("allocations must sum to the refund less the gift-card restoration: %d", sum)
	}

	encoded, err = json.Marshal(event.RefundCompletedEvent{RefundID: "refund-2", OrderNumber: "SO-2", CompletedAt: completedAt})
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(encoded), "payment_allocations") {
		t.Fatalf("a refund without a recorded split must omit payment_allocations: %s", encoded)
	}
}
