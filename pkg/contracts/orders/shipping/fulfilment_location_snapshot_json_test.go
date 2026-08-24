package shipping

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/geography/geography_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/party"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/orders/shipping/shipping_enums"
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
