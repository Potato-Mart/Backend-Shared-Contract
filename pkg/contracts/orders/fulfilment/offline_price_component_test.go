package fulfilment_test

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v54/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v54/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v54/pkg/contracts/common/packaging/packaging_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v54/pkg/contracts/orders/fulfilment"
	"github.com/Potato-Mart/Backend-Shared-Contract/v54/pkg/contracts/pricing/pricebook"
	"github.com/Potato-Mart/Backend-Shared-Contract/v54/pkg/contracts/pricing/pricebook/pricebook_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v54/pkg/contracts/pricing/quote"
	"github.com/Potato-Mart/Backend-Shared-Contract/v54/pkg/contracts/pricing/quote/quote_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v54/pkg/contracts/supply/catalogue/listing"
)

func TestOfflinePricedComponentHasNoInventedOnlineQuote(t *testing.T) {
	at := time.Date(2026, 10, 29, 1, 2, 3, 0, time.UTC)
	amount := func(n int64) money.Money { return money.Money{AmountMinor: n, Currency: "AUD"} }
	want := fulfilment.PricedPackageComponent{
		OfflinePriceSnapshot: &quote.OfflinePriceSnapshot{
			LineID: "device-sale-line-1", LineNumber: 1,
			Merchandise: listing.OfflinePOSMerchandiseEvidence{
				ClientSaleID: "sale-1", SKUCode: "SKU", MarketCode: "AU", DepotCode: "D",
				PackageOption: packaging.PackageOptionRef{SKUCode: "SKU", Code: "CASE6", Version: 1},
				HandlingUnit:  packaging_enums.PackageHandlingUnitCase, PackageCount: 2, UnitsPerPackage: 6, BaseUnits: 12,
				PriceBookVersion: "book-1", PriceBookSHA256: "book-digest", JournalEntryHash: "entry-hash", SoldAt: at,
			},
			PriceBook:        pricebook.OfflinePriceBookEvidence{Version: "book-1", MarketCode: "AU", StoreCode: "D", PayloadSHA256: "book-digest", KeyID: "key-1", IssuedAt: at.Add(-48 * time.Hour), ExpiresAt: at.Add(-24 * time.Hour)},
			CurrencyExponent: money.CurrencyExponent{Currency: "AUD", Exponent: 2},
			SourcePrice:      amount(100), SourcePriceBasis: quote_enums.PriceAmountBasisBaseUnit,
			ChargedPrice: amount(550), ChargedPriceBasis: quote_enums.PriceAmountBasisPackage,
			Tax: quote.OfflineTaxSnapshot{TaxClass: "GST", RateBasisPoints: 1000,
				SourceTaxInclusion: pricebook_enums.PriceTaxInclusionExclusive, ChargedTaxInclusion: pricebook_enums.PriceTaxInclusionInclusive,
				RoundingMethod: quote_enums.TaxRoundingMethodSumExactThenRound, TaxableBase: amount(1000), AllocatedTax: amount(100)},
			CustomOverride: &quote.CustomPriceOverrideEvidence{ActorUserID: "staff", Reason: "manager approved", SourceApprovedPrice: amount(100), SourcePriceBasis: quote_enums.PriceAmountBasisBaseUnit, SourceTaxInclusion: pricebook_enums.PriceTaxInclusionExclusive, OverrideGrossAmount: amount(550), OverridePriceBasis: quote_enums.PriceAmountBasisPackage, OverriddenAt: at},
			CapturedAt:     at,
		},
		RequestedPackageCount: 2, RequestedBaseUnits: 12, PackagePrice: amount(550), TaxAmount: amount(100), DiscountAmount: amount(0), ComponentTotal: amount(1100),
	}
	raw, err := json.Marshal(want)
	if err != nil {
		t.Fatal(err)
	}
	var shape map[string]json.RawMessage
	if err := json.Unmarshal(raw, &shape); err != nil {
		t.Fatal(err)
	}
	if _, present := shape["price_snapshot"]; present {
		t.Fatalf("offline line contains online price snapshot: %s", raw)
	}
	var offline map[string]json.RawMessage
	if err := json.Unmarshal(shape["offline_price_snapshot"], &offline); err != nil {
		t.Fatal(err)
	}
	for _, forbidden := range []string{"quote_id", "fingerprint", "price_entry_id", "price_entry_revision", "tax_rule_id", "tax_rule_revision"} {
		if _, present := offline[forbidden]; present {
			t.Fatalf("offline line contains online identity %s: %s", forbidden, raw)
		}
	}
	var tax map[string]json.RawMessage
	if err := json.Unmarshal(offline["tax"], &tax); err != nil {
		t.Fatal(err)
	}
	for _, forbidden := range []string{"tax_rule_id", "tax_rule_revision"} {
		if _, present := tax[forbidden]; present {
			t.Fatalf("offline tax contains online identity %s: %s", forbidden, raw)
		}
	}
	var got fulfilment.PricedPackageComponent
	if err := json.Unmarshal(raw, &got); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("offline component changed: %#v", got)
	}
}

func TestOnlinePricedComponentHasNoOfflineSnapshot(t *testing.T) {
	want := fulfilment.PricedPackageComponent{PriceSnapshot: &quote.PriceSnapshot{QuoteID: "quote-1", LineID: "line-1"}}
	raw, err := json.Marshal(want)
	if err != nil {
		t.Fatal(err)
	}
	var shape map[string]json.RawMessage
	if err := json.Unmarshal(raw, &shape); err != nil {
		t.Fatal(err)
	}
	if shape["price_snapshot"] == nil || shape["offline_price_snapshot"] != nil {
		t.Fatalf("online line has incorrect price variant: %s", raw)
	}
}
