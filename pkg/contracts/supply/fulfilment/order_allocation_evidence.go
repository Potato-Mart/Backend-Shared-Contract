package fulfilment

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v35/pkg/contracts/supply/inventory"
)

// OrderAllocationEvidence binds one frozen reservation allocation to its
// order line and any subsequent picking and staging records. It carries source
// facts only; readiness and transition policy remain owned by Supply.
type OrderAllocationEvidence struct {
	OrderItemID string                               `json:"order_item_id"`
	Reservation inventory.StockReservation           `json:"reservation"`
	Allocation  inventory.StockReservationAllocation `json:"allocation"`
	Picking     *PickingAllocation                   `json:"picking,omitempty"`
	Staging     *inventory.StockStagingRecord        `json:"staging,omitempty"`
}
