package eventsenum

// EventTopic identifies a Potato Mart Pub/Sub topic. One topic exists per
// aggregate family; every subscriber owns its own subscription and DLQ.
type EventTopic string

const (
	EventTopicOrderEvents      EventTopic = "order-events"
	EventTopicPaymentEvents    EventTopic = "payment-events"
	EventTopicRefundEvents     EventTopic = "refund-events"
	EventTopicStockEvents      EventTopic = "stock-events"
	EventTopicFulfilmentEvents EventTopic = "fulfilment-events"
	EventTopicCustomerEvents   EventTopic = "customer-events"
	EventTopicCatalogEvents    EventTopic = "catalog-events"
	EventTopicEngagementEvents EventTopic = "engagement-events"
	EventTopicProductStats     EventTopic = "product-stats"
)

func (t EventTopic) IsValid() bool {
	switch t {
	case EventTopicOrderEvents, EventTopicPaymentEvents, EventTopicRefundEvents,
		EventTopicStockEvents, EventTopicFulfilmentEvents, EventTopicCustomerEvents,
		EventTopicCatalogEvents, EventTopicEngagementEvents, EventTopicProductStats:
		return true
	default:
		return false
	}
}

func (t EventTopic) String() string { return string(t) }

// EventType names a concrete happening carried inside an EventEnvelope.
// Values are namespaced by aggregate family; new types are additive.
type EventType string

const (
	EventTypeOrderCreated       EventType = "order.created"
	EventTypeOrderPaid          EventType = "order.paid"
	EventTypeOrderStatusChanged EventType = "order.status_changed"
	EventTypeOrderCancelled     EventType = "order.cancelled"
	EventTypePaymentCaptured    EventType = "payment.captured"
	EventTypePaymentFailed      EventType = "payment.failed"
	EventTypeRefundRequested    EventType = "refund.requested"
	EventTypeRefundCompleted    EventType = "refund.completed"
	EventTypeRefundFailed       EventType = "refund.failed"
)

func (t EventType) IsValid() bool {
	switch t {
	case EventTypeOrderCreated, EventTypeOrderPaid, EventTypeOrderStatusChanged,
		EventTypeOrderCancelled, EventTypePaymentCaptured, EventTypePaymentFailed,
		EventTypeRefundRequested, EventTypeRefundCompleted, EventTypeRefundFailed:
		return true
	default:
		return false
	}
}

func (t EventType) String() string { return string(t) }
