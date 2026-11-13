package fulfilment_test

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/orders/order/order_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/fulfilment"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/warehouse/warehouse_enums"
)

func TestPackingRecordRoundTrip(t *testing.T) {
	now := time.Date(2026, 10, 29, 0, 0, 0, 0, time.UTC)
	want := fulfilment.PackingRecord{
		ID: "packing-1", OrderNumber: "ORDER-1", FulfillmentGeneration: 2,
		Revision: 3, Status: order_enums.FulfillmentStatusPacked, Operator: "staff-1",
		Lines: []fulfilment.PackingLine{{ID: "line-1", OrderItemID: "item-1", SKUCode: "SKU-1",
			RequestedComposition: packaging.PackageCompositionSnapshot{TotalBaseUnits: 4},
			PackedComposition:    packaging.PackageCompositionSnapshot{TotalBaseUnits: 3}}},
		Containers:    []fulfilment.OutboundContainerPlan{{ID: "container-1"}},
		Discrepancies: []fulfilment.PackingDiscrepancy{{ID: "shortage-1", OrderItemID: "item-1", SKUCode: "SKU-1"}},
		StartedAt:     &now, PackedAt: &now, CreatedAt: now, UpdatedAt: now,
	}
	raw, err := json.Marshal(want)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(raw), "picking_list") || strings.Contains(string(raw), "picked_composition") {
		t.Fatalf("packing record contains retired picking state: %s", raw)
	}
	var got fulfilment.PackingRecord
	if err := json.Unmarshal(raw, &got); err != nil {
		t.Fatal(err)
	}
	if got.OrderNumber != want.OrderNumber || got.FulfillmentGeneration != 2 || got.Revision != 3 || got.Operator != "staff-1" || got.Status != order_enums.FulfillmentStatusPacked || len(got.Lines) != 1 || got.Lines[0].RequestedComposition.TotalBaseUnits != 4 || got.Lines[0].PackedComposition.TotalBaseUnits != 3 || len(got.Containers) != 1 || len(got.Discrepancies) != 1 || got.Discrepancies[0].OrderItemID != "item-1" || got.StartedAt == nil || got.PackedAt == nil {
		t.Fatalf("packing record lost evidence: %+v", got)
	}
}

// Two lines can share a SKU; the discrepancy must retain its exact line.
func TestPackingDiscrepancyCarriesOrderItemIdentity(t *testing.T) {
	rows := []fulfilment.PackingDiscrepancy{{ID: "d1", OrderItemID: "item-1", SKUCode: "SKU-1"}, {ID: "d2", OrderItemID: "item-2", SKUCode: "SKU-1"}}
	raw, err := json.Marshal(rows)
	if err != nil {
		t.Fatal(err)
	}
	var decoded []fulfilment.PackingDiscrepancy
	if err := json.Unmarshal(raw, &decoded); err != nil {
		t.Fatal(err)
	}
	if len(decoded) != 2 || decoded[0].OrderItemID != "item-1" || decoded[1].OrderItemID != "item-2" || decoded[0].SKUCode != decoded[1].SKUCode {
		t.Fatalf("discrepancy line identity lost: %+v", decoded)
	}
}

