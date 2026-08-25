package warehouse

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/geometry"
)

// WarehouseLayout is the root 3D scene description for a depot.
type WarehouseLayout struct {
	ID              string           `json:"id"`
	DepotCode       string           `json:"depot_code"`
	Name            string           `json:"name,omitempty"`
	Version         int              `json:"version"`
	Origin          geometry.Vector3 `json:"origin"`
	Size            geometry.Size3D  `json:"size"`
	UpAxis          string           `json:"up_axis,omitempty"`    // "Y" (default for Three.js) or "Z"
	UnitScale       float64          `json:"unit_scale,omitempty"` // scene units per metre, default 1.0
	BackgroundColor string           `json:"background_color,omitempty"`
	FloorTextureURL string           `json:"floor_texture_url,omitempty"`
	GridSizeMM      int64            `json:"grid_size_mm,omitempty"`
	PrimaryModel    *ModelAsset      `json:"primary_model,omitempty"` // optional GLB shell of the floor plan
	Walls           []LayoutWall     `json:"walls,omitempty"`
	Cameras         []CameraPreset   `json:"cameras,omitempty"`
	Note            string           `json:"note,omitempty"`
	IsPublished     bool             `json:"is_published"`
	PublishedAt     *time.Time       `json:"published_at,omitempty"`

	audit.AuditFields
}
