package product

import (
	"encoding/json"

	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/commerce/commerce_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/localization"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/packaging/packaging_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/customers/retail/retail_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/supply/classification"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/supply/product/product_enums"
)

func TestProductJSONIncludesTaxed(t *testing.T) {
	effectiveFrom := time.Date(2026, 8, 4, 0, 0, 0, 0, time.UTC)
	body, err := json.Marshal(Product{
		SKUCode:         "A0001",
		CategorySKUCode: "A",
		Name:            "Taxed product",
		Description: []localization.LocalizedDescription{
			{Language: "en", Description: "Localized description"},
		},
		BrandRef:     &classification.BrandRef{ID: "64c13ab08edf48a008793ca1", Slug: "localized-brand", Name: []localization.LocalizedName{{Language: "en", Name: "Localized brand"}}},
		Taxed:        true,
		Collection:   &classification.CollectionRef{ID: "col_frozen", Slug: "frozen", Name: []localization.LocalizedName{{Language: "en", Name: "Frozen"}}},
		CategoryTags: []classification.CategoryTag{{ID: "tag_hotpot", Slug: "hotpot", Name: []localization.LocalizedName{{Language: "en", Name: "Hotpot"}}, CollectionID: "col_frozen", CollectionName: []localization.LocalizedName{{Language: "en", Name: "Frozen"}}}},
		Supply:       &classification.ProductSupply{Supplier: &classification.ProductSupplierRef{Code: "sup_1", Name: "Supplier"}},
		PackageOptions: []ProductPackageOption{
			{ID: "pkg_each", Code: "EACH", ProductSKUCode: "A0001", HandlingUnit: packaging_enums.PackageHandlingUnitEach, UnitsPerPackage: 1, IsCanonical: true, IsActive: true, EffectiveFrom: effectiveFrom},
		},
	})
	if err != nil {
		t.Fatalf("marshal product: %v", err)
	}
	if !strings.Contains(string(body), `"taxed":true`) {
		t.Fatalf("Product JSON = %s, want taxed true field", body)
	}
	if !strings.Contains(string(body), `"description":[`) || !strings.Contains(string(body), `"brand_ref":`) {
		t.Fatalf("Product JSON = %s, want localized description and canonical brand reference", body)
	}

	var got map[string]any
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("unmarshal product JSON: %v", err)
	}
	collection, ok := got["collection"].(map[string]any)
	if !ok || collection["id"] != "col_frozen" {
		t.Fatalf("Product JSON = %s, want collection object with id", body)
	}
	if collection["slug"] != "frozen" {
		t.Fatalf("Product JSON = %s, want collection.slug=frozen", body)
	}
	collectionName, ok := collection["name"].([]any)
	if !ok || len(collectionName) != 1 || collectionName[0].(map[string]any)["name"] != "Frozen" {
		t.Fatalf("Product JSON = %s, want localized collection.name", body)
	}
	tags, ok := got["category_tags"].([]any)
	if !ok || len(tags) != 1 {
		t.Fatalf("Product JSON = %s, want one category tag", body)
	}
	tag, ok := tags[0].(map[string]any)
	if !ok || tag["id"] != "tag_hotpot" || tag["collection_id"] != "col_frozen" {
		t.Fatalf("Product JSON = %s, want category tag id and collection_id", body)
	}
	if tag["slug"] != "hotpot" {
		t.Fatalf("Product JSON = %s, want category tag slug=hotpot", body)
	}
	tagName, ok := tag["name"].([]any)
	if !ok || len(tagName) != 1 || tagName[0].(map[string]any)["name"] != "Hotpot" {
		t.Fatalf("Product JSON = %s, want localized category tag name", body)
	}
	tagCollectionName, ok := tag["collection_name"].([]any)
	if !ok || len(tagCollectionName) != 1 || tagCollectionName[0].(map[string]any)["name"] != "Frozen" {
		t.Fatalf("Product JSON = %s, want localized category tag collection_name", body)
	}
	for _, keyValue := range []string{`"supply":{"supplier":{"code":"sup_1","name":"Supplier"}}`, `"category_sku_code":"A"`, `"handling_unit":"EACH"`} {
		if !strings.Contains(string(body), keyValue) {
			t.Fatalf("Product JSON = %s, want flat %s", body, keyValue)
		}
	}
	for _, legacyTopLevelKey := range []string{"collection_id", "collection_name"} {
		if _, ok := got[legacyTopLevelKey]; ok {
			t.Fatalf("Product JSON = %s, should not include top-level %s", body, legacyTopLevelKey)
		}
	}
	if strings.Contains(string(body), `"identifiers"`) {
		t.Fatalf("Product JSON = %s, should not include nested identifiers", body)
	}
	for _, legacyKey := range []string{`"catalogue"`, `"category_key"`, `"category_path"`, `"merchandising"`} {
		if strings.Contains(string(body), legacyKey) {
			t.Fatalf("Product JSON = %s, should not include legacy %s", body, legacyKey)
		}
	}
}

