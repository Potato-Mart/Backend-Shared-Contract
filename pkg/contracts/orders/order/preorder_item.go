package order

import (
	"time"

	common "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/common/shared"
)

// PreorderItemSnapshot is server-stamped from Supply's active SKU policy.
// Clients cannot choose or alter preorder state.
type PreorderItemSnapshot struct {
	ProductSKUCode         string     `json:"product_sku_code"`
	PolicyVersion          string     `json:"policy_version"`
	ExpectedAvailableAt    *time.Time `json:"expected_available_at,omitempty"`
	ScheduleTimezone       string     `json:"schedule_timezone"`
	MaxQuantityPerOrder    int        `json:"max_quantity_per_order,omitempty"`
	MaxQuantityPerCustomer int        `json:"max_quantity_per_customer,omitempty"`
	CapturedAt             time.Time  `json:"captured_at"`
}

// PreorderItemState is the package-aware preorder state captured on a paid
// sales-order line.
type PreorderItemState struct {
	Snapshot             PreorderItemSnapshot              `json:"snapshot"`
	Status               PreorderAllocationStatus          `json:"status"`
	RequestedComposition common.PackageCompositionSnapshot `json:"requested_composition"`
	AllocatedComposition common.PackageCompositionSnapshot `json:"allocated_composition"`
	ReservationIDs       []string                          `json:"reservation_ids,omitempty"`
	FirstAllocatedAt     *time.Time                        `json:"first_allocated_at,omitempty"`
	FullyAllocatedAt     *time.Time                        `json:"fully_allocated_at,omitempty"`
}
