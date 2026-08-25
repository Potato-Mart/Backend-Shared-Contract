package promotion

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/commerce/commerce_enums"
	geography "github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/geography/geography_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/localization"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/pricing/promotion/promotion_enums"
)

func TestPromotionRoundTripsOpenPromotionAndRelationKinds(t *testing.T) {
	startsAt := time.Date(2026, 8, 9, 10, 0, 0, 0, time.UTC)
	maximumApplications := int64(20)
	value := Promotion{
		ID:        "prm_1",
		SeriesKey: "seasonal-potatoes",
		Kind:      "future_mechanic_without_contract_release",
		Status:    promotion_enums.PromotionStatusActive,
		Revision:  7,
		Content: PromotionContent{
			Names:           []localization.LocalizedName{{Language: "en", Name: "Weekend potato special"}},
			Descriptions:    []localization.LocalizedDescription{{Language: "en", Description: "A future mechanic"}},
			DisplayMessages: []localization.LocalizedText{{Language: "en", Text: "Special applied"}},
			ReceiptMessages: []localization.LocalizedText{{Language: "en", Text: "Weekend special"}},
		},
		Period: PromotionPeriod{StartsAt: &startsAt, Timezone: "Australia/Sydney"},
		Scope: PromotionScope{
			MatchMode: promotion_enums.PromotionMatchModeAll,
			Groups: []PromotionScopeGroup{{
				MatchMode:        promotion_enums.PromotionMatchModeAny,
				CategoryTagCodes: []string{"tag_potato"},
				MinimumBaseUnits: 1,
			}},
		},
		Relations: []PromotionRelation{{
			ID:             "rel_future",
			Kind:           "future_qualifier_to_target",
			QualifierScope: PromotionScope{MatchMode: promotion_enums.PromotionMatchModeAll, Groups: []PromotionScopeGroup{{MatchMode: promotion_enums.PromotionMatchModeAny, SKUCodes: []string{"POTATO-001"}}}},
			TargetScope:    PromotionScope{MatchMode: promotion_enums.PromotionMatchModeAll, Groups: []PromotionScopeGroup{{MatchMode: promotion_enums.PromotionMatchModeAny, SKUCodes: []string{"POTATO-002"}}}},
		}},
		Controls: PromotionControls{
			Priority:            10,
			Stackable:           true,
			MaximumApplications: &maximumApplications,
			Channels:            []commerce_enums.OrderType{commerce_enums.OrderTypeOnline, commerce_enums.OrderTypePOS},
			GeographicScope:     geography.GeographicScope{Mode: geography_enums.GeographicScopeModeGlobal},
		},
		Source: &PromotionSource{Kind: "campaign_import", Ref: "campaign_2026_08"},
	}

	body, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("marshal promotion: %v", err)
	}
	for _, want := range []string{
		`"kind":"future_mechanic_without_contract_release"`,
		`"kind":"future_qualifier_to_target"`,
		`"status":"active"`,
		`"ends_at"`,
		`"receipt_messages"`,
		`"ref":"campaign_2026_08"`,
	} {
		if want == `"ends_at"` {
			if strings.Contains(string(body), want) {
				t.Fatalf("nil period end must mean no scheduled end: %s", body)
			}
			continue
		}
		if !strings.Contains(string(body), want) {
			t.Fatalf("promotion JSON = %s, want %s", body, want)
		}
	}

	var got Promotion
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("unmarshal promotion: %v", err)
	}
	if got.Kind != value.Kind || len(got.Relations) != 1 || got.Relations[0].Kind != value.Relations[0].Kind || got.Period.EndsAt != nil {
		t.Fatalf("open promotion kinds or nil end did not round-trip: %+v", got)
	}
	if got.Content.DisplayMessages[0].Text != "Special applied" || got.Content.ReceiptMessages[0].Text != "Weekend special" {
		t.Fatalf("approved display/receipt content did not round-trip: %+v", got.Content)
	}
}
