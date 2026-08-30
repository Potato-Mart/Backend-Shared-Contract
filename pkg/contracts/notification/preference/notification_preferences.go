package preference

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/audit"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/security"
)

// NotificationPreferences is the latest-state preference-centre aggregate,
// not the legal system of record. UserID is required. When present, AccountID
// and CustomerNumber must resolve to that same user; services enforce that
// identity invariant.
type NotificationPreferences struct {
	ID             string                        `json:"id"`
	UserID         string                        `json:"user_id"`
	AccountID      string                        `json:"account_id,omitempty"`
	CustomerNumber string                        `json:"customer_number,omitempty"`
	Topics         []NotificationTopicPreference `json:"topics,omitempty"`
	Consents       []NotificationChannelConsent  `json:"consents,omitempty"`
	Revision       int64                         `json:"revision"`

	audit.AuditFields
	security.DataProtectionFields
}
