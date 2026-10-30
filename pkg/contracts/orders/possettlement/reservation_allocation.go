package possettlement

import "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/inventory"

// ReservationAllocation retains Supply's native source union and revision.
// The owner must verify it is still the exact allocation for this order and
// generation; a historical reservation receipt alone is insufficient.
type ReservationAllocation struct {
	Allocation             *inventory.StockReservationAllocation `json:"allocation,omitempty"`
	GroupOrderCode         string                                `json:"group_order_code,omitempty"`
	GroupOrderRole         string                                `json:"group_order_role,omitempty"`
	OrderItemID            string                                `json:"order_item_id,omitempty"`
	ParentAllocationLineID string                                `json:"parent_allocation_line_id,omitempty"`
	ParentFulfilmentID     string                                `json:"parent_fulfilment_id,omitempty"`
	ParentOrderNumber      string                                `json:"parent_order_number,omitempty"`
	ReservationID          string                                `json:"reservation_id,omitempty"`
}
