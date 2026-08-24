package warehouse_test

import (
	"encoding/json"

	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/supply/classification/classification_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/supply/product"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/supply/warehouse"

	"reflect"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/packaging/packaging_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/supply/warehouse/warehouse_enums"
)

func TestStockLocationAndBalanceJSONShapes(t *testing.T) {
	location := warehouse.StockLocation{
		ID: "location_1", DepotCode: "AU-VIC-MEL-DC-01",
		LocationCode:    warehouse.StockLocationCodeOnlineStageFrozen,
		StorageType:     classification_enums.StorageFrozen,
		Purpose:         warehouse_enums.StockLocationPurposeOnlineOrderStaging,
		HandlingMode:    warehouse_enums.StockLocationHandlingMixed,
		Access:          warehouse_enums.StockLocationAccessStaffOnly,
		CollectionMode:  warehouse_enums.StockLocationCollectionUnrestricted,
		IsSystemManaged: true, IsActive: true,
	}
	shape := marshalStockLocationObject(t, location)
	for _, key := range []string{"location_code", "storage_type", "purpose", "handling_mode", "access", "collection_mode", "is_system_managed"} {
		if _, ok := shape[key]; !ok {
			t.Fatalf("stock location JSON missing %q: %+v", key, shape)
		}
	}
	for _, key := range []string{"code", "zone"} {
		if _, ok := shape[key]; ok {
			t.Fatalf("stock location JSON retained %q: %+v", key, shape)
		}
	}

	zero := int64(0)
	balanceShape := marshalStockLocationObject(t, warehouse.StockLocationProductBalance{
		AssignmentID: "assignment_1", DepotCode: location.DepotCode,
		LocationCode: location.LocationCode, SKUCode: "A00001",
		PackageComposition: stockLocationComposition(packaging_enums.PackageHandlingUnitEach, "pkg_each", zero, 1),
		OnHandBaseUnits:    zero, ReservedBaseUnits: zero, AvailableBaseUnits: zero,
		IsOutOfStock: true, Revision: 7,
		DepotTimezone: "Australia/Melbourne", AsOf: time.Date(2026, 8, 4, 5, 0, 0, 0, time.UTC),
	})
	for _, key := range []string{"on_hand_base_units", "reserved_base_units", "available_base_units"} {
		if _, ok := balanceShape[key]; !ok {
			t.Fatalf("known zero quantity %q was omitted: %+v", key, balanceShape)
		}
	}
}

func TestRequiredSystemStockLocationCodesJSON(t *testing.T) {
	tests := []struct {
		name        string
		code        string
		wantCode    string
		storageType classification_enums.StorageType
		purpose     warehouse_enums.StockLocationPurpose
	}{
		{name: "quality hold ambient", code: warehouse.StockLocationCodeQualityHoldAmbient, wantCode: "SYS-QH-AMBIENT", storageType: classification_enums.StorageAmbient, purpose: warehouse_enums.StockLocationPurposeQualityHold},
		{name: "quality hold chilled", code: warehouse.StockLocationCodeQualityHoldChilled, wantCode: "SYS-QH-CHILLED", storageType: classification_enums.StorageChilled, purpose: warehouse_enums.StockLocationPurposeQualityHold},
		{name: "quality hold frozen", code: warehouse.StockLocationCodeQualityHoldFrozen, wantCode: "SYS-QH-FROZEN", storageType: classification_enums.StorageFrozen, purpose: warehouse_enums.StockLocationPurposeQualityHold},
		{name: "online stage ambient", code: warehouse.StockLocationCodeOnlineStageAmbient, wantCode: "SYS-ONLINE-STAGE-AMBIENT", storageType: classification_enums.StorageAmbient, purpose: warehouse_enums.StockLocationPurposeOnlineOrderStaging},
		{name: "online stage chilled", code: warehouse.StockLocationCodeOnlineStageChilled, wantCode: "SYS-ONLINE-STAGE-CHILLED", storageType: classification_enums.StorageChilled, purpose: warehouse_enums.StockLocationPurposeOnlineOrderStaging},
		{name: "online stage frozen", code: warehouse.StockLocationCodeOnlineStageFrozen, wantCode: "SYS-ONLINE-STAGE-FROZEN", storageType: classification_enums.StorageFrozen, purpose: warehouse_enums.StockLocationPurposeOnlineOrderStaging},
	}

	seen := make(map[string]struct{}, len(tests))
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.code != tt.wantCode {
				t.Fatalf("system location code = %q, want %q", tt.code, tt.wantCode)
			}
			if _, duplicate := seen[tt.code]; duplicate {
				t.Fatalf("duplicate system location code %q", tt.code)
			}
			seen[tt.code] = struct{}{}

			shape := marshalStockLocationObject(t, warehouse.StockLocation{
				ID:              "location_" + tt.code,
				DepotCode:       "AU-VIC-MEL-DC-01",
				LocationCode:    tt.code,
				StorageType:     tt.storageType,
				Purpose:         tt.purpose,
				HandlingMode:    warehouse_enums.StockLocationHandlingMixed,
				Access:          warehouse_enums.StockLocationAccessStaffOnly,
				CollectionMode:  warehouse_enums.StockLocationCollectionUnrestricted,
				IsSystemManaged: true,
				IsActive:        true,
			})
			if shape["location_code"] != tt.wantCode {
				t.Fatalf("system location JSON code = %v, want %q: %+v", shape["location_code"], tt.wantCode, shape)
			}
			if shape["storage_type"] != tt.storageType.String() || shape["purpose"] != tt.purpose.String() {
				t.Fatalf("system location JSON lost storage type or purpose: %+v", shape)
			}
			if shape["is_system_managed"] != true {
				t.Fatalf("system location JSON was not system managed: %+v", shape)
			}
		})
	}
}

func TestStockLocationAssignmentReplacesProductPlacement(t *testing.T) {
	if _, ok := reflect.TypeOf(product.Product{}).FieldByName("PlacingArea"); ok {
		t.Fatal("Product.PlacingArea survived the stock-location assignment cut-over")
	}

	shape := marshalStockLocationObject(t, warehouse.StockLocationAssignment{
		ID:                       "assignment_1",
		DepotCode:                "AU-VIC-MEL-DC-01",
		LocationCode:             "A-01-03",
		SKUCode:                  "A00001",
		ElectronicShelfLabelCode: "ESL-001",
		IsActive:                 true,
	})
	for _, key := range []string{"id", "depot_code", "location_code", "sku_code", "electronic_shelf_label_code", "is_active"} {
		if _, ok := shape[key]; !ok {
			t.Fatalf("stock location assignment JSON missing %q: %+v", key, shape)
		}
	}
}

func stockLocationComposition(unit packaging_enums.PackageHandlingUnit, optionID string, count, unitsPerPackage int64) packaging.PackageCompositionSnapshot {
	baseUnits := count * unitsPerPackage
	return packaging.PackageCompositionSnapshot{
		TotalBaseUnits: baseUnits,
		Components: []packaging.PackageComponentSnapshot{{
			PackageOptionCode: optionID,
			HandlingUnit:      unit,
			PackageCount:      count,
			UnitsPerPackage:   unitsPerPackage,
			BaseUnits:         baseUnits,
		}},
	}
}

func marshalStockLocationObject(t *testing.T, value any) map[string]any {
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
