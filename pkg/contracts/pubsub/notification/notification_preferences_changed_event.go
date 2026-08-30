package notification

import "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/notification/core/notification_enums"

// NotificationPreferencesChangedEvent is emitted on customer-events after a
// centralized notification preference revision is stored. It carries changed
// identifiers only, never rendered content, raw destinations, or consent data.
type NotificationPreferencesChangedEvent struct {
	UserID              string                                   `json:"user_id"`
	AccountID           string                                   `json:"account_id,omitempty"`
	CustomerNumber      string                                   `json:"customer_number,omitempty"`
	PreferencesRevision int64                                    `json:"preferences_revision"`
	ChangedTopicCodes   []string                                 `json:"changed_topic_codes,omitempty"`
	ChangedChannels     []notification_enums.NotificationChannel `json:"changed_channels,omitempty"`
}
