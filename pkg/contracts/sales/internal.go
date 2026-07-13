package sales

import "time"

// OrderPackingProjection is the durable packing snapshot shared between
// Operations and Commerce. Transport acknowledgement types stay provider-local.
type OrderPackingProjection struct {
	OrderNumber string               `json:"order_number"`
	Revision    int64                `json:"revision"`
	Packing     OrderPackingProgress `json:"packing"`
	UpdatedAt   time.Time            `json:"updated_at"`
}

// PreorderStockArrivalEvent is durably emitted by Operations after a positive
// SKU stock mutation. Commerce stores an inbox receipt before allocating.
type PreorderStockArrivalEvent struct {
	EventID        string    `json:"event_id"`
	ProductSKUCode string    `json:"product_sku_code"`
	ArrivedQty     int       `json:"arrived_qty"`
	AvailableQty   int       `json:"available_qty"`
	OccurredAt     time.Time `json:"occurred_at"`
}

type DemandBucket struct {
	ProductSKUCode string    `json:"product_sku_code"`
	SKU            string    `json:"sku,omitempty"`
	ProductName    string    `json:"product_name,omitempty"`
	Date           time.Time `json:"date"`
	QtyUnits       int       `json:"qty_units"`
	CartonQty      int       `json:"carton_qty,omitempty"`
	CartonSize     int       `json:"carton_size,omitempty"`
}

type OpenDemandLine struct {
	OrderNumber    string    `json:"order_number"`
	ProductSKUCode string    `json:"product_sku_code"`
	SKU            string    `json:"sku,omitempty"`
	ProductName    string    `json:"product_name,omitempty"`
	QtyUnits       int       `json:"qty_units"`
	CartonQty      int       `json:"carton_qty,omitempty"`
	CartonSize     int       `json:"carton_size,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
}
