package template_enums

// EmailFontFamily selects a server-owned font stack without accepting remote fonts.
type EmailFontFamily string

const (
	EmailFontFamilySans  EmailFontFamily = "sans"
	EmailFontFamilySerif EmailFontFamily = "serif"
)

// IsValid reports whether the value is a supported EmailFontFamily.
func (v EmailFontFamily) IsValid() bool {
	switch v {
	case EmailFontFamilySans, EmailFontFamilySerif:
		return true
	}
	return false
}

// String returns the wire value.
func (v EmailFontFamily) String() string { return string(v) }
