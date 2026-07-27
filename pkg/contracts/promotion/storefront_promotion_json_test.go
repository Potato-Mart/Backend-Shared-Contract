package promotion

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v19/pkg/common"
	promotionenum "github.com/Potato-Mart/Backend-Shared-Contract/v19/pkg/enums/promotion"
)

func TestStorefrontPromotionOmitsRuleEngineInternals(t *testing.T) {
	now := time.Date(2026, 7, 16, 1, 2, 3, 0, time.UTC)
	body, err := json.Marshal(StorefrontPromotion{
		ID:               "prm_1",
		Name:             "Weekly special",
		Description:      "Save on pantry favourites",
		Type:             promotionenum.PromotionTypeAutoDiscount,
		Class:            promotionenum.PromotionClassSpecialCampaign,
		TargetScope:      promotionenum.DiscountScopeCategoryTag,
		CategoryTagIDs:   []string{"pantry:rice"},
		CategoryTagNames: []common.LocalizedName{{Language: "en", Name: "Rice"}},
		StartsAt:         &now,
		IsActive:         true,
	})
	if err != nil {
		t.Fatalf("marshal storefront promotion: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("unmarshal storefront promotion: %v", err)
	}
	for _, forbidden := range []string{
		"discount_type",
		"discount_value",
		"max_discount",
		"usage_limit",
		"used_count",
		"per_customer_limit",
		"priority",
		"is_stackable",
		"source",
		"source_ref",
		"history",
		"created_at",
		"updated_at",
	} {
		if _, exists := got[forbidden]; exists {
			t.Fatalf("storefront promotion leaked %q (%s)", forbidden, body)
		}
	}
	if got["is_active"] != true || got["target_scope"] != "category_tag" {
		t.Fatalf("storefront promotion shape mismatch: %s", body)
	}
}
