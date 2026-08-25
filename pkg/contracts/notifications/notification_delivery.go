package notifications

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/notifications/notification_enums"
)

// NotificationDelivery records exactly one channel endpoint and its attempts.
// Exactly one content arm matching Channel must be populated; services enforce
// that one-of invariant before persistence or dispatch. DestinationCode is an
// opaque endpoint reference, never an email address, phone number, device
// token, provider access token, or webhook secret.
type NotificationDelivery struct {
	ID              string                                        `json:"id"`
	Channel         notification_enums.NotificationChannel        `json:"channel"`
	DestinationCode string                                        `json:"destination_code"`
	Status          notification_enums.NotificationDeliveryStatus `json:"status"`
	AttemptCount    int                                           `json:"attempt_count"`
	LastAttemptAt   *time.Time                                    `json:"last_attempt_at,omitempty"`
	DeliveredAt     *time.Time                                    `json:"delivered_at,omitempty"`
	ErrorCode       string                                        `json:"error_code,omitempty"`
	ErrorMessage    string                                        `json:"error_message,omitempty"`
	Email           *EmailNotification                            `json:"email,omitempty"`
	Push            *PushNotification                             `json:"push,omitempty"`
	SMS             *SMSNotification                              `json:"sms,omitempty"`
	InApp           *InAppNotification                            `json:"in_app,omitempty"`
	SocialMedia     *SocialMediaNotification                      `json:"social_media,omitempty"`
}
