package product

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v18/pkg/common"
)

func TestBrandJSONRoundTrip(t *testing.T) {
	createdAt := time.Date(2026, 7, 19, 1, 2, 3, 0, time.UTC)
	updatedAt := createdAt.Add(time.Hour)
	want := Brand{
		ID:       "brand_1",
		BrandKey: "happy-potato",
		Slug:     "happy-potato",
		Name:     []common.LocalizedName{{Language: "en", Name: "Happy Potato"}},
		Aliases:  []common.LocalizedName{{Language: "zh-Hant", Name: "開心馬鈴薯"}},
		AuditFields: common.AuditFields{
			CreatedAt: createdAt,
			CreatedBy: "admin_1",
			UpdatedAt: updatedAt,
			UpdatedBy: "admin_2",
		},
	}

	body, err := json.Marshal(want)
	if err != nil {
		t.Fatalf("marshal brand: %v", err)
	}
	for _, field := range []string{`"id":"brand_1"`, `"brand_key":"happy-potato"`, `"slug":"happy-potato"`, `"name":[`, `"aliases":[`, `"created_at":`, `"updated_at":`} {
		if !strings.Contains(string(body), field) {
			t.Fatalf("Brand JSON = %s, want %s", body, field)
		}
	}

	var got Brand
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("unmarshal brand: %v", err)
	}
	if got.ID != want.ID || got.BrandKey != want.BrandKey || got.Slug != want.Slug || len(got.Name) != 1 || len(got.Aliases) != 1 {
		t.Fatalf("brand did not round-trip: %+v", got)
	}
	if !got.CreatedAt.Equal(createdAt) || !got.UpdatedAt.Equal(updatedAt) {
		t.Fatalf("brand audit timestamps did not round-trip: %+v", got.AuditFields)
	}
}

func TestBrandRefJSONIsLightweight(t *testing.T) {
	body, err := json.Marshal(BrandRef{
		ID:       "brand_1",
		BrandKey: "happy-potato",
		Slug:     "happy-potato",
		Name:     []common.LocalizedName{{Language: "en", Name: "Happy Potato"}},
	})
	if err != nil {
		t.Fatalf("marshal brand ref: %v", err)
	}
	for _, field := range []string{`"id":"brand_1"`, `"brand_key":"happy-potato"`, `"slug":"happy-potato"`, `"name":[`} {
		if !strings.Contains(string(body), field) {
			t.Fatalf("BrandRef JSON = %s, want %s", body, field)
		}
	}
	for _, forbidden := range []string{`"aliases"`, `"created_at"`, `"updated_at"`} {
		if strings.Contains(string(body), forbidden) {
			t.Fatalf("BrandRef JSON = %s, must omit %s", body, forbidden)
		}
	}
}

func TestBrandRefIsAdditiveAcrossProductShapes(t *testing.T) {
	legacyBrand := []common.LocalizedName{{Language: "en", Name: "Legacy Brand"}}
	ref := &BrandRef{
		ID:       "brand_1",
		BrandKey: "legacy-brand",
		Slug:     "legacy-brand",
		Name:     []common.LocalizedName{{Language: "en", Name: "Legacy Brand"}},
	}

	tests := []struct {
		name       string
		withoutRef any
		withRef    any
	}{
		{
			name:       "product",
			withoutRef: Product{ID: "prd_1", SKUCode: "A0001", SKU: "A1", Name: "Product", Brand: legacyBrand},
			withRef:    Product{ID: "prd_1", SKUCode: "A0001", SKU: "A1", Name: "Product", Brand: legacyBrand, BrandRef: ref},
		},
		{
			name:       "snapshot",
			withoutRef: Snapshot{ID: "prd_1", SKUCode: "A0001", Name: "Product", Brand: legacyBrand},
			withRef:    Snapshot{ID: "prd_1", SKUCode: "A0001", Name: "Product", Brand: legacyBrand, BrandRef: ref},
		},
		{
			name:       "storefront product",
			withoutRef: StorefrontProduct{SKUCode: "A0001", SKU: "A1", Name: "Product", Brand: legacyBrand},
			withRef:    StorefrontProduct{SKUCode: "A0001", SKU: "A1", Name: "Product", Brand: legacyBrand, BrandRef: ref},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			legacyBody, err := json.Marshal(tc.withoutRef)
			if err != nil {
				t.Fatalf("marshal legacy-compatible shape: %v", err)
			}
			if strings.Contains(string(legacyBody), `"brand_ref"`) {
				t.Fatalf("optional brand_ref must be omitted when absent: %s", legacyBody)
			}
			if !strings.Contains(string(legacyBody), `"brand":[`) {
				t.Fatalf("legacy localized brand array must remain unchanged: %s", legacyBody)
			}

			body, err := json.Marshal(tc.withRef)
			if err != nil {
				t.Fatalf("marshal shape with brand ref: %v", err)
			}
			if !strings.Contains(string(body), `"brand_ref":{"id":"brand_1","brand_key":"legacy-brand","slug":"legacy-brand","name":[`) {
				t.Fatalf("shape JSON = %s, want additive brand_ref", body)
			}
			if !strings.Contains(string(body), `"brand":[`) {
				t.Fatalf("shape JSON = %s, must retain legacy brand array", body)
			}
		})
	}
}

