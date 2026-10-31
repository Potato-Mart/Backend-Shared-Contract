package promotion

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/localization"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/pricing/promotion/promotion_enums"
)

func TestPromotionSavingOmitsAbsentLineAndNames(t *testing.T) {
	body, err := json.Marshal(PromotionSaving{
		PromotionID: "prm_order", PromotionRevision: 2, PromotionKind: "spend_discount",
		DiscountAmount: money.Money{AmountMinor: 500, Currency: "AUD"},
	})
	if err != nil {
		t.Fatalf("marshal saving: %v", err)
	}
	for _, field := range []string{`"line_id"`, `"names"`} {
		if strings.Contains(string(body), field) {
			t.Fatalf("order-level saving exposed %s: %s", field, body)
		}
	}

	var got PromotionSaving
	line := `{"promotion_id":"prm_line","promotion_revision":3,"promotion_kind":"bogo","names":[{"language":"en","name":"Two for one"}],"line_id":"line_1","discount_amount":{"amount_minor":250,"currency":"AUD"}}`
	if err := json.Unmarshal([]byte(line), &got); err != nil {
		t.Fatalf("unmarshal saving: %v", err)
	}
	if got.LineID != "line_1" || len(got.Names) != 1 || got.DiscountAmount.AmountMinor != 250 {
		t.Fatalf("line saving changed: %+v", got)
	}
}

func TestPromotionProgressRoundTripsShortfallEvidence(t *testing.T) {
	required, current, missing := int64(3), int64(0), int64(3)
	progress := PromotionProgress{
		PromotionID: "prm_bundle", PromotionRevision: 4, PromotionKind: "bundle", RelationID: "rel_bundle",
		Names:  []localization.LocalizedName{{Language: "en", Name: "Pantry bundle"}},
		Status: promotion_enums.PromotionProgressStatusNear, LineIDs: []string{"line_1"},
		RequiredQuantity: &required, CurrentQuantity: &current, MissingQuantity: &missing,
		MissingQualifiers: []PromotionMissingQualifier{{SKUCodes: []string{"POTATO-A", "POTATO-B"}, UnitsMissing: 3}},
		SpendMissing:      &money.Money{AmountMinor: 700, Currency: "AUD"},
		PotentialSaving:   &money.Money{AmountMinor: 300, Currency: "AUD"},
	}

	body, err := json.Marshal(progress)
	if err != nil {
		t.Fatalf("marshal progress: %v", err)
	}
	if !strings.Contains(string(body), `"current_quantity":0`) {
		t.Fatalf("known zero quantity must stay on the wire: %s", body)
	}
	var got PromotionProgress
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("unmarshal progress: %v", err)
	}
	if got.Status != promotion_enums.PromotionProgressStatusNear || got.RelationID != "rel_bundle" ||
		got.CurrentQuantity == nil || *got.MissingQuantity != 3 ||
		got.MissingQualifiers[0].UnitsMissing != 3 || got.MissingQualifiers[0].SKUCodes[1] != "POTATO-B" ||
		got.SpendMissing.AmountMinor != 700 || got.PotentialSaving.AmountMinor != 300 {
		t.Fatalf("progress evidence changed: %+v", got)
	}
}

func TestPromotionProgressOmitsUnknownOptionalEvidence(t *testing.T) {
	body, err := json.Marshal(PromotionProgress{
		PromotionID: "prm_open", PromotionRevision: 1, PromotionKind: "discount", RelationID: "rel_open",
		Status: promotion_enums.PromotionProgressStatusOpen, LineIDs: []string{},
	})
	if err != nil {
		t.Fatalf("marshal progress: %v", err)
	}
	if !strings.Contains(string(body), `"line_ids":[]`) {
		t.Fatalf("line_ids must stay present: %s", body)
	}
	for _, field := range []string{"names", "required_quantity", "current_quantity", "missing_quantity", "missing_qualifiers", "spend_missing", "potential_saving"} {
		if strings.Contains(string(body), `"`+field+`"`) {
			t.Fatalf("unknown evidence exposed %s: %s", field, body)
		}
	}
}

func TestPromotionApplicationNamesAreOptional(t *testing.T) {
	body, err := json.Marshal(PromotionApplication{PromotionID: "prm_a", PromotionKind: "bogo", RelationID: "rel_a"})
	if err != nil {
		t.Fatalf("marshal application: %v", err)
	}
	if strings.Contains(string(body), `"names"`) {
		t.Fatalf("application without names exposed the field: %s", body)
	}

	body, err = json.Marshal(PromotionApplication{Names: []localization.LocalizedName{{Language: "en", Name: "Two for one"}}})
	if err != nil {
		t.Fatalf("marshal named application: %v", err)
	}
	var got PromotionApplication
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("unmarshal application: %v", err)
	}
	if len(got.Names) != 1 || got.Names[0].Name != "Two for one" {
		t.Fatalf("application names changed: %+v", got)
	}
}
