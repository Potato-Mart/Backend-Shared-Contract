package listing

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/supply/warehouse/warehouse_enums"
)

// DamageSaleApproval is the explicit quality decision that allows damaged
// inventory to be sold at a reduced price. Unsafe dispositions are never
// approved.
type DamageSaleApproval struct {
	QualityAssessmentID string                         `json:"quality_assessment_id"`
	Tier                warehouse_enums.DamageSaleTier `json:"tier"`
	ReasonCode          string                         `json:"reason_code"`
	ApprovedBy          string                         `json:"approved_by"`
	ApprovedAt          time.Time                      `json:"approved_at"`
}
