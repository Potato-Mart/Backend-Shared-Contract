package template_enums

// EmailButtonShape selects a bounded button presentation.
type EmailButtonShape string

const (
	EmailButtonShapeSquare  EmailButtonShape = "square"
	EmailButtonShapeRounded EmailButtonShape = "rounded"
)

// IsValid reports whether the value is a supported EmailButtonShape.
func (v EmailButtonShape) IsValid() bool {
	switch v {
	case EmailButtonShapeSquare, EmailButtonShapeRounded:
		return true
	}
	return false
}

// String returns the wire value.
func (v EmailButtonShape) String() string { return string(v) }
