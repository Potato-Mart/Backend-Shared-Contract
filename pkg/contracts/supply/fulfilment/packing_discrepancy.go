package fulfilment

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/warehouse/warehouse_enums"
	"time"
)

// PackingDiscrepancy compares requested and observed package compositions.
//
// Withdrawals record base units later withdrawn from the discrepancy; the
// contract does not restrict them to a discrepancy kind. A shortage's recorded
// base units are requested_composition.total_base_units minus
// observed_composition.total_base_units; its remaining base units are the
// recorded base units minus the withdrawn_base_units of command and scan
// withdrawals plus the withdrawn_base_units of reinstate entries.
// Withdrawals are append-only evidence and do not rewrite the compositions.
//
// Settlements record what Supply did to stock for the units of a shortage the
// packing station refunded (no stock or overweight). They do not change the
// kind, the compositions or the remaining base units.
type PackingDiscrepancy struct {
	ID                   string                                 `json:"id"`
	OrderItemID          string                                 `json:"order_item_id"`
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
	Withdrawals          []PackingDiscrepancyWithdrawal         `json:"withdrawals,omitempty"`
	Settlements          []PackingShortageSettlement            `json:"settlements,omitempty"`
}
