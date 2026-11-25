package classification_enums

// FreightMode is how a supplier's goods usually reach the warehouse: by sea,
// by air, locally, or a mix of sea and air. The empty value means not set.
type FreightMode string

const (
	FreightModeSea   FreightMode = "SEA"
	FreightModeAir   FreightMode = "AIR"
	FreightModeLocal FreightMode = "LOCAL"
	FreightModeMixed FreightMode = "MIXED"
)

// IsValid reports whether m is a known FreightMode.
func (m FreightMode) IsValid() bool {
	switch m {
	case FreightModeSea, FreightModeAir, FreightModeLocal, FreightModeMixed:
		return true
	default:
		return false
	}
}

func (m FreightMode) String() string { return string(m) }
