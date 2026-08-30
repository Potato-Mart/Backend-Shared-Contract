package core

import "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/notification/core/notification_enums"

// NotificationTopicChannel declares how a backend-managed topic supports one
// channel. A social-media channel is never eligible from DefaultEnabled alone:
// a user must also explicitly choose at least one destination code and grant
// consent for that destination.
type NotificationTopicChannel struct {
	Channel         notification_enums.NotificationChannel `json:"channel"`
	DefaultEnabled  bool                                   `json:"default_enabled"`
	OptOutAllowed   bool                                   `json:"opt_out_allowed"`
	ConsentRequired bool                                   `json:"consent_required"`
}
