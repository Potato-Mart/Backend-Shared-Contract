package order

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v48/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v48/pkg/contracts/payments/merchant"
	"github.com/Potato-Mart/Backend-Shared-Contract/v48/pkg/contracts/pricing/pricebook"
)

func TestOfflineJournalPreservesActualSaleAndIssuerEvidence(t *testing.T) {
	issued := time.Date(2026, 10, 1, 0, 0, 0, 0, time.UTC)
	sold := issued.Add(48 * time.Hour) // A sale may use a book beyond its stated expiry.
	want := OfflineSaleJournalEvidence{
		ClientSaleID: "sale-1", DeviceKey: "device-1", ReceiptNumber: "TILL-18446744073709551615", ReceiptPrefix: "TILL", ReceiptCounter: "18446744073709551615",
		MarketCode: "AU", DepotCode: "D", RegisterID: "R", SessionID: "session", BusinessDate: "2026-10-03",
		CashierUserID: "staff-1", CredentialRevision: 3, SoldAt: sold, OfflineSince: issued,
		Sequence: 125, JournalSchema: 1, RecordedAt: sold, PreviousHash: "previous", EntryHash: "entry", CanonicalEntryBase64: "e30=",
		PriceBook:    pricebook.OfflinePriceBookEvidence{Version: "book-1", MarketCode: "AU", StoreCode: "D", PayloadSHA256: "digest", Algorithm: "ed25519", KeyID: "key-1", Signature: "c2ln", IssuedAt: issued, ExpiresAt: issued.Add(24 * time.Hour)},
		CachedIssuer: merchant.MerchantLegalSnapshot{ProfileID: "issuer", ProfileRevision: 7, MarketCode: "AU", LegalName: "Cached merchant", BusinessNumber: "123", FrozenAt: issued},
		CashTender:   OfflineCashTenderEvidence{AmountDue: money.Money{AmountMinor: 500, Currency: "AUD"}, AmountReceived: money.Money{AmountMinor: 1000, Currency: "AUD"}, Change: money.Money{AmountMinor: 500, Currency: "AUD"}},
	}
	raw, err := json.Marshal(want)
	if err != nil {
		t.Fatal(err)
	}
	for _, field := range []string{`"receipt_prefix":"TILL"`, `"receipt_counter":"18446744073709551615"`, `"entry_hash":"entry"`, `"canonical_entry_base64":"e30="`, `"payload_sha256":"digest"`, `"store_code":"D"`, `"legal_name":"Cached merchant"`, `"cash_tender"`} {
		if !strings.Contains(string(raw), field) {
			t.Fatalf("missing %s in %s", field, raw)
		}
	}
	var got OfflineSaleJournalEvidence
	if err := json.Unmarshal(raw, &got); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("offline evidence changed on round trip: %#v", got)
	}
}

func TestOfflineReceiptCounterPreservesFullUnsignedRange(t *testing.T) {
	for _, counter := range []string{"1", "9007199254740993", "18446744073709551615"} {
		raw, err := json.Marshal(OfflineSaleJournalEvidence{ReceiptCounter: counter})
		if err != nil {
			t.Fatal(err)
		}
		var got OfflineSaleJournalEvidence
		if err := json.Unmarshal(raw, &got); err != nil {
			t.Fatal(err)
		}
		if got.ReceiptCounter != counter {
			t.Fatalf("counter %q became %q", counter, got.ReceiptCounter)
		}
	}
}
