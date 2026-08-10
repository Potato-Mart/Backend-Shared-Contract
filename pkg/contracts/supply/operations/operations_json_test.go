package operations_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/packaging/packaging_enums"
	event "github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/pubsub/event"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/supply/classification"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/supply/operations"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/supply/product/product_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/supply/warehouse"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/supply/warehouse/warehouse_enums"
)

func TestReservationAndStagingJSONShapes(t *testing.T) {
	dateMarkAt := time.Date(2026, 12, 1, 0, 0, 0, 0, time.UTC)
	caseComposition := composition(packaging_enums.PackageHandlingUnitCase, "pkg_case_12", 2, 12)

	reservationShape := marshalObject(t, operations.StockReservation{
		ID: "reservation_1", ProductSKUCode: "A00001", DepotCode: "AU-VIC-MEL-DC-01",
		PackagePricingID: "pricing_1", PackagePricingRevision: 9,
		RequestedComposition: caseComposition, ReservedComposition: caseComposition,
		Status:   warehouse_enums.StockReservationStatusReserved,
		Revision: 1, Timezone: "Australia/Melbourne",
	})
	if reservationShape["status"] != "RESERVED" || reservationShape["package_pricing_revision"] != float64(9) {
		t.Fatalf("reservation identity did not marshal: %+v", reservationShape)
	}

	stagingShape := marshalObject(t, operations.StockStagingRecord{
		ID: "staging_1", ReservationID: "reservation_1", AllocationID: "allocation_1",
		OrderNumber: "SO-1", ProductSKUCode: "A00001", PackageOptionID: "pkg_case_12",
		SourceBucketID: "bucket_1", DestinationBucketID: "bucket_stage",
		SourceLocation:      warehouse.StockLocationRef{DepotCode: "AU-VIC-MEL-DC-01", LocationCode: "A-01"},
		DestinationLocation: warehouse.StockLocationRef{DepotCode: "AU-VIC-MEL-DC-01", LocationCode: warehouse.StockLocationCodeOnlineStageAmbient},
		StagedComposition:   caseComposition, MovementID: "movement_1", StagedBy: "operator_1", StagedAt: dateMarkAt,
	})
	for _, key := range []string{"source_location", "destination_location", "staged_composition", "movement_id"} {
		if _, ok := stagingShape[key]; !ok {
			t.Fatalf("staging JSON missing %q: %+v", key, stagingShape)
		}
	}
}

func TestInventoryCategoryTagEvidenceJSONIsLocationQualified(t *testing.T) {
	now := time.Date(2026, 8, 9, 7, 8, 9, 0, time.UTC)
	evidence := operations.InventoryCategoryTagEvidence{
		ProductSKUCode:  "A00001",
		PackageOptionID: "pkg_each",
		CategoryTag:     classification.CategoryTagRef{ID: "tag_soon_expiry", Slug: "soon-expiry"},
		StockLocation:   warehouse.StockLocationRef{DepotCode: "AU-VIC-MEL-DC-01", LocationCode: "A-01-03"},
		Condition:       warehouse_enums.InventoryConditionGood,
		Disposition:     warehouse_enums.InventoryDispositionReducedSellable,
		DateMark:        &warehouse.InventoryDateMark{Kind: warehouse_enums.InventoryDateMarkBestBefore, DateMarkAt: now.Add(48 * time.Hour), Timezone: "Australia/Melbourne"},
		AsOf:            now,
	}

	shape := marshalObject(t, evidence)
	for _, key := range []string{"product_sku_code", "package_option_id", "category_tag", "stock_location", "condition", "disposition", "date_mark", "as_of"} {
		if _, ok := shape[key]; !ok {
			t.Fatalf("inventory category-tag evidence missing %q: %+v", key, shape)
		}
	}
	for _, forbidden := range []string{"promotion", "promotion_kind", "offer"} {
		if _, ok := shape[forbidden]; ok {
			t.Fatalf("inventory category-tag evidence must remain operational: %+v", shape)
		}
	}
}

