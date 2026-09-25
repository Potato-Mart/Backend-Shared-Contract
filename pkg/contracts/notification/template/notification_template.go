package template

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v36/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v36/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v36/pkg/contracts/common/localization"
)

// NotificationTemplate is a mutable country-owned template definition and its
// current authored content. CountryCode and Code form its immutable business
// identity. Revision is the positive optimistic-concurrency revision; it is not
// the document schema or a published version. Active does not authorize sending.
// Notification enforces complete, reviewed content on save and keeps Channel
// fixed after first publication.
type NotificationTemplate struct {
	ID          string                       `json:"id"`
	CountryCode geography.CountryCode        `json:"country_code"`
	Code        string                       `json:"code"`
	Name        []localization.LocalizedName `json:"name"`
	Active      bool                         `json:"active"`
	Revision    int64                        `json:"revision"`
	Content     NotificationContent          `json:"content"`

	audit.AuditFields
}
