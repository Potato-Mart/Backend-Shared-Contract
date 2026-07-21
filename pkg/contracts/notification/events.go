package notification

import (
	"time"

	notificationenum "github.com/Potato-Mart/Backend-Shared-Contract/v18/pkg/enums/notification"
)

// NotificationEngagementEvent is emitted on the engagement-events topic when
// a customer notification is delivered, opened, or clicked. Only delivery has
// a publisher today; opened/clicked activate once tracking exists.
type NotificationEngagementEvent struct {
	NotificationID string                                       `json:"notification_id"`
	Topic          notificationenum.CustomerNotificationTopic   `json:"topic"`
	Channel        notificationenum.CustomerNotificationChannel `json:"channel"`
	CustomerNumber string                                       `json:"customer_number,omitempty"`
	Action         notificationenum.EngagementAction            `json:"action"`
	OccurredAt     time.Time                                    `json:"occurred_at"`
	RequestID      string                                       `json:"request_id,omitempty"`
}
