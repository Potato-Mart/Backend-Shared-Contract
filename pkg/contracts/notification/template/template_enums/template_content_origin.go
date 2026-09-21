package template_enums

// TemplateContentOrigin describes authored source or translated copy, including human refinements.
type TemplateContentOrigin string

const (
	TemplateContentOriginAuthored   TemplateContentOrigin = "authored"
	TemplateContentOriginTranslated TemplateContentOrigin = "translated"
)

// IsValid reports whether the value is a supported TemplateContentOrigin.
func (v TemplateContentOrigin) IsValid() bool {
	switch v {
	case TemplateContentOriginAuthored, TemplateContentOriginTranslated:
		return true
	}
	return false
}

// String returns the wire value.
func (v TemplateContentOrigin) String() string { return string(v) }
