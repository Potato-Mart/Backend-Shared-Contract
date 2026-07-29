package notification

// CustomerNotificationStatus is the customer-visible inbox lifecycle. Reading
// a notification is a terminal dismissal in the customer notification centre.
type CustomerNotificationStatus string

const (
	CustomerNotificationStatusUnread    CustomerNotificationStatus = "unread"
	CustomerNotificationStatusDismissed CustomerNotificationStatus = "dismissed"
)

func (s CustomerNotificationStatus) String() string { return string(s) }

func (s CustomerNotificationStatus) IsValid() bool {
	switch s {
	case CustomerNotificationStatusUnread, CustomerNotificationStatusDismissed:
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
	CustomerNotificationTopicInvoiceAvailable   CustomerNotificationTopic = "invoice_available"
	CustomerNotificationTopicPromotionAvailable CustomerNotificationTopic = "promotion_available"
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
		CustomerNotificationTopicInvoiceAvailable,
		CustomerNotificationTopicPromotionAvailable,
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
)

func (c CustomerNotificationChannel) String() string { return string(c) }

func (c CustomerNotificationChannel) IsValid() bool {
	switch c {
	case CustomerNotificationChannelPortal, CustomerNotificationChannelEmail,
		CustomerNotificationChannelPush:
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
