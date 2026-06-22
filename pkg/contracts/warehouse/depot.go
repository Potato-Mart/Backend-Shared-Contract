package warehouse

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v7/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v7/pkg/enums"
)

type Depot struct {
	ID        string `json:"id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	Address   string `json:"address,omitempty"`
	Phone     string `json:"phone,omitempty"`
	IsActive  bool   `json:"is_active"`
	SortOrder int    `json:"sort_order"`
	// LayoutID points at the published WarehouseLayout for this depot's
	// 3D viewer. Empty when no layout has been built yet.
	LayoutID string `json:"layout_id,omitempty"`

	common.AuditFields `bson:",inline"`
}

type PostcodeRule struct {
	ID       string `json:"id"`
	DepotID  string `json:"depot_id"`
	Postcode string `json:"postcode"`
	Priority int    `json:"priority"`
}

type DepotProduct struct {
	DepotID      string    `json:"depot_id"`
	ProductID    string    `json:"product_id"`
	StockQty     int       `json:"stock_qty"`
	IsAvailable  bool      `json:"is_available"`
	LocationCode string    `json:"location_code,omitempty"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// StockLocation is the inventory bin a product physically lives in.
//
// The 3D fields below are optional: services that do not render the
// warehouse can ignore them entirely. When LayoutNodeID is set, the
// authoritative 3D placement lives on that LayoutNode and the inline
// Transform/Size are a denormalised cache for renderers that fetch
// locations directly without joining the layout tree.
type StockLocation struct {
	ID           string            `json:"id"`
	DepotID      string            `json:"depot_id"`
	Code         string            `json:"code"`
	Name         string            `json:"name,omitempty"`
	Zone         enums.StorageType `json:"zone,omitempty"`
	IsActive     bool              `json:"is_active"`
	LayoutNodeID string            `json:"layout_node_id,omitempty"`
	Transform    *common.Transform `json:"transform,omitempty"`
	Size         *common.Size3D    `json:"size,omitempty"`
	Shape        enums.ShapeType   `json:"shape,omitempty"`
	Color        string            `json:"color,omitempty"`
}
