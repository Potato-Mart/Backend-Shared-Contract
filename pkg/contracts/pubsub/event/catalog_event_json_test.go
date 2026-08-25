package event_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/pubsub/event"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/supply/listing/listing_enums"
)

func TestCatalogBaseCostChangedEventJSONShape(t *testing.T) {
	now := time.Date(2026, 8, 12, 7, 8, 9, 0, time.UTC)
	previous := money.Money{AmountMinor: 500, Currency: "AUD"}
	shape := marshalObject(t, event.CatalogBaseCostChangedEvent{
		SKUCode:          "sku_a00001",
		Currency:         "AUD",
		PreviousAmount:   &previous,
		Amount:           money.Money{AmountMinor: 550, Currency: "AUD"},
		PreviousRevision: 4,
		Revision:         5,
		SourceType:       "supplier_invoice",
		SourceID:         "invoice_1",
		EffectiveFrom:    now,
		OccurredAt:       now,
	})
	for _, key := range []string{"sku_code", "currency", "previous_amount", "amount", "previous_revision", "revision", "effective_from", "occurred_at"} {
		if _, ok := shape[key]; !ok {
			t.Fatalf("base cost changed JSON missing %q: %+v", key, shape)
		}
	}
	if shape["revision"] != float64(5) || shape["occurred_at"] != "2026-08-12T07:08:09Z" {
		t.Fatalf("base cost changed identity did not marshal: %+v", shape)
	}
}

func TestCatalogListingChangedEventCarriesCodeIdentityAndRevision(t *testing.T) {
	now := time.Date(2026, 8, 12, 7, 8, 9, 0, time.UTC)
	leadDays := int32(21)
	value := event.CatalogListingChangedEvent{
		MarketCode:             "market_au",
		SKUCode:                "sku_a00001",
		PreviousStatus:         listing_enums.MarketListingStatusDraft,
		Status:                 listing_enums.MarketListingStatusActive,
		TaxCategoryCode:        "tax_au_gst",
		UnitPricingRequired:    true,
		ExpiryLeadDaysOverride: &leadDays,
		PreviousRevision:       6,
		Revision:               7,
		AvailableFrom:          now,
		OccurredAt:             now,
	}

	payload, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("marshal listing changed event: %v", err)
	}
	for _, want := range []string{
		`"market_code":"market_au"`, `"sku_code":"sku_a00001"`,
		`"previous_status":"draft"`, `"status":"active"`, `"tax_category_code":"tax_au_gst"`,
		`"unit_pricing_required":true`, `"expiry_lead_days_override":21`,
		`"previous_revision":6`, `"revision":7`,
	} {
		if !strings.Contains(string(payload), want) {
			t.Fatalf("listing changed JSON missing %s: %s", want, payload)
		}
	}
	for _, forbidden := range []string{`"listing_id"`, `"price"`, `"amount_minor"`, `"package_pricing_id"`} {
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
}
