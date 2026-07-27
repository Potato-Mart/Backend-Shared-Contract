package warehouse

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v19/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v19/pkg/contracts/shared"
	warehouseenum "github.com/Potato-Mart/Backend-Shared-Contract/v19/pkg/enums/warehouse"
)

type OutboundShipment struct {
	ID             string                               `json:"id"`
	DepotCode      string                               `json:"depot_code,omitempty"`
	PickingListID  string                               `json:"picking_list_id,omitempty"`
	OrderNumber    string                               `json:"order_number,omitempty"`
	CustomerName   string                               `json:"customer_name,omitempty"`
	Address        *common.Address                      `json:"address,omitempty"`
	Operator       string                               `json:"operator"`
	Status         warehouseenum.OutboundShipmentStatus `json:"status"`
	TrackingNumber string                               `json:"tracking_number,omitempty"`
	Note           string                               `json:"note,omitempty"`
	DispatchedAt   *time.Time                           `json:"dispatched_at,omitempty"`
	DeliveredAt    *time.Time                           `json:"delivered_at,omitempty"`
	History        []shared.HistoryEntry                `json:"history,omitempty"`
	CreatedAt      time.Time                            `json:"created_at"`
}

type PackingDiscrepancy struct {
	ID              string                               `json:"id"`
	OrderNumber     string                               `json:"order_number"`
	OrderDate       time.Time                            `json:"order_date"`
	CustomerName    string                               `json:"customer_name,omitempty"`
	ProductSKUCode  string                               `json:"product_sku_code"`
	ProductName     string                               `json:"product_name,omitempty"`
	Kind            warehouseenum.PackingDiscrepancyKind `json:"kind"`
	OrderedQty      int                                  `json:"ordered_qty"`
	ScannedQty      int                                  `json:"scanned_qty"`
	DiffQty         int                                  `json:"diff_qty"`
	UnitPrice       *common.Money                        `json:"unit_price,omitempty"`
	RefundAmount    *common.Money                        `json:"refund_amount,omitempty"`
	ReturnToStock   bool                                 `json:"return_to_stock"`
	Notified        bool                                 `json:"notified"`
	DamageReportID  string                               `json:"damage_report_id,omitempty"`
	StockMovementID string                               `json:"stock_movement_id,omitempty"`
	DamagedQty      int                                  `json:"damaged_qty,omitempty"`
	DamageHandling  string                               `json:"damage_handling,omitempty"`
	RecordedAt      time.Time                            `json:"recorded_at"`
	RecordedBy      string                               `json:"recorded_by,omitempty"`
}
