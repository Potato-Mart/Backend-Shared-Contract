package fulfilment_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v48/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v48/pkg/contracts/orders/order/order_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v48/pkg/contracts/supply/fulfilment"
)

func TestPackingRecordRoundTrip(t *testing.T) {
	now := time.Date(2026, 10, 29, 0, 0, 0, 0, time.UTC)
	want := fulfilment.PackingRecord{
		ID: "packing-1", OrderNumber: "ORDER-1", FulfillmentGeneration: 2,
		Revision: 3, Status: order_enums.FulfillmentStatusPartial, Operator: "staff-1",
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
	if got.OrderNumber != want.OrderNumber || got.FulfillmentGeneration != 2 || got.Revision != 3 || got.Operator != "staff-1" || got.Status != order_enums.FulfillmentStatusPartial || len(got.Lines) != 1 || got.Lines[0].RequestedComposition.TotalBaseUnits != 4 || got.Lines[0].PackedComposition.TotalBaseUnits != 3 || len(got.Containers) != 1 || len(got.Discrepancies) != 1 || got.Discrepancies[0].OrderItemID != "item-1" || got.StartedAt == nil || got.PackedAt == nil {
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
