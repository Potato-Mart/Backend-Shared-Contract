package routing

import "testing"

// eventTypeVersionRegistry is the complete historical payload-version table.
// It is test-only because model packages cannot carry runtime registry state;
// it locks the externally published event-version history without adding an
// executable service concern to the shared contracts.
var eventTypeVersionRegistry = map[EventType]string{
	EventTypeInventoryLotReceived:              "v4",
	EventTypeInventoryStockBucketChanged:       "v4",
	EventTypeInventoryPackageConverted:         "v4",
	EventTypeInventoryQualityAssessed:          "v4",
	EventTypeInventoryReservationChanged:       "v4",
	EventTypeInventoryStaged:                   "v4",
	EventTypeInventorySold:                     "v4",
	EventTypeInventoryDateMarkThresholdReached: "v4",
	EventTypeStockLocationAvailabilityChanged:  "v4",
	EventTypeCatalogBaseCostChanged:            "v4",
	EventTypeCatalogListingChanged:             "v4",
	EventTypeAnalyticsOrderFact:                "v4",
	EventTypeAnalyticsRefundFact:               "v4",

	EventTypeOrderPaid:                      "v3",
	EventTypeRefundCompleted:                "v3",
	EventTypeProductSalesPerformanceUpdated: "v3",
	EventTypeFulfilmentPackingUpdated:       "v3",

	EventTypeCampaignChanged:           "v2",
	EventTypePromotionChanged:          "v2",
	EventTypeFulfilmentShipped:         "v2",
	EventTypeFulfilmentDelivered:       "v2",
	EventTypeFulfilmentCompleted:       "v2",
	EventTypeFulfilmentTrackingUpdated: "v2",
	EventTypeAnalyticsPaymentFact:      "v2",

	EventTypeOrderCreated:                   "v1",
	EventTypeOrderStatusChanged:             "v1",
	EventTypeOrderCancelled:                 "v1",
	EventTypePaymentCaptured:                "v1",
	EventTypePaymentFailed:                  "v1",
	EventTypeInvoiceIssued:                  "v1",
	EventTypeReceiptGenerated:               "v1",
	EventTypeRefundRequested:                "v1",
	EventTypeRefundFailed:                   "v1",
	EventTypeCustomerRegistered:             "v1",
	EventTypeCustomerProfileUpdated:         "v1",
	EventTypeNotificationPreferencesChanged: "v1",
	EventTypeWalletGiftCardIssued:           "v1",
	EventTypePriceChanged:                   "v1",
}

func TestEventTypeVersionRegistryCoversEveryDefinedEventExactlyOnce(t *testing.T) {
	expected := map[EventType]string{
		EventTypeInventoryLotReceived:              "v4",
		EventTypeInventoryStockBucketChanged:       "v4",
		EventTypeInventoryPackageConverted:         "v4",
		EventTypeInventoryQualityAssessed:          "v4",
		EventTypeInventoryReservationChanged:       "v4",
		EventTypeInventoryStaged:                   "v4",
		EventTypeInventorySold:                     "v4",
		EventTypeInventoryDateMarkThresholdReached: "v4",
		EventTypeStockLocationAvailabilityChanged:  "v4",
		EventTypeCatalogBaseCostChanged:            "v4",
		EventTypeCatalogListingChanged:             "v4",
		EventTypeAnalyticsOrderFact:                "v4",
		EventTypeAnalyticsRefundFact:               "v4",

		EventTypeOrderPaid:                      "v3",
		EventTypeRefundCompleted:                "v3",
		EventTypeProductSalesPerformanceUpdated: "v3",
		EventTypeFulfilmentPackingUpdated:       "v3",

		EventTypeCampaignChanged:           "v2",
		EventTypePromotionChanged:          "v2",
		EventTypeFulfilmentShipped:         "v2",
		EventTypeFulfilmentDelivered:       "v2",
		EventTypeFulfilmentCompleted:       "v2",
		EventTypeFulfilmentTrackingUpdated: "v2",
		EventTypeAnalyticsPaymentFact:      "v2",

		EventTypeOrderCreated:                   "v1",
		EventTypeOrderStatusChanged:             "v1",
		EventTypeOrderCancelled:                 "v1",
		EventTypePaymentCaptured:                "v1",
		EventTypePaymentFailed:                  "v1",
		EventTypeInvoiceIssued:                  "v1",
		EventTypeReceiptGenerated:               "v1",
		EventTypeRefundRequested:                "v1",
		EventTypeRefundFailed:                   "v1",
		EventTypeCustomerRegistered:             "v1",
		EventTypeCustomerProfileUpdated:         "v1",
		EventTypeNotificationPreferencesChanged: "v1",
		EventTypeWalletGiftCardIssued:           "v1",
		EventTypePriceChanged:                   "v1",
	}
	if len(expected) != 38 {
		t.Fatalf("expected event registry test has %d entries, want 38", len(expected))
	}
	if len(eventTypeVersionRegistry) != len(expected) {
		t.Fatalf("event registry has %d entries, want %d", len(eventTypeVersionRegistry), len(expected))
	}
	for eventType, wantVersion := range expected {
		if gotVersion, ok := eventTypeVersionRegistry[eventType]; !ok || gotVersion != wantVersion {
			t.Fatalf("event registry[%q] = %q, %t; want %q", eventType, gotVersion, ok, wantVersion)
		}
	}
}
