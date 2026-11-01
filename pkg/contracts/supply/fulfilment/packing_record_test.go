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
