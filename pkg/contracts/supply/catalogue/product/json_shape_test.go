package product

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/commerce/commerce_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/localization"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/money"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/pricing/market/market_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/pricing/pricebook"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/pricing/pricebook/pricebook_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/catalogue/classification"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/catalogue/classification/classification_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/catalogue/product/product_enums"
)

func TestCanonicalProductJSONShape(t *testing.T) {
	product := Product{
		ID: "64c13ab08edf48a008793ca6", SKUCode: "A00001", StorageType: classification_enums.StorageAmbient,
		Status: product_enums.ProductStatusActive,
		Content: ProductContent{
			Name:         localization.LocalizedName{Language: "en", Name: "Product"},
			Localization: &ProductLocalization{OtherNames: []localization.LocalizedName{{Language: "zh-TW", Name: "產品"}, {Language: "zh-CN", Name: "产品"}}},
			Images:       &Images{Cover: &classification.ObjectMediaRef{Code: "MED-SHA256"}},
		},
		Classification: ProductClassification{
			SKUSeriesCode: "A0", Brands: []classification.BrandRef{{Code: "BRD000001"}},
			CollectionRef: &classification.CollectionRef{Code: "COL0001"},
			CategoryTags:  []classification.CategoryTagRef{{Code: "TAG0001"}},
		},
		PackageOptions: []ProductPackageOption{{Code: "PKG-A00001-EACH", UnitsPerPackage: 1}},
		BarcodeAssignments: []ProductBarcodeAssignment{{
			Code: "BAR-A00001", PackageOptionCode: "PKG-A00001-EACH", Value: "A00001", Format: product_enums.BarcodeFormatCode128,
		}},
		Supply: &classification.ProductSupply{Suppliers: []classification.ProductSupplierRef{{Code: "SUP0001"}}},
	}
	body, err := json.Marshal(product)
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{`"sku_code":"A00001"`, `"storage_type":"AMBIENT"`, `"name":{"language":"en","name":"Product"}`, `"sku_series_code":"A0"`, `"brands":[{"code":"BRD000001"}]`, `"code":"MED-SHA256"`, `"package_options":[{"code":"PKG-A00001-EACH"`, `"barcode_assignments":[{"code":"BAR-A00001"`} {
		if !strings.Contains(string(body), want) {
			t.Fatalf("Product JSON = %s, want %s", body, want)
		}
	}
	for _, retired := range []string{`"product_category_code"`, `"brand_ref"`, `"supplier"`, `"media_id"`, `"url"`, `"slug"`} {
		if strings.Contains(string(body), retired) {
			t.Fatalf("Product retained retired %s: %s", retired, body)
		}
	}
}

func TestProductOwnsPackageAndBarcodeFacts(t *testing.T) {
	now := time.Date(2026, 8, 19, 0, 0, 0, 0, time.UTC)
	value := Product{
		SKUCode: "A00001", StorageType: classification_enums.StorageAmbient, Status: product_enums.ProductStatusActive,
		PackageOptions:     []ProductPackageOption{{Code: "PKG-A00001-EACH", UnitsPerPackage: 1, EffectiveFrom: now}},
		BarcodeAssignments: []ProductBarcodeAssignment{{Code: "BAR-A00001", PackageOptionCode: "PKG-A00001-EACH", Value: "A00001", Format: product_enums.BarcodeFormatCode128, EffectiveFrom: now}},
	}
	body, err := json.Marshal(value)
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{`"sku_code":"A00001"`, `"package_option_code":"PKG-A00001-EACH"`, `"code":"BAR-A00001"`} {
		if !strings.Contains(string(body), want) {
			t.Fatalf("Product JSON = %s, want %s", body, want)
		}
	}
	for _, model := range []reflect.Type{reflect.TypeOf(ProductPackageOption{}), reflect.TypeOf(ProductBarcodeAssignment{})} {
		if _, exists := model.FieldByName("ID"); exists {
			t.Fatalf("Product component %s must not expose ID", model)
		}
	}
	for _, model := range []reflect.Type{reflect.TypeOf(ProductPackageOption{}), reflect.TypeOf(ProductBarcodeAssignment{})} {
		field, exists := model.FieldByName("SKUCode")
		if !exists || field.Tag.Get("json") != "sku_code" {
			t.Fatalf("Product component %s must retain its standalone SKU snapshot identity", model)
		}
	}
}

func TestSellingProductIsCustomerSafeAndRenderComplete(t *testing.T) {
	now := time.Date(2026, 8, 24, 0, 0, 0, 0, time.UTC)
	value := SellingProduct{
		SKUCode: "A00001", SKUSeriesCode: "A0", StorageType: classification_enums.StorageAmbient,
		Content: SellingProductContent{
			Name:   localization.LocalizedName{Language: "en", Name: "Product"},
			Images: &SellingProductImages{Cover: &security.ObjectMedia{Code: "MED-SHA256", URL: "https://cdn.example/products/A00001.png"}},
		},
		Classification: SellingProductClassification{
			Brands: []SellingProductClassificationRef{{
				Code: "BRD000001", Name: []localization.LocalizedName{{Language: "en", Name: "Potato Mart"}}, Slug: "potato-mart",
				Media: &security.ObjectMedia{Code: "MED-BRAND", URL: "https://cdn.example/brands/potato-mart.png"},
			}},
		},
		PackageOptions:     []SellingProductPackageOption{{Code: "PKG-A00001-EACH", UnitsPerPackage: 1, IsCanonical: true}},
		BarcodeAssignments: []SellingProductBarcode{{PackageOptionCode: "PKG-A00001-EACH", Value: "1234567890123", Format: product_enums.BarcodeFormatEAN13, IsPrimary: true}},
		Price: pricebook.SellingPrice{
			UnitPrice:        money.Money{AmountMinor: 319, Currency: "AUD"},
			CurrencyExponent: money.CurrencyExponent{Currency: "AUD", Exponent: 2},
			MarketCode:       "AU", Channel: commerce_enums.OrderTypeOnline, Audience: market_enums.PriceAudienceRetail,
			PriceVisibility: pricebook_enums.PriceVisibilityVisible, TaxInclusion: pricebook_enums.PriceTaxInclusionInclusive,
			ValidFrom: now, AsOf: now,
		},
	}
	body, err := json.Marshal(value)
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{`"sku_code":"A00001"`, `"url":"https://cdn.example/products/A00001.png"`, `"slug":"potato-mart"`, `"unit_price":{"amount_minor":319,"currency":"AUD"}`, `"market_code":"AU"`, `"price_visibility":"visible"`} {
		if !strings.Contains(string(body), want) {
			t.Fatalf("SellingProduct JSON = %s, want %s", body, want)
		}
	}
	for _, forbidden := range []string{`"id"`, `"status"`, `"supply"`, `"administration"`, `"price_book_code"`, `"source_base_cost_revision"`, `"manufacturer_code"`, `"effective_from"`} {
		if strings.Contains(string(body), forbidden) {
			t.Fatalf("SellingProduct JSON leaked %s: %s", forbidden, body)
		}
	}
}
