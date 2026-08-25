package shipping_test

import (
	"encoding/json"

	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/orders/shipping"
)

func TestDeliveryScheduleJSONShape(t *testing.T) {
	start := time.Date(2026, 8, 1, 9, 0, 0, 0, time.FixedZone("AEST", 10*60*60))
	end := start.Add(3 * time.Hour)
	schedule := shipping.DeliverySchedule{
		Availability: "available",
		Revision:     7,
		Timezone:     "Australia/Sydney",
		Carrier:      "potato-mart",
		AreaRate: &shipping.DeliveryAreaRate{
			CountryCode: "AU", AdministrativeAreaCode: "AU-NSW", PostalCode: "2000", Locality: "Sydney",
			ZoneID: "zone_au_nsw_sydney", DepotRegionCode: "AU-NSW-SYD",
			DepotCode: "AU-NSW-SYD-DC-01", DepotName: "Sydney",
			ShippingFee:           money.Money{AmountMinor: 1000, Currency: "AUD"},
			FreeShippingThreshold: money.Money{AmountMinor: 10000, Currency: "AUD"},
		},
		DateGroups: []shipping.DeliveryDateGroup{{
			Date: "2026-08-01",
			Slots: []shipping.DeliverySlot{{
				ID: "slot_7", StartAt: start, EndAt: end, Availability: "available",
			}},
		}},
	}

	payload, err := json.Marshal(schedule)
	if err != nil {
		t.Fatalf("marshal delivery schedule: %v", err)
	}
	for _, field := range []string{
		`"revision":7`, `"carrier":"potato-mart"`, `"country_code":"AU"`,
		`"administrative_area_code":"AU-NSW"`, `"postal_code":"2000"`, `"locality":"Sydney"`,
		`"zone_id":"zone_au_nsw_sydney"`, `"depot_region_code":"AU-NSW-SYD"`,
		`"depot_code":"AU-NSW-SYD-DC-01"`, `"depot_name":"Sydney"`,
		`"shipping_fee":{"amount_minor":1000,"currency":"AUD"}`,
		`"free_shipping_threshold":{"amount_minor":10000,"currency":"AUD"}`,
	} {
		if !strings.Contains(string(payload), field) {
			t.Fatalf("delivery schedule missing %s: %s", field, payload)
		}
	}
	for _, removed := range []string{`"postcode"`, `"suburb"`, `"delivery_region"`} {
		if strings.Contains(string(payload), removed) {
			t.Fatalf("delivery schedule retained removed field %s: %s", removed, payload)
		}
	}
}
