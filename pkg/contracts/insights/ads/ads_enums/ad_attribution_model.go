package ads_enums

// AdAttributionModel names which stored ad click a paid order is credited to.
// It is paired with a click window: only a click inside the window counts.
type AdAttributionModel string

const (
	// AdAttributionModelLastClick credits the most recent ad click before
	// the order. It is the agreed default.
	AdAttributionModelLastClick AdAttributionModel = "last_click"
	// AdAttributionModelFirstClick credits the earliest ad click inside the
	// window.
	AdAttributionModelFirstClick AdAttributionModel = "first_click"
)

// IsValid reports whether m is a known AdAttributionModel.
func (m AdAttributionModel) IsValid() bool {
	switch m {
	case AdAttributionModelLastClick, AdAttributionModelFirstClick:
		return true
	}
	return false
}

func (m AdAttributionModel) String() string { return string(m) }
