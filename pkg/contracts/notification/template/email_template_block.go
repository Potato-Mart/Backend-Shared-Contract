package template

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v35/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v35/pkg/contracts/notification/template/template_enums"
)

// EmailTemplateBlock has a stable document-unique ID and a kind-specific payload:
// heading: Text and Level (1..3); paragraph: Text; button: Text and ActionURL;
// image: Media and AltText (explicit empty alt is decorative); divider: no
// payload; system_section: SectionCode only. Style is an optional presentation
// value for editable blocks; protected system sections use server-owned style.
// Notification rejects irrelevant fields and preserves nontext values across
// translations. SectionCode is an open owner-managed identifier with no runtime
// payload; Notification owns required placement/count and protected rendering.
type EmailTemplateBlock struct {
	ID          string                                `json:"id"`
	Kind        template_enums.EmailTemplateBlockKind `json:"kind"`
	Text        *string                               `json:"text,omitempty"`
	Level       *int                                  `json:"level,omitempty"`
	ActionURL   string                                `json:"action_url,omitempty"`
	Media       *security.ObjectMedia                 `json:"media,omitempty"`
	AltText     *string                               `json:"alt_text,omitempty"`
	SectionCode string                                `json:"section_code,omitempty"`
	Style       *EmailBlockStyle                      `json:"style,omitempty"`
}
