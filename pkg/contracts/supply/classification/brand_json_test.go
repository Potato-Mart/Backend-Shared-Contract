package classification_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/localization"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/classification"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/product"
)

func TestBrandRootRetainsIDAndSlugWhileReferenceUsesCode(t *testing.T) {
	brand := classification.Brand{
		ID: "64c13ab08edf48a008793ca1", Code: "BRD000001", Slug: "happy-potato",
		Name: []localization.LocalizedName{{Language: "en", Name: "Happy Potato"}},
	}
	body, err := json.Marshal(brand)
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{`"id":"64c13ab08edf48a008793ca1"`, `"code":"BRD000001"`, `"slug":"happy-potato"`} {
		if !strings.Contains(string(body), want) {
			t.Fatalf("Brand JSON = %s, want %s", body, want)
		}
	}

	refBody, err := json.Marshal(classification.BrandRef{Code: brand.Code})
	if err != nil {
		t.Fatal(err)
	}
	if string(refBody) != `{"code":"BRD000001"}` {
		t.Fatalf("BrandRef must be code-only identity: %s", refBody)
	}
}

func TestBrandLogoUsesCodeOnlyCatalogMediaRef(t *testing.T) {
	body, err := json.Marshal(classification.Brand{Logo: &classification.ObjectMediaRef{Code: "MED-BRAND"}})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(body), `"logo":{"code":"MED-BRAND"}`) || strings.Contains(string(body), `"url"`) {
		t.Fatalf("Brand logo must persist only the media code: %s", body)
	}
}

func TestProductCarriesOrderedBrandReferences(t *testing.T) {
	body, err := json.Marshal(product.Product{Classification: product.ProductClassification{
		SKUSeriesCode: "A0",
		Brands:        []classification.BrandRef{{Code: "BRD000001"}, {Code: "BRD000002"}},
	}})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(body), `"brands":[{"code":"BRD000001"},{"code":"BRD000002"}]`) {
		t.Fatalf("ordered brands missing: %s", body)
	}
}

func TestCategoryTagRefUsesCodeWithoutSlugOrID(t *testing.T) {
	body, err := json.Marshal(classification.CategoryTagRef{Code: "TAG0001"})
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != `{"code":"TAG0001"}` {
		t.Fatalf("CategoryTagRef JSON = %s", body)
	}
}
