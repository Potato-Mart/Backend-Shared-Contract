package order_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
	sales "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/orders/order"
)

func TestGiftCardRedemptionSnapshotGiftCardIDIsOptional(t *testing.T) {
	byCode, err := json.Marshal(sales.GiftCardRedemptionSnapshot{
		GiftCardCode:  "GC000000000001",
		AppliedAmount: money.Money{AmountMinor: 2500, Currency: "AUD"},
	})
	if err != nil {
		t.Fatalf("marshal code redemption: %v", err)
	}
	if strings.Contains(string(byCode), "gift_card_id") {
		t.Fatalf("a redemption tendered by code must omit gift_card_id: %s", byCode)
	}

	payload := []byte(`{"gift_card_code":"GC********1234","gift_card_id":"gcw_01JEXAMPLE","applied_amount":{"amount_minor":1500,"currency":"AUD"},"wallet_transaction_id":"wallet_tx_1"}`)
	var byID sales.GiftCardRedemptionSnapshot
	if err := json.Unmarshal(payload, &byID); err != nil {
		t.Fatalf("unmarshal wallet-id redemption: %v", err)
	}
	if byID.GiftCardID != "gcw_01JEXAMPLE" || byID.GiftCardCode != "GC********1234" {
		t.Fatalf("wallet-id redemption decoded wrongly: %+v", byID)
	}
	round, err := json.Marshal(byID)
	if err != nil {
		t.Fatalf("marshal wallet-id redemption: %v", err)
	}
	if !strings.Contains(string(round), `"gift_card_id":"gcw_01JEXAMPLE"`) {
		t.Fatalf("gift_card_id did not round-trip: %s", round)
	}
}
