package warehouse

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v12/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v12/pkg/contracts/shared"
	warehouseenum "github.com/Potato-Mart/Backend-Shared-Contract/v12/pkg/enums/warehouse"
)

// WMSDraft is an uncommitted batch of inbound or outbound stock movements
// created by a warehouse operator on a PDA or desktop. On submission the draft
// is converted into inbound or outbound warehouse records by the owning service.
type WMSDraft struct {
	ID          string                       `json:"id"`
	Type        warehouseenum.WMSDraftType   `json:"type"`
	Operator    string                       `json:"operator"`
	DepotCode   string                       `json:"depot_code,omitempty"`
	Reference   string                       `json:"reference,omitempty"` // supplier PO / order number
	Items       []WMSDraftItem               `json:"items"`
	ItemCount   int                          `json:"item_count"`
	TotalQty    int                          `json:"total_qty"`
	Status      warehouseenum.WMSDraftStatus `json:"status"`
	Note        string                       `json:"note,omitempty"`
	SubmittedAt *time.Time                   `json:"submitted_at,omitempty"`
	History     []shared.HistoryEntry        `json:"history,omitempty"`

	common.AuditFields
}

// WMSDraftItem is the schema for each element of WMSDraft.Items.
type WMSDraftItem struct {
	ProductSKUCode string `json:"product_sku_code"`
	ProductName    string `json:"product_name,omitempty"`
	Barcode        string `json:"barcode,omitempty"`
	LocationCode   string `json:"location_code,omitempty"`
	Qty            int    `json:"qty"`
	ExpiryYM       string `json:"expiry_ym,omitempty"` // "YYYY-MM"
}

// OrderPackingProgress persists per-SKU scan progress for the packing
// screen so it survives page reloads, status reversions, and re-packs.
// Rows are deleted when an order is marked packed; reopening for re-pack
// intentionally retains them so the operator resumes from where they left off.
type OrderPackingProgress struct {
	ID                 string    `json:"id"`
	OrderNumber        string    `json:"order_number"`
	OrderDate          time.Time `json:"order_date"`
	ProductSKUCode     string    `json:"product_sku_code"`
	ProductName        string    `json:"product_name,omitempty"`
	OrderedQty         int       `json:"ordered_qty"`
	ScannedQty         int       `json:"scanned_qty"`
	MergedToFrozenQty  int       `json:"merged_to_frozen_qty"`
	FrozenConfirmedQty int       `json:"frozen_confirmed_qty"`
	UpdatedAt          time.Time `json:"updated_at"`
}
