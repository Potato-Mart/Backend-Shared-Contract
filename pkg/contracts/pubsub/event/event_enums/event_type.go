package event_enums

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
	EventTypeInvoiceIssued      EventType = "invoice.issued"
	EventTypeReceiptGenerated   EventType = "receipt.generated"
	EventTypeRefundRequested    EventType = "refund.requested"
	EventTypeRefundCompleted    EventType = "refund.completed"
	EventTypeRefundFailed       EventType = "refund.failed"

	EventTypeInventoryLotReceived              EventType = "inventory.lot_received"
	EventTypeInventoryStockBucketChanged       EventType = "inventory.stock_bucket_changed"
	EventTypeInventoryPackageConverted         EventType = "inventory.package_converted"
	EventTypeInventoryQualityAssessed          EventType = "inventory.quality_assessed"
	EventTypeInventoryReservationChanged       EventType = "inventory.reservation_changed"
	EventTypeInventoryStaged                   EventType = "inventory.staged"
	EventTypeInventorySold                     EventType = "inventory.sold"
	EventTypeInventoryDateMarkThresholdReached EventType = "inventory.date_mark_threshold_reached"
	EventTypeStockLocationAvailabilityChanged  EventType = "stock.location_availability_changed"

	EventTypeFulfilmentPackingUpdated  EventType = "fulfilment.packing_updated"
	EventTypeFulfilmentShipped         EventType = "fulfilment.shipped"
	EventTypeFulfilmentDelivered       EventType = "fulfilment.delivered"
	EventTypeFulfilmentCompleted       EventType = "fulfilment.completed"
	EventTypeFulfilmentTrackingUpdated EventType = "fulfilment.tracking_updated"

	EventTypeCustomerRegistered             EventType = "customer.registered"
	EventTypeCustomerProfileUpdated         EventType = "customer.profile_updated"
	EventTypeNotificationPreferencesChanged EventType = "notification.preferences_changed"
	EventTypeWalletGiftCardIssued           EventType = "wallet.gift_card_issued"

	EventTypeCatalogBaseCostChanged EventType = "catalog.base_cost_changed"
	EventTypeCatalogListingChanged  EventType = "catalog.listing_changed"

	EventTypeProductSalesPerformanceUpdated EventType = "product.sales_performance_updated"
	EventTypePromotionChanged               EventType = "promotion.changed"
	EventTypeCampaignChanged                EventType = "campaign.changed"

	EventTypeAnalyticsOrderFact   EventType = "analytics.order_fact"
	EventTypeAnalyticsPaymentFact EventType = "analytics.payment_fact"
	EventTypeAnalyticsRefundFact  EventType = "analytics.refund_fact"
)

func (t EventType) IsValid() bool {
	switch t {
	case EventTypeOrderCreated, EventTypeOrderPaid, EventTypeOrderStatusChanged,
		EventTypeOrderCancelled,
		EventTypePaymentCaptured, EventTypePaymentFailed, EventTypeInvoiceIssued,
		EventTypeReceiptGenerated,
		EventTypeRefundRequested, EventTypeRefundCompleted, EventTypeRefundFailed,
		EventTypeInventoryLotReceived, EventTypeInventoryStockBucketChanged,
		EventTypeInventoryPackageConverted, EventTypeInventoryQualityAssessed,
		EventTypeInventoryReservationChanged, EventTypeInventoryStaged,
		EventTypeInventorySold, EventTypeInventoryDateMarkThresholdReached,
		EventTypeStockLocationAvailabilityChanged,
		EventTypeFulfilmentPackingUpdated, EventTypeFulfilmentShipped,
		EventTypeFulfilmentDelivered, EventTypeFulfilmentCompleted,
		EventTypeFulfilmentTrackingUpdated,
		EventTypeCustomerRegistered, EventTypeCustomerProfileUpdated,
		EventTypeNotificationPreferencesChanged, EventTypeWalletGiftCardIssued,
		EventTypeCatalogBaseCostChanged, EventTypeCatalogListingChanged,
		EventTypeProductSalesPerformanceUpdated,
		EventTypePromotionChanged, EventTypeCampaignChanged,
		EventTypeAnalyticsOrderFact, EventTypeAnalyticsPaymentFact,
		EventTypeAnalyticsRefundFact:
		return true
	default:
		return false
	}
}

func (t EventType) String() string { return string(t) }
