package category

import (
	"encoding/json"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v10/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v10/pkg/enums"
)

func TestSKUJSONIncludesPrimaryName(t *testing.T) {
	body, err := json.Marshal(SKU{
		ID:      "sku_f2",
		Code:    "F2",
		Storage: enums.StorageFrozen,
		PrimaryName: common.LocalizedName{
			Language: "zh-Hant",
			Name:     "冷凍 - 肉品",
		},
		OtherNames: []common.LocalizedName{{Language: "en", Name: "Frozen meat"}},
		SortOrder:  21,
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

	var decoded SKU
	if err := json.Unmarshal(body, &decoded); err != nil {
		t.Fatalf("unmarshal sku: %v", err)
	}
	if decoded.PrimaryName.Language != "zh-Hant" || decoded.PrimaryName.Name != "冷凍 - 肉品" {
		t.Fatalf("primary_name did not round-trip: %+v", decoded.PrimaryName)
	}
}
