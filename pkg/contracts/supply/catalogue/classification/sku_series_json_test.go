package classification

import (
	"encoding/json"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/localization"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/catalogue/classification/classification_enums"
)

func TestSKUSeriesJSONIncludesPrimaryEnglishName(t *testing.T) {
	body, err := json.Marshal(SKUSeries{
		ID: "64c13ab08edf48a008793ca5", Code: "F2", StorageType: classification_enums.StorageFrozen,
		PrimaryName: localization.LocalizedName{Language: "en", Name: "Frozen Meat"},
		OtherNames:  []localization.LocalizedName{{Language: "zh-TW", Name: "冷凍肉品"}, {Language: "zh-CN", Name: "冷冻肉品"}},
	})
	if err != nil {
		t.Fatal(err)
	}
	var got map[string]any
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatal(err)
	}
	if got["code"] != "F2" || got["storage_type"] != "FROZEN" || got["primary_name"].(map[string]any)["language"] != "en" {
		t.Fatalf("SKUSeries JSON = %s", body)
	}
}

func TestStorageTypeEnum(t *testing.T) {
	for _, value := range []classification_enums.StorageType{classification_enums.StorageAmbient, classification_enums.StorageChilled, classification_enums.StorageFrozen} {
		if !value.IsValid() {
			t.Fatalf("storage type %q must validate", value)
		}
	}
	if classification_enums.StorageType("DRY").IsValid() {
		t.Fatal("DRY must not validate")
	}
}
