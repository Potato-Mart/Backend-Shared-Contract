package supply_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v54/pkg/contracts/common/packaging"
	pe "github.com/Potato-Mart/Backend-Shared-Contract/v54/pkg/contracts/common/packaging/packaging_enums"
	event "github.com/Potato-Mart/Backend-Shared-Contract/v54/pkg/contracts/pubsub/supply"
	enums "github.com/Potato-Mart/Backend-Shared-Contract/v54/pkg/contracts/pubsub/supply/supply_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v54/pkg/contracts/supply/inventory"
)

func TestUnidentifiedTrackedSaleDoesNotInventPhysicalSource(t *testing.T) {
	option := packaging.PackageOptionRef{SKUCode: "SKU-1", Code: "EACH", Version: 1}
	quantity := packaging.PackageCompositionSnapshot{TotalBaseUnits: 2, Components: []packaging.PackageComponentSnapshot{{PackageOption: option, HandlingUnit: pe.PackageHandlingUnitEach, PackageCount: 2, UnitsPerPackage: 1, BaseUnits: 2}}}
	value := event.InventorySaleCommittedEvent{MovementID: "movement-1", OrderNumber: "order-1", DepotCode: "depot-1", ReservationID: "reservation-1", AllocationID: "allocation-1", SKUCode: "SKU-1", PackageOption: option, CommittedComposition: quantity, CostStatus: enums.InventorySaleCostStatusUnresolved, UnidentifiedTrackedPhysical: &inventory.UnidentifiedTrackedPhysicalEvidence{}}
	if err := value.ValidatePhysicalEvidence(); err != nil {
		t.Fatal(err)
	}
	raw, err := json.Marshal(value)
	if err != nil {
		t.Fatal(err)
	}
	var fields map[string]json.RawMessage
	if err = json.Unmarshal(raw, &fields); err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"bucket_id", "inventory_revision", "known_lot_physical", "unknown_identity_physical", "unallocated_no_lot_physical", "lot_id", "stock_unit_ids"} {
		if _, ok := fields[name]; ok {
			t.Fatalf("invented %s: %s", name, raw)
		}
	}
	if string(fields["unidentified_tracked_physical"]) != "{}" {
		t.Fatalf("missing explicit unresolved variant: %s", raw)
	}
	var roundtrip event.InventorySaleCommittedEvent
	if err = json.Unmarshal(raw, &roundtrip); err != nil || roundtrip.ValidatePhysicalEvidence() != nil {
		t.Fatalf("round trip: %+v %v", roundtrip, err)
	}
	for name, mutate := range map[string]func(*event.InventorySaleCommittedEvent){
		"bucket":   func(v *event.InventorySaleCommittedEvent) { v.BucketID = "invented" },
		"revision": func(v *event.InventorySaleCommittedEvent) { v.InventoryRevision = 1 },
		"value":    func(v *event.InventorySaleCommittedEvent) { v.CostStatus = enums.InventorySaleCostStatusValued },
		"second variant": func(v *event.InventorySaleCommittedEvent) {
			v.KnownLotPhysical = &inventory.KnownLotPhysicalEvidence{LotID: "lot"}
		},
		"zero quantity": func(v *event.InventorySaleCommittedEvent) { v.CommittedComposition.TotalBaseUnits = 0 },
	} {
		t.Run(name, func(t *testing.T) {
			bad := value
			mutate(&bad)
			if bad.ValidatePhysicalEvidence() == nil {
				t.Fatal("invalid evidence accepted")
			}
		})
	}
	forged := strings.Replace(string(raw), `"unidentified_tracked_physical":{}`, `"unidentified_tracked_physical":{"unit_id":"made-up"}`, 1)
	if err = json.Unmarshal([]byte(forged), &roundtrip); err == nil {
		t.Fatal("forged identity field accepted")
	}
}

func TestSaleVerifiedSourceRequiresBucketAndRevision(t *testing.T) {
	option := packaging.PackageOptionRef{SKUCode: "SKU-1", Code: "EACH", Version: 1}
	value := event.InventorySaleCommittedEvent{MovementID: "m", OrderNumber: "o", DepotCode: "d", ReservationID: "r", AllocationID: "a", SKUCode: "SKU-1", PackageOption: option, CommittedComposition: packaging.PackageCompositionSnapshot{TotalBaseUnits: 1, Components: []packaging.PackageComponentSnapshot{{PackageOption: option, HandlingUnit: pe.PackageHandlingUnitEach, PackageCount: 1, UnitsPerPackage: 1, BaseUnits: 1}}}, CostStatus: enums.InventorySaleCostStatusUnresolved, KnownLotPhysical: &inventory.KnownLotPhysicalEvidence{LotID: "lot"}}
	if value.ValidatePhysicalEvidence() == nil {
		t.Fatal("missing source accepted")
	}
	value.BucketID = "bucket"
	value.InventoryRevision = 1
	if err := value.ValidatePhysicalEvidence(); err != nil {
		t.Fatal(err)
	}
}
