package template_enums

// EmailTextSize selects a bounded text-size token resolved by the renderer.
type EmailTextSize string

const (
	EmailTextSizeSmall  EmailTextSize = "small"
	EmailTextSizeNormal EmailTextSize = "normal"
	EmailTextSizeLarge  EmailTextSize = "large"
)

// IsValid reports whether the value is a supported EmailTextSize.
func (v EmailTextSize) IsValid() bool {
	switch v {
	case EmailTextSizeSmall, EmailTextSizeNormal, EmailTextSizeLarge:
		return true
	}
	return false
}

// String returns the wire value.
func (v EmailTextSize) String() string { return string(v) }
