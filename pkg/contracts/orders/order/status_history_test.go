package order

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/orders/order/order_enums"
)

func TestStatusHistoryPreservesActorIdentity(t *testing.T) {
	raw, err := json.Marshal(StatusHistory{ToValue: "fulfillment:picking_in_progress", ActorID: "staff-1", ActorEmail: "staff@example.test"})
	if err != nil || !strings.Contains(string(raw), `"actor_id":"staff-1"`) {
		t.Fatalf("history=%s error=%v", raw, err)
	}
	var got StatusHistory
	if err := json.Unmarshal(raw, &got); err != nil || got.ActorID != "staff-1" {
		t.Fatalf("history=%+v error=%v", got, err)
	}
}

func TestStatusHistoryOverrideCarriesKindReasonAndChanges(t *testing.T) {
	const fixture = `{"from_value":"shipped","to_value":"processing","note":"Customer asked to add an item","kind":"override","reason_code":"STAFF_STATUS_OVERRIDE","changes":[{"field":"status","from_value":"shipped","to_value":"processing"},{"field":"shipping_amount","from_value":"990","to_value":"0"}],"actor_id":"staff-1","actor_email":"staff@example.test","created_at":"2026-11-12T01:02:03Z"}`
	var history StatusHistory
	if err := json.Unmarshal([]byte(fixture), &history); err != nil {
		t.Fatal(err)
	}
	if history.Kind != order_enums.StatusHistoryKindOverride || !history.Kind.IsValid() || history.ReasonCode != "STAFF_STATUS_OVERRIDE" {
		t.Fatalf("override classification lost: %+v", history)
	}
	if len(history.Changes) != 2 || history.Changes[0].Field != "status" || history.Changes[0].FromValue != "shipped" || history.Changes[0].ToValue != "processing" || history.Changes[1].ToValue != "0" {
		t.Fatalf("field changes lost: %+v", history.Changes)
	}
	encoded, err := json.Marshal(history)
	if err != nil {
		t.Fatal(err)
	}
	var got, want map[string]any
	if err := json.Unmarshal(encoded, &got); err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal([]byte(fixture), &want); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("override history JSON changed:\n got: %s\nwant: %s", encoded, fixture)
	}
}

func TestOrdinaryStatusHistoryOmitsOverrideFields(t *testing.T) {
	raw, err := json.Marshal(StatusHistory{FromValue: "paid", ToValue: "processing"})
	if err != nil {
		t.Fatal(err)
	}
	for _, key := range []string{`"kind"`, `"reason_code"`, `"changes"`} {
		if strings.Contains(string(raw), key) {
			t.Fatalf("ordinary history JSON = %s, want %s omitted", raw, key)
		}
	}
}
