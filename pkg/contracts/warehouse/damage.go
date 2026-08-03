package warehouse

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/contracts/shared"
	warehouseenum "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/enums/warehouse"
)

// DamageReport references the canonical inventory and quality assessment for
// damage observed during warehouse handling.
type DamageReport struct {
	ID                  string                            `json:"id"`
	ProductSKUCode      string                            `json:"product_sku_code"`
	BucketID            string                            `json:"bucket_id"`
	StockUnitID         string                            `json:"stock_unit_id,omitempty"`
	QualityAssessmentID string                            `json:"quality_assessment_id"`
	AffectedComposition common.PackageCompositionSnapshot `json:"affected_composition"`
	LossValue           *common.Money                     `json:"loss_value,omitempty"`
	Stage               warehouseenum.DamageStage         `json:"stage"`
	ReferenceID         string                            `json:"reference_id,omitempty"`
	Note                string                            `json:"note,omitempty"`
	ReportedBy          string                            `json:"reported_by"`
	ReportedAt          time.Time                         `json:"reported_at"`
	History             []shared.HistoryEntry             `json:"history,omitempty"`
}
