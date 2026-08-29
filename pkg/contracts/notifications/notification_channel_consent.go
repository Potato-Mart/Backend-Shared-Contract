package notifications

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/notifications/notification_enums"
)

// NotificationChannelConsent records legal or policy consent evidence. Social
// media consent is destination-scoped and requires DestinationCode; non-social
// channel consent leaves DestinationCode empty.
type NotificationChannelConsent struct {
	Channel         notification_enums.NotificationChannel `json:"channel"`
	DestinationCode string                                 `json:"destination_code,omitempty"`
	Granted         bool                                   `json:"granted"`
	Source          string                                 `json:"source,omitempty"`
	ChangedAt       time.Time                              `json:"changed_at"`
}
