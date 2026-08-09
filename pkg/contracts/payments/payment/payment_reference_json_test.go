package payment_test

import (
	"encoding/json"
	"strings"
	"testing"

	payment "github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/payments/payment"
)

func TestWalletPaymentReferenceRoundTrip(t *testing.T) {
	reference := payment.PaymentReference{
		Wallet: &payment.WalletPaymentReference{
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

	var decoded payment.PaymentReference
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal wallet payment reference: %v", err)
	}
	if decoded.Wallet == nil || decoded.Wallet.WalletTransactionID != "wallet_tx_1" {
		t.Fatalf("wallet payment reference did not round-trip: %+v", decoded)
	}
}
