package template

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v36/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v36/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v36/pkg/contracts/notification/core/notification_enums"
)

// NotificationTemplateBinding records a country/topic/channel/purpose selection
// of an immutable template publication. The owner permits at most one active
// binding for that tuple, and requires the reference country and content channel
// to match. PurposeCode is backend-managed; absence means the ordinary purpose,
// never a wildcard or an authorization grant. Revision is a mutable CAS revision.
type NotificationTemplateBinding struct {
	ID          string                                 `json:"id"`
	CountryCode geography.CountryCode                  `json:"country_code"`
	TopicCode   string                                 `json:"topic_code"`
	PurposeCode string                                 `json:"purpose_code,omitempty"`
	Channel     notification_enums.NotificationChannel `json:"channel"`
	Template    NotificationTemplateReference          `json:"template"`
	Active      bool                                   `json:"active"`
	Revision    int64                                  `json:"revision"`

	audit.AuditFields
}
