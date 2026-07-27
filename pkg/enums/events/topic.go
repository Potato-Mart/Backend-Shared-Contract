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
	EventTopicProductStats     EventTopic = "product-stats"
)

func (t EventTopic) IsValid() bool {
	switch t {
	case EventTopicOrderEvents, EventTopicPaymentEvents, EventTopicRefundEvents,
		EventTopicStockEvents, EventTopicFulfilmentEvents, EventTopicCustomerEvents,
		EventTopicProductStats:
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
	EventTypeOrderCreated                  EventType = "order.created"
	EventTypeOrderPaid                     EventType = "order.paid"
	EventTypeOrderStatusChanged            EventType = "order.status_changed"
	EventTypeOrderCancelled                EventType = "order.cancelled"
	EventTypeOrderPreorderAvailable        EventType = "order.preorder_available"
	EventTypeCheckoutCompensationRequested EventType = "checkout.compensation_requested"
	EventTypePaymentCaptured               EventType = "payment.captured"
	EventTypePaymentFailed                 EventType = "payment.failed"
	EventTypeInvoiceIssued                 EventType = "invoice.issued"
	EventTypeInvoiceDeliveryRequested      EventType = "invoice.delivery_requested"
	EventTypeRefundRequested               EventType = "refund.requested"
	EventTypeRefundCompleted               EventType = "refund.completed"
	EventTypeRefundFailed                  EventType = "refund.failed"

	EventTypeStockArrived  EventType = "stock.arrived"
	EventTypeStockAdjusted EventType = "stock.adjusted"

	EventTypeFulfilmentPackingUpdated  EventType = "fulfilment.packing_updated"
	EventTypeFulfilmentShipped         EventType = "fulfilment.shipped"
	EventTypeFulfilmentDelivered       EventType = "fulfilment.delivered"
	EventTypeFulfilmentCompleted       EventType = "fulfilment.completed"
	EventTypeFulfilmentTrackingUpdated EventType = "fulfilment.tracking_updated"

	EventTypeCustomerRegistered     EventType = "customer.registered"
	EventTypeCustomerProfileUpdated EventType = "customer.profile_updated"
	EventTypeCustomerConsentChanged EventType = "customer.consent_changed"
	EventTypeWalletGiftCardIssued   EventType = "wallet.gift_card_issued"

	EventTypeProductSalesPerformanceUpdated EventType = "product.sales_performance_updated"

	EventTypeAnalyticsOrderFact   EventType = "analytics.order_fact"
	EventTypeAnalyticsPaymentFact EventType = "analytics.payment_fact"
	EventTypeAnalyticsRefundFact  EventType = "analytics.refund_fact"
)

func (t EventType) IsValid() bool {
	switch t {
	case EventTypeOrderCreated, EventTypeOrderPaid, EventTypeOrderStatusChanged,
		EventTypeOrderCancelled, EventTypeOrderPreorderAvailable, EventTypeCheckoutCompensationRequested,
		EventTypePaymentCaptured, EventTypePaymentFailed, EventTypeInvoiceIssued, EventTypeInvoiceDeliveryRequested,
		EventTypeRefundRequested, EventTypeRefundCompleted, EventTypeRefundFailed,
		EventTypeStockArrived, EventTypeStockAdjusted,
		EventTypeFulfilmentPackingUpdated, EventTypeFulfilmentShipped,
		EventTypeFulfilmentDelivered, EventTypeFulfilmentCompleted,
		EventTypeFulfilmentTrackingUpdated,
		EventTypeCustomerRegistered, EventTypeCustomerProfileUpdated,
		EventTypeCustomerConsentChanged, EventTypeWalletGiftCardIssued,
		EventTypeProductSalesPerformanceUpdated,
		EventTypeAnalyticsOrderFact, EventTypeAnalyticsPaymentFact,
		EventTypeAnalyticsRefundFact:
		return true
	default:
		return false
	}
}

func (t EventType) String() string { return string(t) }
