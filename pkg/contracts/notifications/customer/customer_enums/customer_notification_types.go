package customer_enums

// CustomerNotificationStatus is the customer-visible inbox lifecycle. Read
// notifications remain visible until the customer explicitly dismisses them.
type CustomerNotificationStatus string

const (
	CustomerNotificationStatusUnread    CustomerNotificationStatus = "unread"
	CustomerNotificationStatusRead      CustomerNotificationStatus = "read"
	CustomerNotificationStatusDismissed CustomerNotificationStatus = "dismissed"
)

func (s CustomerNotificationStatus) String() string { return string(s) }

func (s CustomerNotificationStatus) IsValid() bool {
	switch s {
	case CustomerNotificationStatusUnread, CustomerNotificationStatusRead,
		CustomerNotificationStatusDismissed:
		return true
	default:
		return false
	}
}

// CustomerNotificationTopic identifies a Customers-owned customer message.
// The topic selects server-owned copy and delivery policy; callers cannot
// provide arbitrary email subjects or bodies.
type CustomerNotificationTopic string

const (
	CustomerNotificationTopicPreorderAvailable  CustomerNotificationTopic = "preorder_available"
	CustomerNotificationTopicBackInStock        CustomerNotificationTopic = "back_in_stock"
	CustomerNotificationTopicOrderPlaced        CustomerNotificationTopic = "order_placed"
	CustomerNotificationTopicOrderConfirmed     CustomerNotificationTopic = "order_confirmed"
	CustomerNotificationTopicOrderCancelled     CustomerNotificationTopic = "order_cancelled"
	CustomerNotificationTopicPaymentReceived    CustomerNotificationTopic = "payment_received"
	CustomerNotificationTopicPaymentFailed      CustomerNotificationTopic = "payment_failed"
	CustomerNotificationTopicPaymentRefunded    CustomerNotificationTopic = "payment_refunded"
	CustomerNotificationTopicPackingStarted     CustomerNotificationTopic = "packing_started"
	CustomerNotificationTopicOrderPacked        CustomerNotificationTopic = "order_packed"
	CustomerNotificationTopicOrderDispatched    CustomerNotificationTopic = "order_dispatched"
	CustomerNotificationTopicOrderDelivered     CustomerNotificationTopic = "order_delivered"
	CustomerNotificationTopicOrderCompleted     CustomerNotificationTopic = "order_completed"
	CustomerNotificationTopicInvoiceAvailable   CustomerNotificationTopic = "invoice_available"
	CustomerNotificationTopicReceiptAvailable   CustomerNotificationTopic = "receipt_available"
	CustomerNotificationTopicPromotionAvailable CustomerNotificationTopic = "promotion_available"
	CustomerNotificationTopicNewProduct         CustomerNotificationTopic = "new_product"
	CustomerNotificationTopicAnnouncement       CustomerNotificationTopic = "announcement"
)

func (t CustomerNotificationTopic) String() string { return string(t) }

func (t CustomerNotificationTopic) IsValid() bool {
	switch t {
	case CustomerNotificationTopicPreorderAvailable,
		CustomerNotificationTopicBackInStock,
		CustomerNotificationTopicOrderPlaced,
		CustomerNotificationTopicOrderConfirmed,
		CustomerNotificationTopicOrderCancelled,
		CustomerNotificationTopicPaymentReceived,
		CustomerNotificationTopicPaymentFailed,
		CustomerNotificationTopicPaymentRefunded,
		CustomerNotificationTopicPackingStarted,
		CustomerNotificationTopicOrderPacked,
		CustomerNotificationTopicOrderDispatched,
		CustomerNotificationTopicOrderDelivered,
		CustomerNotificationTopicOrderCompleted,
		CustomerNotificationTopicInvoiceAvailable,
		CustomerNotificationTopicReceiptAvailable,
		CustomerNotificationTopicPromotionAvailable,
		CustomerNotificationTopicNewProduct,
		CustomerNotificationTopicAnnouncement:
		return true
	default:
		return false
	}
}

// CustomerNotificationChannel is a delivery route owned by Customers.
type CustomerNotificationChannel string

const (
	CustomerNotificationChannelPortal CustomerNotificationChannel = "portal"
	CustomerNotificationChannelEmail  CustomerNotificationChannel = "email"
	CustomerNotificationChannelPush   CustomerNotificationChannel = "push"
	CustomerNotificationChannelSMS    CustomerNotificationChannel = "sms"
)

func (c CustomerNotificationChannel) String() string { return string(c) }

func (c CustomerNotificationChannel) IsValid() bool {
	switch c {
	case CustomerNotificationChannelPortal, CustomerNotificationChannelEmail,
		CustomerNotificationChannelPush, CustomerNotificationChannelSMS:
		return true
	default:
		return false
	}
}

// CustomerNotificationDeliveryStatus is the durable state of one channel.
type CustomerNotificationDeliveryStatus string

const (
	CustomerNotificationDeliveryStatusPending   CustomerNotificationDeliveryStatus = "pending"
	CustomerNotificationDeliveryStatusDelivered CustomerNotificationDeliveryStatus = "delivered"
	CustomerNotificationDeliveryStatusFailed    CustomerNotificationDeliveryStatus = "failed"
)

func (s CustomerNotificationDeliveryStatus) String() string { return string(s) }

func (s CustomerNotificationDeliveryStatus) IsValid() bool {
	switch s {
	case CustomerNotificationDeliveryStatusPending,
		CustomerNotificationDeliveryStatusDelivered,
		CustomerNotificationDeliveryStatusFailed:
		return true
	default:
		return false
	}
}
