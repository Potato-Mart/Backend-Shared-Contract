package promotion

import (
	"encoding/json"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v17/pkg/common"
	promotionenum "github.com/Potato-Mart/Backend-Shared-Contract/v17/pkg/enums/promotion"
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

func TestPromotionReceiptMessagesUseExplicitCustomerFacingShape(t *testing.T) {
	body, err := json.Marshal(Promotion{
		ID:              "prm_receipt",
		Name:            "Internal campaign name",
		Type:            promotionenum.PromotionTypeAutoDiscount,
		ReceiptEnabled:  true,
		ReceiptMessages: []common.LocalizedName{{Language: "en", Name: "Save 10% this weekend"}},
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
	messages, ok := got["receipt_messages"].([]any)
	if !ok || len(messages) != 1 {
		t.Fatalf("receipt_messages = %#v, want one localized message (%s)", got["receipt_messages"], body)
	}
	message, ok := messages[0].(map[string]any)
	if !ok || message["language"] != "en" || message["name"] != "Save 10% this weekend" {
		t.Fatalf("receipt_messages[0] = %#v, want en receipt copy (%s)", messages[0], body)
	}
}
