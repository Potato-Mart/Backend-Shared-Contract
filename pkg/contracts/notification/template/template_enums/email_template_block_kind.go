package template_enums

// EmailTemplateBlockKind identifies a constrained email block's payload shape.
type EmailTemplateBlockKind string

const (
	EmailTemplateBlockKindHeading       EmailTemplateBlockKind = "heading"
	EmailTemplateBlockKindParagraph     EmailTemplateBlockKind = "paragraph"
	EmailTemplateBlockKindButton        EmailTemplateBlockKind = "button"
	EmailTemplateBlockKindImage         EmailTemplateBlockKind = "image"
	EmailTemplateBlockKindDivider       EmailTemplateBlockKind = "divider"
	EmailTemplateBlockKindSystemSection EmailTemplateBlockKind = "system_section"
)

// IsValid reports whether the value is a supported EmailTemplateBlockKind.
func (v EmailTemplateBlockKind) IsValid() bool {
	switch v {
	case EmailTemplateBlockKindHeading, EmailTemplateBlockKindParagraph, EmailTemplateBlockKindButton, EmailTemplateBlockKindImage, EmailTemplateBlockKindDivider, EmailTemplateBlockKindSystemSection:
		return true
	}
	return false
}

// String returns the wire value.
func (v EmailTemplateBlockKind) String() string { return string(v) }
