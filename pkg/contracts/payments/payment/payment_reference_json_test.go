package payment_test

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/party"
	payment "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/payments/payment"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/payments/payment/payment_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/payments/provider"
)

func TestWalletPaymentReferenceRoundTrip(t *testing.T) {
	reference := payment.PaymentReference{
		Wallet: &provider.WalletPaymentReference{
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

func TestBankTransferPaymentReferenceRoundTrip(t *testing.T) {
	confirmedAt := time.Date(2026, 11, 3, 1, 2, 3, 0, time.UTC)
	capturedAt := confirmedAt.Add(-time.Minute)
	reference := payment.PaymentReference{
		BankTransfer: &provider.BankTransferPaymentReference{
			PayerReference: "INV-42",
			Customer: &provider.BankTransferCustomer{
				PartyRef:       party.PartyRef{ID: "cust_1", Name: "Ada Lovelace", Phone: "+61400000000"},
				CustomerNumber: "RC000123",
			},
			TransferAmount:     money.Money{AmountMinor: 12345, Currency: "AUD"},
			ConfirmationStatus: payment_enums.BankTransferConfirmationStatusConfirmedPartial,
			Evidence: []provider.BankTransferEvidence{{
				Code:       "pev_1",
				MimeType:   "image/jpeg",
				SHA256:     "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
				CapturedAt: capturedAt,
				CapturedBy: "user_1",
			}},
			ConfirmedAt: confirmedAt,
			ConfirmedBy: "user_1",
		},
	}

	payload, err := json.Marshal(reference)
	if err != nil {
		t.Fatalf("marshal bank-transfer payment reference: %v", err)
	}
	text := string(payload)
	for _, want := range []string{
		`"bank_transfer":{`,
		`"payer_reference":"INV-42"`,
		`"customer":{"id":"cust_1","name":"Ada Lovelace","phone":"+61400000000","customer_number":"RC000123"}`,
		`"transfer_amount":{"amount_minor":12345,"currency":"AUD"}`,
		`"confirmation_status":"confirmed_partial"`,
		`"evidence":[{"code":"pev_1","mime_type":"image/jpeg","sha256":"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08","captured_at":"2026-11-03T01:01:03Z","captured_by":"user_1"}]`,
		`"confirmed_at":"2026-11-03T01:02:03Z"`,
		`"confirmed_by":"user_1"`,
	} {
		if !strings.Contains(text, want) {
			t.Fatalf("bank-transfer payment reference JSON missing %s: %s", want, text)
		}
	}
	if strings.Contains(text, "media_code") {
		t.Fatalf("bank-transfer evidence must use code, not media_code: %s", text)
	}
	if strings.Contains(text, `"stripe"`) || strings.Contains(text, `"mx51"`) || strings.Contains(text, `"wallet"`) {
		t.Fatalf("absent provider references must be omitted: %s", text)
	}

	var decoded payment.PaymentReference
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal bank-transfer payment reference: %v", err)
	}
	if !reflect.DeepEqual(decoded, reference) {
		t.Fatalf("bank-transfer payment reference did not round-trip:\n got %+v\nwant %+v", decoded.BankTransfer, reference.BankTransfer)
	}
}

func TestPaymentReferenceOmitsAbsentBankTransfer(t *testing.T) {
	payload, err := json.Marshal(payment.PaymentReference{Stripe: &provider.StripePaymentReference{PaymentIntentID: "pi_1"}})
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(payload), "bank_transfer") {
		t.Fatalf("absent bank transfer must be omitted: %s", payload)
	}
}
