package enums_test

import (
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v42/pkg/contracts/pricing/listing/listing_enums"
)

func TestListingEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "listingenum.AgeComparison", valid: []stringEnum{listing_enums.AgeComparisonBelow, listing_enums.AgeComparisonAbove}, invalid: listing_enums.AgeComparison("__invalid__")},
		{name: "listingenum.MarketListingStatus", valid: []stringEnum{listing_enums.MarketListingStatusDraft, listing_enums.MarketListingStatusComingSoon, listing_enums.MarketListingStatusActive, listing_enums.MarketListingStatusSuspended, listing_enums.MarketListingStatusUnavailable, listing_enums.MarketListingStatusDelisted}, invalid: listing_enums.MarketListingStatus("__invalid__")},
		{name: "listingenum.SaleRestrictionKind", valid: []stringEnum{listing_enums.SaleRestrictionKindAge, listing_enums.SaleRestrictionKindQuantityLimit, listing_enums.SaleRestrictionKindChannelExcluded, listing_enums.SaleRestrictionKindDeliveryExcluded, listing_enums.SaleRestrictionKindPrescription}, invalid: listing_enums.SaleRestrictionKind("__invalid__")},
	})
}
