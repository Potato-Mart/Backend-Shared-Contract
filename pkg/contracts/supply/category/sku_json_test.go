package category

import (
	"encoding/json"
	common "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/common/shared"
	warehouseenum "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/supply/warehouse"
	"testing"
)

func TestSKUJSONIncludesPrimaryName(t *testing.T) {
	body, err := json.Marshal(SKU{
		ID:          "sku_f2",
		Code:        "F2",
		StorageType: warehouseenum.StorageFrozen,
		PrimaryName: common.LocalizedName{
			Language: "zh-Hant",
			Name:     "冷凍 - 肉品",
		},
		OtherNames: []common.LocalizedName{{Language: "en", Name: "Frozen meat"}},
	})
	if err != nil {
		t.Fatalf("marshal sku: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("unmarshal sku JSON: %v", err)
	}
	primary, ok := got["primary_name"].(map[string]any)
	if !ok {
		t.Fatalf("SKU JSON missing primary_name object: %s", body)
	}
	if primary["language"] != "zh-Hant" || primary["name"] != "冷凍 - 肉品" {
		t.Fatalf("primary_name = %#v, want zh-Hant/冷凍 - 肉品 (%s)", primary, body)
	}
	if _, exists := got["products"]; exists {
		t.Fatalf("SKU JSON must not embed products: %s", body)
	}
	if got["storage_type"] != "FROZEN" {
		t.Fatalf("SKU JSON storage_type = %v, want FROZEN: %s", got["storage_type"], body)
	}
	if _, exists := got["storage"]; exists {
		t.Fatalf("SKU JSON retained vague storage key: %s", body)
	}

	var decoded SKU
	if err := json.Unmarshal(body, &decoded); err != nil {
		t.Fatalf("unmarshal sku: %v", err)
	}
	if decoded.PrimaryName.Language != "zh-Hant" || decoded.PrimaryName.Name != "冷凍 - 肉品" {
		t.Fatalf("primary_name did not round-trip: %+v", decoded.PrimaryName)
	}
}
