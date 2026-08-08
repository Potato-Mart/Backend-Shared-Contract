package promotion

import (
	"encoding/json"

	geography "github.com/Potato-Mart/Backend-Shared-Contract/v25/pkg/contracts/common/geography"

	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v25/pkg/contracts/common/geography/geography_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v25/pkg/contracts/common/localization"
	"github.com/Potato-Mart/Backend-Shared-Contract/v25/pkg/contracts/pricing/promotion/promotion_enums"
)

func TestStorefrontPromotionOmitsRuleEngineInternals(t *testing.T) {
	now := time.Date(2026, 7, 16, 1, 2, 3, 0, time.UTC)
	body, err := json.Marshal(StorefrontPromotion{
		ID:               "prm_1",
		SeriesKey:        "series_weekly",
		Name:             "Weekly special",
		Description:      "Save on pantry favourites",
		Type:             promotion_enums.PromotionTypeAutoDiscount,
		Class:            promotion_enums.PromotionClassSpecialCampaign,
		TargetScope:      promotion_enums.DiscountScopeCategoryTag,
		CategoryTagIDs:   []string{"pantry:rice"},
		CategoryTagNames: []localization.LocalizedName{{Language: "en", Name: "Rice"}},
		StartsAt:         &now,
		ScheduleTimezone: "Australia/Sydney",
		GeographicScope: geography.GeographicScope{
			Mode:    geography_enums.GeographicScopeModeTargeted,
			Targets: []geography.GeographicTarget{{Kind: geography_enums.GeographicTargetSubdivision, Code: "AU-NSW"}},
		},
		IsActive: true,
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
	scope, ok := got["geographic_scope"].(map[string]any)
	if !ok || scope["mode"] != "TARGETED" || got["schedule_timezone"] != "Australia/Sydney" {
		t.Fatalf("storefront promotion geographic schedule mismatch: %s", body)
	}
}
