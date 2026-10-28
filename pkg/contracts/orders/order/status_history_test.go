package order

import (
	"encoding/json"
	"strings"
	"testing"
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
