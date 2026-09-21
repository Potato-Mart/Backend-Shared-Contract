package template_enums

// EmailSpacing selects a bounded spacing token resolved by the renderer.
type EmailSpacing string

const (
	EmailSpacingCompact EmailSpacing = "compact"
	EmailSpacingNormal  EmailSpacing = "normal"
	EmailSpacingRelaxed EmailSpacing = "relaxed"
)

// IsValid reports whether the value is a supported EmailSpacing.
func (v EmailSpacing) IsValid() bool {
	switch v {
	case EmailSpacingCompact, EmailSpacingNormal, EmailSpacingRelaxed:
		return true
	}
	return false
}

// String returns the wire value.
func (v EmailSpacing) String() string { return string(v) }
