package template

import "github.com/Potato-Mart/Backend-Shared-Contract/v35/pkg/contracts/common/geography"

// NotificationTemplateReference identifies one immutable country-scoped
// publication. All fields are required; zero never means latest, global or
// fallback. References cannot resolve across countries.
type NotificationTemplateReference struct {
	CountryCode      geography.CountryCode `json:"country_code"`
	TemplateCode     string                `json:"template_code"`
	PublishedVersion int64                 `json:"published_version"`
}
