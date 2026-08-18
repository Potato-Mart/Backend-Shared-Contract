package classification_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/localization"
	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/party"
	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/supply/classification"
)

func TestV29SupplierExpansionAndCodeOnlyManufacturing(t *testing.T) {
	supplier := classification.Supplier{
		OrganisationDetail:   party.OrganisationDetail{PartyRef: party.PartyRef{ID: "64c13ab08edf48a008793ca4", Code: "SUP0001", Name: "Supplier"}},
		GeographicLocation:   &geography.Address{Line1: "1 Supply Rd", Locality: "Sydney", PostalCode: "2000", Country: geography.CountryRef{Code: "AU"}},
		AvailableMarketCodes: []string{"AU"},
		AvailableProducts: []classification.SupplierAvailableProduct{{
			SKUCode: "A00001", SupplierSKUCode: "THEIR-1",
			ProductNames: []localization.LocalizedName{{Language: "en", Name: "Product"}},
			OfferedPrice: &money.Money{AmountMinor: 100, Currency: "AUD"}, MinimumPurchaseQuantity: 2,
			Manufacturing: &classification.ProductManufacturing{CountryRef: &classification.CountryCodeRef{Code: "TW"}},
		}},
		AvailablePromotions: []classification.SupplierAvailablePromotion{{Names: []localization.LocalizedName{{Language: "en", Name: "Launch"}}}},
	}
	body, err := json.Marshal(supplier)
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{`"available_market_codes":["AU"]`, `"sku_code":"A00001"`, `"country_ref":{"code":"TW"}`, `"available_promotions"`} {
		if !strings.Contains(string(body), want) {
			t.Fatalf("Supplier JSON = %s, want %s", body, want)
		}
	}
	if strings.Contains(string(body), `"location"`) || strings.Contains(string(body), `"supplier_id"`) {
		t.Fatalf("Supplier retained legacy reference: %s", body)
	}
}

func TestV29ProductSupplyUsesMultipleSupplierCodes(t *testing.T) {
	body, err := json.Marshal(classification.ProductSupply{Suppliers: []classification.ProductSupplierRef{{Code: "SUP0001"}, {Code: "SUP0002"}}})
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != `{"suppliers":[{"code":"SUP0001"},{"code":"SUP0002"}]}` {
		t.Fatalf("ProductSupply JSON = %s", body)
	}
}

func TestV2901ManufacturingCountryReferenceIsCodeOnly(t *testing.T) {
	body, err := json.Marshal(classification.ProductManufacturing{
		CompanyName: "Maker",
		CountryRef:  &classification.CountryCodeRef{Code: "TW"},
	})
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != `{"company_name":"Maker","country_ref":{"code":"TW"}}` {
		t.Fatalf("ProductManufacturing JSON = %s", body)
	}
}
