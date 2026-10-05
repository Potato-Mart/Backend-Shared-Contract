package giftcard_test

import (
	"encoding/json"
	"github.com/Potato-Mart/Backend-Shared-Contract/v43/pkg/contracts/pricing/wallet/giftcard"
	"reflect"
	"testing"
)

func TestGiftCardBindingAndSecretBoundary(t *testing.T) {
	card := giftcard.GiftCard{ID: "stable-id", Code: "GC000000000001", CustomerNumber: "RC-1"}
	payload, err := json.Marshal(card)
	if err != nil {
		t.Fatal(err)
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(payload, &fields); err != nil {
		t.Fatal(err)
	}
	if string(fields["code"]) != `"GC000000000001"` || string(fields["customer_number"]) != `"RC-1"` {
		t.Fatalf("binding lost: %s", payload)
	}
	for _, key := range []string{"pin", "pin_hash", "claim_code"} {
		if _, ok := fields[key]; ok {
			t.Fatalf("unexpected secret or duplicate code: %s", key)
		}
	}
	var decoded giftcard.GiftCard
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(decoded, card) {
		t.Fatalf("round trip changed card: %#v", decoded)
	}
	empty, err := json.Marshal(giftcard.GiftCard{ID: "stable-id", Code: "GC000000000001"})
	if err != nil {
		t.Fatal(err)
	}
	var unbound map[string]json.RawMessage
	if err := json.Unmarshal(empty, &unbound); err != nil {
		t.Fatal(err)
	}
	if _, ok := unbound["customer_number"]; ok {
		t.Fatal("unbound card emitted customer number")
	}
}
