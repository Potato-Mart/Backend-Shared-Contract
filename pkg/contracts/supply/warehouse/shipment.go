package warehouse

import (
	security "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/common/security"
	"time"

	common "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/common/shared"
)

type OutboundShipment struct {
	ID             string                  `json:"id"`
	DepotCode      string                  `json:"depot_code"`
	PickingListID  string                  `json:"picking_list_id,omitempty"`
	OrderNumber    string                  `json:"order_number"`
	CustomerName   string                  `json:"customer_name,omitempty"`
	Address        *common.Address         `json:"address,omitempty"`
	Operator       string                  `json:"operator"`
	Status         OutboundShipmentStatus  `json:"status"`
	Containers     []OutboundContainerPlan `json:"containers,omitempty"`
	TrackingNumber string                  `json:"tracking_number,omitempty"`
	Note           string                  `json:"note,omitempty"`
	DispatchedAt   *time.Time              `json:"dispatched_at,omitempty"`
	DeliveredAt    *time.Time              `json:"delivered_at,omitempty"`
	History        []security.HistoryEntry `json:"history,omitempty"`
	CreatedAt      time.Time               `json:"created_at"`
}

// PackingDiscrepancy compares requested and observed package compositions.
type PackingDiscrepancy struct {
	ID                   string                            `json:"id"`
	OrderNumber          string                            `json:"order_number"`
	OrderDate            time.Time                         `json:"order_date"`
	CustomerName         string                            `json:"customer_name,omitempty"`
	ProductSKUCode       string                            `json:"product_sku_code"`
	ProductName          string                            `json:"product_name,omitempty"`
	Kind                 PackingDiscrepancyKind            `json:"kind"`
	RequestedComposition common.PackageCompositionSnapshot `json:"requested_composition"`
	ObservedComposition  common.PackageCompositionSnapshot `json:"observed_composition"`
	RefundAmount         *common.Money                     `json:"refund_amount,omitempty"`
	Notified             bool                              `json:"notified"`
	QualityAssessmentID  string                            `json:"quality_assessment_id,omitempty"`
	DamageHandling       PackingDamageHandling             `json:"damage_handling,omitempty"`
	RecordedAt           time.Time                         `json:"recorded_at"`
	RecordedBy           string                            `json:"recorded_by,omitempty"`
}
