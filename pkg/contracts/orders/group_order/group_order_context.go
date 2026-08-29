package group_order

import order_enums "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/orders/group_order/group_order_enums"

// GroupOrderContext identifies an order as the consolidated fulfilment owner
// or as a participant referencing that parent fulfilment.
type GroupOrderContext struct {
	GroupOrderCode          string                     `json:"group_order_code"`
	Role                    order_enums.GroupOrderRole `json:"role"`
	ParentOrderNumber       string                     `json:"parent_order_number,omitempty"`
	ParentFulfilmentID      string                     `json:"parent_fulfilment_id,omitempty"`
	ParentAllocationLineIDs []string                   `json:"parent_allocation_line_ids,omitempty"`
}
