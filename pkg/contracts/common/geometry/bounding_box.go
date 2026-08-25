package geometry

// BoundingBox is an axis-aligned bounding box defined by two corners.
// Use it for collision tests, frustum culling hints, and model bounds.
type BoundingBox struct {
	Min Vector3 `json:"min"`
	Max Vector3 `json:"max"`
}
