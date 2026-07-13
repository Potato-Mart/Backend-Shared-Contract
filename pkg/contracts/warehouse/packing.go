package warehouse

import (
	"time"

	warehouseenum "github.com/Potato-Mart/Backend-Shared-Contract/v16/pkg/enums/warehouse"
)

// PackingLine is the order-line projection needed by the packing UI and
// backend settlement logic.
type PackingLine struct {
	ProductSKUCode string `json:"product_sku_code"`
	SKU            string `json:"sku,omitempty"`
	ProductName    string `json:"product_name,omitempty"`
	OrderedQty     int    `json:"ordered_qty"`
	ScannedQty     int    `json:"scanned_qty"`
	DamagedQty     int    `json:"damaged_qty,omitempty"`
}

// PackingDamage is an auditable damage event discovered while packing.
type PackingDamage struct {
	ID              string                              `json:"id"`
	ProductSKUCode  string                              `json:"product_sku_code"`
	SKU             string                              `json:"sku,omitempty"`
	ProductName     string                              `json:"product_name,omitempty"`
	DamagedQty      int                                 `json:"damaged_qty"`
	Handling        warehouseenum.PackingDamageHandling `json:"handling"`
	Note            string                              `json:"note,omitempty"`
	DamageReportID  string                              `json:"damage_report_id,omitempty"`
	StockMovementID string                              `json:"stock_movement_id,omitempty"`
	CreatedAt       time.Time                           `json:"created_at"`
	CreatedBy       string                              `json:"created_by,omitempty"`
}

// PackingBoxPlan persists box counts and optional contents so label reprints
// stay consistent across sessions.
type PackingBoxPlan struct {
	AmbientBoxes       int                 `json:"ambient_boxes,omitempty"`
	FrozenBoxes        int                 `json:"frozen_boxes,omitempty"`
	ManualAmbientCount bool                `json:"manual_ambient_count,omitempty"`
	ManualFrozenCount  bool                `json:"manual_frozen_count,omitempty"`
	Contents           []PackingBoxContent `json:"contents,omitempty"`
	UpdatedAt          time.Time           `json:"updated_at,omitempty"`
}

// PackingBoxContent describes one product quantity assigned to one box label.
type PackingBoxContent struct {
	BoxNo          int    `json:"box_no"`
	Zone           string `json:"zone,omitempty"`
	ProductSKUCode string `json:"product_sku_code"`
	SKU            string `json:"sku,omitempty"`
	ProductName    string `json:"product_name,omitempty"`
	Qty            int    `json:"qty"`
}
