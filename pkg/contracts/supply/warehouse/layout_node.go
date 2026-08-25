package warehouse

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/geometry"

	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/supply/classification/classification_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/supply/warehouse/warehouse_enums"
)

// LayoutNode is a single element in the warehouse 3D hierarchy:
// Zone > Aisle > Rack > Shelf > Bin.
type LayoutNode struct {
	ID            string                           `json:"id"`
	DepotCode     string                           `json:"depot_code"`
	LayoutVersion int                              `json:"layout_version"`
	ParentCode    string                           `json:"parent_code,omitempty"`
	PathCodes     []string                         `json:"path_codes,omitempty"` // ancestor codes from root to immediate parent
	Type          warehouse_enums.LayoutNodeType   `json:"type"`
	Code          string                           `json:"code"` // human code, e.g. "A-12-3-2"
	Name          string                           `json:"name,omitempty"`
	Storage       classification_enums.StorageType `json:"storage,omitempty"`
	Shape         warehouse_enums.ShapeType        `json:"shape,omitempty"`
	Transform     geometry.Transform               `json:"transform"`
	Size          geometry.Size3D                  `json:"size"`
	Color         string                           `json:"color,omitempty"`         // hex like "#3b82f6"
	Model         *ModelAsset                      `json:"model,omitempty"`         // optional model override for this node
	LocationCode  string                           `json:"location_code,omitempty"` // links a BIN node to a StockLocation
	IsActive      bool                             `json:"is_active"`

	audit.AuditFields
}
