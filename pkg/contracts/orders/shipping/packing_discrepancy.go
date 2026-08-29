package shipping

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/warehouse/warehouse_enums"
	"time"
)

// PackingDiscrepancy compares requested and observed package compositions.
type PackingDiscrepancy struct {
	ID                   string                                 `json:"id"`
	OrderNumber          string                                 `json:"order_number"`
	OrderDate            time.Time                              `json:"order_date"`
	CustomerName         string                                 `json:"customer_name,omitempty"`
	SKUCode              string                                 `json:"sku_code"`
	ProductName          string                                 `json:"product_name,omitempty"`
	Kind                 warehouse_enums.PackingDiscrepancyKind `json:"kind"`
	RequestedComposition packaging.PackageCompositionSnapshot   `json:"requested_composition"`
	ObservedComposition  packaging.PackageCompositionSnapshot   `json:"observed_composition"`
	RefundAmount         *money.Money                           `json:"refund_amount,omitempty"`
	Notified             bool                                   `json:"notified"`
	QualityAssessmentID  string                                 `json:"quality_assessment_id,omitempty"`
	DamageHandling       warehouse_enums.PackingDamageHandling  `json:"damage_handling,omitempty"`
	RecordedAt           time.Time                              `json:"recorded_at"`
	RecordedBy           string                                 `json:"recorded_by,omitempty"`
}
