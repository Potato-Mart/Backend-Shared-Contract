package template

import "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/notification/template/template_enums"

// EmailBlockStyle is a bounded block presentation override. Omitted tokens
// inherit the enclosing email theme/renderer defaults. All locales retain the
// same style for a given block ID; a translator cannot change it.
type EmailBlockStyle struct {
	Alignment template_enums.EmailAlignment `json:"alignment,omitempty"`
	Spacing   template_enums.EmailSpacing   `json:"spacing,omitempty"`
}
