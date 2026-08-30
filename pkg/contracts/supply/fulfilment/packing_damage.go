package fulfilment

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/warehouse/warehouse_enums"
)

// PackingDamage links packing evidence to canonical inventory assessment and
// movement identities.
type PackingDamage struct {
	ID                   string                                `json:"id"`
	SKUCode              string                                `json:"sku_code"`
	SourceBucketID       string                                `json:"source_bucket_id"`
	StockUnitID          string                                `json:"stock_unit_id,omitempty"`
	QualityAssessmentID  string                                `json:"quality_assessment_id"`
	AffectedComposition  packaging.PackageCompositionSnapshot  `json:"affected_composition"`
	Handling             warehouse_enums.PackingDamageHandling `json:"handling"`
	Note                 string                                `json:"note,omitempty"`
	ResultingMovementIDs []string                              `json:"resulting_movement_ids,omitempty"`
	CreatedAt            time.Time                             `json:"created_at"`
	CreatedBy            string                                `json:"created_by,omitempty"`
}
