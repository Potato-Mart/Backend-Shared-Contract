package fulfilment_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v35/pkg/contracts/common/packaging"
	contractfulfilment "github.com/Potato-Mart/Backend-Shared-Contract/v35/pkg/contracts/supply/fulfilment"
	"github.com/Potato-Mart/Backend-Shared-Contract/v35/pkg/contracts/supply/inventory"
	"github.com/Potato-Mart/Backend-Shared-Contract/v35/pkg/contracts/supply/warehouse/warehouse_enums"
)

func TestOrderAllocationEvidenceRoundTripsAuthoritativeRecords(t *testing.T) {
	composition := packaging.PackageCompositionSnapshot{TotalBaseUnits: 2}
	evidence := contractfulfilment.OrderAllocationEvidence{
		OrderItemID: "item_1",
		Reservation: inventory.StockReservation{
			ID: "reservation_1", OrderNumber: "SO-1", SKUCode: "A00001",
			Status: warehouse_enums.StockReservationStatusStaged, Revision: 4,
			RequestedComposition: composition, ReservedComposition: composition,
		},
		Allocation: inventory.StockReservationAllocation{
			ID: "allocation_1", ReservationID: "reservation_1", BucketID: "bucket_1",
			AllocatedComposition: composition, Revision: 2,
		},
		Picking: &contractfulfilment.PickingAllocation{
			ReservationAllocationID: "allocation_1", AllocatedComposition: composition,
			PickedComposition: composition,
		},
		Staging: &inventory.StockStagingRecord{
			ID: "staging_1", ReservationID: "reservation_1", AllocationID: "allocation_1",
			OrderNumber: "SO-1", SKUCode: "A00001", StagedComposition: composition,
		},
	}

	raw, err := json.Marshal(evidence)
	if err != nil {
		t.Fatalf("marshal order allocation evidence: %v", err)
	}
	for _, key := range []string{`"order_item_id":"item_1"`, `"reservation"`, `"allocation"`, `"picking"`, `"staging"`} {
		if !strings.Contains(string(raw), key) {
			t.Fatalf("order allocation evidence JSON missing %s: %s", key, raw)
		}
	}
	var decoded contractfulfilment.OrderAllocationEvidence
	if err := json.Unmarshal(raw, &decoded); err != nil {
		t.Fatalf("unmarshal order allocation evidence: %v", err)
	}
	if decoded.OrderItemID != evidence.OrderItemID || decoded.Reservation.ID != evidence.Reservation.ID ||
		decoded.Allocation.ID != evidence.Allocation.ID || decoded.Picking == nil ||
		decoded.Picking.ReservationAllocationID != evidence.Allocation.ID || decoded.Staging == nil ||
		decoded.Staging.AllocationID != evidence.Allocation.ID {
		t.Fatalf("order allocation evidence lost its cross-record links: %+v", decoded)
	}
}

func TestPickingListAllocationFingerprintIsOptionalAndRoundTrips(t *testing.T) {
	list := contractfulfilment.PickingList{AllocationFingerprint: "frozen-set-1"}
	raw, err := json.Marshal(list)
	if err != nil {
		t.Fatalf("marshal picking list: %v", err)
	}
	if !strings.Contains(string(raw), `"allocation_fingerprint":"frozen-set-1"`) {
		t.Fatalf("picking list JSON missing allocation fingerprint: %s", raw)
	}
	var decoded contractfulfilment.PickingList
	if err := json.Unmarshal(raw, &decoded); err != nil {
		t.Fatalf("unmarshal picking list: %v", err)
	}
	if decoded.AllocationFingerprint != list.AllocationFingerprint {
		t.Fatalf("allocation fingerprint = %q, want %q", decoded.AllocationFingerprint, list.AllocationFingerprint)
	}

	legacyRaw, err := json.Marshal(contractfulfilment.PickingList{})
	if err != nil {
		t.Fatalf("marshal legacy picking list: %v", err)
	}
	if strings.Contains(string(legacyRaw), `"allocation_fingerprint"`) {
		t.Fatalf("empty allocation fingerprint must be omitted: %s", legacyRaw)
	}
}