func TestSnapshotJSONIncludesTaxed(t *testing.T) {
	body, err := json.Marshal(Snapshot{
		SKUCode: "A0001",
		Name:    "Taxed snapshot",
		Description: []localization.LocalizedDescription{
			{Language: "en", Description: "Snapshot description"},
		},
		BrandRef: &classification.BrandRef{ID: "64c13ab08edf48a008793ca1", Slug: "snapshot-brand", Name: []localization.LocalizedName{{Language: "en", Name: "Snapshot brand"}}},
		Supply:   &classification.ProductSupply{Supplier: &classification.ProductSupplierRef{Code: "sup_1", Name: "Supplier"}},
		Taxed:    true,
	})
	if err != nil {
		t.Fatalf("marshal snapshot: %v", err)
	}
	if !strings.Contains(string(body), `"taxed":true`) {
		t.Fatalf("Snapshot JSON = %s, want taxed true field", body)
	}
	if strings.Contains(string(body), `"supplier_code"`) || !strings.Contains(string(body), `"supply":{"supplier":{"code":"sup_1","name":"Supplier"}}`) {
		t.Fatalf("Snapshot JSON = %s, want canonical nested supply only", body)
	}
}

func TestProductSellingJSONGroup(t *testing.T) {
	body, err := json.Marshal(Product{
		SKUCode:         "A0001",
		CategorySKUCode: "A",
		Name:            "Wholesale-only product",
		Selling: &Selling{
			Channels:   []commerce_enums.OrderType{commerce_enums.OrderTypeB2B, commerce_enums.OrderTypePOS},
			BuyerTypes: []retail_enums.BuyerType{retail_enums.BuyerTypeWholesaleOrganisation},
			Visibility: product_enums.PriceVisibilityWholesaleApprovedOnly,
		},
	})
	if err != nil {
		t.Fatalf("marshal product: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("unmarshal product JSON: %v", err)
	}
	selling, ok := got["selling"].(map[string]any)
	if !ok {
		t.Fatalf("Product JSON missing selling group: %s", body)
	}
	if selling["visibility"] != "wholesale_approved_only" {
		t.Fatalf("selling.visibility = %v, want wholesale_approved_only (%s)", selling["visibility"], body)
	}

	var decoded Product
	if err := json.Unmarshal(body, &decoded); err != nil {
		t.Fatalf("unmarshal product: %v", err)
	}
	if decoded.Selling == nil || len(decoded.Selling.Channels) != 2 || decoded.Selling.BuyerTypes[0] != retail_enums.BuyerTypeWholesaleOrganisation {
		t.Fatalf("selling rules did not round-trip: %+v", decoded.Selling)
	}
}

func TestProductOmitsEmptySelling(t *testing.T) {
	body, err := json.Marshal(Product{SKUCode: "A0001", CategorySKUCode: "A", Name: "No selling rules"})
	if err != nil {
		t.Fatalf("marshal product: %v", err)
	}
	if strings.Contains(string(body), `"selling"`) {
		t.Fatalf("empty selling should be omitted, got %s", body)
	}
}
