package orders_test

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/payments/payment/payment_enums"
	events "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/pubsub/orders"
)

// Method stays the primary external tender; Tenders lists every captured
// part of the split in sequence order.
func TestOrderPaidEventListsSplitTenders(t *testing.T) {
	want := events.OrderPaidEvent{
		OrderID: "order-1", OrderNumber: "POS-1", PaymentID: "pay-2",
		Method:     payment_enums.PaymentMethodEFTPOS,
		AmountPaid: money.Money{AmountMinor: 10003, Currency: "AUD"},
		PaidAt:     time.Date(2026, 11, 15, 1, 3, 0, 0, time.UTC),
		Tenders: []events.OrderPaidTender{
			{Method: payment_enums.PaymentMethodEFTPOS, Amount: money.Money{AmountMinor: 6000, Currency: "AUD"}},
			{Method: payment_enums.PaymentMethodCash, Amount: money.Money{AmountMinor: 4003, Currency: "AUD"}},
		},
	}
	encoded, err := json.Marshal(want)
	if err != nil {
		t.Fatal(err)
	}
	for _, fragment := range []string{
		`"method":"eftpos"`,
		`"tenders":[{"method":"eftpos","amount":{"amount_minor":6000,"currency":"AUD"}},{"method":"cash","amount":{"amount_minor":4003,"currency":"AUD"}}]`,
	} {
		if !strings.Contains(string(encoded), fragment) {
			t.Fatalf("order paid tenders wire shape changed: missing %s in %s", fragment, encoded)
		}
	}
	var got events.OrderPaidEvent
	if err := json.Unmarshal(encoded, &got); err != nil || !reflect.DeepEqual(got, want) {
		t.Fatalf("order paid round trip: %+v %v", got, err)
	}

	encoded, err = json.Marshal(events.OrderPaidEvent{OrderID: "order-2", OrderNumber: "SO-2", Method: payment_enums.PaymentMethodCard})
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(encoded), `"tenders"`) {
		t.Fatalf("an order that was not split must omit tenders: %s", encoded)
	}
}
