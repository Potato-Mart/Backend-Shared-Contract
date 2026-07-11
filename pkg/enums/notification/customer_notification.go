package notification

// CustomerNotificationTopic identifies a Management-owned customer message.
// The topic selects server-owned copy and delivery policy; callers cannot
// provide arbitrary email subjects or bodies.
type CustomerNotificationTopic string

const (
	CustomerNotificationTopicPreorderAvailable CustomerNotificationTopic = "preorder_available"
	CustomerNotificationTopicBackInStock       CustomerNotificationTopic = "back_in_stock"
)

func (t CustomerNotificationTopic) String() string { return string(t) }

func (t CustomerNotificationTopic) IsValid() bool {
	switch t {
	case CustomerNotificationTopicPreorderAvailable, CustomerNotificationTopicBackInStock:
		return true
	default:
		return false
	}
}

// CustomerNotificationChannel is a delivery route owned by Management.
type CustomerNotificationChannel string

const (
	CustomerNotificationChannelPortal CustomerNotificationChannel = "portal"
	CustomerNotificationChannelEmail  CustomerNotificationChannel = "email"
)

func (c CustomerNotificationChannel) String() string { return string(c) }

func (c CustomerNotificationChannel) IsValid() bool {
	switch c {
	case CustomerNotificationChannelPortal, CustomerNotificationChannelEmail:
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
