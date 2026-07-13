package promotion

import (
	"encoding/json"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v16/pkg/common"
	promotionenum "github.com/Potato-Mart/Backend-Shared-Contract/v16/pkg/enums/promotion"
)

func TestPromotionCategoryTagTargetNameIsLocalized(t *testing.T) {
	body, err := json.Marshal(Promotion{
		ID:                    "prm_1",
		Name:                  "Hotpot tag discount",
		Type:                  promotionenum.PromotionTypeAutoDiscount,
		TargetScope:           promotionenum.DiscountScopeCategoryTag,
		TargetCategoryTagID:   "tag_hotpot",
		TargetCategoryTagName: []common.LocalizedName{{Language: "en", Name: "Hotpot"}},
	})
	if err != nil {
		t.Fatalf("marshal promotion: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("unmarshal promotion JSON: %v", err)
	}
	names, ok := got["target_category_tag_name"].([]any)
	if !ok || len(names) != 1 {
		t.Fatalf("target_category_tag_name = %#v, want localized name array (%s)", got["target_category_tag_name"], body)
	}
	name, ok := names[0].(map[string]any)
	if !ok || name["language"] != "en" || name["name"] != "Hotpot" {
		t.Fatalf("target_category_tag_name[0] = %#v, want en/Hotpot (%s)", names[0], body)
	}
}
