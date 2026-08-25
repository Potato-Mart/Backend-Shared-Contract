package order

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/orders/order/order_enums"
	"time"
)

type Fulfillment struct {
	OrderNumber      string                        `json:"order_number"`
	OrderDate        time.Time                     `json:"order_date"`
	Status           order_enums.FulfillmentStatus `json:"status"`
	PickingPrintedAt *time.Time                    `json:"picking_printed_at,omitempty"`
	PackedAt         *time.Time                    `json:"packed_at,omitempty"`
	UpdatedAt        time.Time                     `json:"updated_at"`
}
