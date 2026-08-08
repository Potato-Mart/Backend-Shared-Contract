package enums_test

import (
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v25/pkg/contracts/pubsub/event/event_enums"
)

func TestV25InventoryAndCommerceEventTypes(t *testing.T) {
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
				event_enums.EventTypeInventorySellableOfferAvailable,
				event_enums.EventTypeInventorySellableOfferWithdrawn,
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
}
