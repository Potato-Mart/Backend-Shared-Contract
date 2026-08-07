package warehouse_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/contracts/warehouse"
	warehouseenum "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/enums/warehouse"
)

func TestLotBucketReservationAndStagingJSONShapes(t *testing.T) {
	dateMarkAt := time.Date(2026, 12, 1, 0, 0, 0, 0, time.UTC)
	dateMark := warehouse.InventoryDateMark{
		Kind:       warehouseenum.InventoryDateMarkExpiry,
		DateMarkAt: dateMarkAt,
		Timezone:   "Australia/Melbourne",
	}
	lotShape := marshalObject(t, warehouse.InventoryLot{
		ID: "lot_1", ProductSKUCode: "A00001", ReceivedAt: dateMarkAt.Add(-30 * 24 * time.Hour),
		DateMark: &dateMark,
	})
	nestedDateMark := lotShape["date_mark"].(map[string]any)
	if nestedDateMark["date_mark_at"] != "2026-12-01T00:00:00Z" || nestedDateMark["timezone"] != "Australia/Melbourne" {
		t.Fatalf("lot date mark lost UTC instant or timezone: %+v", nestedDateMark)
	}

	caseComposition := composition(common.PackageHandlingUnitCase, "pkg_case_12", 2, 12)
	bucketShape := marshalObject(t, warehouse.InventoryStockBucket{
		ID: "bucket_1", Location: warehouse.StockLocationRef{DepotCode: "AU-VIC-MEL-DC-01", LocationCode: "A-01"},
		ProductSKUCode: "A00001", LotID: "lot_1", PackageOptionID: "pkg_case_12",
		HandlingUnit:       common.PackageHandlingUnitCase,
		Condition:          warehouseenum.InventoryConditionGood,
		Disposition:        warehouseenum.InventoryDispositionStandardSellable,
		PackageComposition: caseComposition,
		OnHandBaseUnits:    24, ReservedBaseUnits: 12, AvailableBaseUnits: 12,
		Revision: 3, DepotTimezone: "Australia/Melbourne", AsOf: dateMarkAt,
	})
	if bucketShape["handling_unit"] != "CASE" {
		t.Fatalf("case bucket did not preserve physical package form: %+v", bucketShape)
	}
	if _, ok := bucketShape["available_base_units"]; !ok {
		t.Fatalf("bucket omitted derived availability projection: %+v", bucketShape)
	}
	unitShape := marshalObject(t, warehouse.InventoryStockUnit{
		ID: "unit_1", BucketID: "bucket_1", ProductSKUCode: "A00001", LotID: "lot_1",
		PackageOptionID: "pkg_each", HandlingUnit: common.PackageHandlingUnitEach, BaseUnits: 1,
		Condition:     warehouseenum.InventoryConditionPackagingDamagedMinor,
		Disposition:   warehouseenum.InventoryDispositionReducedSellable,
		UnitLabelCode: "UNIT-1", ClearanceLabelCode: "CLEARANCE-1",
	})
	if unitShape["clearance_label_code"] != "CLEARANCE-1" {
		t.Fatalf("stock unit did not keep clearance label identity separate: %+v", unitShape)
	}
	if _, ok := unitShape["electronic_shelf_label_code"]; ok {
		t.Fatalf("stock unit conflated clearance and electronic shelf labels: %+v", unitShape)
	}

	reservationShape := marshalObject(t, warehouse.StockReservation{
		ID: "reservation_1", ProductSKUCode: "A00001", DepotCode: "AU-VIC-MEL-DC-01",
		OfferID: "offer_1", OfferRevision: 9,
		RequestedComposition: caseComposition, ReservedComposition: caseComposition,
		Status:   warehouseenum.StockReservationStatusReserved,
		Revision: 1, Timezone: "Australia/Melbourne",
	})
	if reservationShape["status"] != "RESERVED" || reservationShape["offer_revision"] != float64(9) {
		t.Fatalf("reservation identity did not marshal: %+v", reservationShape)
	}

	stagingShape := marshalObject(t, warehouse.StockStagingRecord{
		ID: "staging_1", ReservationID: "reservation_1", AllocationID: "allocation_1",
		OrderNumber: "SO-1", ProductSKUCode: "A00001", PackageOptionID: "pkg_case_12",
		SourceBucketID: "bucket_1", DestinationBucketID: "bucket_stage",
		SourceLocation:      warehouse.StockLocationRef{DepotCode: "AU-VIC-MEL-DC-01", LocationCode: "A-01"},
		DestinationLocation: warehouse.StockLocationRef{DepotCode: "AU-VIC-MEL-DC-01", LocationCode: warehouse.StockLocationCodeOnlineStageDry},
		StagedComposition:   caseComposition, MovementID: "movement_1", StagedBy: "operator_1", StagedAt: dateMarkAt,
	})
	for _, key := range []string{"source_location", "destination_location", "staged_composition", "movement_id"} {
		if _, ok := stagingShape[key]; !ok {
			t.Fatalf("staging JSON missing %q: %+v", key, stagingShape)
		}
	}
}

