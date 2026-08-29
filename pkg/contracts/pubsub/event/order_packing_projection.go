package event

import (
	"time"

	order "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/orders/order"
)

// OrderPackingProjection is the durable packing snapshot shared between
// Supply and Orders. Transport acknowledgement types stay provider-local.
type OrderPackingProjection struct {
	OrderNumber string                     `json:"order_number"`
	Revision    int64                      `json:"revision"`
	Packing     order.OrderPackingProgress `json:"packing"`
	UpdatedAt   time.Time                  `json:"updated_at"`
}
