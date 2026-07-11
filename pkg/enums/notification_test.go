package enums_test

import (
	"testing"

	notificationenum "github.com/Potato-Mart/Backend-Shared-Contract/v15/pkg/enums/notification"
)

func TestNotificationEnums(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{
			name: "notification.BackInStockChannel",
			valid: []stringEnum{
				notificationenum.BackInStockChannelEmail,
				notificationenum.BackInStockChannelSMS,
			},
			invalid: notificationenum.BackInStockChannel("__invalid__"),
		},
		{
			name: "notification.BackInStockStatus",
			valid: []stringEnum{
				notificationenum.BackInStockStatusPending,
				notificationenum.BackInStockStatusNotified,
				notificationenum.BackInStockStatusCancelled,
			},
			invalid: notificationenum.BackInStockStatus("__invalid__"),
		},
		{
			name: "notification.BackInStockCustomerType",
			valid: []stringEnum{
				notificationenum.BackInStockCustomerTypeRetail,
				notificationenum.BackInStockCustomerTypeWholesale,
			},
			invalid: notificationenum.BackInStockCustomerType("__invalid__"),
		},
		{
			name: "notification.CustomerNotificationTopic",
			valid: []stringEnum{
				notificationenum.CustomerNotificationTopicPreorderAvailable,
				notificationenum.CustomerNotificationTopicBackInStock,
			},
			invalid: notificationenum.CustomerNotificationTopic("__invalid__"),
		},
		{
			name: "notification.CustomerNotificationChannel",
			valid: []stringEnum{
				notificationenum.CustomerNotificationChannelPortal,
				notificationenum.CustomerNotificationChannelEmail,
			},
			invalid: notificationenum.CustomerNotificationChannel("__invalid__"),
		},
		{
			name: "notification.CustomerNotificationDeliveryStatus",
			valid: []stringEnum{
				notificationenum.CustomerNotificationDeliveryStatusPending,
				notificationenum.CustomerNotificationDeliveryStatusDelivered,
				notificationenum.CustomerNotificationDeliveryStatusFailed,
			},
			invalid: notificationenum.CustomerNotificationDeliveryStatus("__invalid__"),
		},
	})
}
