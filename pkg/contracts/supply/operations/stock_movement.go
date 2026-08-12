package operations

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/supply/warehouse/warehouse_enums"
)

// InventoryCauseRef identifies the contract record that caused an inventory
// change.
type InventoryCauseRef struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

// StockMovement represents a physical base-unit transfer or state change.
// Logical reservation lifecycle changes are represented by StockReservation.
type StockMovement struct {
	ID                               string                                `json:"id"`
	SKUID                            string                                `json:"sku_id"`
	Type                             warehouse_enums.StockMovementType     `json:"type"`
	SourceBucketID                   string                                `json:"source_bucket_id,omitempty"`
	DestinationBucketID              string                                `json:"destination_bucket_id,omitempty"`
	LotID                            string                                `json:"lot_id,omitempty"`
	SourcePackageOptionID            string                                `json:"source_package_option_id,omitempty"`
	DestinationPackageOptionID       string                                `json:"destination_package_option_id,omitempty"`
	BaseUnits                        int64                                 `json:"base_units"`
	SourcePackageComposition         *packaging.PackageCompositionSnapshot `json:"source_package_composition,omitempty"`
	DestinationPackageComposition    *packaging.PackageCompositionSnapshot `json:"destination_package_composition,omitempty"`
	SourceBalanceAfterBaseUnits      *int64                                `json:"source_balance_after_base_units,omitempty"`
	DestinationBalanceAfterBaseUnits *int64                                `json:"destination_balance_after_base_units,omitempty"`
	Cause                            *InventoryCauseRef                    `json:"cause,omitempty"`
	PurchaseOrderNumber              string                                `json:"purchase_order_number,omitempty"`
	PurchaseReceiptID                string                                `json:"purchase_receipt_id,omitempty"`
	OrderNumber                      string                                `json:"order_number,omitempty"`
	ReservationID                    string                                `json:"reservation_id,omitempty"`
	AllocationID                     string                                `json:"allocation_id,omitempty"`
	StagingRecordID                  string                                `json:"staging_record_id,omitempty"`
	QualityAssessmentID              string                                `json:"quality_assessment_id,omitempty"`
	DamageReportID                   string                                `json:"damage_report_id,omitempty"`
	ReasonCode                       string                                `json:"reason_code,omitempty"`
	Note                             string                                `json:"note,omitempty"`
	PerformedBy                      string                                `json:"performed_by,omitempty"`
	OccurredAt                       time.Time                             `json:"occurred_at"`
}
