package warehouse

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/geometry"

	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/supply/warehouse/warehouse_enums"
)

// CameraPreset is a saved viewpoint a user can jump to in the 3D viewer.
type CameraPreset struct {
	ID         string                           `json:"id"`
	Name       string                           `json:"name"`
	Position   geometry.Vector3                 `json:"position"`
	Target     geometry.Vector3                 `json:"target"`
	FOV        float64                          `json:"fov,omitempty"`
	OrthoZoom  float64                          `json:"ortho_zoom,omitempty"`
	Projection warehouse_enums.CameraProjection `json:"projection,omitempty"`
	IsDefault  bool                             `json:"is_default,omitempty"`
}
