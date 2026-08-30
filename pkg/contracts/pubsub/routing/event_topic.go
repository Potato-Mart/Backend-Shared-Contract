package routing

// EventTopic identifies a Potato Mart Pub/Sub topic. Domain-fact topics are
// aggregate-family scoped; storefront-events carries only customer-safe
// invalidation facts. Every subscriber owns its own subscription and DLQ.
type EventTopic string

const (
	EventTopicOrderEvents      EventTopic = "order-events"
	EventTopicPaymentEvents    EventTopic = "payment-events"
	EventTopicRefundEvents     EventTopic = "refund-events"
	EventTopicStockEvents      EventTopic = "stock-events"
	EventTopicFulfilmentEvents EventTopic = "fulfilment-events"
	EventTopicCustomerEvents   EventTopic = "customer-events"
	EventTopicProductStats     EventTopic = "product-stats"
	EventTopicStorefrontEvents EventTopic = "storefront-events"
	EventTopicCatalogEvents    EventTopic = "catalog-events"
)

func (t EventTopic) IsValid() bool {
	switch t {
	case EventTopicOrderEvents, EventTopicPaymentEvents, EventTopicRefundEvents,
		EventTopicStockEvents, EventTopicFulfilmentEvents, EventTopicCustomerEvents,
		EventTopicProductStats, EventTopicStorefrontEvents, EventTopicCatalogEvents:
		return true
	default:
		return false
	}
}

func (t EventTopic) String() string { return string(t) }
