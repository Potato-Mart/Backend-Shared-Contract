package identity

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v18/pkg/common"
)

// UserNotificationPreferences captures the account-level notification choices
// for a user. Services may treat a missing preference bundle as "use defaults".
type UserNotificationPreferences struct {
	Channels   UserNotificationChannels `json:"channels"`
	Topics     UserNotificationTopics   `json:"topics"`
	QuietHours *NotificationQuietHours  `json:"quiet_hours,omitempty"`
	UpdatedAt  *time.Time               `json:"updated_at,omitempty"`
}

// UserNotificationChannels controls which delivery routes the user allows.
type UserNotificationChannels struct {
	Email      bool                           `json:"email"`
	SMS        bool                           `json:"sms"`
	Push       bool                           `json:"push"`
	InApp      bool                           `json:"in_app"`
	SocialApps []NotificationSocialAppChannel `json:"social_apps,omitempty"`
}

// NotificationSocialAppChannel stores preferences for social or messaging app
// providers. AppKey should be a stable lowercase key such as "line", "wechat",
// "whatsapp", "telegram", or a tenant-specific integration key.
type NotificationSocialAppChannel struct {
	AppKey      string `json:"app_key"`
	DisplayName string `json:"display_name,omitempty"`
	Enabled     bool   `json:"enabled"`
}

// UserNotificationTopics controls which kinds of notifications the user wants.
type UserNotificationTopics struct {
	AccountUpdates  bool `json:"account_updates"`
	SecurityAlerts  bool `json:"security_alerts"`
	OrderUpdates    bool `json:"order_updates"`
	DeliveryUpdates bool `json:"delivery_updates"`
	InvoiceUpdates  bool `json:"invoice_updates"`
	Promotions      bool `json:"promotions"`
	SystemAlerts    bool `json:"system_alerts"`
}

// NotificationQuietHours lets clients mute non-urgent notifications during a
// local time window. Timezone should be an IANA name such as "Australia/Sydney".
type NotificationQuietHours struct {
	Enabled   bool             `json:"enabled"`
	StartTime common.TimeOfDay `json:"start_time,omitempty"`
	EndTime   common.TimeOfDay `json:"end_time,omitempty"`
	Timezone  string           `json:"timezone,omitempty"` // IANA name, e.g. "Australia/Sydney"
}
