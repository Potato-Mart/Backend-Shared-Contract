package pricing_test

import (
	"encoding/json"
	"github.com/Potato-Mart/Backend-Shared-Contract/v42/pkg/contracts/pricing/listing/listing_enums"
	event "github.com/Potato-Mart/Backend-Shared-Contract/v42/pkg/contracts/pubsub/pricing"
	"strings"
	"testing"
	"time"
)

func TestCatalogListingChangedEventCarriesCodeIdentityAndRevision(t *testing.T) {
	now := time.Date(2026, 8, 12, 7, 8, 9, 0, time.UTC)
	leadDays := int32(21)
	value := event.CatalogListingChangedEvent{
		MarketCode:             "market_au",
		SKUCode:                "sku_a00001",
		PreviousStatus:         listing_enums.MarketListingStatusDraft,
		Status:                 listing_enums.MarketListingStatusActive,
		TaxCategoryCode:        "tax_au_gst",
		ExpiryLeadDaysOverride: &leadDays,
		PreviousRevision:       6,
		Revision:               7,
		AvailableFrom:          now,
	}

	payload, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("marshal listing changed event: %v", err)
	}
	for _, want := range []string{
		`"market_code":"market_au"`, `"sku_code":"sku_a00001"`,
		`"previous_status":"draft"`, `"status":"active"`, `"tax_category_code":"tax_au_gst"`,
		`"expiry_lead_days_override":21`,
		`"previous_revision":6`, `"revision":7`,
	} {
		if !strings.Contains(string(payload), want) {
			t.Fatalf("listing changed JSON missing %s: %s", want, payload)
		}
	}
	for _, forbidden := range []string{`"listing_id"`, `"price"`, `"amount_minor"`, `"package_pricing_id"`, `"unit_pricing_required"`, `"display_name"`} {
		if strings.Contains(string(payload), forbidden) {
			t.Fatalf("listing changed JSON leaked %s: %s", forbidden, payload)
		}
	}

	var got event.CatalogListingChangedEvent
	if err := json.Unmarshal(payload, &got); err != nil {
		t.Fatalf("unmarshal listing changed event: %v", err)
	}
	if got.Revision != 7 || got.PreviousRevision != 6 || got.ExpiryLeadDaysOverride == nil || *got.ExpiryLeadDaysOverride != 21 {
		t.Fatalf("listing revision evidence did not round-trip: %+v", got)
	}
	if strings.Contains(string(payload), `"occurred_at"`) {
		t.Fatalf("listing payload must rely on envelope occurred_at: %s", payload)
	}
}
