package payments_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v17/pkg/contracts/payments"
)

func TestWalletPaymentReferenceRoundTrip(t *testing.T) {
	reference := payments.PaymentReference{
		Wallet: &payments.WalletPaymentReference{
			GiftCardCode:        "GC-1",
			WalletTransactionID: "wallet_tx_1",
		},
	}

	payload, err := json.Marshal(reference)
	if err != nil {
		t.Fatalf("marshal wallet payment reference: %v", err)
	}
	if !strings.Contains(string(payload), `"wallet":{"gift_card_code":"GC-1","wallet_transaction_id":"wallet_tx_1"}`) {
		t.Fatalf("wallet payment reference JSON is incomplete: %s", payload)
	}

	var decoded payments.PaymentReference
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal wallet payment reference: %v", err)
	}
	if decoded.Wallet == nil || decoded.Wallet.WalletTransactionID != "wallet_tx_1" {
		t.Fatalf("wallet payment reference did not round-trip: %+v", decoded)
	}
}
