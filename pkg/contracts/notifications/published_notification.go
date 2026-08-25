package notifications

import (
	"time"

	security "github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/notifications/notification_enums"
)

// PublishedNotification is the customer-safe in-app inbox projection. It
// deliberately excludes recipient identities, endpoint destinations, provider
// data, delivery attempts/failures, and every non-in-app content arm.
type PublishedNotification struct {
	ID          string                                     `json:"id"`
	TopicCode   string                                     `json:"topic_code"`
	Title       string                                     `json:"title"`
	Body        string                                     `json:"body"`
	ActionURL   string                                     `json:"action_url,omitempty"`
	Image       *security.ObjectMedia                      `json:"image,omitempty"`
	Status      notification_enums.InAppNotificationStatus `json:"status"`
	CreatedAt   time.Time                                  `json:"created_at"`
	ReadAt      *time.Time                                 `json:"read_at,omitempty"`
	DismissedAt *time.Time                                 `json:"dismissed_at,omitempty"`
	ExpiresAt   *time.Time                                 `json:"expires_at,omitempty"`
}
