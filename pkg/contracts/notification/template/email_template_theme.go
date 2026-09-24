package template

import "github.com/Potato-Mart/Backend-Shared-Contract/v34/pkg/contracts/notification/template/template_enums"

// EmailTemplateTheme contains bounded presentation tokens shared by every
// locale. Colors are #RRGGBB strings. Notification validates colors, resolves
// explicit defaults at publication and maps tokens to accessible rendering.
// These fields affect the full preview/publication digest, not translation
// fingerprints. No arbitrary CSS, remote font or renderer configuration is held.
type EmailTemplateTheme struct {
	BackgroundColor string                          `json:"background_color"`
	ContentColor    string                          `json:"content_color"`
	TextColor       string                          `json:"text_color"`
	AccentColor     string                          `json:"accent_color"`
	FontFamily      template_enums.EmailFontFamily  `json:"font_family"`
	TextSize        template_enums.EmailTextSize    `json:"text_size"`
	Spacing         template_enums.EmailSpacing     `json:"spacing"`
	ButtonShape     template_enums.EmailButtonShape `json:"button_shape"`
}
