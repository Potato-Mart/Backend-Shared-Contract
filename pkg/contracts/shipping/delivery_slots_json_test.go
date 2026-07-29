package shipping_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v21/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v21/pkg/contracts/shipping"
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
			Postcode: "2000", Suburb: "Sydney", DeliveryRegion: "NSW",
			DepotCode: "SYD", DepotName: "Sydney",
			ShippingFee:           common.Money{AmountMinor: 1000, Currency: "AUD"},
			FreeShippingThreshold: common.Money{AmountMinor: 10000, Currency: "AUD"},
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
		`"revision":7`, `"carrier":"potato-mart"`, `"postcode":"2000"`, `"suburb":"Sydney"`,
		`"delivery_region":"NSW"`, `"depot_code":"SYD"`, `"depot_name":"Sydney"`,
		`"shipping_fee":{"amount_minor":1000,"currency":"AUD"}`,
		`"free_shipping_threshold":{"amount_minor":10000,"currency":"AUD"}`,
	} {
		if !strings.Contains(string(payload), field) {
			t.Fatalf("delivery schedule missing %s: %s", field, payload)
		}
	}
}
