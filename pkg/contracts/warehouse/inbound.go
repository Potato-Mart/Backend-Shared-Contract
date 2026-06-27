package warehouse

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v9/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v9/pkg/contracts/shared"
	"github.com/Potato-Mart/Backend-Shared-Contract/v9/pkg/enums"
)

type InboundReceipt struct {
	ID        string `json:"id"`
	DepotCode string `json:"depot_code"`
	// Deprecated: use DepotCode.
	DepotID      string `json:"depot_id,omitempty"`
	Reference    string `json:"reference,omitempty"`
	SupplierCode string `json:"supplier_code,omitempty"`
	// Deprecated: use SupplierCode.
	Supplier    string                     `json:"supplier,omitempty"`
	ETA         *time.Time                 `json:"eta,omitempty"`
	Operator    string                     `json:"operator,omitempty"`
	Status      enums.InboundReceiptStatus `json:"status"`
	Note        string                     `json:"note,omitempty"`
	ConfirmedAt *time.Time                 `json:"confirmed_at,omitempty"`
	History     []shared.HistoryEntry      `json:"history,omitempty"`

	common.AuditFields `bson:",inline"`
}

type InboundItem struct {
	ID               string `json:"id"`
	InboundReceiptID string `json:"inbound_receipt_id"`
	ProductSKUCode   string `json:"product_sku_code"`
	// Deprecated: use ProductSKUCode.
	ProductID    string            `json:"product_id,omitempty"`
	Barcode      string            `json:"barcode,omitempty"`
	ProductName  string            `json:"product_name,omitempty"`
	Storage      enums.StorageType `json:"storage,omitempty"`
	ExpectedQty  int               `json:"expected_qty"`
	ReceivedQty  int               `json:"received_qty"`
	LocationCode string            `json:"location_code,omitempty"`
	CreatedAt    time.Time         `json:"created_at"`
}
