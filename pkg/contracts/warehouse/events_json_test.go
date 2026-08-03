package warehouse_test

import (
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/contracts/warehouse"
	warehouseenum "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/enums/warehouse"
)

func TestV22InventoryEventJSONShapes(t *testing.T) {
	now := time.Date(2026, 8, 4, 7, 8, 9, 0, time.UTC)
	caseComposition := composition(common.PackageHandlingUnitCase, "pkg_case_12", 1, 12)
	eachComposition := composition(common.PackageHandlingUnitEach, "pkg_each", 12, 1)
	cause := warehouse.InventoryCauseRef{Type: "SALE_COMMIT", ID: "movement_1"}
	location := warehouse.StockLocationRef{DepotCode: "AU-VIC-MEL-DC-01", LocationCode: "A-01-03"}

	cases := []struct {
		name     string
		value    any
		required []string
	}{
		{
			name: "lot received",
			value: warehouse.InventoryLotReceivedEvent{
				LotID: "lot_1", ProductSKUCode: "A00001", DepotCode: location.DepotCode,
				DestinationBucketID: "bucket_case", ReceivedComposition: caseComposition,
				MovementID: "movement_receipt", LotRevision: 2, ReceivedAt: now, OccurredAt: now,
			},
			required: []string{"lot_id", "received_composition", "lot_revision", "occurred_at"},
		},
		{
			name: "bucket changed",
			value: warehouse.InventoryStockBucketChangedEvent{
				BucketID: "bucket_case", Location: location, ProductSKUCode: "A00001", LotID: "lot_1",
				PackageOptionID: "pkg_case_12", HandlingUnit: common.PackageHandlingUnitCase,
				Condition: warehouseenum.InventoryConditionGood, Disposition: warehouseenum.InventoryDispositionStandardSellable,
				OnHandBeforeBaseUnits: 12, OnHandAfterBaseUnits: 0, AvailableBeforeBaseUnits: 12,
				AvailableAfterBaseUnits: 0, Cause: cause, Revision: 3, OccurredAt: now,
			},
			required: []string{"bucket_id", "location", "available_before_base_units", "available_after_base_units", "revision"},
		},
		{
			name: "package converted",
			value: warehouse.InventoryPackageConvertedEvent{
				MovementID: "movement_conversion", ProductSKUCode: "A00001", LotID: "lot_1",
				SourceBucketID: "bucket_case", DestinationBucketID: "bucket_each",
				SourcePackageOptionID: "pkg_case_12", DestinationPackageOptionID: "pkg_each",
				BaseUnits: 12, SourcePackageComposition: caseComposition, DestinationPackageComposition: eachComposition,
				SourceBucketRevision: 4, DestinationBucketRevision: 2, OccurredAt: now,
			},
			required: []string{"source_bucket_id", "destination_bucket_id", "source_bucket_revision", "destination_bucket_revision", "occurred_at"},
		},
		{
			name: "quality assessed",
			value: warehouse.InventoryQualityAssessedEvent{
				QualityAssessmentID: "assessment_1", ProductSKUCode: "A00001", BucketID: "bucket_each",
				AssessedComposition: eachComposition, PreviousCondition: warehouseenum.InventoryConditionGood,
				ResultCondition:     warehouseenum.InventoryConditionPackagingDamagedMinor,
				PreviousDisposition: warehouseenum.InventoryDispositionStandardSellable,
				ResultDisposition:   warehouseenum.InventoryDispositionReducedSellable,
				MovementIDs:         []string{"movement_quality"}, Revision: 2, OccurredAt: now,
			},
			required: []string{"quality_assessment_id", "previous_condition", "result_disposition", "revision"},
		},
		{
			name: "reservation changed",
			value: warehouse.InventoryReservationChangedEvent{
				ReservationID: "reservation_1", ProductSKUCode: "A00001", DepotCode: location.DepotCode,
				PreviousStatus: warehouseenum.StockReservationStatusReserved, Status: warehouseenum.StockReservationStatusStaged,
				RequestedComposition: caseComposition, ReservedComposition: caseComposition, Revision: 4, OccurredAt: now,
			},
			required: []string{"reservation_id", "previous_status", "status", "requested_composition", "revision"},
		},
		{
			name: "staging changed",
			value: warehouse.StockStagingChangedEvent{
				StagingRecordID: "staging_1", ReservationID: "reservation_1", AllocationID: "allocation_1",
				OrderNumber: "SO-1", ProductSKUCode: "A00001", SourceLocation: location,
				DestinationLocation: warehouse.StockLocationRef{DepotCode: location.DepotCode, LocationCode: warehouse.StockLocationCodeOnlineStageDry},
				StagedComposition:   caseComposition, MovementID: "movement_stage", Revision: 1, OccurredAt: now,
			},
			required: []string{"staging_record_id", "source_location", "destination_location", "revision"},
		},
		{
			name: "sale committed",
			value: warehouse.InventorySaleCommittedEvent{
				MovementID: "movement_sale", OrderNumber: "SO-1", ReservationID: "reservation_1",
				AllocationID: "allocation_1", BucketID: "bucket_case", ProductSKUCode: "A00001",
				LotID: "lot_1", PackageOptionID: "pkg_case_12", CommittedComposition: caseComposition,
				InventoryRevision: 5, OccurredAt: now,
			},
			required: []string{"order_number", "allocation_id", "committed_composition", "inventory_revision"},
		},
		{
			name: "date mark threshold",
			value: warehouse.InventoryDateMarkThresholdEvent{
				LotID: "lot_1", ProductSKUCode: "A00001", DepotCode: location.DepotCode,
				DateMark:  warehouse.InventoryDateMark{Kind: warehouseenum.InventoryDateMarkBestBefore, DateMarkAt: now, Timezone: "Australia/Melbourne"},
				Threshold: warehouseenum.InventoryDateMarkThresholdApproaching, ThresholdAt: now, LotRevision: 3, OccurredAt: now,
			},
			required: []string{"date_mark", "threshold", "threshold_at", "lot_revision"},
		},
		{
			name: "offer availability changed",
			value: warehouse.SellableOfferAvailabilityChangedEvent{
				OfferID: "offer_1", ProductSKUCode: "A00001", DepotCode: location.DepotCode,
				SourceBucketID: "bucket_case", AvailableBeforeBaseUnits: 12, AvailableAfterBaseUnits: 0,
				InventoryRevision: 5, OfferRevision: 6, OccurredAt: now,
			},
			required: []string{"offer_id", "available_before_base_units", "available_after_base_units", "offer_revision"},
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
			if occurredAt, ok := shape["occurred_at"]; ok && occurredAt != "2026-08-04T07:08:09Z" {
				t.Fatalf("%T occurred_at = %v, want UTC instant", testCase.value, occurredAt)
			}
		})
	}
}
