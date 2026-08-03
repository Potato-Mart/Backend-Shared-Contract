package warehouse

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/common"
	warehouseenum "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/enums/warehouse"
)

// StockLocationAvailabilityChangedEvent represents a customer-accessible
// standard-location zero crossing. AssignmentID is its ordering key.
type StockLocationAvailabilityChangedEvent struct {
	AssignmentID             string                                   `json:"assignment_id"`
	DepotCode                string                                   `json:"depot_code"`
	LocationCode             string                                   `json:"location_code"`
	ProductSKUCode           string                                   `json:"product_sku_code"`
	AvailableBeforeBaseUnits int64                                    `json:"available_before_base_units"`
	AvailableAfterBaseUnits  int64                                    `json:"available_after_base_units"`
	Direction                warehouseenum.StockAvailabilityDirection `json:"direction"`
	ElectronicShelfLabelCode string                                   `json:"electronic_shelf_label_code,omitempty"`
	Cause                    InventoryCauseRef                        `json:"cause"`
	Revision                 int64                                    `json:"revision"`
	OccurredAt               time.Time                                `json:"occurred_at"`
	AsOf                     time.Time                                `json:"as_of"`
}

type InventoryLotReceivedEvent struct {
	LotID               string                            `json:"lot_id"`
	ProductSKUCode      string                            `json:"product_sku_code"`
	DepotCode           string                            `json:"depot_code"`
	DestinationBucketID string                            `json:"destination_bucket_id"`
	ReceivedComposition common.PackageCompositionSnapshot `json:"received_composition"`
	MovementID          string                            `json:"movement_id"`
	LotRevision         int64                             `json:"lot_revision"`
	ReceivedAt          time.Time                         `json:"received_at"`
	OccurredAt          time.Time                         `json:"occurred_at"`
}

type InventoryStockBucketChangedEvent struct {
	BucketID                 string                             `json:"bucket_id"`
	Location                 StockLocationRef                   `json:"location"`
	ProductSKUCode           string                             `json:"product_sku_code"`
	LotID                    string                             `json:"lot_id,omitempty"`
	PackageOptionID          string                             `json:"package_option_id"`
	HandlingUnit             common.PackageHandlingUnit         `json:"handling_unit"`
	Condition                warehouseenum.InventoryCondition   `json:"condition"`
	Disposition              warehouseenum.InventoryDisposition `json:"disposition"`
	OnHandBeforeBaseUnits    int64                              `json:"on_hand_before_base_units"`
	OnHandAfterBaseUnits     int64                              `json:"on_hand_after_base_units"`
	ReservedBeforeBaseUnits  int64                              `json:"reserved_before_base_units"`
	ReservedAfterBaseUnits   int64                              `json:"reserved_after_base_units"`
	AvailableBeforeBaseUnits int64                              `json:"available_before_base_units"`
	AvailableAfterBaseUnits  int64                              `json:"available_after_base_units"`
	Cause                    InventoryCauseRef                  `json:"cause"`
	Revision                 int64                              `json:"revision"`
	OccurredAt               time.Time                          `json:"occurred_at"`
}

type InventoryPackageConvertedEvent struct {
	MovementID                    string                            `json:"movement_id"`
	ProductSKUCode                string                            `json:"product_sku_code"`
	DepotCode                     string                            `json:"depot_code"`
	LotID                         string                            `json:"lot_id"`
	SourceBucketID                string                            `json:"source_bucket_id"`
	DestinationBucketID           string                            `json:"destination_bucket_id"`
	SourcePackageOptionID         string                            `json:"source_package_option_id"`
	DestinationPackageOptionID    string                            `json:"destination_package_option_id"`
	BaseUnits                     int64                             `json:"base_units"`
	SourcePackageComposition      common.PackageCompositionSnapshot `json:"source_package_composition"`
	DestinationPackageComposition common.PackageCompositionSnapshot `json:"destination_package_composition"`
	SourceBucketRevision          int64                             `json:"source_bucket_revision"`
	DestinationBucketRevision     int64                             `json:"destination_bucket_revision"`
	OccurredAt                    time.Time                         `json:"occurred_at"`
}

