package enums_test

import (
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/pubsub/event/event_enums"
)

func TestV27InventoryAndCommerceEventTypes(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{
			name: "events.EventType",
			valid: []stringEnum{
				event_enums.EventTypeOrderCreated,
				event_enums.EventTypeOrderPaid,
				event_enums.EventTypeOrderStatusChanged,
				event_enums.EventTypeOrderCancelled,
				event_enums.EventTypeCheckoutCompensationRequested,
				event_enums.EventTypePaymentCaptured,
				event_enums.EventTypePaymentFailed,
				event_enums.EventTypeInvoiceIssued,
				event_enums.EventTypeInvoiceDeliveryRequested,
				event_enums.EventTypeReceiptGenerated,
				event_enums.EventTypeRefundRequested,
				event_enums.EventTypeRefundCompleted,
				event_enums.EventTypeRefundFailed,
				event_enums.EventTypeVoucherClaimIssued,
				event_enums.EventTypeInventoryLotReceived,
				event_enums.EventTypeInventoryStockBucketChanged,
				event_enums.EventTypeInventoryPackageConverted,
				event_enums.EventTypeInventoryQualityAssessed,
				event_enums.EventTypeInventoryReservationChanged,
				event_enums.EventTypeInventoryStaged,
				event_enums.EventTypeInventorySold,
				event_enums.EventTypeInventoryDateMarkThresholdReached,
				event_enums.EventTypeStockLocationAvailabilityChanged,
				event_enums.EventTypeFulfilmentPackingUpdated,
				event_enums.EventTypeFulfilmentShipped,
				event_enums.EventTypeFulfilmentDelivered,
				event_enums.EventTypeFulfilmentCompleted,
				event_enums.EventTypeFulfilmentTrackingUpdated,
				event_enums.EventTypeCustomerRegistered,
				event_enums.EventTypeCustomerProfileUpdated,
				event_enums.EventTypeCustomerConsentChanged,
				event_enums.EventTypeWalletGiftCardIssued,
				event_enums.EventTypeCatalogBaseCostChanged,
				event_enums.EventTypeCatalogListingChanged,
				event_enums.EventTypeProductSalesPerformanceUpdated,
				event_enums.EventTypePromotionChanged,
				event_enums.EventTypeCampaignChanged,
				event_enums.EventTypeAnalyticsOrderFact,
				event_enums.EventTypeAnalyticsPaymentFact,
				event_enums.EventTypeAnalyticsRefundFact,
			},
			invalid: event_enums.EventType("__invalid__"),
		},
	})
	if got := event_enums.EventTypeStockLocationAvailabilityChanged.String(); got != "stock.location_availability_changed" {
		t.Fatalf("location availability event type = %q", got)
	}
	for _, retired := range []event_enums.EventType{"inventory.package_pricing_available", "inventory.package_pricing_withdrawn"} {
		if retired.IsValid() {
			t.Fatalf("retired v26 event type %q must not validate", retired)
		}
	}
	if got := event_enums.EventTypeCatalogBaseCostChanged.String(); got != "catalog.base_cost_changed" {
		t.Fatalf("catalog base cost event type = %q", got)
	}
	if got := event_enums.EventTypeCatalogListingChanged.String(); got != "catalog.listing_changed" {
		t.Fatalf("catalog listing event type = %q", got)
	}
	if got := event_enums.EventTopicCatalogEvents.String(); got != "catalog-events" || !event_enums.EventTopicCatalogEvents.IsValid() {
		t.Fatalf("catalog topic = %q", got)
	}
}
