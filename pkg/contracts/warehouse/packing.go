package warehouse

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v9/pkg/enums"
)

// PackingLine is the order-line projection needed by the packing UI and
// backend settlement logic.
type PackingLine struct {
	ProductSKUCode string `bson:"product_sku_code" json:"product_sku_code"`
	// Deprecated: use ProductSKUCode.
	ProductID   string `bson:"product_id,omitempty" json:"product_id,omitempty"`
	SKU         string `bson:"sku,omitempty" json:"sku,omitempty"`
	ProductName string `bson:"product_name,omitempty" json:"product_name,omitempty"`
	OrderedQty  int    `bson:"ordered_qty" json:"ordered_qty"`
	ScannedQty  int    `bson:"scanned_qty" json:"scanned_qty"`
	DamagedQty  int    `bson:"damaged_qty,omitempty" json:"damaged_qty,omitempty"`
}

// PackingDamage is an auditable damage event discovered while packing.
type PackingDamage struct {
	ID             string `bson:"id" json:"id"`
	ProductSKUCode string `bson:"product_sku_code" json:"product_sku_code"`
	// Deprecated: use ProductSKUCode.
	ProductID       string                      `bson:"product_id,omitempty" json:"product_id,omitempty"`
	SKU             string                      `bson:"sku,omitempty" json:"sku,omitempty"`
	ProductName     string                      `bson:"product_name,omitempty" json:"product_name,omitempty"`
	DamagedQty      int                         `bson:"damaged_qty" json:"damaged_qty"`
	Handling        enums.PackingDamageHandling `bson:"handling" json:"handling"`
	Note            string                      `bson:"note,omitempty" json:"note,omitempty"`
	DamageReportID  string                      `bson:"damage_report_id,omitempty" json:"damage_report_id,omitempty"`
	StockMovementID string                      `bson:"stock_movement_id,omitempty" json:"stock_movement_id,omitempty"`
	CreatedAt       time.Time                   `bson:"created_at" json:"created_at"`
	CreatedBy       string                      `bson:"created_by,omitempty" json:"created_by,omitempty"`
}

// PackingBoxPlan persists box counts and optional contents so label reprints
// stay consistent across sessions.
type PackingBoxPlan struct {
	AmbientBoxes       int                 `bson:"ambient_boxes,omitempty" json:"ambient_boxes,omitempty"`
	FrozenBoxes        int                 `bson:"frozen_boxes,omitempty" json:"frozen_boxes,omitempty"`
	ManualAmbientCount bool                `bson:"manual_ambient_count,omitempty" json:"manual_ambient_count,omitempty"`
	ManualFrozenCount  bool                `bson:"manual_frozen_count,omitempty" json:"manual_frozen_count,omitempty"`
	Contents           []PackingBoxContent `bson:"contents,omitempty" json:"contents,omitempty"`
	UpdatedAt          time.Time           `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}

// PackingBoxContent describes one product quantity assigned to one box label.
type PackingBoxContent struct {
	BoxNo          int    `bson:"box_no" json:"box_no"`
	Zone           string `bson:"zone,omitempty" json:"zone,omitempty"`
	ProductSKUCode string `bson:"product_sku_code" json:"product_sku_code"`
	// Deprecated: use ProductSKUCode.
	ProductID   string `bson:"product_id,omitempty" json:"product_id,omitempty"`
	SKU         string `bson:"sku,omitempty" json:"sku,omitempty"`
	ProductName string `bson:"product_name,omitempty" json:"product_name,omitempty"`
	Qty         int    `bson:"qty" json:"qty"`
}
