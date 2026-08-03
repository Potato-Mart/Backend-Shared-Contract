package enums_test

import (
	"testing"

	eventsenum "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/enums/events"
)

func TestV22InventoryAndCommerceEventTypes(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{
			name: "events.EventType",
			valid: []stringEnum{
				eventsenum.EventTypeVoucherClaimIssued,
				eventsenum.EventTypeInventoryLotReceived,
				eventsenum.EventTypeInventoryStockBucketChanged,
				eventsenum.EventTypeInventoryPackageConverted,
				eventsenum.EventTypeInventoryQualityAssessed,
				eventsenum.EventTypeInventoryReservationChanged,
				eventsenum.EventTypeInventoryStaged,
				eventsenum.EventTypeInventorySold,
				eventsenum.EventTypeInventoryDateMarkThresholdReached,
				eventsenum.EventTypeInventorySellableOfferAvailable,
				eventsenum.EventTypeInventorySellableOfferWithdrawn,
				eventsenum.EventTypeStockLocationAvailabilityChanged,
				eventsenum.EventTypePromotionChanged,
				eventsenum.EventTypeCampaignChanged,
			},
			invalid: eventsenum.EventType("__invalid__"),
		},
	})
	if got := eventsenum.EventTypeStockLocationAvailabilityChanged.String(); got != "stock.location_availability_changed" {
		t.Fatalf("location availability event type = %q", got)
	}
}