func TestPackageAwarePackingAndAvailabilityEventJSONShapes(t *testing.T) {
	now := time.Date(2026, 8, 4, 6, 0, 0, 0, time.UTC)
	requested := composition(common.PackageHandlingUnitCase, "pkg_case_12", 1, 12)
	replacement := composition(common.PackageHandlingUnitEach, "pkg_each", 12, 1)
	substitution := warehouse.PackageSubstitutionSnapshot{
		ID: "sub_1", RequestedCasePackageOptionID: "pkg_case_12", RequestedCaseCount: 1,
		RequestedUnitsPerCase: 12, FulfilledSealedCaseCount: 0,
		ReplacementEachPackageOptionID: "pkg_each", ReplacementBaseUnits: 12,
		LotID: "lot_1", SourceBucketID: "bucket_each", ReasonCode: "NO_SEALED_CASE",
		Operator: "operator_1", CapturedAt: now,
	}
	containerShape := marshalObject(t, warehouse.OutboundContainerPlan{
		ID: "container_1", ContainerCode: "OUT-1", StorageType: warehouseenum.StorageDry,
		Contents: []warehouse.OutboundContainerContent{{
			OrderItemID: "item_1", ProductSKUCode: "A00001", AllocationID: "allocation_1",
			BucketID: "bucket_each", LotID: "lot_1", PackageOptionID: "pkg_each",
			PackedComposition: replacement, Substitutions: []warehouse.PackageSubstitutionSnapshot{substitution},
		}},
		UpdatedAt: now,
	})
	if containerShape["storage_type"] != "DRY" {
		t.Fatalf("outbound container omitted storage type: %+v", containerShape)
	}
	if _, ok := containerShape["contents"]; !ok {
		t.Fatalf("outbound container omitted package contents: %+v", containerShape)
	}

	lineShape := marshalObject(t, warehouse.PackingLine{
		ID: "packing_line_1", OrderItemID: "item_1", ProductSKUCode: "A00001",
		RequestedComposition: requested, AllocatedComposition: replacement,
		PickedComposition: replacement, PackedComposition: replacement,
		SubstitutedComposition: replacement, ReturnedComposition: composition(common.PackageHandlingUnitEach, "pkg_each", 0, 1),
		RefundedComposition: composition(common.PackageHandlingUnitEach, "pkg_each", 0, 1),
		Substitutions:       []warehouse.PackageSubstitutionSnapshot{substitution},
	})
	for _, key := range []string{"requested_composition", "allocated_composition", "picked_composition", "packed_composition", "substituted_composition", "returned_composition", "refunded_composition"} {
		if _, ok := lineShape[key]; !ok {
			t.Fatalf("packing line JSON missing %q: %+v", key, lineShape)
		}
	}

	pickingShape := marshalObject(t, warehouse.PickingListItem{
		ID:                     "picking_item_1",
		PickingListID:          "picking_list_1",
		OrderItemID:            "item_1",
		ProductSKUCode:         "A00001",
		RequestedComposition:   requested,
		AllocatedComposition:   replacement,
		PickedComposition:      replacement,
		PackedComposition:      replacement,
		SubstitutedComposition: replacement,
		ReturnedComposition:    composition(common.PackageHandlingUnitEach, "pkg_each", 0, 1),
		RefundedComposition:    composition(common.PackageHandlingUnitEach, "pkg_each", 0, 1),
		Status:                 warehouseenum.PickingItemStatusComplete,
		CreatedAt:              now,
	})
	for _, key := range []string{"requested_composition", "allocated_composition", "picked_composition", "packed_composition", "substituted_composition", "returned_composition", "refunded_composition"} {
		compositionShape, ok := pickingShape[key].(map[string]any)
		if !ok {
			t.Fatalf("picking line JSON missing composition object %q: %+v", key, pickingShape)
		}
		if _, ok := compositionShape["total_base_units"]; !ok {
			t.Fatalf("picking line composition %q omitted total_base_units: %+v", key, compositionShape)
		}
	}

	eventShape := marshalObject(t, warehouse.StockLocationAvailabilityChangedEvent{
		AssignmentID: "assignment_1", DepotCode: "AU-VIC-MEL-DC-01", LocationCode: "A-01",
		ProductSKUCode: "A00001", AvailableBeforeBaseUnits: 12, AvailableAfterBaseUnits: 0,
		Direction: warehouseenum.StockAvailabilityOutOfStock,
		Cause:     warehouse.InventoryCauseRef{Type: "SALE_COMMIT", ID: "movement_1"},
		Revision:  8, OccurredAt: now, AsOf: now,
	})
	if eventShape["direction"] != "OUT_OF_STOCK" || eventShape["revision"] != float64(8) {
		t.Fatalf("availability event lost direction or revision: %+v", eventShape)
	}
}

func composition(unit common.PackageHandlingUnit, optionID string, count, unitsPerPackage int64) common.PackageCompositionSnapshot {
	baseUnits := count * unitsPerPackage
	return common.PackageCompositionSnapshot{
		TotalBaseUnits: baseUnits,
		Components: []common.PackageComponentSnapshot{{
			PackageOptionID: optionID,
			HandlingUnit:    unit,
			PackageCount:    count,
			UnitsPerPackage: unitsPerPackage,
			BaseUnits:       baseUnits,
		}},
	}
}

func marshalObject(t *testing.T, value any) map[string]any {
	t.Helper()
	raw, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("marshal %T: %v", value, err)
	}
	var shape map[string]any
	if err := json.Unmarshal(raw, &shape); err != nil {
		t.Fatalf("unmarshal %T JSON: %v", value, err)
	}
	return shape
}