// A shortage's remaining base units are its recorded base units minus the sum
// of its withdrawals.
func TestPackingDiscrepancyWithdrawalsRoundTrip(t *testing.T) {
	recordedAt := time.Date(2026, 11, 1, 9, 0, 0, 0, time.UTC)
	want := fulfilment.PackingDiscrepancy{
		ID: "shortage-1", OrderItemID: "item-1", SKUCode: "SKU-1",
		Kind:                 warehouse_enums.PackingDiscrepancyKindShortage,
		RequestedComposition: packaging.PackageCompositionSnapshot{TotalBaseUnits: 5},
		ObservedComposition:  packaging.PackageCompositionSnapshot{TotalBaseUnits: 1},
		RecordedAt:           recordedAt,
		Withdrawals: []fulfilment.PackingDiscrepancyWithdrawal{
			{WithdrawnBaseUnits: 1, Method: warehouse_enums.PackingDiscrepancyWithdrawalMethodCommand, WithdrawnBy: "staff-1", WithdrawnAt: recordedAt.Add(time.Minute)},
			{WithdrawnBaseUnits: 2, Method: warehouse_enums.PackingDiscrepancyWithdrawalMethodScan, WithdrawnBy: "staff-2", WithdrawnAt: recordedAt.Add(2 * time.Minute)},
		},
	}
	raw, err := json.Marshal(want)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(raw), `"withdrawals":[{"withdrawn_base_units":1,"method":"command","withdrawn_by":"staff-1","withdrawn_at":"2026-11-01T09:01:00Z"},{"withdrawn_base_units":2,"method":"scan","withdrawn_by":"staff-2","withdrawn_at":"2026-11-01T09:02:00Z"}]`) {
		t.Fatalf("withdrawals wire shape changed: %s", raw)
	}
	var got fulfilment.PackingDiscrepancy
	if err := json.Unmarshal(raw, &got); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got.Withdrawals, want.Withdrawals) {
		t.Fatalf("withdrawals lost: %+v", got.Withdrawals)
	}
	remaining := got.RequestedComposition.TotalBaseUnits - got.ObservedComposition.TotalBaseUnits
	for _, withdrawal := range got.Withdrawals {
		if !withdrawal.Method.IsValid() {
			t.Fatalf("withdrawal method %q must validate", withdrawal.Method)
		}
		remaining -= withdrawal.WithdrawnBaseUnits
	}
	if remaining != 1 {
		t.Fatalf("remaining shortage base units = %d, want 1", remaining)
	}

	raw, err = json.Marshal(fulfilment.PackingDiscrepancy{ID: "shortage-2"})
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(raw), "withdrawals") {
		t.Fatalf("empty withdrawals must be omitted: %s", raw)
	}
}

// A reinstate entry re-records previously withdrawn units as missing again.
// withdrawn_base_units stays positive; the method gives the direction.
func TestPackingDiscrepancyReinstateRestoresRemainingShortage(t *testing.T) {
	recordedAt := time.Date(2026, 11, 1, 9, 0, 0, 0, time.UTC)
	discrepancy := fulfilment.PackingDiscrepancy{
		ID: "shortage-3", OrderItemID: "item-1", SKUCode: "SKU-1",
		Kind:                 warehouse_enums.PackingDiscrepancyKindShortage,
		RequestedComposition: packaging.PackageCompositionSnapshot{TotalBaseUnits: 6},
		ObservedComposition:  packaging.PackageCompositionSnapshot{TotalBaseUnits: 2},
		RecordedAt:           recordedAt,
		Withdrawals: []fulfilment.PackingDiscrepancyWithdrawal{
			{WithdrawnBaseUnits: 3, Method: warehouse_enums.PackingDiscrepancyWithdrawalMethodScan, WithdrawnBy: "staff-1", WithdrawnAt: recordedAt.Add(time.Minute)},
			{WithdrawnBaseUnits: 1, Method: warehouse_enums.PackingDiscrepancyWithdrawalMethodCommand, WithdrawnBy: "staff-1", WithdrawnAt: recordedAt.Add(2 * time.Minute)},
			{WithdrawnBaseUnits: 2, Method: warehouse_enums.PackingDiscrepancyWithdrawalMethodReinstate, WithdrawnBy: "staff-2", WithdrawnAt: recordedAt.Add(3 * time.Minute)},
		},
	}
	raw, err := json.Marshal(discrepancy)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(raw), `{"withdrawn_base_units":2,"method":"reinstate","withdrawn_by":"staff-2","withdrawn_at":"2026-11-01T09:03:00Z"}`) {
		t.Fatalf("reinstate wire shape changed: %s", raw)
	}
	var got fulfilment.PackingDiscrepancy
	if err := json.Unmarshal(raw, &got); err != nil {
		t.Fatal(err)
	}
	remaining := got.RequestedComposition.TotalBaseUnits - got.ObservedComposition.TotalBaseUnits
	for _, entry := range got.Withdrawals {
		if !entry.Method.IsValid() || entry.WithdrawnBaseUnits <= 0 {
			t.Fatalf("withdrawal entry must have a valid method and positive units: %+v", entry)
		}
		switch entry.Method {
		case warehouse_enums.PackingDiscrepancyWithdrawalMethodReinstate:
			remaining += entry.WithdrawnBaseUnits
		default:
			remaining -= entry.WithdrawnBaseUnits
		}
	}
	if remaining != 2 {
		t.Fatalf("remaining shortage base units = %d, want 2", remaining)
	}
}

