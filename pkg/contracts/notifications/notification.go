package notifications

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/audit"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/notifications/notification_enums"
)

// Notification is the protected durable aggregate. Customer-facing callers
// receive PublishedNotification, never this record or its delivery endpoints.
type Notification struct {
	ID          string                                `json:"id"`
	EventID     string                                `json:"event_id,omitempty"`
	TopicCode   string                                `json:"topic_code"`
	Recipient   NotificationRecipient                 `json:"recipient"`
	Locale      string                                `json:"locale,omitempty"`
	References  []NotificationReference               `json:"references,omitempty"`
	Deliveries  []NotificationDelivery                `json:"deliveries"`
	Status      notification_enums.NotificationStatus `json:"status"`
	ScheduledAt *time.Time                            `json:"scheduled_at,omitempty"`
	ExpiresAt   *time.Time                            `json:"expires_at,omitempty"`

	// AuditFields provides the required created_at timestamp plus mutation
	// provenance for this protected durable record.
	audit.AuditFields
	security.DataProtectionFields
}
