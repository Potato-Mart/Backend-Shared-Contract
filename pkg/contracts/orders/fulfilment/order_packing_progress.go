package fulfilment

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/orders/order/order_enums"
	supplyfulfilment "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/fulfilment"
)

// OrderPackingProgress is the order-owned packing state shown to both staff
// and customer order views. It replaces admin-facing dependence on a separate
// packing-session aggregate for current packing progress.
type OrderPackingProgress struct {
	Status        order_enums.FulfillmentStatus            `json:"status,omitempty"`
	Operator      string                                   `json:"operator,omitempty"`
	Lines         []supplyfulfilment.PackingLine           `json:"lines,omitempty"`
	Containers    []supplyfulfilment.OutboundContainerPlan `json:"containers,omitempty"`
	Damages       []supplyfulfilment.PackingDamage         `json:"damages,omitempty"`
	Discrepancies []supplyfulfilment.PackingDiscrepancy    `json:"discrepancies,omitempty"`
	StartedAt     *time.Time                               `json:"started_at,omitempty"`
	UpdatedAt     *time.Time                               `json:"updated_at,omitempty"`
	PackedAt      *time.Time                               `json:"packed_at,omitempty"`
	FulfilledAt   *time.Time                               `json:"fulfilled_at,omitempty"`
}
