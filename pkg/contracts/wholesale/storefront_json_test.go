package wholesale_test

import (
	"encoding/json"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v10/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v10/pkg/contracts/product"
	"github.com/Potato-Mart/Backend-Shared-Contract/v10/pkg/contracts/wholesale"
	"github.com/Potato-Mart/Backend-Shared-Contract/v10/pkg/enums"
)

func TestApprovedStorefrontProductJSONShape(t *testing.T) {
	p := wholesale.ApprovedStorefrontProduct{
		ID:                "prd_1",
		SKUCode:           "SKU-1",
		Name:              "Bulk Potatoes",
		Description:       []common.LocalizedDescription{{Language: "en", Description: "Foodservice carton"}},
		Brand:             []common.LocalizedName{{Language: "en", Name: "Potato Mart"}},
		Supplier:          "SUP-1",
		Storage:           enums.StorageDry,
		DisplayStatus:     "active",
		CoverURL:          "https://example.com/potatoes.jpg",
		Collection:        &product.CollectionRef{ID: "col_produce", Name: []common.LocalizedName{{Language: "en", Name: "Produce"}}},
		CategoryTags:      []product.CategoryTag{{ID: "tag_potatoes", Name: []common.LocalizedName{{Language: "en", Name: "Potatoes"}}, CollectionID: "col_produce", CollectionName: []common.LocalizedName{{Language: "en", Name: "Produce"}}}},
		Price:             common.Money{AmountMinor: 2500, Currency: "AUD"},
		StockAvailable:    true,
		AvailabilityState: "available",
	}

	raw, err := json.Marshal(p)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(raw, &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	for _, key := range []string{
		"id",
		"sku_code",
		"name",
		"description",
		"brand",
		"supplier",
		"storage",
		"display_status",
		"cover_url",
		"collection",
		"category_tags",
		"price",
		"stock_available",
		"availability_state",
	} {
		if _, ok := got[key]; !ok {
			t.Fatalf("missing json key %q in %s", key, raw)
		}
	}

	collection, ok := got["collection"].(map[string]any)
	if !ok || collection["id"] != "col_produce" {
		t.Fatalf("collection = %#v, want collection object with id (%s)", got["collection"], raw)
	}
	collectionName, ok := collection["name"].([]any)
	if !ok || len(collectionName) != 1 || collectionName[0].(map[string]any)["name"] != "Produce" {
		t.Fatalf("collection.name = %#v, want localized Produce (%s)", collection["name"], raw)
	}

	tags, ok := got["category_tags"].([]any)
	if !ok || len(tags) != 1 {
		t.Fatalf("category_tags = %#v, want one tag (%s)", got["category_tags"], raw)
	}
	tag, ok := tags[0].(map[string]any)
	if !ok || tag["id"] != "tag_potatoes" || tag["collection_id"] != "col_produce" {
		t.Fatalf("category_tags[0] = %#v, want id and collection_id (%s)", tags[0], raw)
	}
	tagName, ok := tag["name"].([]any)
	if !ok || len(tagName) != 1 || tagName[0].(map[string]any)["name"] != "Potatoes" {
		t.Fatalf("category_tags[0].name = %#v, want localized Potatoes (%s)", tag["name"], raw)
	}
	tagCollectionName, ok := tag["collection_name"].([]any)
	if !ok || len(tagCollectionName) != 1 || tagCollectionName[0].(map[string]any)["name"] != "Produce" {
		t.Fatalf("category_tags[0].collection_name = %#v, want localized Produce (%s)", tag["collection_name"], raw)
	}

	for _, forbidden := range []string{"pricing", "online", "original", "cost", "current_stock", "history", "catalogue", "category_key", "category_path", "collection_id", "collection_name"} {
		if _, ok := got[forbidden]; ok {
			t.Fatalf("approved wholesale storefront product leaked %q in %s", forbidden, raw)
		}
	}
}
