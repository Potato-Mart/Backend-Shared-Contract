package template

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v35/pkg/contracts/common/geography"
)

// NotificationTemplateVersion freezes content and its review evidence at
// publication. PublishedVersion increases within (CountryCode, TemplateCode).
// SourceRevision identifies the mutable definition revision that was published.
// Restoring old content creates a new definition revision and publication;
// this record is never edited. Publication does not imply an active binding.
type NotificationTemplateVersion struct {
	CountryCode      geography.CountryCode `json:"country_code"`
	TemplateCode     string                `json:"template_code"`
	PublishedVersion int64                 `json:"published_version"`
	SourceRevision   int64                 `json:"source_revision"`
	Content          NotificationContent   `json:"content"`
	PublishedAt      time.Time             `json:"published_at"`
	PublishedBy      string                `json:"published_by"`
}
