package supply_test

import (
	"encoding/json"
	"reflect"

	event "github.com/Potato-Mart/Backend-Shared-Contract/v51/pkg/contracts/pubsub/supply"
	"github.com/Potato-Mart/Backend-Shared-Contract/v51/pkg/contracts/pubsub/supply/supply_enums"
	operations "github.com/Potato-Mart/Backend-Shared-Contract/v51/pkg/contracts/supply/inventory"
	"github.com/Potato-Mart/Backend-Shared-Contract/v51/pkg/contracts/supply/warehouse"

	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v51/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v51/pkg/contracts/common/packaging/packaging_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v51/pkg/contracts/supply/warehouse/warehouse_enums"
)

func TestInventoryEventJSONShapes(t *testing.T) {
	now := time.Date(2026, 8, 4, 7, 8, 9, 0, time.UTC)
	caseComposition := composition(packaging_enums.PackageHandlingUnitCase, "pkg_case_12", 1, 12)
	eachComposition := composition(packaging_enums.PackageHandlingUnitEach, "pkg_each", 12, 1)
	cause := operations.InventoryCauseRef{Type: "SALE_COMMIT", ID: "movement_1"}
	location := warehouse.StockLocationRef{DepotCode: "AU-VIC-MEL-DC-01", LocationCode: "A-01-03"}

	cases := []struct {
		name     string
		value    any
		required []string
	}{
		{
			name: "lot received",
			value: event.InventoryLotReceivedEvent{
				LotID: "lot_1", SKUCode: "A00001", DepotCode: location.DepotCode,
				DestinationBucketID: "bucket_case", ReceivedComposition: caseComposition,
				MovementID: "movement_receipt", LotRevision: 2, ReceivedAt: now,
			},
			required: []string{"lot_id", "received_composition", "lot_revision"},
		},
		{
			name: "bucket changed",
			value: event.InventoryStockBucketChangedEvent{
				BucketID: "bucket_case", Location: location, SKUCode: "A00001", LotID: "lot_1",
				PackageOption: packaging.PackageOptionRef{SKUCode: "A00001", Code: "CASE6", Version: 1}, HandlingUnit: packaging_enums.PackageHandlingUnitCase,
				Condition: warehouse_enums.InventoryConditionGood, Disposition: warehouse_enums.InventoryDispositionStandardSellable,
				OnHandBeforeBaseUnits: 12, OnHandAfterBaseUnits: 0, AvailableBeforeBaseUnits: 12,
				AvailableAfterBaseUnits: 0, Cause: cause, Revision: 3,
			},
			required: []string{"bucket_id", "location", "available_before_base_units", "available_after_base_units", "revision"},
		},
		{
			name: "package converted",
			value: event.InventoryPackageConvertedEvent{
				MovementID: "movement_conversion", SKUCode: "A00001", DepotCode: location.DepotCode, LotID: "lot_1",
				SourceBucketID: "bucket_case", DestinationBucketID: "bucket_each",
				SourcePackageOption: packaging.PackageOptionRef{SKUCode: "A00001", Code: "CASE6", Version: 1}, DestinationPackageOption: packaging.PackageOptionRef{SKUCode: "A00001", Code: "EACH", Version: 1},
				BaseUnits: 12, SourcePackageComposition: caseComposition, DestinationPackageComposition: eachComposition,
				SourceBucketRevision: 4, DestinationBucketRevision: 2,
			},
			required: []string{"depot_code", "source_bucket_id", "destination_bucket_id", "source_bucket_revision", "destination_bucket_revision"},
		},
		{
			name: "quality assessed",
			value: event.InventoryQualityAssessedEvent{
				QualityAssessmentID: "assessment_1", SKUCode: "A00001", DepotCode: location.DepotCode, BucketID: "bucket_each",
				AssessedComposition: eachComposition, PreviousCondition: warehouse_enums.InventoryConditionGood,
				ResultCondition:     warehouse_enums.InventoryConditionPackagingDamagedMinor,
				PreviousDisposition: warehouse_enums.InventoryDispositionStandardSellable,
				ResultDisposition:   warehouse_enums.InventoryDispositionReducedSellable,
				MovementIDs:         []string{"movement_quality"}, Revision: 2,
			},
			required: []string{"quality_assessment_id", "depot_code", "previous_condition", "result_disposition", "revision"},
		},
		{
			name: "reservation changed",
			value: event.InventoryReservationChangedEvent{
				ReservationID: "reservation_1", SKUCode: "A00001", DepotCode: location.DepotCode,
				PreviousStatus: warehouse_enums.StockReservationStatusReserved, Status: warehouse_enums.StockReservationStatusStaged,
				RequestedComposition: caseComposition, ReservedComposition: caseComposition, Revision: 4,
			},
			required: []string{"reservation_id", "previous_status", "status", "requested_composition", "revision"},
		},
		{
			name: "staging changed",
			value: event.StockStagingChangedEvent{
				StagingRecordID: "staging_1", ReservationID: "reservation_1", AllocationID: "allocation_1",
				OrderNumber: "SO-1", SKUCode: "A00001", SourceLocation: location,
				DestinationLocation: warehouse.StockLocationRef{DepotCode: location.DepotCode, LocationCode: "SYS-ONLINE-STAGE-AMBIENT"},
				StagedComposition:   caseComposition, MovementID: "movement_stage", Revision: 1,
			},
			required: []string{"staging_record_id", "source_location", "destination_location", "revision"},
		},
		{
			name: "sale committed",
			value: event.InventorySaleCommittedEvent{
				MovementID: "movement_sale", OrderNumber: "SO-1", DepotCode: location.DepotCode, ReservationID: "reservation_1",
				AllocationID: "allocation_1", BucketID: "bucket_case", SKUCode: "A00001",
				KnownLotPhysical: &operations.KnownLotPhysicalEvidence{LotID: "lot_1"}, PackageOption: packaging.PackageOptionRef{SKUCode: "A00001", Code: "CASE6", Version: 1}, CommittedComposition: caseComposition,
				InventoryRevision: 5, CostStatus: supply_enums.InventorySaleCostStatusValued,
			},
			required: []string{"order_number", "depot_code", "allocation_id", "committed_composition", "inventory_revision"},
		},
		{
			name: "date mark threshold",
			value: event.InventoryDateMarkThresholdEvent{
				LotID: "lot_1", SKUCode: "A00001", DepotCode: location.DepotCode,
				DateMark:  warehouse.InventoryDateMark{Kind: warehouse_enums.InventoryDateMarkBestBefore, DateMarkAt: now, Timezone: "Australia/Melbourne"},
				Threshold: warehouse_enums.InventoryDateMarkThresholdApproaching, ThresholdAt: now, LotRevision: 3,
			},
			required: []string{"date_mark", "threshold", "threshold_at", "lot_revision"},
		},
		{
			name: "stock location availability changed",
			value: event.StockLocationAvailabilityChangedEvent{
				AssignmentID: "assignment_1", DepotCode: location.DepotCode, LocationCode: location.LocationCode,
				SKUCode: "A00001", AvailableBeforeBaseUnits: 12, AvailableAfterBaseUnits: 0,
				Direction: warehouse_enums.StockAvailabilityOutOfStock, Cause: cause, Revision: 4, AsOf: now,
			},
			required: []string{"assignment_id", "available_before_base_units", "available_after_base_units", "as_of"},
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			shape := marshalObject(t, testCase.value)
			for _, key := range testCase.required {
				if _, ok := shape[key]; !ok {
					t.Fatalf("%T JSON missing %q: %+v", testCase.value, key, shape)
				}
			}
			if _, ok := shape["occurred_at"]; ok {
				t.Fatalf("%T payload must rely on envelope occurred_at: %+v", testCase.value, shape)
			}
		})
	}
}

