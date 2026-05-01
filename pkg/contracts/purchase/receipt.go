package purchase

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v3/pkg/enums"
)

type Receipt struct {
	ID          string                    `json:"id"`
	OrderID     string                    `json:"order_id"`
	DepotID     string                    `json:"depot_id,omitempty"`
	Reference   string                    `json:"reference,omitempty"`
	Supplier    string                    `json:"supplier,omitempty"`
	Operator    string                    `json:"operator,omitempty"`
	Status      enums.PurchaseOrderStatus `json:"status"`
	ReceivedAt  *time.Time                `json:"received_at,omitempty"`
	ConfirmedAt *time.Time                `json:"confirmed_at,omitempty"`
	Note        string                    `json:"note,omitempty"`
	Items       []ReceiptItem             `json:"items"`
	CreatedAt   time.Time                 `json:"created_at"`
	UpdatedAt   time.Time                 `json:"updated_at"`
}

type ReceiptItem struct {
	ID           string `json:"id,omitempty"`
	ProductID    string `json:"product_id,omitempty"`
	SKU          string `json:"sku,omitempty"`
	ProductName  string `json:"product_name,omitempty"`
	OrderedQty   int    `json:"ordered_qty"`
	ReceivedQty  int    `json:"received_qty"`
	RejectedQty  int    `json:"rejected_qty,omitempty"`
	LocationCode string `json:"location_code,omitempty"`
	Note         string `json:"note,omitempty"`
}
