package fulfilment_test

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/fulfilment"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/warehouse/warehouse_enums"
)

// Units written off beyond the release are recorded apart from the released
// write-off, and a whole-settlement reversal records what it restored of them.
func TestPackingShortageSettlementBeyondReleaseRoundTrip(t *testing.T) {
	settledAt := time.Date(2026, 11, 14, 9, 0, 0, 0, time.UTC)
	before, after := int64(3), int64(0)
	want := fulfilment.PackingShortageSettlement{
		ID: "settlement-1", Reason: warehouse_enums.PackingShortageSettlementReasonNoStock, RefundRequestID: "refund-request-1",
		BaseUnits: 2, SettledBaseUnits: 2, Status: warehouse_enums.PackingShortageSettlementStatusReversed,
		ReleasedBaseUnits: 2, WrittenOffBaseUnits: 2, WriteOffMovementIDs: []string{"movement-1", "movement-2"},
		WrittenOffBeyondReleaseBaseUnits: 3,
		DepotAvailableBeforeBaseUnits:    &before, DepotAvailableAfterBaseUnits: &after,
		SettledBy: "staff-1", SettledAt: settledAt, ReversedBaseUnits: 2, ReleasedReversedBaseUnits: 2,
		Reversals: []fulfilment.PackingShortageSettlementReversal{{
			BaseUnits: 2, Cause: warehouse_enums.PackingShortageSettlementReversalCauseRefundRequestClosed, ReleasedBaseUnits: 2,
			WriteOffRestoredBaseUnits: 2, WriteOffBeyondReleaseRestoredBaseUnits: 3,
			WriteOffReversalMovementIDs: []string{"movement-3", "movement-4"},
			ReversedBy:                  "payments", ReversedAt: settledAt.Add(time.Hour),
		}},
	}
	raw, err := json.Marshal(want)
	if err != nil {
		t.Fatal(err)
	}
	for _, fragment := range []string{
		`"write_off_movement_ids":["movement-1","movement-2"],"written_off_beyond_release_base_units":3,`,
		`"cause":"refund_request_closed"`,
		`"write_off_restored_base_units":2,"write_off_beyond_release_restored_base_units":3,`,
	} {
		if !strings.Contains(string(raw), fragment) {
			t.Fatalf("beyond-release wire shape changed: missing %s in %s", fragment, raw)
		}
	}
	var got fulfilment.PackingShortageSettlement
	if err := json.Unmarshal(raw, &got); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("settlement round trip = %+v, want %+v", got, want)
	}
	if !got.Reversals[0].Cause.IsValid() {
		t.Fatalf("refund_request_closed must validate: %+v", got.Reversals[0])
	}
}

// Settlements recorded before v53.36.0 decode with zero beyond-release units,
// and zero values stay off the wire.
func TestPackingShortageSettlementBeyondReleaseIsOptional(t *testing.T) {
	var legacy fulfilment.PackingShortageSettlement
	if err := json.Unmarshal([]byte(`{"id":"settlement-1","reason":"no_stock","refund_request_id":"rr-1","base_units":1,"settled_base_units":1,"status":"applied","released_base_units":1,"written_off_base_units":1,"settled_by":"staff-1","settled_at":"2026-11-13T09:00:00Z","reversals":[{"base_units":1,"cause":"scan","reversed_by":"staff-1","reversed_at":"2026-11-13T10:00:00Z"}]}`), &legacy); err != nil {
		t.Fatal(err)
	}
	if legacy.WrittenOffBeyondReleaseBaseUnits != 0 || legacy.Reversals[0].WriteOffBeyondReleaseRestoredBaseUnits != 0 {
		t.Fatalf("legacy settlement must decode with no beyond-release units: %+v", legacy)
	}
	raw, err := json.Marshal(legacy)
	if err != nil {
		t.Fatal(err)
	}
	for _, key := range []string{"written_off_beyond_release_base_units", "write_off_beyond_release_restored_base_units"} {
		if strings.Contains(string(raw), key) {
			t.Fatalf("zero %s must be omitted: %s", key, raw)
		}
	}
}