func composition(unit packaging_enums.PackageHandlingUnit, optionID string, count, unitsPerPackage int64) packaging.PackageCompositionSnapshot {
	return packaging.PackageCompositionSnapshot{
		TotalBaseUnits: count * unitsPerPackage,
		Components: []packaging.PackageComponentSnapshot{{
			PackageOption:   packaging.PackageOptionRef{SKUCode: "A00001", Code: optionID, Version: 1},
			HandlingUnit:    unit,
			PackageCount:    count,
			UnitsPerPackage: unitsPerPackage,
			BaseUnits:       count * unitsPerPackage,
		}},
	}
}

func marshalObject(t *testing.T, value any) map[string]any {
	t.Helper()
	payload, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("marshal %T: %v", value, err)
	}
	var object map[string]any
	if err := json.Unmarshal(payload, &object); err != nil {
		t.Fatalf("unmarshal %T: %v", value, err)
	}
	return object
}

func TestPOSSaleEventMovementEvidenceVariants(t *testing.T) {
	tests := []struct {
		name    string
		variant string
		cost    supply_enums.POSSaleCostStatus
		known   *event.POSKnownLotPhysicalEvidence
		unknown *event.POSUnknownIdentityPhysicalEvidence
		deficit *event.POSDeficitEvidence
	}{
		{name: "known lot valued", variant: "known_lot_physical", cost: supply_enums.POSSaleCostStatusValued, known: &event.POSKnownLotPhysicalEvidence{LotID: "lot-1"}},
		{name: "known lot unresolved", variant: "known_lot_physical", cost: supply_enums.POSSaleCostStatusUnresolved, known: &event.POSKnownLotPhysicalEvidence{LotID: "lot-1"}},
		{name: "unidentified return stock", variant: "unknown_identity_physical", cost: supply_enums.POSSaleCostStatusUnresolved, unknown: &event.POSUnknownIdentityPhysicalEvidence{ReturnSources: []operations.UnknownReturnSource{{ReceivingID: "return-1", BaseUnits: 1}, {ReceivingID: "return-2", BaseUnits: 1}}}},
		{name: "unidentified deficit", variant: "deficit", cost: supply_enums.POSSaleCostStatusUnresolved, deficit: &event.POSDeficitEvidence{}},
		{name: "evidenced lot deficit", variant: "deficit", cost: supply_enums.POSSaleCostStatusUnresolved, deficit: &event.POSDeficitEvidence{LotID: "lot-1"}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			value := event.InventoryPOSSaleCommittedEvent{
				SaleID: "sale-1", OrderNumber: "order-1", FulfillmentGeneration: 1, OrderItemID: "item-1",
				MovementID: "movement-1", DepotCode: "depot-1", LocationCode: "location-1", BucketID: "bucket-1", SKUCode: "A00001",
				PackageOption:        packaging.PackageOptionRef{SKUCode: "A00001", Code: "EACH", Version: 1},
				CommittedComposition: composition(packaging_enums.PackageHandlingUnitEach, "EACH", 2, 1), BaseUnits: 2, InventoryRevision: 7,
				CostStatus: tc.cost, KnownLotPhysical: tc.known, UnknownIdentityPhysical: tc.unknown, Deficit: tc.deficit,
			}
			raw, err := json.Marshal(value)
			if err != nil {
				t.Fatal(err)
			}
			var fields map[string]json.RawMessage
			if err := json.Unmarshal(raw, &fields); err != nil {
				t.Fatal(err)
			}
			for _, name := range []string{"sale_id", "order_number", "fulfillment_generation", "order_item_id", "movement_id", "depot_code", "location_code", "bucket_id", "sku_code", "package_option", "committed_composition", "base_units", "inventory_revision", "cost_status", tc.variant} {
				if _, ok := fields[name]; !ok {
					t.Errorf("missing %s: %s", name, raw)
				}
			}
			for _, name := range []string{"reservation_id", "allocation_id", "is_deficit", "lot_id"} {
				if _, ok := fields[name]; ok {
					t.Errorf("retired/invented %s: %s", name, raw)
				}
			}
			for _, name := range []string{"known_lot_physical", "unknown_identity_physical", "deficit"} {
				if _, ok := fields[name]; ok != (name == tc.variant) {
					t.Errorf("incorrect variant %s: %s", name, raw)
				}
			}
			var decoded event.InventoryPOSSaleCommittedEvent
			if err := json.Unmarshal(raw, &decoded); err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(value, decoded) {
				t.Fatalf("round trip changed movement: %+v", decoded)
			}
			if tc.unknown != nil {
				var arm map[string]json.RawMessage
				if err := json.Unmarshal(fields[tc.variant], &arm); err != nil {
					t.Fatal(err)
				}
				if _, ok := arm["lot_id"]; ok {
					t.Fatalf("unknown identity invented a lot: %s", raw)
				}
				if _, ok := arm["unit_id"]; ok {
					t.Fatalf("unknown identity invented a unit: %s", raw)
				}
			}
			if tc.deficit != nil {
				var arm map[string]json.RawMessage
				if err := json.Unmarshal(fields[tc.variant], &arm); err != nil {
					t.Fatal(err)
				}
				_, exists := arm["lot_id"]
				if exists != (tc.deficit.LotID != "") {
					t.Fatalf("deficit lot was not preserved accurately: %s", raw)
				}
			}
		})
	}
}

