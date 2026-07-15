package category

import (
	"encoding/json"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v17/pkg/common"
	warehouseenum "github.com/Potato-Mart/Backend-Shared-Contract/v17/pkg/enums/warehouse"
)

func TestSKUJSONIncludesPrimaryName(t *testing.T) {
	body, err := json.Marshal(SKU{
		ID:      "sku_f2",
		Code:    "F2",
		Storage: warehouseenum.StorageFrozen,
		PrimaryName: common.LocalizedName{
			Language: "zh-Hant",
			Name:     "å†·å‡ - è‚‰å“",
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
	if primary["language"] != "zh-Hant" || primary["name"] != "å†·å‡ - è‚‰å“" {
		t.Fatalf("primary_name = %#v, want zh-Hant/å†·å‡ - è‚‰å“ (%s)", primary, body)
	}

	var decoded SKU
	if err := json.Unmarshal(body, &decoded); err != nil {
		t.Fatalf("unmarshal sku: %v", err)
	}
	if decoded.PrimaryName.Language != "zh-Hant" || decoded.PrimaryName.Name != "å†·å‡ - è‚‰å“" {
		t.Fatalf("primary_name did not round-trip: %+v", decoded.PrimaryName)
	}
}
