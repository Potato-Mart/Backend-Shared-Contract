package notifications

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/localization"
)

// NotificationTopic is a backend-managed, provider-neutral notification
// category. Code is intentionally open: services create topics without a
// shared-contract release.
type NotificationTopic struct {
	ID          string                              `json:"id"`
	Code        string                              `json:"code"`
	Name        []localization.LocalizedName        `json:"name"`
	Description []localization.LocalizedDescription `json:"description,omitempty"`
	Channels    []NotificationTopicChannel          `json:"channels"`
	Active      bool                                `json:"active"`
	Revision    int64                               `json:"revision"`

	audit.AuditFields
}
