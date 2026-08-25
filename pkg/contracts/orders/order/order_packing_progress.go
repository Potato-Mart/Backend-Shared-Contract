package order

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/orders/order/order_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/orders/shipping"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/supply/operations"
)

// OrderPackingProgress is the order-owned packing state shown to both staff
// and customer order views. It replaces admin-facing dependence on a separate
// packing-session aggregate for current packing progress.
type OrderPackingProgress struct {
	Status        order_enums.FulfillmentStatus      `json:"status,omitempty"`
	Operator      string                             `json:"operator,omitempty"`
	Lines         []operations.PackingLine           `json:"lines,omitempty"`
	Containers    []operations.OutboundContainerPlan `json:"containers,omitempty"`
	Damages       []operations.PackingDamage         `json:"damages,omitempty"`
	Discrepancies []shipping.PackingDiscrepancy      `json:"discrepancies,omitempty"`
	StartedAt     *time.Time                         `json:"started_at,omitempty"`
	UpdatedAt     *time.Time                         `json:"updated_at,omitempty"`
	PackedAt      *time.Time                         `json:"packed_at,omitempty"`
	FulfilledAt   *time.Time                         `json:"fulfilled_at,omitempty"`
}
