package product

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v21/pkg/common"
	customerenum "github.com/Potato-Mart/Backend-Shared-Contract/v21/pkg/enums/customer"
	productenum "github.com/Potato-Mart/Backend-Shared-Contract/v21/pkg/enums/product"
	salesenum "github.com/Potato-Mart/Backend-Shared-Contract/v21/pkg/enums/sales"
)

func TestProductJSONIncludesTaxed(t *testing.T) {
	body, err := json.Marshal(Product{
		ID:      "prd_1",
		SKUCode: "A0001",
		SKU:     "A",
		Name:    "Taxed product",
		Description: []common.LocalizedDescription{
			{Language: "en", Description: "Localized description"},
		},
		BrandRef:     &BrandRef{ID: "64c13ab08edf48a008793ca1", Slug: "localized-brand", Name: []common.LocalizedName{{Language: "en", Name: "Localized brand"}}},
		Taxed:        true,
		Collection:   &CollectionRef{ID: "col_frozen", Slug: "frozen", Name: []common.LocalizedName{{Language: "en", Name: "Frozen"}}},
		CategoryTags: []CategoryTag{{ID: "tag_hotpot", Slug: "hotpot", Name: []common.LocalizedName{{Language: "en", Name: "Hotpot"}}, CollectionID: "col_frozen", CollectionName: []common.LocalizedName{{Language: "en", Name: "Frozen"}}}},
		Supply:       &ProductSupply{Supplier: &ProductSupplierRef{Code: "sup_1", Name: "Supplier"}},
		PlacingArea:  &ProductPlacement{DepotCode: "MEL", LocationCode: "A1"},
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
	for _, keyValue := range []string{`"supply":{"supplier":{"code":"sup_1","name":"Supplier"}}`, `"placing_area":{"depot_code":"MEL","location_code":"A1"}`} {
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

func TestProductCollectionAndCategorySlugsAreOptional(t *testing.T) {
	body, err := json.Marshal(Collection{
		ID:   "col_frozen",
		Slug: "frozen",
		Name: []common.LocalizedName{{Language: "en", Name: "Frozen"}},
	})
	if err != nil {
		t.Fatalf("marshal collection with slug: %v", err)
	}
	if !strings.Contains(string(body), `"slug":"frozen"`) {
		t.Fatalf("Collection JSON = %s, want slug when present", body)
	}

	body, err = json.Marshal(Collection{
		ID:   "col_frozen",
		Name: []common.LocalizedName{{Language: "en", Name: "Frozen"}},
		CategoryTags: []CategoryTag{
			{ID: "tag_hotpot", Name: []common.LocalizedName{{Language: "en", Name: "Hotpot"}}, CollectionID: "col_frozen"},
		},
	})
	if err != nil {
		t.Fatalf("marshal collection: %v", err)
	}
	if strings.Contains(string(body), `"slug"`) {
		t.Fatalf("empty collection/category slugs should be omitted, got %s", body)
	}

	body, err = json.Marshal(CollectionRef{
		ID:   "col_frozen",
		Name: []common.LocalizedName{{Language: "en", Name: "Frozen"}},
	})
	if err != nil {
		t.Fatalf("marshal collection ref: %v", err)
	}
	if strings.Contains(string(body), `"slug"`) {
		t.Fatalf("empty collection ref slug should be omitted, got %s", body)
	}
}

func TestSnapshotJSONIncludesTaxed(t *testing.T) {
	body, err := json.Marshal(Snapshot{
		ID:      "prd_1",
		SKUCode: "A0001",
		Name:    "Taxed snapshot",
		Description: []common.LocalizedDescription{
			{Language: "en", Description: "Snapshot description"},
		},
		BrandRef: &BrandRef{ID: "64c13ab08edf48a008793ca1", Slug: "snapshot-brand", Name: []common.LocalizedName{{Language: "en", Name: "Snapshot brand"}}},
		Supply:   &ProductSupply{Supplier: &ProductSupplierRef{Code: "sup_1", Name: "Supplier"}},
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
		ID:      "prd_1",
		SKUCode: "A0001",
		SKU:     "A",
		Name:    "Wholesale-only product",
		Selling: &Selling{
			Channels:   []salesenum.OrderType{salesenum.OrderTypeB2B, salesenum.OrderTypePOS},
			BuyerTypes: []customerenum.BuyerType{customerenum.BuyerTypeWholesaleOrganisation},
			Visibility: productenum.PriceVisibilityWholesaleApprovedOnly,
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
	if decoded.Selling == nil || len(decoded.Selling.Channels) != 2 || decoded.Selling.BuyerTypes[0] != customerenum.BuyerTypeWholesaleOrganisation {
		t.Fatalf("selling rules did not round-trip: %+v", decoded.Selling)
	}
}

func TestProductOmitsEmptySelling(t *testing.T) {
	body, err := json.Marshal(Product{ID: "prd_1", SKUCode: "A0001", SKU: "A", Name: "No selling rules"})
	if err != nil {
		t.Fatalf("marshal product: %v", err)
	}
	if strings.Contains(string(body), `"selling"`) {
		t.Fatalf("empty selling should be omitted, got %s", body)
	}
}
