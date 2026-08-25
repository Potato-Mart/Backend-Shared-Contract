package enums_test

import (
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/notifications/notification_enums"
)

func TestNotificationEnums(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{
			name: "notification.InAppNotificationStatus",
			valid: []stringEnum{
				notification_enums.InAppNotificationStatusUnread,
				notification_enums.InAppNotificationStatusRead,
				notification_enums.InAppNotificationStatusDismissed,
			},
			invalid: notification_enums.InAppNotificationStatus("__invalid__"),
		},
		{
			name: "notification.NotificationChannel",
			valid: []stringEnum{
				notification_enums.NotificationChannelEmail,
				notification_enums.NotificationChannelPush,
				notification_enums.NotificationChannelSMS,
				notification_enums.NotificationChannelInApp,
				notification_enums.NotificationChannelSocialMedia,
			},
			invalid: notification_enums.NotificationChannel("__invalid__"),
		},
		{
			name: "notification.NotificationStatus",
			valid: []stringEnum{
				notification_enums.NotificationStatusPending,
				notification_enums.NotificationStatusScheduled,
				notification_enums.NotificationStatusDispatching,
				notification_enums.NotificationStatusPartiallyDelivered,
				notification_enums.NotificationStatusDelivered,
				notification_enums.NotificationStatusFailed,
				notification_enums.NotificationStatusCancelled,
				notification_enums.NotificationStatusExpired,
			},
			invalid: notification_enums.NotificationStatus("__invalid__"),
		},
		{
			name: "notification.NotificationDeliveryStatus",
			valid: []stringEnum{
				notification_enums.NotificationDeliveryStatusPending,
				notification_enums.NotificationDeliveryStatusDispatching,
				notification_enums.NotificationDeliveryStatusDelivered,
				notification_enums.NotificationDeliveryStatusFailed,
				notification_enums.NotificationDeliveryStatusCancelled,
			},
			invalid: notification_enums.NotificationDeliveryStatus("__invalid__"),
		},
	})
}
