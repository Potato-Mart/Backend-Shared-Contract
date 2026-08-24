package geography_enums

// GeographicScopeMode distinguishes rules that apply everywhere from rules
// that apply only to their inclusive geographic targets.
type GeographicScopeMode string

const (
	GeographicScopeModeGlobal   GeographicScopeMode = "GLOBAL"
	GeographicScopeModeTargeted GeographicScopeMode = "TARGETED"
)

func (m GeographicScopeMode) IsValid() bool {
	switch m {
	case GeographicScopeModeGlobal, GeographicScopeModeTargeted:
		return true
	default:
		return false
	}
}

func (m GeographicScopeMode) String() string { return string(m) }
