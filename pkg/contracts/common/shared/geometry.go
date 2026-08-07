package shared

// Vector3 is a 3D point or direction in scene-space.
//
// Coordinates are stored as float64 in metres unless a contract that uses
// them documents otherwise. Use Size3D for whole-number millimetre sizes.
type Vector3 struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

// Rotation3 expresses orientation as Euler angles in radians, applied in
// XYZ order. Use this for axis-aligned racks/shelves where a quaternion
// would be overkill. If a future renderer needs quaternions, a separate
// Quaternion type should be introduced rather than overloading this one.
type Rotation3 struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

// Size3D records width (X axis), height (Y axis), and depth (Z axis) in
// millimetres so it can be stored as integers and round-trips losslessly
// through JSON. Mirror of common.Dimensions for cases where depth matters.
type Size3D struct {
	WidthMM  int64 `json:"width_mm"`
	HeightMM int64 `json:"height_mm"`
	DepthMM  int64 `json:"depth_mm"`
}

// BoundingBox is an axis-aligned bounding box defined by two corners.
// Use it for collision tests, frustum culling hints, and model bounds.
type BoundingBox struct {
	Min Vector3 `json:"min"`
	Max Vector3 `json:"max"`
}

// Transform is the position/rotation/scale of an object in scene-space.
//
// Position is required. Rotation and Scale are pointers so callers can
// omit them and let the renderer assume zero rotation and unit scale.
type Transform struct {
	Position Vector3    `json:"position"`
	Rotation *Rotation3 `json:"rotation,omitempty"`
	Scale    *Vector3   `json:"scale,omitempty"`
}
