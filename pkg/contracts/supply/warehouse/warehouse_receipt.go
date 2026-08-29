package warehouse

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/audit"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/warehouse/warehouse_enums"
)

// WarehouseReceipt is the operational inbound-receiving root. The owning
// service allows an empty PurchaseReceiptID only while Status is DRAFT and
// requires it before confirmation or inbound WMS submission.
type WarehouseReceipt struct {
	ID                  string                               `json:"id"`
	PurchaseReceiptID   string                               `json:"purchase_receipt_id,omitempty"`
	PurchaseOrderNumber string                               `json:"purchase_order_number,omitempty"`
	DepotCode           string                               `json:"depot_code"`
	MarketCode          string                               `json:"market_code,omitempty"`
	Reference           string                               `json:"reference,omitempty"`
	SupplierCode        string                               `json:"supplier_code,omitempty"`
	ETA                 *time.Time                           `json:"eta,omitempty"`
	Operator            string                               `json:"operator,omitempty"`
	Status              warehouse_enums.InboundReceiptStatus `json:"status"`
	Items               []WarehouseReceiptItem               `json:"items"`
	Note                string                               `json:"note,omitempty"`
	ConfirmedAt         *time.Time                           `json:"confirmed_at,omitempty"`
	History             []security.HistoryEntry              `json:"history,omitempty"`

	audit.AuditFields
}
