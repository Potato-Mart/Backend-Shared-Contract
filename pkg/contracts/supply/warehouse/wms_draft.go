package warehouse

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/packaging"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/security"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/warehouse/warehouse_enums"
)

// WMSDraft is an uncommitted package-aware warehouse operation.
type WMSDraft struct {
	ID        string                       `json:"id"`
	Type      warehouse_enums.WMSDraftType `json:"type"`
	Operator  string                       `json:"operator"`
	DepotCode string                       `json:"depot_code"`
	// PurchaseReceiptID links inbound WMS work to the procurement receipt.
	// It may be empty only while the draft remains unsubmitted.
	PurchaseReceiptID   string                               `json:"purchase_receipt_id,omitempty"`
	PurchaseOrderNumber string                               `json:"purchase_order_number,omitempty"`
	Items               []WMSDraftItem                       `json:"items"`
	ItemCount           int64                                `json:"item_count"`
	TotalComposition    packaging.PackageCompositionSnapshot `json:"total_composition"`
	Status              warehouse_enums.WMSDraftStatus       `json:"status"`
	Note                string                               `json:"note,omitempty"`
	SubmittedAt         *time.Time                           `json:"submitted_at,omitempty"`
	History             []security.HistoryEntry              `json:"history,omitempty"`

	audit.AuditFields
}
