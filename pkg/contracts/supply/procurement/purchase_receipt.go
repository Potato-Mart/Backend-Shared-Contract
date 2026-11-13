package procurement

import (
	"time"

	security "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/security"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/procurement/purchase_enums"
)

// PurchaseReceipt is the supplier-facing receipt root. Reference is retained
// unchanged because its historical meaning is not source-proven.
//
// ReversedAt, ReversedBy and ReversalReason describe the reversal that
// completed the receipt (status REVERSED). A partial reversal leaves them
// empty and the receipt PARTIALLY_REVERSED; its details are recorded per line
// on PurchaseReceiptItem.ReversedComposition and PurchaseReceiptItem.Reversals.
type PurchaseReceipt struct {
	ID                  string                               `json:"id"`
	PurchaseOrderNumber string                               `json:"purchase_order_number,omitempty"`
	DepotCode           string                               `json:"depot_code,omitempty"`
	MarketCode          string                               `json:"market_code,omitempty"`
	CountryCode         geography.CountryCode                `json:"country_code,omitempty"`
	Reference           string                               `json:"reference,omitempty"`
	SupplierCode        string                               `json:"supplier_code,omitempty"`
	Operator            string                               `json:"operator,omitempty"`
	Status              purchase_enums.PurchaseReceiptStatus `json:"status"`
	ReceivedAt          *time.Time                           `json:"received_at,omitempty"`
	ConfirmedAt         *time.Time                           `json:"confirmed_at,omitempty"`
	ReversedAt          *time.Time                           `json:"reversed_at,omitempty"`
	ReversedBy          string                               `json:"reversed_by,omitempty"`
	ReversalReason      string                               `json:"reversal_reason,omitempty"`
	Note                string                               `json:"note,omitempty"`
	Items               []PurchaseReceiptItem                `json:"items"`
	History             []security.HistoryEntry              `json:"history,omitempty"`

	audit.AuditFields
}
