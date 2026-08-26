package shipping

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/geography/geography_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/party"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/orders/shipping/shipping_enums"
)

func TestFulfilmentLocationSnapshotJSONShape(t *testing.T) {
	capturedAt := time.Date(2026, 8, 24, 1, 2, 3, 0, time.UTC)
	snapshot := FulfilmentLocationSnapshot{
		Intent: shipping_enums.FulfilmentIntentDelivery,
		DeliveryAddress: &party.ContactAddress{
			Address: &geography.Address{
				Line1: "1 Example Street", Locality: "Dandenong", PostalCode: "3175", Country: geography.CountryRef{Code: "AU"},
			},
		},
		GeographicContext: geography.GeographicContext{
			Source:      geography_enums.GeographicContextSourceDeliveryAddress,
			MarketCode:  "mkt_au_vic",
			CountryCode: "AU",
		},
		LocationFingerprint: "locfp_01",
		CapturedAt:          capturedAt,
	}

	payload, err := json.Marshal(snapshot)
	if err != nil {
		t.Fatalf("marshal fulfilment location: %v", err)
	}
	for _, field := range []string{
		`"intent":"delivery"`,
		`"delivery_address"`,
		`"geographic_context"`,
		`"market_code":"mkt_au_vic"`,
		`"location_fingerprint":"locfp_01"`,
		`"captured_at":"2026-08-24T01:02:03Z"`,
	} {
		if !strings.Contains(string(payload), field) {
			t.Fatalf("fulfilment location JSON missing %s: %s", field, payload)
		}
	}
	if strings.Contains(string(payload), `"selected_depot_code"`) {
		t.Fatalf("delivery snapshot must omit selected depot: %s", payload)
	}
}

func TestFulfilmentLocationSnapshotPickupOmitsDeliveryAddress(t *testing.T) {
	payload, err := json.Marshal(FulfilmentLocationSnapshot{
		Intent:            shipping_enums.FulfilmentIntentPickup,
		SelectedDepotCode: "AU-VIC-MEL-DC-01",
	})
	if err != nil {
		t.Fatalf("marshal pickup fulfilment location: %v", err)
	}
	if strings.Contains(string(payload), `"delivery_address"`) || !strings.Contains(string(payload), `"selected_depot_code":"AU-VIC-MEL-DC-01"`) {
		t.Fatalf("pickup fulfilment location JSON = %s", payload)
	}
}

func TestFulfilmentLocationSnapshotDigitalOmitsPhysicalLocationAndRetainsEvidence(t *testing.T) {
	capturedAt := time.Date(2026, 8, 24, 1, 2, 3, 0, time.UTC)
	snapshot := FulfilmentLocationSnapshot{
		Intent: shipping_enums.FulfilmentIntentDigital,
		GeographicContext: geography.GeographicContext{
			Source:      geography_enums.GeographicContextSourceGlobalFallback,
			MarketCode:  "mkt_au",
			CountryCode: "AU",
		},
		LocationFingerprint: "locfp_digital_01",
		CapturedAt:          capturedAt,
	}

	if !fulfilmentLocationHasValidLocationInvariant(snapshot) {
		t.Fatal("digital snapshot with no physical location was classified as invalid")
	}
	payload, err := json.Marshal(snapshot)
	if err != nil {
		t.Fatalf("marshal digital fulfilment location: %v", err)
	}
	text := string(payload)
	for _, field := range []string{
		`"intent":"digital"`,
		`"source":"GLOBAL_FALLBACK"`,
		`"geographic_context"`,
		`"location_fingerprint":"locfp_digital_01"`,
		`"captured_at":"2026-08-24T01:02:03Z"`,
	} {
		if !strings.Contains(text, field) {
			t.Fatalf("digital fulfilment location JSON missing %s: %s", field, payload)
		}
	}
	for _, physicalField := range []string{`"delivery_address"`, `"selected_depot_code"`} {
		if strings.Contains(text, physicalField) {
			t.Fatalf("digital fulfilment location retained %s: %s", physicalField, payload)
		}
	}
}

func TestFulfilmentLocationSnapshotLocksAddressOrDepotInvariant(t *testing.T) {
	capturedAt := time.Date(2026, 8, 24, 1, 2, 3, 0, time.UTC)
	context := geography.GeographicContext{MarketCode: "mkt_au_vic", CountryCode: "AU"}
	for name, snapshot := range map[string]FulfilmentLocationSnapshot{
		"delivery-requires-address": {Intent: shipping_enums.FulfilmentIntentDelivery, GeographicContext: context, LocationFingerprint: "locfp_1", CapturedAt: capturedAt},
		"delivery-forbids-depot":    {Intent: shipping_enums.FulfilmentIntentDelivery, DeliveryAddress: &party.ContactAddress{}, SelectedDepotCode: "AU-VIC-01", GeographicContext: context, LocationFingerprint: "locfp_1", CapturedAt: capturedAt},
		"pickup-requires-depot":     {Intent: shipping_enums.FulfilmentIntentPickup, GeographicContext: context, LocationFingerprint: "locfp_1", CapturedAt: capturedAt},
		"pickup-forbids-address":    {Intent: shipping_enums.FulfilmentIntentPickup, DeliveryAddress: &party.ContactAddress{}, SelectedDepotCode: "AU-VIC-01", GeographicContext: context, LocationFingerprint: "locfp_1", CapturedAt: capturedAt},
		"digital-forbids-address":   {Intent: shipping_enums.FulfilmentIntentDigital, DeliveryAddress: &party.ContactAddress{}, GeographicContext: context, LocationFingerprint: "locfp_1", CapturedAt: capturedAt},
		"digital-forbids-depot":     {Intent: shipping_enums.FulfilmentIntentDigital, SelectedDepotCode: "AU-VIC-01", GeographicContext: context, LocationFingerprint: "locfp_1", CapturedAt: capturedAt},
	} {
		t.Run(name, func(t *testing.T) {
			if fulfilmentLocationHasValidLocationInvariant(snapshot) {
				t.Fatal("invalid fulfilment location fixture was classified as valid")
			}
		})
	}
}

func fulfilmentLocationHasValidLocationInvariant(snapshot FulfilmentLocationSnapshot) bool {
	switch snapshot.Intent {
	case shipping_enums.FulfilmentIntentDelivery:
		return snapshot.DeliveryAddress != nil && snapshot.SelectedDepotCode == ""
	case shipping_enums.FulfilmentIntentPickup, shipping_enums.FulfilmentIntentInStoreCarry:
		return snapshot.DeliveryAddress == nil && snapshot.SelectedDepotCode != ""
	case shipping_enums.FulfilmentIntentDigital:
		return snapshot.DeliveryAddress == nil && snapshot.SelectedDepotCode == ""
	default:
		return false
	}
}
