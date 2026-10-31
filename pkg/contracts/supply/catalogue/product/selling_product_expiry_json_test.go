package product

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/warehouse"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/warehouse/warehouse_enums"
)

func TestSellingProductExpiryDateIsOmittedWhenAbsent(t *testing.T) {
	body, err := json.Marshal(SellingProduct{SKUCode: "A00001"})
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(body), `"expiry_date"`) {
		t.Fatalf("SellingProduct JSON without a date mark = %s, want expiry_date omitted", body)
	}

	legacy := `{"sku_code":"A00001","sku_series_code":"A0","storage_type":"AMBIENT","content":{"name":{"language":"en","name":"Product"}},"classification":{},"package_options":[],"price":{"unit_price":{"amount_minor":319,"currency":"AUD"},"currency_exponent":{"currency":"AUD","exponent":2},"market_code":"AU","channel":"online","audience":"retail","price_visibility":"visible","tax_inclusion":"tax_inclusive","valid_from":"2026-08-24T00:00:00Z","as_of":"2026-08-24T00:00:00Z"}}`
	var decoded SellingProduct
	if err := json.Unmarshal([]byte(legacy), &decoded); err != nil {
		t.Fatalf("decode legacy SellingProduct: %v", err)
	}
	if decoded.SKUCode != "A00001" || decoded.Price.MarketCode != "AU" || decoded.ExpiryDate != nil {
		t.Fatalf("legacy SellingProduct decoded as %+v, want SKU and price kept and no expiry date", decoded)
	}
}

func TestSellingProductExpiryDateRoundTripsDateMarkKinds(t *testing.T) {
	sydney, err := time.LoadLocation("Australia/Sydney")
	if err != nil {
		t.Skipf("tzdata unavailable: %v", err)
	}
	for _, kind := range []warehouse_enums.InventoryDateMarkKind{
		warehouse_enums.InventoryDateMarkBestBefore,
		warehouse_enums.InventoryDateMarkExpiry,
	} {
		t.Run(kind.String(), func(t *testing.T) {
			value := SellingProduct{
				SKUCode: "A00001",
				ExpiryDate: &warehouse.InventoryDateMark{
					Kind:       kind,
					DateMarkAt: time.Date(2026, 12, 15, 0, 0, 0, 0, sydney),
					Timezone:   "Australia/Sydney",
				},
			}
			body, err := json.Marshal(value)
			if err != nil {
				t.Fatal(err)
			}
			want := `"expiry_date":{"kind":"` + kind.String() + `","date_mark_at":"2026-12-15T00:00:00+11:00","timezone":"Australia/Sydney"}`
			if !strings.Contains(string(body), want) {
				t.Fatalf("SellingProduct JSON = %s, want %s", body, want)
			}

			var decoded SellingProduct
			if err := json.Unmarshal(body, &decoded); err != nil {
				t.Fatal(err)
			}
			if decoded.ExpiryDate == nil || !decoded.ExpiryDate.Kind.IsValid() || decoded.ExpiryDate.Kind != kind ||
				!decoded.ExpiryDate.DateMarkAt.Equal(value.ExpiryDate.DateMarkAt) || decoded.ExpiryDate.Timezone != "Australia/Sydney" {
				t.Fatalf("decoded expiry date = %+v, want %+v", decoded.ExpiryDate, value.ExpiryDate)
			}
		})
	}
}

func TestSellingProductExpiryDateKindRejectsUnknownValues(t *testing.T) {
	var decoded SellingProduct
	body := `{"sku_code":"A00001","expiry_date":{"kind":"USE_BY","date_mark_at":"2026-12-15T00:00:00+11:00","timezone":"Australia/Sydney"}}`
	if err := json.Unmarshal([]byte(body), &decoded); err != nil {
		t.Fatal(err)
	}
	if decoded.ExpiryDate == nil || decoded.ExpiryDate.Kind.IsValid() {
		t.Fatalf("expiry date kind %q must be invalid; use EXPIRY for use-by dates", decoded.ExpiryDate.Kind)
	}
}
