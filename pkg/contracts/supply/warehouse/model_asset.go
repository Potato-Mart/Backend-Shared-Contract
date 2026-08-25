package warehouse

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/geometry"

	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/metadata"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/supply/warehouse/warehouse_enums"
)

// ModelAsset is the JSON reference and rendering metadata for a 3D model file.
type ModelAsset struct {
	ID        string                      `json:"id,omitempty"`
	URL       string                      `json:"url"`
	Format    warehouse_enums.ModelFormat `json:"format,omitempty"`
	SizeBytes int64                       `json:"size_bytes,omitempty"`
	SHA256    string                      `json:"sha256,omitempty"`
	Anchor    *geometry.Vector3           `json:"anchor,omitempty"` // offset to apply when placing the model
	Bounds    *geometry.BoundingBox       `json:"bounds,omitempty"`
	Metadata  metadata.Metadata           `json:"metadata,omitempty"`
}
