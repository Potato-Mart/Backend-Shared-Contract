package common

// Dimensions records physical size using millimetres as the shared base unit.
// Use it for packaged product dimensions unless a contract explicitly says
// otherwise.
type Dimensions struct {
	WidthMM  int64 `json:"width_mm"`
	LengthMM int64 `json:"length_mm"`
	HeightMM int64 `json:"height_mm"`
}

// Weight records physical weight using grams as the shared base unit.
type Weight struct {
	Grams int64 `json:"grams"`
}

// PhysicalPackage is a package-level shipping measurement snapshot.
type PhysicalPackage struct {
	Dimensions *Dimensions `json:"dimensions,omitempty"`
	Weight     *Weight     `json:"weight,omitempty"`
	Quantity   int         `json:"quantity,omitempty"`
}
