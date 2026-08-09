package promotion

import (
	"encoding/json"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/pricing/promotion/promotion_enums"
)

func TestPromotionScopeRepresentsPerProductAndCombinedQuantityRequirements(t *testing.T) {
	maximum := int64(8)
	perProduct := PromotionScope{
		MatchMode: promotion_enums.PromotionMatchModeAll,
		Groups: []PromotionScopeGroup{
			{MatchMode: promotion_enums.PromotionMatchModeAny, ProductSKUCodes: []string{"POTATO-A"}, MinimumBaseUnits: 2},
			{MatchMode: promotion_enums.PromotionMatchModeAny, ProductSKUCodes: []string{"POTATO-B"}, MinimumBaseUnits: 1, MaximumBaseUnits: &maximum},
		},
	}
	combined := PromotionScope{
		MatchMode: promotion_enums.PromotionMatchModeAll,
		Groups: []PromotionScopeGroup{{
			MatchMode:        promotion_enums.PromotionMatchModeAny,
			ProductSKUCodes:  []string{"POTATO-A", "POTATO-B"},
			MinimumBaseUnits: 3,
		}},
	}

	body, err := json.Marshal(struct {
		PerProduct PromotionScope `json:"per_product"`
		Combined   PromotionScope `json:"combined"`
	}{PerProduct: perProduct, Combined: combined})
	if err != nil {
		t.Fatalf("marshal scopes: %v", err)
	}
	var got map[string]map[string]any
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("unmarshal scopes: %v", err)
	}
	perGroups := got["per_product"]["groups"].([]any)
	if got["per_product"]["match_mode"] != "all" || len(perGroups) != 2 {
		t.Fatalf("per-product ALL scope changed: %s", body)
	}
	first := perGroups[0].(map[string]any)
	if first["minimum_base_units"] != float64(2) || first["product_sku_codes"].([]any)[0] != "POTATO-A" {
		t.Fatalf("per-product scope group changed: %#v", first)
	}
	combinedGroups := got["combined"]["groups"].([]any)
	if len(combinedGroups) != 1 || len(combinedGroups[0].(map[string]any)["product_sku_codes"].([]any)) != 2 {
		t.Fatalf("combined quantity pool changed: %s", body)
	}
}

func TestPromotionScopeRepresentsUnrestrictedSelection(t *testing.T) {
	body, err := json.Marshal(PromotionScope{Unrestricted: true, MatchMode: promotion_enums.PromotionMatchModeAny})
	if err != nil {
		t.Fatalf("marshal unrestricted scope: %v", err)
	}
	var got map[string]any
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("unmarshal unrestricted scope: %v", err)
	}
	if got["unrestricted"] != true || got["match_mode"] != "any" {
		t.Fatalf("unrestricted scope changed: %s", body)
	}
	if _, present := got["groups"]; present {
		t.Fatalf("unrestricted scope should not require groups: %s", body)
	}
}

func TestPromotionScopeRoundTripsCollectionTagAndPackageSelectors(t *testing.T) {
	value := PromotionScope{
		MatchMode: promotion_enums.PromotionMatchModeAny,
		Groups: []PromotionScopeGroup{{
			MatchMode:        promotion_enums.PromotionMatchModeAll,
			CollectionIDs:    []string{"collection_root", "collection_seasonal"},
			CategoryTagIDs:   []string{"tag_potato", "tag_local"},
			PackageOptionIDs: []string{"pkg_each", "pkg_case"},
		}},
	}
	body, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("marshal selector scope: %v", err)
	}
	var got PromotionScope
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("unmarshal selector scope: %v", err)
	}
	group := got.Groups[0]
	if len(group.CollectionIDs) != 2 || group.CollectionIDs[1] != "collection_seasonal" || len(group.CategoryTagIDs) != 2 || group.CategoryTagIDs[0] != "tag_potato" || len(group.PackageOptionIDs) != 2 || group.PackageOptionIDs[1] != "pkg_case" {
		t.Fatalf("collection/tag/package selectors changed: %+v", group)
	}
}
