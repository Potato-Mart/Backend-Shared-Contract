package warehouse_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/packaging/packaging_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/supply/warehouse"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/supply/warehouse/warehouse_enums"
)

func TestLotBucketJSONShapes(t *testing.T) {
	dateMarkAt := time.Date(2026, 12, 1, 0, 0, 0, 0, time.UTC)
	dateMark := warehouse.InventoryDateMark{
		Kind:       warehouse_enums.InventoryDateMarkExpiry,
		DateMarkAt: dateMarkAt,
		Timezone:   "Australia/Melbourne",
	}
	lotShape := marshalObject(t, warehouse.InventoryLot{
		ID: "lot_1", SKUID: "A00001", ReceivedAt: dateMarkAt.Add(-30 * 24 * time.Hour),
		DateMark: &dateMark,
	})
	nestedDateMark := lotShape["date_mark"].(map[string]any)
	if nestedDateMark["date_mark_at"] != "2026-12-01T00:00:00Z" || nestedDateMark["timezone"] != "Australia/Melbourne" {
		t.Fatalf("lot date mark lost UTC instant or timezone: %+v", nestedDateMark)
	}

	caseComposition := composition(packaging_enums.PackageHandlingUnitCase, "pkg_case_12", 2, 12)
	bucketShape := marshalObject(t, warehouse.InventoryStockBucket{
		ID: "bucket_1", Location: warehouse.StockLocationRef{DepotCode: "AU-VIC-MEL-DC-01", LocationCode: "A-01"},
		SKUID: "A00001", LotID: "lot_1", PackageOptionID: "pkg_case_12",
		HandlingUnit:       packaging_enums.PackageHandlingUnitCase,
		Condition:          warehouse_enums.InventoryConditionGood,
		Disposition:        warehouse_enums.InventoryDispositionStandardSellable,
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
		ID: "unit_1", BucketID: "bucket_1", SKUID: "A00001", LotID: "lot_1",
		PackageOptionID: "pkg_each", HandlingUnit: packaging_enums.PackageHandlingUnitEach, BaseUnits: 1,
		Condition:     warehouse_enums.InventoryConditionPackagingDamagedMinor,
		Disposition:   warehouse_enums.InventoryDispositionReducedSellable,
		UnitLabelCode: "UNIT-1", ClearanceLabelCode: "CLEARANCE-1",
	})
	if unitShape["clearance_label_code"] != "CLEARANCE-1" {
		t.Fatalf("stock unit did not keep clearance label identity separate: %+v", unitShape)
	}
	if _, ok := unitShape["electronic_shelf_label_code"]; ok {
		t.Fatalf("stock unit conflated clearance and electronic shelf labels: %+v", unitShape)
	}
}

func composition(unit packaging_enums.PackageHandlingUnit, optionID string, count, unitsPerPackage int64) packaging.PackageCompositionSnapshot {
	baseUnits := count * unitsPerPackage
	return packaging.PackageCompositionSnapshot{
		TotalBaseUnits: baseUnits,
		Components: []packaging.PackageComponentSnapshot{{
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
