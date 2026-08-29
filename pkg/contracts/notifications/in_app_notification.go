package notifications

import (
	"time"

	security "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/notifications/notification_enums"
)

// InAppNotification is the renderable content and read state for one in-app
// delivery.
type InAppNotification struct {
	Title       string                                     `json:"title"`
	Body        string                                     `json:"body"`
	ActionURL   string                                     `json:"action_url,omitempty"`
	Image       *security.ObjectMedia                      `json:"image,omitempty"`
	Status      notification_enums.InAppNotificationStatus `json:"status"`
	ReadAt      *time.Time                                 `json:"read_at,omitempty"`
	DismissedAt *time.Time                                 `json:"dismissed_at,omitempty"`
}