type InventoryQualityAssessedEvent struct {
	QualityAssessmentID string                             `json:"quality_assessment_id"`
	ProductSKUCode      string                             `json:"product_sku_code"`
	DepotCode           string                             `json:"depot_code"`
	BucketID            string                             `json:"bucket_id"`
	StockUnitID         string                             `json:"stock_unit_id,omitempty"`
	AssessedComposition common.PackageCompositionSnapshot  `json:"assessed_composition"`
	PreviousCondition   warehouseenum.InventoryCondition   `json:"previous_condition"`
	ResultCondition     warehouseenum.InventoryCondition   `json:"result_condition"`
	PreviousDisposition warehouseenum.InventoryDisposition `json:"previous_disposition"`
	ResultDisposition   warehouseenum.InventoryDisposition `json:"result_disposition"`
	MovementIDs         []string                           `json:"movement_ids,omitempty"`
	Revision            int64                              `json:"revision"`
	OccurredAt          time.Time                          `json:"occurred_at"`
}

type InventoryReservationChangedEvent struct {
	ReservationID        string                               `json:"reservation_id"`
	ProductSKUCode       string                               `json:"product_sku_code"`
	DepotCode            string                               `json:"depot_code"`
	PreviousStatus       warehouseenum.StockReservationStatus `json:"previous_status,omitempty"`
	Status               warehouseenum.StockReservationStatus `json:"status"`
	RequestedComposition common.PackageCompositionSnapshot    `json:"requested_composition"`
	ReservedComposition  common.PackageCompositionSnapshot    `json:"reserved_composition"`
	Revision             int64                                `json:"revision"`
	OccurredAt           time.Time                            `json:"occurred_at"`
}

type StockStagingChangedEvent struct {
	StagingRecordID     string                            `json:"staging_record_id"`
	ReservationID       string                            `json:"reservation_id"`
	AllocationID        string                            `json:"allocation_id"`
	OrderNumber         string                            `json:"order_number"`
	ProductSKUCode      string                            `json:"product_sku_code"`
	SourceLocation      StockLocationRef                  `json:"source_location"`
	DestinationLocation StockLocationRef                  `json:"destination_location"`
	StagedComposition   common.PackageCompositionSnapshot `json:"staged_composition"`
	MovementID          string                            `json:"movement_id"`
	Revision            int64                             `json:"revision"`
	OccurredAt          time.Time                         `json:"occurred_at"`
}

type InventorySaleCommittedEvent struct {
	MovementID           string                            `json:"movement_id"`
	OrderNumber          string                            `json:"order_number"`
	DepotCode            string                            `json:"depot_code"`
	ReservationID        string                            `json:"reservation_id"`
	AllocationID         string                            `json:"allocation_id"`
	BucketID             string                            `json:"bucket_id"`
	ProductSKUCode       string                            `json:"product_sku_code"`
	LotID                string                            `json:"lot_id,omitempty"`
	PackageOptionID      string                            `json:"package_option_id"`
	CommittedComposition common.PackageCompositionSnapshot `json:"committed_composition"`
	InventoryRevision    int64                             `json:"inventory_revision"`
	OccurredAt           time.Time                         `json:"occurred_at"`
}

type InventoryDateMarkThresholdEvent struct {
	LotID          string                                   `json:"lot_id"`
	ProductSKUCode string                                   `json:"product_sku_code"`
	DepotCode      string                                   `json:"depot_code"`
	DateMark       InventoryDateMark                        `json:"date_mark"`
	Threshold      warehouseenum.InventoryDateMarkThreshold `json:"threshold"`
	ThresholdAt    time.Time                                `json:"threshold_at"`
	LotRevision    int64                                    `json:"lot_revision"`
	OccurredAt     time.Time                                `json:"occurred_at"`
}

type SellableOfferAvailabilityChangedEvent struct {
	OfferID                  string    `json:"offer_id"`
	ProductSKUCode           string    `json:"product_sku_code"`
	DepotCode                string    `json:"depot_code"`
	SourceBucketID           string    `json:"source_bucket_id,omitempty"`
	SourceStockUnitID        string    `json:"source_stock_unit_id,omitempty"`
	AvailableBeforeBaseUnits int64     `json:"available_before_base_units"`
	AvailableAfterBaseUnits  int64     `json:"available_after_base_units"`
	InventoryRevision        int64     `json:"inventory_revision"`
	OfferRevision            int64     `json:"offer_revision"`
	OccurredAt               time.Time `json:"occurred_at"`
}