func TestLegacyProductJSONDecodesWithoutBrandRef(t *testing.T) {
	const payload = `{"id":"prd_1","sku_code":"A0001","sku":"A1","name":"Product","brand":[{"language":"en","name":"Legacy Brand"}],"taxed":false,"current_stock":0,"pricing":{}}`

	var got Product
	if err := json.Unmarshal([]byte(payload), &got); err != nil {
		t.Fatalf("unmarshal legacy product JSON: %v", err)
	}
	if got.BrandRef != nil || len(got.Brand) != 1 || got.Brand[0].Name != "Legacy Brand" {
		t.Fatalf("legacy product compatibility changed: %+v", got)
	}
}

func TestBrandSummaryJSONShape(t *testing.T) {
	summary := BrandSummary{
		BrandKey: "happy-potato",
		Names: []common.LocalizedName{
			{Language: "en", Name: "Happy Potato"},
			{Language: "zh-TW", Name: "開心薯仔"},
		},
		Featured:           false,
		SortOrder:          0,
		ActiveProductCount: 0,
	}

	body, err := json.Marshal(summary)
	if err != nil {
		t.Fatalf("marshal brand summary: %v", err)
	}
	for _, field := range []string{
		`"brand_key":"happy-potato"`,
		`"names":[{"language":"en","name":"Happy Potato"},{"language":"zh-TW","name":"開心薯仔"}]`,
		`"featured":false`,
		`"sort_order":0`,
		`"active_product_count":0`,
	} {
		if !strings.Contains(string(body), field) {
			t.Fatalf("BrandSummary JSON = %s, want %s", body, field)
		}
	}
	if strings.Contains(string(body), `"logo_url"`) {
		t.Fatalf("BrandSummary JSON = %s, logo_url must be omitted when absent", body)
	}
	var decoded BrandSummary
	if err := json.Unmarshal(body, &decoded); err != nil {
		t.Fatalf("unmarshal brand summary: %v", err)
	}
	if decoded.BrandKey != summary.BrandKey || len(decoded.Names) != 2 || decoded.Names[1].Language != "zh-TW" {
		t.Fatalf("BrandSummary did not round-trip: %+v", decoded)
	}

	summary.LogoURL = "https://cdn.example.test/brands/happy-potato.png"
	body, err = json.Marshal(summary)
	if err != nil {
		t.Fatalf("marshal brand summary with logo: %v", err)
	}
	if !strings.Contains(string(body), `"logo_url":"https://cdn.example.test/brands/happy-potato.png"`) {
		t.Fatalf("BrandSummary JSON = %s, want logo_url", body)
	}
}

func TestBrandKeyIsAdditiveAndLegacyPayloadsRemainCompatible(t *testing.T) {
	legacyBrandJSON := `{"id":"brand_1","slug":"legacy-brand","name":[{"language":"en","name":"Legacy Brand"}],"created_at":"2026-07-19T01:02:03Z","updated_at":"2026-07-19T01:02:03Z"}`
	var brand Brand
	if err := json.Unmarshal([]byte(legacyBrandJSON), &brand); err != nil {
		t.Fatalf("decode legacy Brand: %v", err)
	}
	if brand.BrandKey != "" || brand.Slug != "legacy-brand" || len(brand.Name) != 1 {
		t.Fatalf("legacy Brand changed: %+v", brand)
	}
	body, err := json.Marshal(brand)
	if err != nil {
		t.Fatalf("re-encode legacy Brand: %v", err)
	}
	if strings.Contains(string(body), `"brand_key"`) {
		t.Fatalf("legacy Brand unexpectedly gained brand_key: %s", body)
	}

	legacyRefJSON := `{"id":"brand_1","slug":"legacy-brand","name":[{"language":"en","name":"Legacy Brand"}]}`
	var ref BrandRef
	if err := json.Unmarshal([]byte(legacyRefJSON), &ref); err != nil {
		t.Fatalf("decode legacy BrandRef: %v", err)
	}
	if ref.BrandKey != "" || ref.Slug != "legacy-brand" {
		t.Fatalf("legacy BrandRef changed: %+v", ref)
	}

	legacyStorefrontJSON := `{"sku_code":"A0001","sku":"A1","name":"Product","brand":[{"language":"en","name":"Legacy Brand"}],"current_stock":0,"pricing":{"audience":"retail"},"storefront_display":{}}`
	var storefront StorefrontProduct
	if err := json.Unmarshal([]byte(legacyStorefrontJSON), &storefront); err != nil {
		t.Fatalf("decode legacy StorefrontProduct: %v", err)
	}
	if storefront.BrandKey != "" || len(storefront.Brand) != 1 {
		t.Fatalf("legacy StorefrontProduct changed: %+v", storefront)
	}
	storefront.BrandKey = "legacy-brand"
	body, err = json.Marshal(storefront)
	if err != nil {
		t.Fatalf("marshal StorefrontProduct with brand_key: %v", err)
	}
	if !strings.Contains(string(body), `"brand_key":"legacy-brand"`) || !strings.Contains(string(body), `"brand":[`) {
		t.Fatalf("StorefrontProduct additive brand_key JSON = %s", body)
	}
}
