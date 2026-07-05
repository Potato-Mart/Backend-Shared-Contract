package purchase

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v11/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v11/pkg/contracts/shared"
	"github.com/Potato-Mart/Backend-Shared-Contract/v11/pkg/enums"
)

type Receipt struct {
	ID           string                    `json:"id"`
	OrderNumber  string                    `json:"order_number"`
	DepotCode    string                    `json:"depot_code,omitempty"`
	Reference    string                    `json:"reference,omitempty"`
	SupplierCode string                    `json:"supplier_code,omitempty"`
	Operator     string                    `json:"operator,omitempty"`
	Status       enums.PurchaseOrderStatus `json:"status"`
	ReceivedAt   *time.Time                `json:"received_at,omitempty"`
	ConfirmedAt  *time.Time                `json:"confirmed_at,omitempty"`
	Note         string                    `json:"note,omitempty"`
	Items        []ReceiptItem             `json:"items"`
	History      []shared.HistoryEntry     `json:"history,omitempty"`

	common.AuditFields
}

type ReceiptItem struct {
	ID             string `json:"id,omitempty"`
	ProductSKUCode string `json:"product_sku_code,omitempty"`
	SKU            string `json:"sku,omitempty"`
	ProductName    string `json:"product_name,omitempty"`
	OrderedQty     int    `json:"ordered_qty"`
	ReceivedQty    int    `json:"received_qty"`
	RejectedQty    int    `json:"rejected_qty,omitempty"`
	LocationCode   string `json:"location_code,omitempty"`
	Note           string `json:"note,omitempty"`
}
