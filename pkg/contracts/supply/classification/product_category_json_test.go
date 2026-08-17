package classification

import (
	"encoding/json"

	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/localization"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/supply/warehouse/warehouse_enums"
)

func TestProductCategoryJSONIncludesPrimaryName(t *testing.T) {
	body, err := json.Marshal(ProductCategory{
		ID:          "product_category_f2",
		Code:        "F2",
		StorageType: warehouse_enums.StorageFrozen,
		PrimaryName: localization.LocalizedName{
			Language: "zh-Hant",
			Name:     "冷凍 - 肉品",
		},
		OtherNames: []localization.LocalizedName{{Language: "en", Name: "Frozen meat"}},
	})
	if err != nil {
		t.Fatalf("marshal product category: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("unmarshal product category JSON: %v", err)
	}
	primary, ok := got["primary_name"].(map[string]any)
	if !ok {
		t.Fatalf("ProductCategory JSON missing primary_name object: %s", body)
	}
	if primary["language"] != "zh-Hant" || primary["name"] != "冷凍 - 肉品" {
		t.Fatalf("primary_name = %#v, want zh-Hant/冷凍 - 肉品 (%s)", primary, body)
	}
	if _, exists := got["products"]; exists {
		t.Fatalf("ProductCategory JSON must not embed products: %s", body)
	}
	if got["storage_type"] != "FROZEN" {
		t.Fatalf("ProductCategory JSON storage_type = %v, want FROZEN: %s", got["storage_type"], body)
	}
	if _, exists := got["storage"]; exists {
		t.Fatalf("ProductCategory JSON retained vague storage key: %s", body)
	}

	var decoded ProductCategory
	if err := json.Unmarshal(body, &decoded); err != nil {
		t.Fatalf("unmarshal product category: %v", err)
	}
	if decoded.PrimaryName.Language != "zh-Hant" || decoded.PrimaryName.Name != "冷凍 - 肉品" {
		t.Fatalf("primary_name did not round-trip: %+v", decoded.PrimaryName)
	}
}
