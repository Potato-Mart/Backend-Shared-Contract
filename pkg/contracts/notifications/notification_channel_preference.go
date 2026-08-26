package notifications

import "github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/notifications/notification_enums"

// NotificationChannelPreference records a topic/channel choice. For social
// media, DestinationCodes is an explicit allow-list: an empty or absent list
// never means all configured destinations and makes social delivery ineligible.
type NotificationChannelPreference struct {
	Channel          notification_enums.NotificationChannel `json:"channel"`
	Enabled          bool                                   `json:"enabled"`
	DestinationCodes []string                               `json:"destination_codes,omitempty"`
}
