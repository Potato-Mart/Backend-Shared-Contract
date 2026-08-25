package geometry

// Vector3 is a 3D point or direction in scene-space.
//
// Coordinates are stored as float64 in metres unless a contract that uses
// them documents otherwise. Use Size3D for whole-number millimetre sizes.
type Vector3 struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}
