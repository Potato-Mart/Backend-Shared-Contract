package warehouse

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v14/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v14/pkg/contracts/shared"
	warehouseenum "github.com/Potato-Mart/Backend-Shared-Contract/v14/pkg/enums/warehouse"
)

type InboundReceipt struct {
	ID           string                             `json:"id"`
	DepotCode    string                             `json:"depot_code"`
	Reference    string                             `json:"reference,omitempty"`
	SupplierCode string                             `json:"supplier_code,omitempty"`
	ETA          *time.Time                         `json:"eta,omitempty"`
	Operator     string                             `json:"operator,omitempty"`
	Status       warehouseenum.InboundReceiptStatus `json:"status"`
	Note         string                             `json:"note,omitempty"`
	ConfirmedAt  *time.Time                         `json:"confirmed_at,omitempty"`
	History      []shared.HistoryEntry              `json:"history,omitempty"`

	common.AuditFields
}

type InboundItem struct {
	ID               string                    `json:"id"`
	InboundReceiptID string                    `json:"inbound_receipt_id"`
	ProductSKUCode   string                    `json:"product_sku_code"`
	Barcode          string                    `json:"barcode,omitempty"`
	ProductName      string                    `json:"product_name,omitempty"`
	Storage          warehouseenum.StorageType `json:"storage,omitempty"`
	ExpectedQty      int                       `json:"expected_qty"`
	ReceivedQty      int                       `json:"received_qty"`
	LocationCode     string                    `json:"location_code,omitempty"`
	CreatedAt        time.Time                 `json:"created_at"`
}