func TestOnlineSalePhysicalEvidenceVariants(t *testing.T) {
	tests := []struct {
		name, variant string
		cost          supply_enums.InventorySaleCostStatus
		known         *operations.KnownLotPhysicalEvidence
		unknown       *operations.UnknownReturnPhysicalEvidence
		unallocated   *operations.UnallocatedNoLotPhysicalEvidence
	}{
		{name: "known lot valued", variant: "known_lot_physical", cost: supply_enums.InventorySaleCostStatusValued, known: &operations.KnownLotPhysicalEvidence{LotID: "lot-1", StockUnitIDs: []string{"unit-1", "unit-2"}}},
		{name: "known lot unresolved", variant: "known_lot_physical", cost: supply_enums.InventorySaleCostStatusUnresolved, known: &operations.KnownLotPhysicalEvidence{LotID: "lot-1"}},
		{name: "unknown returns", variant: "unknown_identity_physical", cost: supply_enums.InventorySaleCostStatusUnresolved, unknown: &operations.UnknownReturnPhysicalEvidence{ReturnSources: []operations.UnknownReturnSource{{ReceivingID: "return-1", BaseUnits: 1}, {ReceivingID: "return-2", BaseUnits: 1}}}},
		{name: "unallocated no lot unresolved", variant: "unallocated_no_lot_physical", cost: supply_enums.InventorySaleCostStatusUnresolved, unallocated: &operations.UnallocatedNoLotPhysicalEvidence{SourceBucketID: "original-bucket"}},
		{name: "unallocated no lot valued", variant: "unallocated_no_lot_physical", cost: supply_enums.InventorySaleCostStatusValued, unallocated: &operations.UnallocatedNoLotPhysicalEvidence{SourceBucketID: "original-bucket"}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			value := event.InventorySaleCommittedEvent{MovementID: "movement", OrderNumber: "order", DepotCode: "depot", ReservationID: "reservation", AllocationID: "allocation", BucketID: "staged-bucket", SKUCode: "A00001", PackageOption: packaging.PackageOptionRef{SKUCode: "A00001", Code: "EACH", Version: 1}, CommittedComposition: composition(packaging_enums.PackageHandlingUnitEach, "EACH", 2, 1), InventoryRevision: 7, CostStatus: tc.cost, KnownLotPhysical: tc.known, UnknownIdentityPhysical: tc.unknown, UnallocatedNoLotPhysical: tc.unallocated}
			raw, err := json.Marshal(value)
			if err != nil {
				t.Fatal(err)
			}
			var fields map[string]json.RawMessage
			if err := json.Unmarshal(raw, &fields); err != nil {
				t.Fatal(err)
			}
			for _, name := range []string{"movement_id", "order_number", "reservation_id", "allocation_id", "bucket_id", "committed_composition", "cost_status", "inventory_revision", tc.variant} {
				if _, ok := fields[name]; !ok {
					t.Errorf("missing %s: %s", name, raw)
				}
			}
			for _, name := range []string{"lot_id", "stock_unit_ids", "deficit", "is_deficit"} {
				if _, ok := fields[name]; ok {
					t.Errorf("obsolete or invented %s: %s", name, raw)
				}
			}
			for _, name := range []string{"known_lot_physical", "unknown_identity_physical", "unallocated_no_lot_physical"} {
				if _, ok := fields[name]; ok != (name == tc.variant) {
					t.Errorf("incorrect variant %s: %s", name, raw)
				}
			}
			var got event.InventorySaleCommittedEvent
			if err := json.Unmarshal(raw, &got); err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(value, got) {
				t.Fatalf("round trip changed evidence: %+v", got)
			}
		})
	}
}
