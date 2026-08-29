package warehouse

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/geometry"
)

// LayoutWall is a vertical wall segment used when a depot does not have
// a baked-in PrimaryModel and the renderer should draw walls from data.
type LayoutWall struct {
	ID          string           `json:"id,omitempty"`
	Start       geometry.Vector3 `json:"start"`
	End         geometry.Vector3 `json:"end"`
	HeightMM    int64            `json:"height_mm"`
	ThicknessMM int64            `json:"thickness_mm,omitempty"`
	Color       string           `json:"color,omitempty"`
}
