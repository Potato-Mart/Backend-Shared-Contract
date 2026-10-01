package preference

import "github.com/Potato-Mart/Backend-Shared-Contract/v41/pkg/contracts/notification/core/notification_enums"

// NotificationChannelPreference groups topic choices under a channel. Services preserve mandatory transactional delivery
// policy, consent and unsubscribe rules independently of these choices.
type NotificationChannelPreference struct {
	Channel notification_enums.NotificationChannel `json:"channel"`
	Topics  []NotificationTopicPreference          `json:"topics,omitempty"`
}
