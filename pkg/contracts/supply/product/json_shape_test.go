package product

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/localization"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/measurement"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/packaging/packaging_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/supply/classification"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/supply/product/product_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/supply/warehouse/warehouse_enums"
)

func TestProductJSONUsesCanonicalComponents(t *testing.T) {
	value := Product{
		ID:     "product_a0001",
		Status: product_enums.ProductStatusActive,
		Content: ProductContent{
			Name:         "Global product",
			Descriptions: []localization.LocalizedDescription{{Language: "en", Description: "Localized description"}},
		},
		Classification: ProductClassification{
			ProductCategoryCode: "A",
			BrandRef:            &classification.BrandRef{ID: "brand_1", Slug: "localized-brand", Name: []localization.LocalizedName{{Language: "en", Name: "Localized brand"}}},
			CollectionRef:       &classification.CollectionRef{ID: "col_frozen", Slug: "frozen", Name: []localization.LocalizedName{{Language: "en", Name: "Frozen"}}},
			CategoryTags:        []classification.CategoryTagRef{{ID: "tag_hotpot", Slug: "hotpot", Name: []localization.LocalizedName{{Language: "en", Name: "Hotpot"}}}},
		},
		Supply: &classification.ProductSupply{Supplier: &classification.ProductSupplierRef{Code: "sup_1", Name: "Supplier"}},
	}

	body, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("marshal product: %v", err)
	}
	for _, want := range []string{`"id":"product_a0001"`, `"status":"active"`, `"content":{"name":"Global product"`, `"classification":{"product_category_code":"A"`, `"supply":{"supplier":{"code":"sup_1","name":"Supplier"}}`} {
		if !strings.Contains(string(body), want) {
			t.Fatalf("Product JSON = %s, want %s", body, want)
		}
	}

	var got map[string]any
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("unmarshal product JSON: %v", err)
	}
	for _, component := range []string{"content", "classification"} {
		if _, ok := got[component].(map[string]any); !ok {
			t.Fatalf("Product JSON = %s, missing %s component", body, component)
		}
	}
	for _, retired := range []string{
		"sku_code", "packaging", "commerce", "metrics", "selling", "taxed",
		"package_options", "barcode_assignments", "storage_type", "package_price",
		"tax_amount", "stock_state", "first_listed_at", "sales_performance",
	} {
		if _, ok := got[retired]; ok {
			t.Fatalf("Product JSON = %s, retained retired key %s", body, retired)
		}
	}
	if strings.Contains(string(body), "price") || strings.Contains(string(body), "tax") {
		t.Fatalf("Product JSON must carry no price or tax facts: %s", body)
	}
}

func TestSKUJSONCarriesSellableIdentityWithoutMarketFacts(t *testing.T) {
	effectiveFrom := time.Date(2026, 8, 12, 0, 0, 0, 0, time.UTC)
	body, err := json.Marshal(SKU{
		ID:        "sku_a0001_375",
		ProductID: "product_a0001",
		Code:      "A0001-375",
		PackageOptions: []ProductPackageOption{{
			ID: "pkg_each", Code: "EACH", SKUID: "sku_a0001_375",
			HandlingUnit: packaging_enums.PackageHandlingUnitEach, UnitsPerPackage: 1,
			IsCanonical: true, IsActive: true, EffectiveFrom: effectiveFrom,
		}},
		BarcodeAssignments: []ProductBarcodeAssignment{{
			ID: "bar_1", SKUID: "sku_a0001_375", PackageOptionID: "pkg_each",
			Value: "9312345678907", Format: product_enums.BarcodeFormatEAN13,
			IsPrimary: true, EffectiveFrom: effectiveFrom,
		}},
		NetContent: &measurement.NetContent{
			NetQuantity:     measurement.Measure{Amount: 375, Unit: "mL"},
			StandardMeasure: measurement.Measure{Amount: 100, Unit: "mL"},
		},
		StorageType: warehouse_enums.StorageAmbient,
		Status:      product_enums.SKUStatusActive,
	})
	if err != nil {
		t.Fatalf("marshal sku: %v", err)
	}
	for _, want := range []string{
		`"id":"sku_a0001_375"`, `"product_id":"product_a0001"`, `"code":"A0001-375"`,
		`"package_options":[{`, `"sku_id":"sku_a0001_375"`, `"barcode_assignments":[{`,
		`"net_content":{"net_quantity":{"amount":375,"exponent":0,"unit":"mL"}`,
		`"standard_measure":{"amount":100,"exponent":0,"unit":"mL"}`,
		`"status":"active"`,
	} {
		if !strings.Contains(string(body), want) {
			t.Fatalf("SKU JSON = %s, want %s", body, want)
		}
	}
	for _, forbidden := range []string{`"market_id"`, `"price"`, `"amount_minor"`, `"taxed"`, `"tax_category_id"`, `"sku_code"`} {
		if strings.Contains(string(body), forbidden) {
			t.Fatalf("SKU JSON leaked market/commercial fact %s: %s", forbidden, body)
		}
	}

	var decoded SKU
	if err := json.Unmarshal(body, &decoded); err != nil {
		t.Fatalf("unmarshal sku: %v", err)
	}
	if decoded.NetContent == nil || decoded.NetContent.NetQuantity.Amount != 375 || decoded.NetContent.StandardMeasure.Unit != "mL" {
		t.Fatalf("net content did not round-trip: %+v", decoded.NetContent)
	}
}

func TestSKUOmitsOptionalMeasurementAndBarcodes(t *testing.T) {
	body, err := json.Marshal(SKU{ID: "sku_1", ProductID: "product_1", Code: "S1", Status: product_enums.SKUStatusDraft})
	if err != nil {
		t.Fatalf("marshal sku: %v", err)
	}
	for _, omitted := range []string{`"net_content"`, `"barcode_assignments"`, `"storage_type"`} {
		if strings.Contains(string(body), omitted) {
			t.Fatalf("empty %s should be omitted, got %s", omitted, body)
		}
	}
}
