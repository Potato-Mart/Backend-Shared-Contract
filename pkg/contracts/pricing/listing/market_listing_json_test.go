package listing

import (
	"encoding/json"
	"github.com/Potato-Mart/Backend-Shared-Contract/v43/pkg/contracts/common/commerce/commerce_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v43/pkg/contracts/pricing/listing/listing_enums"
	"strings"
	"testing"
	"time"
)

func TestMarketListingCarriesAvailabilityWithoutPrice(t *testing.T) {
	availableFrom := time.Date(2026, 8, 12, 0, 0, 0, 0, time.UTC)
	leadDays := int32(21)
	ageYears := int32(18)
	payload, err := json.Marshal(MarketListing{
		ID: "listing_1", MarketCode: "market_au", SKUCode: "sku_a00001",
		Status: listing_enums.MarketListingStatusActive, TaxCategoryCode: "tax_au_gst",
		Restrictions: []SaleRestriction{
			{Kind: listing_enums.SaleRestrictionKindAge, AgeYears: &ageYears, AgeComparison: listing_enums.AgeComparisonBelow, Channels: []commerce_enums.OrderType{commerce_enums.OrderTypeOnline}},
		},
		ExpiryLeadDaysOverride: &leadDays,
		AvailableFrom:          availableFrom, Revision: 4,
	})
	if err != nil {
		t.Fatalf("marshal market listing: %v", err)
	}
	for _, want := range []string{
		`"market_code":"market_au"`, `"sku_code":"sku_a00001"`, `"status":"active"`,
		`"tax_category_code":"tax_au_gst"`, `"kind":"age"`,
		`"expiry_lead_days_override":21`, `"revision":4`,
	} {
		if !strings.Contains(string(payload), want) {
			t.Fatalf("MarketListing JSON = %s, want %s", payload, want)
		}
	}
	for _, forbidden := range []string{`"amount_minor"`, `"price"`, `"currency"`, `"price_book_code"`, `"display_name"`, `"unit_pricing_required"`} {
		if strings.Contains(string(payload), forbidden) {
			t.Fatalf("MarketListing must never carry a commercial price, leaked %s: %s", forbidden, payload)
		}
	}
}
