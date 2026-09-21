package template_enums

// EmailAlignment selects logical block alignment.
type EmailAlignment string

const (
	EmailAlignmentStart  EmailAlignment = "start"
	EmailAlignmentCenter EmailAlignment = "center"
	EmailAlignmentEnd    EmailAlignment = "end"
)

// IsValid reports whether the value is a supported EmailAlignment.
func (v EmailAlignment) IsValid() bool {
	switch v {
	case EmailAlignmentStart, EmailAlignmentCenter, EmailAlignmentEnd:
		return true
	}
	return false
}

// String returns the wire value.
func (v EmailAlignment) String() string { return string(v) }
