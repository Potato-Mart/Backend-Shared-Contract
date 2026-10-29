package inventory_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v52/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v52/pkg/contracts/supply/inventory"
)

func TestReservationAllocationPhysicalEvidenceWire(t *testing.T) {
	tests := []struct {
		name       string
		allocation inventory.StockReservationAllocation
	}{
		{name: "known lot", allocation: inventory.StockReservationAllocation{KnownLotPhysical: &inventory.KnownLotPhysicalEvidence{LotID: "lot-1", StockUnitIDs: []string{"unit-1", "unit-2"}}}},
		{name: "unknown returns", allocation: inventory.StockReservationAllocation{UnknownIdentityPhysical: &inventory.UnknownReturnPhysicalEvidence{ReturnSources: []inventory.UnknownReturnSource{{ReceivingID: "return-1", BaseUnits: 1}, {ReceivingID: "return-2", BaseUnits: 1}}}}},
		{name: "unallocated no lot", allocation: inventory.StockReservationAllocation{UnallocatedNoLotPhysical: &inventory.UnallocatedNoLotPhysicalEvidence{SourceBucketID: "source"}}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.allocation.ID, tc.allocation.ReservationID, tc.allocation.BucketID = "allocation", "reservation", "source"
			tc.allocation.AllocatedComposition = packaging.PackageCompositionSnapshot{TotalBaseUnits: 2}
			tc.allocation.Revision = 1
			raw, err := json.Marshal(tc.allocation)
			if err != nil {
				t.Fatal(err)
			}
			var shape map[string]json.RawMessage
			if err := json.Unmarshal(raw, &shape); err != nil {
				t.Fatal(err)
			}
			for _, key := range []string{"lot_id", "stock_unit_ids", "deficit"} {
				if _, ok := shape[key]; ok {
					t.Fatalf("unexpected top-level %s: %s", key, raw)
				}
			}
			count := 0
			for _, key := range []string{"known_lot_physical", "unknown_identity_physical", "unallocated_no_lot_physical"} {
				if _, ok := shape[key]; ok {
					count++
				}
			}
			if count != 1 {
				t.Fatalf("expected one physical variant: %s", raw)
			}
			var got inventory.StockReservationAllocation
			if err := json.Unmarshal(raw, &got); err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(tc.allocation, got) {
				t.Fatalf("round trip changed allocation: %+v", got)
			}
		})
	}
}
