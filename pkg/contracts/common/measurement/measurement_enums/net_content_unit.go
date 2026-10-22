package measurement_enums

// NetContentUnit is the closed vocabulary Supply accepts for a product's
// declared net content. Units are case-sensitive. Both measures of one net
// content declaration must share a dimension: mass (g, kg), volume (mL, L),
// length (m) or count (ea). measurement.MeasurementUnit remains an open
// string for other measurement contexts.
type NetContentUnit string

const (
	NetContentUnitGram       NetContentUnit = "g"
	NetContentUnitKilogram   NetContentUnit = "kg"
	NetContentUnitMillilitre NetContentUnit = "mL"
	NetContentUnitLitre      NetContentUnit = "L"
	NetContentUnitMetre      NetContentUnit = "m"
	NetContentUnitEach       NetContentUnit = "ea"
)

// IsValid reports whether u is a supported net content unit.
func (u NetContentUnit) IsValid() bool {
	switch u {
	case NetContentUnitGram, NetContentUnitKilogram, NetContentUnitMillilitre,
		NetContentUnitLitre, NetContentUnitMetre, NetContentUnitEach:
		return true
	default:
		return false
	}
}

func (u NetContentUnit) String() string { return string(u) }
