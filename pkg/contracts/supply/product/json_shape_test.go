package product

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/localization"
	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/supply/classification"
	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/supply/classification/classification_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/supply/product/product_enums"
)

func TestV29CanonicalProductJSONShape(t *testing.T) {
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
		Supply: &classification.ProductSupply{Suppliers: []classification.ProductSupplierRef{{Code: "SUP0001"}}},
	}
	body, err := json.Marshal(product)
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{`"sku_code":"A00001"`, `"storage_type":"AMBIENT"`, `"name":{"language":"en","name":"Product"}`, `"sku_series_code":"A0"`, `"brands":[{"code":"BRD000001"}]`, `"code":"MED-SHA256"`} {
		if !strings.Contains(string(body), want) {
			t.Fatalf("Product JSON = %s, want %s", body, want)
		}
	}
	for _, legacy := range []string{`"product_category_code"`, `"brand_ref"`, `"supplier"`, `"media_id"`, `"url"`, `"slug"`} {
		if strings.Contains(string(body), legacy) {
			t.Fatalf("Product retained legacy %s: %s", legacy, body)
		}
	}
}

func TestV29SKUAndEmbeddedPackageUseCodesOnly(t *testing.T) {
	now := time.Date(2026, 8, 19, 0, 0, 0, 0, time.UTC)
	sku := SKU{
		SKUCode: "A00001", StorageType: classification_enums.StorageAmbient, Status: product_enums.SKUStatusActive,
		PackageOptions:     []ProductPackageOption{{Code: "PKG-A00001-EACH", SKUCode: "A00001", UnitsPerPackage: 1, EffectiveFrom: now}},
		BarcodeAssignments: []ProductBarcodeAssignment{{Code: "BAR-A00001", SKUCode: "A00001", PackageOptionCode: "PKG-A00001-EACH", Value: "A00001", Format: product_enums.BarcodeFormatCode128, EffectiveFrom: now}},
	}
	body, err := json.Marshal(sku)
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{`"sku_code":"A00001"`, `"package_option_code":"PKG-A00001-EACH"`, `"code":"BAR-A00001"`} {
		if !strings.Contains(string(body), want) {
			t.Fatalf("SKU JSON = %s, want %s", body, want)
		}
	}
	for _, model := range []reflect.Type{reflect.TypeOf(SKU{}), reflect.TypeOf(ProductPackageOption{}), reflect.TypeOf(ProductBarcodeAssignment{})} {
		if _, exists := model.FieldByName("ID"); exists {
			t.Fatalf("embedded model %s must not expose ID", model)
		}
	}
}

func TestV29ProductAndSKUStorageTypeCanBeEnforcedEqual(t *testing.T) {
	product := Product{StorageType: classification_enums.StorageFrozen}
	sku := SKU{StorageType: classification_enums.StorageFrozen}
	if product.StorageType != sku.StorageType {
		t.Fatal("product and SKU storage types differ")
	}
}
