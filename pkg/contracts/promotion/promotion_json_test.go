package promotion

import (
	"encoding/json"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/common"
	geographyenum "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/enums/geography"
	promotionenum "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/enums/promotion"
)

func TestPromotionCategoryTagTargetNameIsLocalized(t *testing.T) {
	body, err := json.Marshal(Promotion{
		ID:                    "prm_1",
		SeriesKey:             "series_hotpot",
		Name:                  "Hotpot tag discount",
		Type:                  promotionenum.PromotionTypeAutoDiscount,
		Class:                 promotionenum.PromotionClassNormal,
		TargetScope:           promotionenum.DiscountScopeCategoryTag,
		TargetCategoryTagID:   "tag_hotpot",
		TargetCategoryTagName: []common.LocalizedName{{Language: "en", Name: "Hotpot"}},
		ActiveWindow:          ActiveWindow{ScheduleTimezone: "Australia/Sydney"},
		GeographicScope:       common.GeographicScope{Mode: geographyenum.GeographicScopeModeTargeted, Targets: []common.GeographicTarget{{Kind: geographyenum.GeographicTargetCountry, Code: "AU"}}},
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

func TestPromotionReceiptMessagesUseExplicitCustomerFacingShape(t *testing.T) {
	body, err := json.Marshal(Promotion{
		ID:              "prm_receipt",
		SeriesKey:       "series_receipt",
		Name:            "Internal campaign name",
		Type:            promotionenum.PromotionTypeAutoDiscount,
		Class:           promotionenum.PromotionClassNormal,
		TargetScope:     promotionenum.DiscountScopeAll,
		ReceiptEnabled:  true,
		ReceiptMessages: []common.LocalizedName{{Language: "en", Name: "Save 10% this weekend"}},
		ActiveWindow:    ActiveWindow{ScheduleTimezone: "Etc/UTC"},
		GeographicScope: common.GeographicScope{Mode: geographyenum.GeographicScopeModeGlobal},
	})
	if err != nil {
		t.Fatalf("marshal promotion: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("unmarshal promotion JSON: %v", err)
	}
	if got["receipt_enabled"] != true {
		t.Fatalf("receipt_enabled = %#v, want true (%s)", got["receipt_enabled"], body)
	}
	scope, ok := got["geographic_scope"].(map[string]any)
	if !ok || scope["mode"] != "GLOBAL" || got["schedule_timezone"] != "Etc/UTC" {
		t.Fatalf("promotion geographic schedule mismatch: %s", body)
	}
	messages, ok := got["receipt_messages"].([]any)
	if !ok || len(messages) != 1 {
		t.Fatalf("receipt_messages = %#v, want one localized message (%s)", got["receipt_messages"], body)
	}
	message, ok := messages[0].(map[string]any)
	if !ok || message["language"] != "en" || message["name"] != "Save 10% this weekend" {
		t.Fatalf("receipt_messages[0] = %#v, want en receipt copy (%s)", messages[0], body)
	}
}
