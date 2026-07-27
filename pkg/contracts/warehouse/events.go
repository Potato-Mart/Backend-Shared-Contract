package warehouse

import (
	"time"

	warehouseenum "github.com/Potato-Mart/Backend-Shared-Contract/v19/pkg/enums/warehouse"
)

// StockAdjustedEvent is emitted on the stock-events topic for every stock
// movement (receipt, adjustment, damage, transfer). AggregateID is the SKU.
// Positive arrivals additionally emit stock.arrived with the preorder
// stock-arrival model in contracts/sales.
type StockAdjustedEvent struct {
	MovementID       string                          `json:"movement_id"`
	MovementType     warehouseenum.StockMovementType `json:"movement_type"`
	ProductSKUCode   string                          `json:"product_sku_code"`
	DepotCode        string                          `json:"depot_code,omitempty"`
	LocationCode     string                          `json:"location_code,omitempty"`
	QtyDelta         int64                           `json:"qty_delta"`
	BalanceAfter     int64                           `json:"balance_after"`
	ReasonCode       string                          `json:"reason_code,omitempty"`
	ReferenceType    string                          `json:"reference_type,omitempty"`
	ReferenceID      string                          `json:"reference_id,omitempty"`
	SalesOrderNumber string                          `json:"sales_order_number,omitempty"`
	OccurredAt       time.Time                       `json:"occurred_at"`
	RequestID        string                          `json:"request_id,omitempty"`
}
