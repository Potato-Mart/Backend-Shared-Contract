package warehouse

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v24/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v24/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v24/pkg/contracts/common/packaging"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v24/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v24/pkg/contracts/supply/warehouse/warehouse_enums"
)

type OutboundShipment struct {
	ID             string                                 `json:"id"`
	DepotCode      string                                 `json:"depot_code"`
	PickingListID  string                                 `json:"picking_list_id,omitempty"`
	OrderNumber    string                                 `json:"order_number"`
	CustomerName   string                                 `json:"customer_name,omitempty"`
	Address        *geography.Address                     `json:"address,omitempty"`
	Operator       string                                 `json:"operator"`
	Status         warehouse_enums.OutboundShipmentStatus `json:"status"`
	Containers     []OutboundContainerPlan                `json:"containers,omitempty"`
	TrackingNumber string                                 `json:"tracking_number,omitempty"`
	Note           string                                 `json:"note,omitempty"`
	DispatchedAt   *time.Time                             `json:"dispatched_at,omitempty"`
	DeliveredAt    *time.Time                             `json:"delivered_at,omitempty"`
	History        []security.HistoryEntry                `json:"history,omitempty"`
	CreatedAt      time.Time                              `json:"created_at"`
}

// PackingDiscrepancy compares requested and observed package compositions.
type PackingDiscrepancy struct {
	ID                   string                                 `json:"id"`
	OrderNumber          string                                 `json:"order_number"`
	OrderDate            time.Time                              `json:"order_date"`
	CustomerName         string                                 `json:"customer_name,omitempty"`
	ProductSKUCode       string                                 `json:"product_sku_code"`
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
