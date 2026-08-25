package geometry

// Size3D records width (X axis), height (Y axis), and depth (Z axis) in
// millimetres so it can be stored as integers and round-trips losslessly
// through JSON. Mirror of measurement.Dimensions for cases where depth matters.
type Size3D struct {
	WidthMM  int64 `json:"width_mm"`
	HeightMM int64 `json:"height_mm"`
	DepthMM  int64 `json:"depth_mm"`
}
