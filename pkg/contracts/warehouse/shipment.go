package warehouse

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v2/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v2/pkg/enums"
)

type OutboundShipment struct {
	ID             string                       `json:"id"`
	DepotID        string                       `json:"depot_id,omitempty"`
	PickingListID  string                       `json:"picking_list_id,omitempty"`
	OrderID        string                       `json:"order_id,omitempty"`
	OrderNumber    string                       `json:"order_number,omitempty"`
	CustomerName   string                       `json:"customer_name,omitempty"`
	Address        string                       `json:"address,omitempty"`
	State          string                       `json:"state,omitempty"`
	Operator       string                       `json:"operator"`
	Status         enums.OutboundShipmentStatus `json:"status"`
	TrackingNumber string                       `json:"tracking_number,omitempty"`
	Note           string                       `json:"note,omitempty"`
	DispatchedAt   *time.Time                   `json:"dispatched_at,omitempty"`
	CreatedAt      time.Time                    `json:"created_at"`
}

type PackingDiscrepancy struct {
	ID            string                       `json:"id"`
	OrderNumber   string                       `json:"order_number"`
	OrderDate     time.Time                    `json:"order_date"`
	CustomerName  string                       `json:"customer_name,omitempty"`
	SKU           string                       `json:"sku"`
	ProductName   string                       `json:"product_name,omitempty"`
	Kind          enums.PackingDiscrepancyKind `json:"kind"`
	OrderedQty    int                          `json:"ordered_qty"`
	ScannedQty    int                          `json:"scanned_qty"`
	DiffQty       int                          `json:"diff_qty"`
	UnitPrice     *common.Money                `json:"unit_price,omitempty"`
	RefundAmount  *common.Money                `json:"refund_amount,omitempty"`
	ReturnToStock bool                         `json:"return_to_stock"`
	Notified      bool                         `json:"notified"`
	RecordedAt    time.Time                    `json:"recorded_at"`
	RecordedBy    string                       `json:"recorded_by,omitempty"`
}
