package receipt_test

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
	pos "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/payments/receipt"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/pricing/quote/quote_enums"
)

func TestReceiptSnapshotShowsOneCartManualDiscountWithoutActor(t *testing.T) {
	receipt := pos.ReceiptSnapshot{
		OrderNumber: "SO-1",
		ManualDiscount: &pos.ReceiptManualDiscount{
			Mode:       quote_enums.ManualDiscountModePercent,
			ReasonCode: "staff_discretion",
			Reason:     "Regular customer goodwill",
			Amount:     money.Money{AmountMinor: 450, Currency: "AUD"},
		},
	}
	payload, err := json.Marshal(receipt)
	if err != nil {
		t.Fatalf("marshal receipt: %v", err)
	}
	want := `"manual_discount":{"mode":"percent","reason_code":"staff_discretion","reason":"Regular customer goodwill","amount":{"amount_minor":450,"currency":"AUD"}}`
	if !strings.Contains(string(payload), want) {
		t.Fatalf("receipt JSON = %s, want %s", payload, want)
	}
	for _, private := range []string{"actor", "user_id", "cashier"} {
		if strings.Contains(string(payload), private) {
			t.Fatalf("receipt manual discount exposes staff identity %q: %s", private, payload)
		}
	}
	var decoded pos.ReceiptSnapshot
	if err := json.Unmarshal(payload, &decoded); err != nil || !reflect.DeepEqual(decoded.ManualDiscount, receipt.ManualDiscount) {
		t.Fatalf("receipt manual discount round trip: %+v %v", decoded.ManualDiscount, err)
	}
	model := reflect.TypeOf(pos.ReceiptManualDiscount{})
	for index := 0; index < model.NumField(); index++ {
		if name := model.Field(index).Name; strings.Contains(name, "Actor") || strings.Contains(name, "User") {
			t.Fatalf("ReceiptManualDiscount must not carry staff identity field %s", name)
		}
	}
}

func TestReceiptSnapshotOmitsAbsentManualDiscount(t *testing.T) {
	payload, err := json.Marshal(pos.ReceiptSnapshot{OrderNumber: "SO-1"})
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(payload), `"manual_discount"`) {
		t.Fatalf("absent manual discount must be omitted: %s", payload)
	}
}
