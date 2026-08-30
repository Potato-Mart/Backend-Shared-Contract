package order

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/orders/order/order_enums"
	"time"
)

// PreorderItemState is the package-aware preorder state captured on a paid
// sales-order line.
type PreorderItemState struct {
	Snapshot             PreorderItemSnapshot                 `json:"snapshot"`
	Status               order_enums.PreorderAllocationStatus `json:"status"`
	RequestedComposition packaging.PackageCompositionSnapshot `json:"requested_composition"`
	AllocatedComposition packaging.PackageCompositionSnapshot `json:"allocated_composition"`
	ReservationIDs       []string                             `json:"reservation_ids,omitempty"`
	FirstAllocatedAt     *time.Time                           `json:"first_allocated_at,omitempty"`
	FullyAllocatedAt     *time.Time                           `json:"fully_allocated_at,omitempty"`
}
