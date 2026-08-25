package order_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/geography/geography_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/orders/order"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/orders/shipping"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/orders/shipping/shipping_enums"
)

func TestCartAndOrderUseFulfilmentLocationSnapshot(t *testing.T) {
	capturedAt := time.Date(2026, 8, 24, 1, 2, 3, 0, time.UTC)
	location := shipping.FulfilmentLocationSnapshot{
		Intent: shipping_enums.FulfilmentIntentPickup,
		GeographicContext: geography.GeographicContext{
			Source:      geography_enums.GeographicContextSourceFulfilmentDepot,
			MarketCode:  "mkt_au_vic",
			CountryCode: "AU",
			DepotCode:   "AU-VIC-MEL-DC-01",
		},
		SelectedDepotCode:   "AU-VIC-MEL-DC-01",
		LocationFingerprint: "locfp_01",
		CapturedAt:          capturedAt,
	}
	for name, value := range map[string]any{
		"cart": order.Cart{
			MarketCode:         "mkt_au_vic",
			CountryCode:        "AU",
			FulfilmentLocation: location,
		},
		"order": order.Order{
			MarketCode:         "mkt_au_vic",
			CountryCode:        "AU",
			FulfilmentLocation: location,
		},
	} {
		t.Run(name, func(t *testing.T) {
			payload, err := json.Marshal(value)
			if err != nil {
				t.Fatalf("marshal %s: %v", name, err)
			}
			var root map[string]any
			if err := json.Unmarshal(payload, &root); err != nil {
				t.Fatalf("unmarshal %s: %v", name, err)
			}
			if _, exists := root["geographic_context"]; exists {
				t.Fatalf("%s retained top-level geographic_context: %s", name, payload)
			}
			if _, exists := root["shipping"]; exists {
				t.Fatalf("%s retained removed shipping address: %s", name, payload)
			}
			fulfilment, ok := root["fulfilment_location"].(map[string]any)
			if !ok {
				t.Fatalf("%s missing fulfilment location: %s", name, payload)
			}
			context, ok := fulfilment["geographic_context"].(map[string]any)
			if !ok || context["market_code"] != root["market_code"] || context["country_code"] != root["country_code"] {
				t.Fatalf("%s does not carry matching market/country resolution: %s", name, payload)
			}
		})
	}
}

func TestCartAndOrderRequireMarketAndCountryToMatchFulfilmentLocation(t *testing.T) {
	capturedAt := time.Date(2026, 8, 24, 1, 2, 3, 0, time.UTC)
	location := shipping.FulfilmentLocationSnapshot{
		Intent:            shipping_enums.FulfilmentIntentPickup,
		GeographicContext: geography.GeographicContext{MarketCode: "mkt_au_vic", CountryCode: "AU"},
		SelectedDepotCode: "AU-VIC-MEL-DC-01", LocationFingerprint: "locfp_01", CapturedAt: capturedAt,
	}
	wrongMarketCart := order.Cart{MarketCode: "mkt_au_nsw", CountryCode: "AU", FulfilmentLocation: location}
	if wrongMarketCart.MarketCode == wrongMarketCart.FulfilmentLocation.GeographicContext.MarketCode && wrongMarketCart.CountryCode == wrongMarketCart.FulfilmentLocation.GeographicContext.CountryCode {
		t.Fatal("mismatched cart fixture was classified as consistent")
	}
	wrongCountryOrder := order.Order{MarketCode: "mkt_au_vic", CountryCode: "NZ", FulfilmentLocation: location}
	if wrongCountryOrder.MarketCode == wrongCountryOrder.FulfilmentLocation.GeographicContext.MarketCode && wrongCountryOrder.CountryCode == wrongCountryOrder.FulfilmentLocation.GeographicContext.CountryCode {
		t.Fatal("mismatched order fixture was classified as consistent")
	}
	matchingCart := order.Cart{MarketCode: "mkt_au_vic", CountryCode: "AU", FulfilmentLocation: location}
	if matchingCart.MarketCode != matchingCart.FulfilmentLocation.GeographicContext.MarketCode || matchingCart.CountryCode != matchingCart.FulfilmentLocation.GeographicContext.CountryCode {
		t.Fatal("matching cart fixture was classified as inconsistent")
	}
}
