package geometry

// Transform is the position/rotation/scale of an object in scene-space.
//
// Position is required. Rotation and Scale are pointers so callers can
// omit them and let the renderer assume zero rotation and unit scale.
type Transform struct {
	Position Vector3    `json:"position"`
	Rotation *Rotation3 `json:"rotation,omitempty"`
	Scale    *Vector3   `json:"scale,omitempty"`
}
