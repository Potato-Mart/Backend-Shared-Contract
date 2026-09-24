package template

import "github.com/Potato-Mart/Backend-Shared-Contract/v35/pkg/contracts/notification/core/notification_enums"

// NotificationContent is one channel's unrendered localized authoring document,
// reusable by template definitions and Notification-owned manual messages.
// SchemaVersion identifies document/token interpretation (initially 1), not
// draft or publication order. The owning record supplies country identity.
// SourceLanguage and variant languages are BCP 47 tags; schema 1 uses en,
// zh-TW and zh-CN, with any one as source. Notification enforces exactly those
// three complete variants and current target reviews before completed save.
// SourceFingerprint is server-calculated translation evidence, not a revision.
// EmailTheme is shared across locales and only applicable to email.
type NotificationContent struct {
	SchemaVersion     int                                    `json:"schema_version"`
	Channel           notification_enums.NotificationChannel `json:"channel"`
	SourceLanguage    string                                 `json:"source_language"`
	SourceFingerprint string                                 `json:"source_fingerprint"`
	Placeholders      []TemplatePlaceholder                  `json:"placeholders"`
	Variants          []LocalizedNotificationContent         `json:"variants"`
	EmailTheme        *EmailTemplateTheme                    `json:"email_theme,omitempty"`
}
