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
		ID:      "brand_1",
		Slug:    "happy-potato",
		Name:    []common.LocalizedName{{Language: "en", Name: "Happy Potato"}},
		Aliases: []common.LocalizedName{{Language: "zh-Hant", Name: "開心馬鈴薯"}},
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
	for _, field := range []string{`"id":"brand_1"`, `"slug":"happy-potato"`, `"name":[`, `"aliases":[`, `"created_at":`, `"updated_at":`} {
		if !strings.Contains(string(body), field) {
			t.Fatalf("Brand JSON = %s, want %s", body, field)
		}
	}

	var got Brand
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("unmarshal brand: %v", err)
	}
	if got.ID != want.ID || got.Slug != want.Slug || len(got.Name) != 1 || len(got.Aliases) != 1 {
		t.Fatalf("brand did not round-trip: %+v", got)
	}
	if !got.CreatedAt.Equal(createdAt) || !got.UpdatedAt.Equal(updatedAt) {
		t.Fatalf("brand audit timestamps did not round-trip: %+v", got.AuditFields)
	}
}

func TestBrandRefJSONIsLightweight(t *testing.T) {
	body, err := json.Marshal(BrandRef{
		ID:   "brand_1",
		Slug: "happy-potato",
		Name: []common.LocalizedName{{Language: "en", Name: "Happy Potato"}},
	})
	if err != nil {
		t.Fatalf("marshal brand ref: %v", err)
	}
	for _, field := range []string{`"id":"brand_1"`, `"slug":"happy-potato"`, `"name":[`} {
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
		ID:   "brand_1",
		Slug: "legacy-brand",
		Name: []common.LocalizedName{{Language: "en", Name: "Legacy Brand"}},
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
			if !strings.Contains(string(body), `"brand_ref":{"id":"brand_1","slug":"legacy-brand","name":[`) {
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
