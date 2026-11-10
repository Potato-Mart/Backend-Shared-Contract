package procurement_test

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
	"time"

	purchase "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/procurement"
)

func TestPurchaseOrderWarehouseConfirmationAndCasePlanRoundTrip(t *testing.T) {
	bare, err := json.Marshal(purchase.PurchaseOrder{ID: "po_1", Items: []purchase.PurchaseOrderItem{{SKUCode: "A00001"}}})
	if err != nil {
		t.Fatalf("marshal bare order: %v", err)
	}
	for _, key := range []string{`"warehouse_confirmed_at"`, `"warehouse_confirmed_by"`, `"case_plan"`} {
		if strings.Contains(string(bare), key) {
			t.Fatalf("empty %s must be omitted: %s", key, bare)
		}
	}

	confirmedAt := time.Date(2026, 11, 11, 1, 30, 0, 0, time.UTC)
	plan := []purchase.PurchaseOrderCasePlanEntry{
		{LocationCode: "C-01-01", Cartons: 6},
		{DepotCode: "D3", LocationCode: "C-02-04", Cartons: 2},
	}
	raw, err := json.Marshal(purchase.PurchaseOrder{
		ID: "po_1", WarehouseConfirmedAt: &confirmedAt, WarehouseConfirmedBy: "staff_1",
		Items: []purchase.PurchaseOrderItem{{SKUCode: "A00001", CasePlan: plan}},
	})
	if err != nil {
		t.Fatalf("marshal order: %v", err)
	}
	var back purchase.PurchaseOrder
	if err := json.Unmarshal(raw, &back); err != nil {
		t.Fatalf("unmarshal order: %v", err)
	}
	if back.WarehouseConfirmedAt == nil || !back.WarehouseConfirmedAt.Equal(confirmedAt) || back.WarehouseConfirmedBy != "staff_1" {
		t.Fatalf("warehouse confirmation did not round-trip: %s", raw)
	}
	if back.ConfirmedAt != nil {
		t.Fatalf("warehouse confirmation must not set confirmed_at: %s", raw)
	}
	if len(back.Items) != 1 || !reflect.DeepEqual(back.Items[0].CasePlan, plan) {
		t.Fatalf("case plan did not round-trip: %s", raw)
	}

	entry, err := json.Marshal(plan[0])
	if err != nil {
		t.Fatalf("marshal case plan entry: %v", err)
	}
	if string(entry) != `{"location_code":"C-01-01","cartons":6}` {
		t.Fatalf("case plan entry = %s", entry)
	}
}
