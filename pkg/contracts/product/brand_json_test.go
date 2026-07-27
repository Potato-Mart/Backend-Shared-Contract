package product

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v19/pkg/common"
	productenum "github.com/Potato-Mart/Backend-Shared-Contract/v19/pkg/enums/product"
)

func TestBrandJSONRoundTripRequiresCanonicalKey(t *testing.T) {
	createdAt := time.Date(2026, 7, 27, 1, 2, 3, 0, time.UTC)
	want := Brand{
		ID: "brand_1", BrandKey: "happy-potato", Slug: "happy-potato",
		Name:        []common.LocalizedName{{Language: "en", Name: "Happy Potato"}},
		AuditFields: common.AuditFields{CreatedAt: createdAt, UpdatedAt: createdAt},
	}
	body, err := json.Marshal(want)
	if err != nil {
		t.Fatal(err)
	}
	for _, field := range []string{`"id":"brand_1"`, `"brand_key":"happy-potato"`, `"name":[`} {
		if !strings.Contains(string(body), field) {
			t.Fatalf("Brand JSON = %s, want %s", body, field)
		}
	}
	var got Brand
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatal(err)
	}
	if got.BrandKey != want.BrandKey || len(got.Name) != 1 {
		t.Fatalf("brand did not round-trip: %+v", got)
	}
	if strings.Contains(string(body), `"aliases"`) {
		t.Fatalf("Brand JSON retains removed localized aliases: %s", body)
	}
}

func TestBrandRefUsesBrandKeyOnly(t *testing.T) {
	body, err := json.Marshal(BrandRef{
		BrandKey: "happy-potato", Slug: "happy-potato",
		Name: []common.LocalizedName{{Language: "en", Name: "Happy Potato"}},
	})
	if err != nil {
		t.Fatal(err)
	}
	for _, field := range []string{`"brand_key":"happy-potato"`, `"slug":"happy-potato"`, `"name":[`} {
		if !strings.Contains(string(body), field) {
			t.Fatalf("BrandRef JSON = %s, want %s", body, field)
		}
	}
	for _, removed := range []string{`"id"`, `"brand"`, `"sort_order"`} {
		if strings.Contains(string(body), removed) {
			t.Fatalf("BrandRef JSON = %s, contains removed %s", body, removed)
		}
	}
}

func TestCanonicalBrandReferenceAcrossProductShapes(t *testing.T) {
	ref := &BrandRef{BrandKey: "happy-potato", Slug: "happy-potato", Name: []common.LocalizedName{{Language: "en", Name: "Happy Potato"}}}
	for name, value := range map[string]any{
		"product":    Product{ID: "prd_1", SKUCode: "A0001", SKU: "A1", Name: "Product", BrandRef: ref},
		"snapshot":   Snapshot{ID: "prd_1", SKUCode: "A0001", Name: "Product", BrandRef: ref},
		"storefront": StorefrontProduct{SKUCode: "A0001", SKU: "A1", Name: "Product", BrandKey: ref.BrandKey, BrandRef: ref},
	} {
		t.Run(name, func(t *testing.T) {
			body, err := json.Marshal(value)
			if err != nil {
				t.Fatal(err)
			}
			if !strings.Contains(string(body), `"brand_ref":{"brand_key":"happy-potato"`) {
				t.Fatalf("canonical brand missing: %s", body)
			}
			if strings.Contains(string(body), `"brand":[`) {
				t.Fatalf("removed localized brand array returned: %s", body)
			}
		})
	}
}

func TestBrandSummaryHasNoManualRank(t *testing.T) {
	summary := BrandSummary{
		BrandKey: "happy-potato",
		Names:    []common.LocalizedName{{Language: "en", Name: "Happy Potato"}},
		Audience: productenum.PriceAudienceWholesale, Featured: true, ActiveProductCount: 2,
	}
	body, err := json.Marshal(summary)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(body), "sort_order") || !strings.Contains(string(body), `"featured":true`) {
		t.Fatalf("BrandSummary JSON = %s", body)
	}
}
