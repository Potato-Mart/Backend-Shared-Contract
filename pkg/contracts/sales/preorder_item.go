package sales

import (
	"time"

	salesenum "github.com/Potato-Mart/Backend-Shared-Contract/v16/pkg/enums/sales"
)

// PreorderItemSnapshot is server-stamped from Operations' active SKU policy.
// Clients cannot choose or alter preorder state.
type PreorderItemSnapshot struct {
	ProductSKUCode         string     `json:"product_sku_code"`
	PolicyVersion          string     `json:"policy_version"`
	ExpectedAvailableAt    *time.Time `json:"expected_available_at,omitempty"`
	MaxQuantityPerOrder    int        `json:"max_quantity_per_order,omitempty"`
	MaxQuantityPerCustomer int        `json:"max_quantity_per_customer,omitempty"`
	CapturedAt             time.Time  `json:"captured_at"`
}

// PreorderItemState is persisted on the paid sales-order line. Allocations may
// be partial, but FIFO prevents a later line from bypassing an older line for
// the same SKU.
type PreorderItemState struct {
	Snapshot          PreorderItemSnapshot               `json:"snapshot"`
	Status            salesenum.PreorderAllocationStatus `json:"status"`
	OrderedQuantity   int                                `json:"ordered_quantity"`
	AllocatedQuantity int                                `json:"allocated_quantity"`
	ReservationIDs    []string                           `json:"reservation_ids,omitempty"`
	FirstAllocatedAt  *time.Time                         `json:"first_allocated_at,omitempty"`
	FullyAllocatedAt  *time.Time                         `json:"fully_allocated_at,omitempty"`
}
