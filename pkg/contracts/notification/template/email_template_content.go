package template

// EmailTemplateContent is a constrained email document, not HTML or CSS.
// Subject, PreviewText and textual block leaves are translatable. Stable block
// identities, order, kinds, links, media and system sections are preserved.
type EmailTemplateContent struct {
	Subject     string               `json:"subject"`
	PreviewText string               `json:"preview_text,omitempty"`
	Blocks      []EmailTemplateBlock `json:"blocks"`
}
