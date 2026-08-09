package event

import (
	"time"

	order "github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/orders/order"
)

// OrderPackingProjection is the durable packing snapshot shared between
// Supply and Orders. Transport acknowledgement types stay provider-local.
type OrderPackingProjection struct {
	OrderNumber string                     `json:"order_number"`
	Revision    int64                      `json:"revision"`
	Packing     order.OrderPackingProgress `json:"packing"`
	UpdatedAt   time.Time                  `json:"updated_at"`
}

// FulfilmentShippedEvent is emitted on the fulfilment-events topic when a
// shipment leaves the warehouse. AggregateID is the order number.
type FulfilmentShippedEvent struct {
	OrderNumber    string    `json:"order_number"`
	ShipmentID     string    `json:"shipment_id,omitempty"`
	Carrier        string    `json:"carrier,omitempty"`
	TrackingNumber string    `json:"tracking_number,omitempty"`
	TrackingURL    string    `json:"tracking_url,omitempty"`
	Note           string    `json:"note,omitempty"`
	OccurredAt     time.Time `json:"occurred_at"`
	RequestID      string    `json:"request_id,omitempty"`
}

// FulfilmentDeliveredEvent is emitted on the fulfilment-events topic when a
// shipment is confirmed delivered.
type FulfilmentDeliveredEvent struct {
	OrderNumber string    `json:"order_number"`
	ShipmentID  string    `json:"shipment_id,omitempty"`
	Note        string    `json:"note,omitempty"`
	OccurredAt  time.Time `json:"occurred_at"`
	RequestID   string    `json:"request_id,omitempty"`
}

// FulfilmentCompletedEvent is emitted on the fulfilment-events topic when
// fulfilment finishes end-to-end for an order.
type FulfilmentCompletedEvent struct {
	OrderNumber string    `json:"order_number"`
	Note        string    `json:"note,omitempty"`
	OccurredAt  time.Time `json:"occurred_at"`
	RequestID   string    `json:"request_id,omitempty"`
}

// FulfilmentTrackingEvent is emitted on the fulfilment-events topic when
// carrier tracking details are assigned or corrected.
type FulfilmentTrackingEvent struct {
	OrderNumber    string    `json:"order_number"`
	TrackingNumber string    `json:"tracking_number"`
	TrackingURL    string    `json:"tracking_url,omitempty"`
	OccurredAt     time.Time `json:"occurred_at"`
	RequestID      string    `json:"request_id,omitempty"`
}
