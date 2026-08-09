package classification_test

import (
	"encoding/json"

	"strings"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/localization"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/supply/classification"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/supply/product"
)

func TestBrandJSONUsesV25PublicShape(t *testing.T) {
	want := classification.Brand{
		ID:      "64c13ab08edf48a008793ca1",
		Slug:    "happy-potato",
		Name:    []localization.LocalizedName{{Language: "en", Name: "Happy Potato"}},
		LogoURL: "https://cdn.example.com/brands/happy-potato.png",
	}
	body, err := json.Marshal(want)
	if err != nil {
		t.Fatal(err)
	}
	for _, field := range []string{`"id":"64c13ab08edf48a008793ca1"`, `"slug":"happy-potato"`, `"name":[`, `"logo_url":"https://cdn.example.com/brands/happy-potato.png"`} {
		if !strings.Contains(string(body), field) {
			t.Fatalf("Brand JSON = %s, want %s", body, field)
		}
	}
	for _, removed := range []string{`"brand_key"`, `"audience"`, `"featured"`, `"active_product_count"`, `"wholesale_product_count"`, `"created_at"`, `"updated_at"`, `"names"`} {
		if strings.Contains(string(body), removed) {
			t.Fatalf("Brand JSON = %s, contains removed %s", body, removed)
		}
	}
	var got classification.Brand
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatal(err)
	}
	if got.ID != want.ID || got.Slug != want.Slug || got.LogoURL != want.LogoURL || len(got.Name) != 1 {
		t.Fatalf("brand did not round-trip: %+v", got)
	}
}

func TestBrandRefUsesV25IdentityAndDisplayShape(t *testing.T) {
	body, err := json.Marshal(classification.BrandRef{
		ID:      "64c13ab08edf48a008793ca1",
		Slug:    "happy-potato",
		Name:    []localization.LocalizedName{{Language: "en", Name: "Happy Potato"}},
		LogoURL: "https://cdn.example.com/brands/happy-potato.png",
	})
	if err != nil {
		t.Fatal(err)
	}
	for _, field := range []string{`"id":"64c13ab08edf48a008793ca1"`, `"slug":"happy-potato"`, `"name":[`, `"logo_url":"https://cdn.example.com/brands/happy-potato.png"`} {
		if !strings.Contains(string(body), field) {
			t.Fatalf("BrandRef JSON = %s, want %s", body, field)
		}
	}
	for _, removed := range []string{`"brand_key"`, `"brand"`, `"sort_order"`, `"audience"`, `"active_product_count"`} {
		if strings.Contains(string(body), removed) {
			t.Fatalf("BrandRef JSON = %s, contains removed %s", body, removed)
		}
	}
}

func TestCanonicalBrandReferenceAcrossProductShapes(t *testing.T) {
	ref := &classification.BrandRef{ID: "64c13ab08edf48a008793ca1", Slug: "happy-potato", Name: []localization.LocalizedName{{Language: "en", Name: "Happy Potato"}}}
	for name, value := range map[string]any{
		"product": product.Product{SKUCode: "A0001", Classification: product.ProductClassification{CategorySKUCode: "A1", BrandRef: ref}},
	} {
		t.Run(name, func(t *testing.T) {
			body, err := json.Marshal(value)
			if err != nil {
				t.Fatal(err)
			}
			if !strings.Contains(string(body), `"brand_ref":{"id":"64c13ab08edf48a008793ca1"`) {
				t.Fatalf("canonical brand missing: %s", body)
			}
			for _, removed := range []string{`"brand_key"`, `"brand":[`} {
				if strings.Contains(string(body), removed) {
					t.Fatalf("removed brand representation returned: %s", body)
				}
			}
		})
	}
}

func TestCategoryTagRefUsesLightweightDisplayShape(t *testing.T) {
	body, err := json.Marshal(classification.CategoryTagRef{
		ID:   "tag_hotpot",
		Slug: "hotpot",
		Name: []localization.LocalizedName{{Language: "en", Name: "Hotpot"}},
	})
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{`"id":"tag_hotpot"`, `"slug":"hotpot"`, `"name":[`} {
		if !strings.Contains(string(body), want) {
			t.Fatalf("CategoryTagRef JSON = %s, want %s", body, want)
		}
	}
	for _, forbidden := range []string{`"collection_id"`, `"collection_name"`, `"created_at"`, `"updated_at"`} {
		if strings.Contains(string(body), forbidden) {
			t.Fatalf("CategoryTagRef JSON = %s, must not include %s", body, forbidden)
		}
	}
}

func TestBrandAndBrandRefOmitEmptyLogoURL(t *testing.T) {
	for name, value := range map[string]any{
		"brand": classification.Brand{ID: "64c13ab08edf48a008793ca1", Slug: "happy-potato", Name: []localization.LocalizedName{{Language: "en", Name: "Happy Potato"}}},
		"ref":   classification.BrandRef{ID: "64c13ab08edf48a008793ca1", Slug: "happy-potato", Name: []localization.LocalizedName{{Language: "en", Name: "Happy Potato"}}},
	} {
		t.Run(name, func(t *testing.T) {
			body, err := json.Marshal(value)
			if err != nil {
				t.Fatal(err)
			}
			if strings.Contains(string(body), `"logo_url"`) {
				t.Fatalf("empty logo_url should be omitted: %s", body)
			}
		})
	}
}
