package preference

import "time"

// NotificationPreferences is the centralized preference-centre aggregate.
// UserID is required. When present, AccountID and CustomerNumber must resolve
// to that same user; services enforce that identity invariant.
type NotificationPreferences struct {
	ID             string                        `json:"id"`
	UserID         string                        `json:"user_id"`
	AccountID      string                        `json:"account_id,omitempty"`
	CustomerNumber string                        `json:"customer_number,omitempty"`
	Topics         []NotificationTopicPreference `json:"topics,omitempty"`
	Consents       []NotificationChannelConsent  `json:"consents,omitempty"`
	Revision       int64                         `json:"revision"`
	UpdatedAt      time.Time                     `json:"updated_at"`
}
