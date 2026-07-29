package warehouse

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v20/pkg/common"
	warehouseenum "github.com/Potato-Mart/Backend-Shared-Contract/v20/pkg/enums/warehouse"
)

// StockMovement is the shared read model for every stock balance change.
//
// Product.CurrentStock and DepotProduct.StockQty are projections; this ledger
// explains how the balance changed and links the movement to the business
// document that caused it.
type StockMovement struct {
	ID             string                          `json:"id"`
	ProductSKUCode string                          `json:"product_sku_code"`
	SKU            string                          `json:"sku,omitempty"`
	ProductName    string                          `json:"product_name,omitempty"`
	DepotCode      string                          `json:"depot_code,omitempty"`
	LocationCode   string                          `json:"location_code,omitempty"`
	Type           warehouseenum.StockMovementType `json:"type"`
	QtyDelta       int                             `json:"qty_delta"`
	BalanceAfter   int                             `json:"balance_after"`
	OccurredAt     time.Time                       `json:"occurred_at"`
	CreatedBy      string                          `json:"created_by,omitempty"`
	ReasonCode     string                          `json:"reason_code,omitempty"`
	Note           string                          `json:"note,omitempty"`

	PurchaseOrderNumber string `json:"purchase_order_number,omitempty"`
	PurchaseReceiptID   string `json:"purchase_receipt_id,omitempty"`
	SalesOrderNumber    string `json:"sales_order_number,omitempty"`
	DamageReportID      string `json:"damage_report_id,omitempty"`
	ReferenceType       string `json:"reference_type,omitempty"`
	ReferenceID         string `json:"reference_id,omitempty"`

	Metadata common.Metadata `json:"metadata,omitempty"`
}