// A settlement records the stock outcome of refunded shortage units without
// changing the discrepancy kind or its remaining base units.
func TestPackingDiscrepancySettlementsRoundTrip(t *testing.T) {
	settledAt := time.Date(2026, 11, 13, 9, 0, 0, 0, time.UTC)
	before, after := int64(-1), int64(0)
	want := fulfilment.PackingDiscrepancy{
		ID: "shortage-1", OrderItemID: "item-1", SKUCode: "SKU-1",
		Kind:                 warehouse_enums.PackingDiscrepancyKindShortage,
		RequestedComposition: packaging.PackageCompositionSnapshot{TotalBaseUnits: 2},
		RecordedAt:           settledAt.Add(-time.Hour),
		Settlements: []fulfilment.PackingShortageSettlement{{
			ID: "settlement-1", Reason: warehouse_enums.PackingShortageSettlementReasonNoStock, RefundRequestID: "refund-request-1",
			BaseUnits: 2, SettledBaseUnits: 2, Status: warehouse_enums.PackingShortageSettlementStatusApplied,
			ReleasedBaseUnits: 2, WrittenOffBaseUnits: 1, WriteOffMovementIDs: []string{"movement-1"},
			DepotAvailableBeforeBaseUnits: &before, DepotAvailableAfterBaseUnits: &after,
			SettledBy: "staff-1", SettledAt: settledAt, ReversedBaseUnits: 1, ReleasedReversedBaseUnits: 1,
			Reversals: []fulfilment.PackingShortageSettlementReversal{{
				BaseUnits: 1, Cause: warehouse_enums.PackingShortageSettlementReversalCauseScan, ReleasedBaseUnits: 1,
				WriteOffRestoredBaseUnits: 1, WriteOffReversalMovementIDs: []string{"movement-2"},
				ReversedBy: "staff-2", ReversedAt: settledAt.Add(time.Minute),
			}},
		}},
	}
	raw, err := json.Marshal(want)
	if err != nil {
		t.Fatal(err)
	}
	for _, fragment := range []string{
		`"kind":"shortage"`,
		`"settlements":[{"id":"settlement-1","reason":"no_stock","refund_request_id":"refund-request-1","base_units":2,"settled_base_units":2,"status":"applied","released_base_units":2,"written_off_base_units":1,"write_off_movement_ids":["movement-1"],"depot_available_before_base_units":-1,"depot_available_after_base_units":0,"settled_by":"staff-1","settled_at":"2026-11-13T09:00:00Z","reversed_base_units":1,"released_reversed_base_units":1,`,
		`"reversals":[{"base_units":1,"cause":"scan","released_base_units":1,"write_off_restored_base_units":1,"write_off_reversal_movement_ids":["movement-2"],"reversed_by":"staff-2","reversed_at":"2026-11-13T09:01:00Z"}]`,
	} {
		if !strings.Contains(string(raw), fragment) {
			t.Fatalf("settlement wire shape changed: missing %s in %s", fragment, raw)
		}
	}
	var got fulfilment.PackingDiscrepancy
	if err := json.Unmarshal(raw, &got); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got.Settlements, want.Settlements) {
		t.Fatalf("settlements lost: %+v", got.Settlements)
	}
	settlement := got.Settlements[0]
	if !settlement.Reason.IsValid() || !settlement.Status.IsValid() || !settlement.Reversals[0].Cause.IsValid() {
		t.Fatalf("settlement enums must validate: %+v", settlement)
	}

	// A deferred settlement carries its zero stock figures and no optional
	// evidence; a discrepancy without settlements omits the list.
	raw, err = json.Marshal(fulfilment.PackingShortageSettlement{ID: "settlement-2", Reason: warehouse_enums.PackingShortageSettlementReasonOverweight, RefundRequestID: "refund-request-2", BaseUnits: 1, SettledBaseUnits: 1, Status: warehouse_enums.PackingShortageSettlementStatusDeferred, DeferredReason: "tracked_units", SettledBy: "staff-1", SettledAt: settledAt})
	if err != nil {
		t.Fatal(err)
	}
	if string(raw) != `{"id":"settlement-2","reason":"overweight","refund_request_id":"refund-request-2","base_units":1,"settled_base_units":1,"status":"deferred","deferred_reason":"tracked_units","released_base_units":0,"written_off_base_units":0,"settled_by":"staff-1","settled_at":"2026-11-13T09:00:00Z"}` {
		t.Fatalf("deferred settlement wire shape changed: %s", raw)
	}
	raw, err = json.Marshal(fulfilment.PackingDiscrepancy{ID: "shortage-2"})
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(raw), "settlements") {
		t.Fatalf("an empty settlements list must be omitted: %s", raw)
	}
}