func TestPackingPickingAndAvailabilityEventJSONShapes(t *testing.T) {
	now := time.Date(2026, 8, 4, 6, 0, 0, 0, time.UTC)
	requested := composition(packaging_enums.PackageHandlingUnitCase, "pkg_case_12", 1, 12)
	replacement := composition(packaging_enums.PackageHandlingUnitEach, "pkg_each", 12, 1)
	substitution := operations.PackageSubstitutionSnapshot{
		ID: "sub_1", RequestedCasePackageOptionID: "pkg_case_12", RequestedCaseCount: 1,
		RequestedUnitsPerCase: 12, FulfilledSealedCaseCount: 0,
		ReplacementEachPackageOptionID: "pkg_each", ReplacementBaseUnits: 12,
		LotID: "lot_1", SourceBucketID: "bucket_each", ReasonCode: "NO_SEALED_CASE",
		Operator: "operator_1", CapturedAt: now,
	}
	containerShape := marshalObject(t, operations.OutboundContainerPlan{
		ID: "container_1", ContainerCode: "OUT-1", StorageType: warehouse_enums.StorageAmbient,
		Contents: []operations.OutboundContainerContent{{
			OrderItemID: "item_1", ProductSKUCode: "A00001", AllocationID: "allocation_1",
			BucketID: "bucket_each", LotID: "lot_1", PackageOptionID: "pkg_each",
			PackedComposition: replacement, Substitutions: []operations.PackageSubstitutionSnapshot{substitution},
		}},
		UpdatedAt: now,
	})
	if containerShape["storage_type"] != "AMBIENT" {
		t.Fatalf("outbound container omitted storage type: %+v", containerShape)
	}
	if _, ok := containerShape["contents"]; !ok {
		t.Fatalf("outbound container omitted package contents: %+v", containerShape)
	}

	lineShape := marshalObject(t, operations.PackingLine{
		ID: "packing_line_1", OrderItemID: "item_1", ProductSKUCode: "A00001",
		RequestedComposition: requested, AllocatedComposition: replacement,
		PickedComposition: replacement, PackedComposition: replacement,
		SubstitutedComposition: replacement, ReturnedComposition: composition(packaging_enums.PackageHandlingUnitEach, "pkg_each", 0, 1),
		RefundedComposition: composition(packaging_enums.PackageHandlingUnitEach, "pkg_each", 0, 1),
		Substitutions:       []operations.PackageSubstitutionSnapshot{substitution},
	})
	for _, key := range []string{"requested_composition", "allocated_composition", "picked_composition", "packed_composition", "substituted_composition", "returned_composition", "refunded_composition"} {
		if _, ok := lineShape[key]; !ok {
			t.Fatalf("packing line JSON missing %q: %+v", key, lineShape)
		}
	}

	pickingShape := marshalObject(t, operations.PickingListItem{
		ID:                     "picking_item_1",
		PickingListID:          "picking_list_1",
		OrderItemID:            "item_1",
		ProductSKUCode:         "A00001",
		RequestedComposition:   requested,
		AllocatedComposition:   replacement,
		PickedComposition:      replacement,
		PackedComposition:      replacement,
		SubstitutedComposition: replacement,
		ReturnedComposition:    composition(packaging_enums.PackageHandlingUnitEach, "pkg_each", 0, 1),
		RefundedComposition:    composition(packaging_enums.PackageHandlingUnitEach, "pkg_each", 0, 1),
		Status:                 warehouse_enums.PickingItemStatusComplete,
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

	eventShape := marshalObject(t, event.StockLocationAvailabilityChangedEvent{
		AssignmentID: "assignment_1", DepotCode: "AU-VIC-MEL-DC-01", LocationCode: "A-01",
		ProductSKUCode: "A00001", AvailableBeforeBaseUnits: 12, AvailableAfterBaseUnits: 0,
		Direction: warehouse_enums.StockAvailabilityOutOfStock,
		Cause:     operations.InventoryCauseRef{Type: "SALE_COMMIT", ID: "movement_1"},
		Revision:  8, OccurredAt: now, AsOf: now,
	})
	if eventShape["direction"] != "OUT_OF_STOCK" || eventShape["revision"] != float64(8) {
		t.Fatalf("availability event lost direction or revision: %+v", eventShape)
	}
}

func TestStorefrontPlaceAvailabilityJSONIsCustomerSafe(t *testing.T) {
	now := time.Date(2026, 8, 11, 3, 4, 5, 0, time.UTC)
	availability := operations.StorefrontPlaceAvailability{
		DepotCode:              "AU-VIC-MEL-DC-01",
		DepotName:              "Melbourne Distribution Centre",
		RegionCode:             "AU-VIC-MEL",
		RegionName:             "Greater Melbourne",
		CountryCode:            geography.CountryCode("AU"),
		AdministrativeAreaCode: geography.SubdivisionCode("AU-VIC"),
		Locality:               "Melbourne",
		StockState:             product_enums.StorefrontStockStateInStock,
		AvailableBaseUnits:     24,
		AsOf:                   now,
	}

	shape := marshalObject(t, availability)
	for _, key := range []string{"depot_code", "depot_name", "region_code", "region_name", "country_code", "administrative_area_code", "locality", "stock_state", "available_base_units", "as_of"} {
		if _, ok := shape[key]; !ok {
			t.Fatalf("storefront place availability missing %q: %+v", key, shape)
		}
	}
	if shape["stock_state"] != "in_stock" || shape["available_base_units"] != float64(24) {
		t.Fatalf("storefront place availability lost stock state or quantity: %+v", shape)
	}
	for _, forbidden := range []string{"location_code", "lot_id", "bucket_id", "reserved_base_units", "staged_base_units", "quality_hold_base_units", "on_hand_base_units"} {
		if _, ok := shape[forbidden]; ok {
			t.Fatalf("storefront place availability must stay customer-safe: %+v", shape)
		}
	}

	raw, err := json.Marshal(availability)
	if err != nil {
		t.Fatalf("marshal storefront place availability: %v", err)
	}
	var decoded operations.StorefrontPlaceAvailability
	if err := json.Unmarshal(raw, &decoded); err != nil {
		t.Fatalf("unmarshal storefront place availability: %v", err)
	}
	if decoded.DepotCode != availability.DepotCode || decoded.DepotName != availability.DepotName ||
		decoded.RegionCode != availability.RegionCode || decoded.RegionName != availability.RegionName ||
		decoded.CountryCode != availability.CountryCode ||
		decoded.AdministrativeAreaCode != availability.AdministrativeAreaCode ||
		decoded.Locality != availability.Locality || decoded.StockState != availability.StockState ||
		decoded.AvailableBaseUnits != availability.AvailableBaseUnits || !decoded.AsOf.Equal(now) {
		t.Fatalf("storefront place availability did not round-trip: %+v", decoded)
	}

	minimalShape := marshalObject(t, operations.StorefrontPlaceAvailability{
		DepotCode:  "AU-VIC-MEL-DC-01",
		DepotName:  "Melbourne Distribution Centre",
		StockState: product_enums.StorefrontStockStateOutOfStock,
		AsOf:       now,
	})
	for _, optional := range []string{"region_code", "region_name", "country_code", "administrative_area_code", "locality"} {
		if _, ok := minimalShape[optional]; ok {
			t.Fatalf("zero-value optional field %q must be omitted: %+v", optional, minimalShape)
		}
	}
	for _, required := range []string{"depot_code", "depot_name", "stock_state", "available_base_units", "as_of"} {
		if _, ok := minimalShape[required]; !ok {
			t.Fatalf("required field %q must always marshal: %+v", required, minimalShape)
		}
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
