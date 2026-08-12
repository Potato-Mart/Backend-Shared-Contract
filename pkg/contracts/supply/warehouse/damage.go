package warehouse

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/packaging"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/security"

	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/supply/warehouse/warehouse_enums"
)

// DamageReport references the canonical inventory and quality assessment for
// damage observed during warehouse handling.
type DamageReport struct {
	ID                  string                               `json:"id"`
	ProductSKUCode      string                               `json:"product_sku_code"`
	BucketID            string                               `json:"bucket_id"`
	StockUnitID         string                               `json:"stock_unit_id,omitempty"`
	QualityAssessmentID string                               `json:"quality_assessment_id"`
	AffectedComposition packaging.PackageCompositionSnapshot `json:"affected_composition"`
	LossValue           *money.Money                         `json:"loss_value,omitempty"`
	Stage               warehouse_enums.DamageStage          `json:"stage"`
	ReferenceID         string                               `json:"reference_id,omitempty"`
	Note                string                               `json:"note,omitempty"`
	ReportedBy          string                               `json:"reported_by"`
	ReportedAt          time.Time                            `json:"reported_at"`
	History             []security.HistoryEntry              `json:"history,omitempty"`
}
