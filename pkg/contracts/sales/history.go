package sales

import (
	"time"

	salesenum "github.com/Potato-Mart/Backend-Shared-Contract/v21/pkg/enums/sales"
)

type StatusHistory struct {
	FromValue  string    `json:"from_value,omitempty"`
	ToValue    string    `json:"to_value"`
	Note       string    `json:"note,omitempty"`
	ActorEmail string    `json:"actor_email,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
}

type Fulfillment struct {
	OrderNumber      string                      `json:"order_number"`
	OrderDate        time.Time                   `json:"order_date"`
	Status           salesenum.FulfillmentStatus `json:"status"`
	PickingPrintedAt *time.Time                  `json:"picking_printed_at,omitempty"`
	PackedAt         *time.Time                  `json:"packed_at,omitempty"`
	UpdatedAt        time.Time                   `json:"updated_at"`
}
