package warehouse

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v3/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v3/pkg/enums"
)

// WarehouseLayout is the root 3D scene description for a depot. There is
// at most one published layout per depot (Version is bumped per edit).
//
// MongoDB collection: "warehouse_layouts" with a unique index on
// (depot_id, version) and a partial index on (depot_id) where
// is_published=true to fetch the live scene cheaply.
type WarehouseLayout struct {
	ID              string         `json:"id"`
	DepotID         string         `json:"depot_id"`
	Name            string         `json:"name,omitempty"`
	Version         int            `json:"version"`
	Origin          common.Vector3 `json:"origin"`
	Size            common.Size3D  `json:"size"`
	UpAxis          string         `json:"up_axis,omitempty"`        // "Y" (default for Three.js) or "Z"
	UnitScale       float64        `json:"unit_scale,omitempty"`     // scene units per metre, default 1.0
	BackgroundColor string         `json:"background_color,omitempty"`
	FloorTextureURL string         `json:"floor_texture_url,omitempty"`
	GridSizeMM      int64          `json:"grid_size_mm,omitempty"`
	PrimaryModel    *ModelAsset    `json:"primary_model,omitempty"`  // optional GLB shell of the floor plan
	Walls           []LayoutWall   `json:"walls,omitempty"`
	Cameras         []CameraPreset `json:"cameras,omitempty"`
	Note            string         `json:"note,omitempty"`
	IsPublished     bool           `json:"is_published"`
	PublishedAt     *time.Time     `json:"published_at,omitempty"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
}

// LayoutNode is a single element in the warehouse 3D hierarchy:
// Zone > Aisle > Rack > Shelf > Bin. Each node references its layout
// and its parent, with a materialised Path so a subtree can be fetched
// in one query.
//
// MongoDB collection: "warehouse_layout_nodes" with indexes on
// (layout_id, parent_id), (layout_id, type), and a multikey index on
// "path" for "give me everything under aisle X" queries.
type LayoutNode struct {
	ID         string               `json:"id"`
	LayoutID   string               `json:"layout_id"`
	DepotID    string               `json:"depot_id"`
	ParentID   string               `json:"parent_id,omitempty"`
	Path       []string             `json:"path,omitempty"` // ancestor IDs from root to immediate parent
	Type       enums.LayoutNodeType `json:"type"`
	Code       string               `json:"code"` // human code, e.g. "A-12-3-2"
	Name       string               `json:"name,omitempty"`
	Storage    enums.StorageType    `json:"storage,omitempty"`
	Shape      enums.ShapeType      `json:"shape,omitempty"`
	Transform  common.Transform     `json:"transform"`
	Size       common.Size3D        `json:"size"`
	Color      string               `json:"color,omitempty"`     // hex like "#3b82f6"
	Model      *ModelAsset          `json:"model,omitempty"`     // optional model override for this node
	LocationID string               `json:"location_id,omitempty"` // links a BIN node to a StockLocation
	SortOrder  int                  `json:"sort_order,omitempty"`
	IsActive   bool                 `json:"is_active"`
	CreatedAt  time.Time            `json:"created_at"`
	UpdatedAt  time.Time            `json:"updated_at"`
}

// ModelAsset references a 3D model file stored elsewhere (object storage,
// CDN). Never embed mesh data here — keep documents small to stay well
// under the 16MB MongoDB document limit and keep reads fast.
type ModelAsset struct {
	ID        string              `json:"id,omitempty"`
	URL       string              `json:"url"`
	Format    enums.ModelFormat   `json:"format,omitempty"`
	SizeBytes int64               `json:"size_bytes,omitempty"`
	SHA256    string              `json:"sha256,omitempty"`
	Anchor    *common.Vector3     `json:"anchor,omitempty"` // offset to apply when placing the model
	Bounds    *common.BoundingBox `json:"bounds,omitempty"`
	Metadata  common.Metadata     `json:"metadata,omitempty"`
}

// CameraPreset is a saved viewpoint a user can jump to in the 3D viewer.
type CameraPreset struct {
	ID         string                 `json:"id"`
	Name       string                 `json:"name"`
	Position   common.Vector3         `json:"position"`
	Target     common.Vector3         `json:"target"`
	FOV        float64                `json:"fov,omitempty"`
	OrthoZoom  float64                `json:"ortho_zoom,omitempty"`
	Projection enums.CameraProjection `json:"projection,omitempty"`
	IsDefault  bool                   `json:"is_default,omitempty"`
}

// LayoutWall is a vertical wall segment used when a depot does not have
// a baked-in PrimaryModel and the renderer should draw walls from data.
type LayoutWall struct {
	ID          string         `json:"id,omitempty"`
	Start       common.Vector3 `json:"start"`
	End         common.Vector3 `json:"end"`
	HeightMM    int64          `json:"height_mm"`
	ThicknessMM int64          `json:"thickness_mm,omitempty"`
	Color       string         `json:"color,omitempty"`
}
