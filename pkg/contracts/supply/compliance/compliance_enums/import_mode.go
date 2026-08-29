package compliance_enums

type ImportMode string

const (
	ImportModeAirCargo   ImportMode = "air_cargo"
	ImportModeSeaAmbient ImportMode = "sea_ambient"
	ImportModeSeaFrozen  ImportMode = "sea_frozen"
)

func (m ImportMode) IsValid() bool {
	switch m {
	case ImportModeAirCargo, ImportModeSeaAmbient, ImportModeSeaFrozen:
		return true
	}
	return false
}

func (m ImportMode) String() string { return string(m) }
