package ads_enums

// AdHealthSeverity grades one tracking health check. The contract carries the
// grade and its numbers; the wording shown to staff stays admin-owned.
type AdHealthSeverity string

const (
	// AdHealthSeverityOK means the check passed.
	AdHealthSeverityOK AdHealthSeverity = "ok"
	// AdHealthSeverityWarning means tracking is degraded but still usable.
	AdHealthSeverityWarning AdHealthSeverity = "warning"
	// AdHealthSeverityError means tracking is broken for the checked scope.
	AdHealthSeverityError AdHealthSeverity = "error"
)

// IsValid reports whether s is a known AdHealthSeverity.
func (s AdHealthSeverity) IsValid() bool {
	switch s {
	case AdHealthSeverityOK, AdHealthSeverityWarning, AdHealthSeverityError:
		return true
	}
	return false
}

func (s AdHealthSeverity) String() string { return string(s) }
