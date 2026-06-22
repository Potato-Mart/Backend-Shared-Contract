package warehouse

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v7/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v7/pkg/enums"
)

// StockMovement is the shared read model for every stock balance change.
//
// Product.CurrentStock and DepotProduct.StockQty are projections; this ledger
// explains how the balance changed and links the movement to the business
// document that caused it.
type StockMovement struct {
	ID           string                  `json:"id"`
	ProductID    string                  `json:"product_id"`
	SKU          string                  `json:"sku,omitempty"`
	ProductName  string                  `json:"product_name,omitempty"`
	DepotID      string                  `json:"depot_id,omitempty"`
	LocationCode string                  `json:"location_code,omitempty"`
	Type         enums.StockMovementType `json:"type"`
	QtyDelta     int                     `json:"qty_delta"`
	BalanceAfter int                     `json:"balance_after"`
	OccurredAt   time.Time               `json:"occurred_at"`
	CreatedBy    string                  `json:"created_by,omitempty"`
	ReasonCode   string                  `json:"reason_code,omitempty"`
	Note         string                  `json:"note,omitempty"`

	PurchaseOrderID     string `json:"purchase_order_id,omitempty"`
	PurchaseOrderNumber string `json:"purchase_order_number,omitempty"`
	PurchaseReceiptID   string `json:"purchase_receipt_id,omitempty"`
	SalesOrderID        string `json:"sales_order_id,omitempty"`
	SalesOrderNumber    string `json:"sales_order_number,omitempty"`
	DamageReportID      string `json:"damage_report_id,omitempty"`
	ReferenceType       string `json:"reference_type,omitempty"`
	ReferenceID         string `json:"reference_id,omitempty"`

	Metadata common.Metadata `json:"metadata,omitempty"`
}
