package template_enums

// TemplatePlaceholderType identifies the scalar runtime value a placeholder requires.
type TemplatePlaceholderType string

const (
	TemplatePlaceholderTypeText      TemplatePlaceholderType = "text"
	TemplatePlaceholderTypeInteger   TemplatePlaceholderType = "integer"
	TemplatePlaceholderTypeMoney     TemplatePlaceholderType = "money"
	TemplatePlaceholderTypeTimestamp TemplatePlaceholderType = "timestamp"
	TemplatePlaceholderTypeURL       TemplatePlaceholderType = "url"
)

// IsValid reports whether the value is a supported TemplatePlaceholderType.
func (v TemplatePlaceholderType) IsValid() bool {
	switch v {
	case TemplatePlaceholderTypeText, TemplatePlaceholderTypeInteger, TemplatePlaceholderTypeMoney, TemplatePlaceholderTypeTimestamp, TemplatePlaceholderTypeURL:
		return true
	}
	return false
}

// String returns the wire value.
func (v TemplatePlaceholderType) String() string { return string(v) }
