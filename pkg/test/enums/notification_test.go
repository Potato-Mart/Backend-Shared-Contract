package enums_test

import (
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v24/pkg/contracts/notifications/backinstock/backinstock_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v24/pkg/contracts/notifications/customer/customer_enums"
)

func TestNotificationEnums(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{
			name: "notification.CustomerNotificationStatus",
			valid: []stringEnum{
				customer_enums.CustomerNotificationStatusUnread,
				customer_enums.CustomerNotificationStatusRead,
				customer_enums.CustomerNotificationStatusDismissed,
			},
			invalid: customer_enums.CustomerNotificationStatus("__invalid__"),
		},
		{
			name: "notification.BackInStockChannel",
			valid: []stringEnum{
				backinstock_enums.BackInStockChannelEmail,
				backinstock_enums.BackInStockChannelSMS,
			},
			invalid: backinstock_enums.BackInStockChannel("__invalid__"),
		},
		{
			name: "notification.BackInStockStatus",
			valid: []stringEnum{
				backinstock_enums.BackInStockStatusPending,
				backinstock_enums.BackInStockStatusNotified,
				backinstock_enums.BackInStockStatusCancelled,
			},
			invalid: backinstock_enums.BackInStockStatus("__invalid__"),
		},
		{
			name: "notification.BackInStockCustomerType",
			valid: []stringEnum{
				backinstock_enums.BackInStockCustomerTypeRetail,
				backinstock_enums.BackInStockCustomerTypeWholesale,
			},
			invalid: backinstock_enums.BackInStockCustomerType("__invalid__"),
		},
		{
			name: "notification.CustomerNotificationTopic",
			valid: []stringEnum{
				customer_enums.CustomerNotificationTopicPreorderAvailable,
				customer_enums.CustomerNotificationTopicBackInStock,
				customer_enums.CustomerNotificationTopicOrderPlaced,
				customer_enums.CustomerNotificationTopicOrderConfirmed,
				customer_enums.CustomerNotificationTopicOrderCancelled,
				customer_enums.CustomerNotificationTopicPaymentReceived,
				customer_enums.CustomerNotificationTopicPaymentFailed,
				customer_enums.CustomerNotificationTopicPaymentRefunded,
				customer_enums.CustomerNotificationTopicPackingStarted,
				customer_enums.CustomerNotificationTopicOrderPacked,
				customer_enums.CustomerNotificationTopicOrderDispatched,
				customer_enums.CustomerNotificationTopicOrderDelivered,
				customer_enums.CustomerNotificationTopicInvoiceAvailable,
				customer_enums.CustomerNotificationTopicPromotionAvailable,
				customer_enums.CustomerNotificationTopicAnnouncement,
			},
			invalid: customer_enums.CustomerNotificationTopic("__invalid__"),
		},
		{
			name: "notification.CustomerNotificationChannel",
			valid: []stringEnum{
				customer_enums.CustomerNotificationChannelPortal,
				customer_enums.CustomerNotificationChannelEmail,
				customer_enums.CustomerNotificationChannelPush,
			},
			invalid: customer_enums.CustomerNotificationChannel("__invalid__"),
		},
		{
			name: "notification.CustomerNotificationDeliveryStatus",
			valid: []stringEnum{
				customer_enums.CustomerNotificationDeliveryStatusPending,
				customer_enums.CustomerNotificationDeliveryStatusDelivered,
				customer_enums.CustomerNotificationDeliveryStatusFailed,
			},
			invalid: customer_enums.CustomerNotificationDeliveryStatus("__invalid__"),
		},
	})
}
