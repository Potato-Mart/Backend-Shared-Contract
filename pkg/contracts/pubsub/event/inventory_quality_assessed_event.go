package event

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/packaging"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/warehouse/warehouse_enums"
)

type InventoryQualityAssessedEvent struct {
	QualityAssessmentID string                               `json:"quality_assessment_id"`
	SKUCode             string                               `json:"sku_code"`
	DepotCode           string                               `json:"depot_code"`
	BucketID            string                               `json:"bucket_id"`
	StockUnitID         string                               `json:"stock_unit_id,omitempty"`
	AssessedComposition packaging.PackageCompositionSnapshot `json:"assessed_composition"`
	PreviousCondition   warehouse_enums.InventoryCondition   `json:"previous_condition"`
	ResultCondition     warehouse_enums.InventoryCondition   `json:"result_condition"`
	PreviousDisposition warehouse_enums.InventoryDisposition `json:"previous_disposition"`
	ResultDisposition   warehouse_enums.InventoryDisposition `json:"result_disposition"`
	MovementIDs         []string                             `json:"movement_ids,omitempty"`
	Revision            int64                                `json:"revision"`
	OccurredAt          time.Time                            `json:"occurred_at"`
}
