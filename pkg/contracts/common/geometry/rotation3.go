package geometry

// Rotation3 expresses orientation as Euler angles in radians, applied in
// XYZ order. Use this for axis-aligned racks/shelves where a quaternion
// would be overkill. If a future renderer needs quaternions, a separate
// Quaternion type should be introduced rather than overloading this one.
type Rotation3 struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}
