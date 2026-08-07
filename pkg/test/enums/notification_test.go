package enums_test

import (
	"testing"

	backinstock "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/notifications/backinstock"
	notificationenum "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/notifications/customer"
)

func TestNotificationEnums(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{
			name: "notification.CustomerNotificationStatus",
			valid: []stringEnum{
				notificationenum.CustomerNotificationStatusUnread,
				notificationenum.CustomerNotificationStatusRead,
				notificationenum.CustomerNotificationStatusDismissed,
			},
			invalid: notificationenum.CustomerNotificationStatus("__invalid__"),
		},
		{
			name: "notification.BackInStockChannel",
			valid: []stringEnum{
				backinstock.BackInStockChannelEmail,
				backinstock.BackInStockChannelSMS,
			},
			invalid: backinstock.BackInStockChannel("__invalid__"),
		},
		{
			name: "notification.BackInStockStatus",
			valid: []stringEnum{
				backinstock.BackInStockStatusPending,
				backinstock.BackInStockStatusNotified,
				backinstock.BackInStockStatusCancelled,
			},
			invalid: backinstock.BackInStockStatus("__invalid__"),
		},
		{
			name: "notification.BackInStockCustomerType",
			valid: []stringEnum{
				backinstock.BackInStockCustomerTypeRetail,
				backinstock.BackInStockCustomerTypeWholesale,
			},
			invalid: backinstock.BackInStockCustomerType("__invalid__"),
		},
		{
			name: "notification.CustomerNotificationTopic",
			valid: []stringEnum{
				notificationenum.CustomerNotificationTopicPreorderAvailable,
				notificationenum.CustomerNotificationTopicBackInStock,
				notificationenum.CustomerNotificationTopicOrderPlaced,
				notificationenum.CustomerNotificationTopicOrderConfirmed,
				notificationenum.CustomerNotificationTopicOrderCancelled,
				notificationenum.CustomerNotificationTopicPaymentReceived,
				notificationenum.CustomerNotificationTopicPaymentFailed,
				notificationenum.CustomerNotificationTopicPaymentRefunded,
				notificationenum.CustomerNotificationTopicPackingStarted,
				notificationenum.CustomerNotificationTopicOrderPacked,
				notificationenum.CustomerNotificationTopicOrderDispatched,
				notificationenum.CustomerNotificationTopicOrderDelivered,
				notificationenum.CustomerNotificationTopicInvoiceAvailable,
				notificationenum.CustomerNotificationTopicPromotionAvailable,
				notificationenum.CustomerNotificationTopicAnnouncement,
			},
			invalid: notificationenum.CustomerNotificationTopic("__invalid__"),
		},
		{
			name: "notification.CustomerNotificationChannel",
			valid: []stringEnum{
				notificationenum.CustomerNotificationChannelPortal,
				notificationenum.CustomerNotificationChannelEmail,
				notificationenum.CustomerNotificationChannelPush,
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
