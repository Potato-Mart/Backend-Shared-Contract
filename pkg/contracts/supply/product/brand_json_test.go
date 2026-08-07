package product

import (
	"encoding/json"
	common "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/common/shared"
	"strings"
	"testing"
)

func TestBrandJSONUsesV23PublicShape(t *testing.T) {
	want := Brand{
		ID:      "64c13ab08edf48a008793ca1",
		Slug:    "happy-potato",
		Name:    []common.LocalizedName{{Language: "en", Name: "Happy Potato"}},
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
	var got Brand
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatal(err)
	}
	if got.ID != want.ID || got.Slug != want.Slug || got.LogoURL != want.LogoURL || len(got.Name) != 1 {
		t.Fatalf("brand did not round-trip: %+v", got)
	}
}

func TestBrandRefUsesV23IdentityAndDisplayShape(t *testing.T) {
	body, err := json.Marshal(BrandRef{
		ID:      "64c13ab08edf48a008793ca1",
		Slug:    "happy-potato",
		Name:    []common.LocalizedName{{Language: "en", Name: "Happy Potato"}},
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
	ref := &BrandRef{ID: "64c13ab08edf48a008793ca1", Slug: "happy-potato", Name: []common.LocalizedName{{Language: "en", Name: "Happy Potato"}}}
	for name, value := range map[string]any{
		"product":    Product{SKUCode: "A0001", CategorySKUCode: "A1", Name: "Product", BrandRef: ref},
		"snapshot":   Snapshot{SKUCode: "A0001", Name: "Product", BrandRef: ref},
		"storefront": StorefrontProduct{SKUCode: "A0001", CategorySKUCode: "A1", Name: "Product", BrandRef: ref},
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

func TestBrandAndBrandRefOmitEmptyLogoURL(t *testing.T) {
	for name, value := range map[string]any{
		"brand": Brand{ID: "64c13ab08edf48a008793ca1", Slug: "happy-potato", Name: []common.LocalizedName{{Language: "en", Name: "Happy Potato"}}},
		"ref":   BrandRef{ID: "64c13ab08edf48a008793ca1", Slug: "happy-potato", Name: []common.LocalizedName{{Language: "en", Name: "Happy Potato"}}},
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
