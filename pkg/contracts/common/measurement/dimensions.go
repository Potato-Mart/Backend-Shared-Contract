package measurement

// Dimensions records physical size using millimetres as the shared base unit.
// Use it for packaged product dimensions unless a contract explicitly says
// otherwise.
type Dimensions struct {
	WidthMM  int64 `json:"width_mm"`
	LengthMM int64 `json:"length_mm"`
	HeightMM int64 `json:"height_mm"`
}
