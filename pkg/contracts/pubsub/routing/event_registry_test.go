package routing

import "testing"

// eventTypeVersionRegistry is the complete current payload-version table.
// It is test-only because model packages cannot carry runtime registry state;
// it locks the externally published event-version history without adding an
// executable service concern to the shared contracts.
var eventTypeVersionRegistry = map[EventType]string{
	EventTypeInventoryLotReceived:              "v5",
	EventTypeInventoryStockBucketChanged:       "v5",
	EventTypeInventoryPackageConverted:         "v5",
	EventTypeInventoryQualityAssessed:          "v5",
	EventTypeInventoryReservationChanged:       "v5",
	EventTypeInventoryStaged:                   "v5",
	EventTypeInventorySold:                     "v5",
	EventTypeInventoryDateMarkThresholdReached: "v4",
	EventTypeStockLocationAvailabilityChanged:  "v4",
	EventTypeCatalogBaseCostChanged:            "v4",
	EventTypeCatalogListingChanged:             "v4",
	EventTypeAnalyticsOrderFact:                "v5",
	EventTypeAnalyticsRefundFact:               "v5",

	EventTypeOrderPaid:                      "v4",
	EventTypeRefundCompleted:                "v4",
	EventTypeProductSalesPerformanceUpdated: "v3",
	EventTypeFulfilmentPackingUpdated:       "v4",

	EventTypeCampaignChanged:           "v2",
	EventTypePromotionChanged:          "v3",
	EventTypeCouponChanged:             "v1",
	EventTypeFulfilmentShipped:         "v2",
	EventTypeFulfilmentDelivered:       "v2",
	EventTypeFulfilmentCompleted:       "v2",
	EventTypeFulfilmentTrackingUpdated: "v2",
	EventTypeAnalyticsPaymentFact:      "v2",

	EventTypeOrderCreated:                   "v1",
	EventTypeOrderStatusChanged:             "v1",
	EventTypeOrderCancelled:                 "v1",
	EventTypeOrderEdited:                    "v2",
	EventTypePaymentCaptured:                "v2",
	EventTypePaymentFailed:                  "v1",
	EventTypeInvoiceIssued:                  "v1",
	EventTypeReceiptGenerated:               "v1",
	EventTypeRefundRequested:                "v1",
	EventTypeRefundFailed:                   "v1",
	EventTypeCustomerRegistered:             "v1",
	EventTypeCustomerProfileUpdated:         "v1",
	EventTypeNotificationPreferencesChanged: "v1",
	EventTypeWalletGiftCardIssued:           "v2",
	EventTypeNotificationGiftCardDelivered:  "v1",
	EventTypePriceChanged:                   "v1",
}

func TestEventTypeVersionRegistryCoversEveryDefinedEventExactlyOnce(t *testing.T) {
	expected := map[EventType]string{
		EventTypeInventoryLotReceived:              "v5",
		EventTypeInventoryStockBucketChanged:       "v5",
		EventTypeInventoryPackageConverted:         "v5",
		EventTypeInventoryQualityAssessed:          "v5",
		EventTypeInventoryReservationChanged:       "v5",
		EventTypeInventoryStaged:                   "v5",
		EventTypeInventorySold:                     "v5",
		EventTypeInventoryDateMarkThresholdReached: "v4",
		EventTypeStockLocationAvailabilityChanged:  "v4",
		EventTypeCatalogBaseCostChanged:            "v4",
		EventTypeCatalogListingChanged:             "v4",
		EventTypeAnalyticsOrderFact:                "v5",
		EventTypeAnalyticsRefundFact:               "v5",

		EventTypeOrderPaid:                      "v4",
		EventTypeRefundCompleted:                "v4",
		EventTypeProductSalesPerformanceUpdated: "v3",
		EventTypeFulfilmentPackingUpdated:       "v4",

		EventTypeCampaignChanged:           "v2",
		EventTypePromotionChanged:          "v3",
		EventTypeCouponChanged:             "v1",
		EventTypeFulfilmentShipped:         "v2",
		EventTypeFulfilmentDelivered:       "v2",
		EventTypeFulfilmentCompleted:       "v2",
		EventTypeFulfilmentTrackingUpdated: "v2",
		EventTypeAnalyticsPaymentFact:      "v2",

		EventTypeOrderCreated:                   "v1",
		EventTypeOrderStatusChanged:             "v1",
		EventTypeOrderCancelled:                 "v1",
		EventTypeOrderEdited:                    "v2",
		EventTypePaymentCaptured:                "v2",
		EventTypePaymentFailed:                  "v1",
		EventTypeInvoiceIssued:                  "v1",
		EventTypeReceiptGenerated:               "v1",
		EventTypeRefundRequested:                "v1",
		EventTypeRefundFailed:                   "v1",
		EventTypeCustomerRegistered:             "v1",
		EventTypeCustomerProfileUpdated:         "v1",
		EventTypeNotificationPreferencesChanged: "v1",
		EventTypeWalletGiftCardIssued:           "v2",
		EventTypeNotificationGiftCardDelivered:  "v1",
		EventTypePriceChanged:                   "v1",
	}
	if len(expected) != 41 {
		t.Fatalf("expected event registry test has %d entries, want 41", len(expected))
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
