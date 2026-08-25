package warehouse

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/packaging"

	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/supply/warehouse/warehouse_enums"
)

// QualityAssessment captures an observed condition/disposition decision and
// the resulting physical inventory movements.
type QualityAssessment struct {
	ID                   string                               `json:"id"`
	SKUCode              string                               `json:"sku_code"`
	BucketID             string                               `json:"bucket_id"`
	StockUnitID          string                               `json:"stock_unit_id,omitempty"`
	AssessedComposition  packaging.PackageCompositionSnapshot `json:"assessed_composition"`
	PreviousCondition    warehouse_enums.InventoryCondition   `json:"previous_condition"`
	ResultCondition      warehouse_enums.InventoryCondition   `json:"result_condition"`
	PreviousDisposition  warehouse_enums.InventoryDisposition `json:"previous_disposition"`
	ResultDisposition    warehouse_enums.InventoryDisposition `json:"result_disposition"`
	AssessedBy           string                               `json:"assessed_by"`
	ReasonCode           string                               `json:"reason_code"`
	Note                 string                               `json:"note,omitempty"`
	EvidenceMediaURLs    []string                             `json:"evidence_media_urls,omitempty"`
	AssessedAt           time.Time                            `json:"assessed_at"`
	ResultingMovementIDs []string                             `json:"resulting_movement_ids,omitempty"`
}
