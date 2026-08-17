package analytics

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/commerce/commerce_enums"
)

func TestSKUDemandForecastJSONUsesQualifiedAvailabilitySnapshot(t *testing.T) {
	computedAt := time.Date(2026, 8, 4, 3, 2, 1, 0, time.UTC)
	forecast := SKUDemandForecast{
		SKUID:                "A00001",
		DepotCode:            "AU-VIC-MEL-DC-01",
		Channel:              commerce_enums.OrderTypeOnline,
		Timezone:             "Australia/Melbourne",
		ComputedAt:           computedAt,
		HistoryDays:          30,
		ForecastHorizonDays:  7,
		PredictedDemandTotal: 18,
		PredictedDaily:       []DailyPrediction{},
		AvailabilityAtRun: ForecastAvailabilitySnapshot{
			SellableAvailableBaseUnits: 0,
			ReservedBaseUnits:          12,
			StagedBaseUnits:            6,
			QualityHoldBaseUnits:       2,
			Revision:                   17,
			AsOf:                       computedAt,
		},
		Algorithm: "seasonal-v2",
	}

	payload, err := json.Marshal(forecast)
	if err != nil {
		t.Fatalf("marshal forecast: %v", err)
	}
	text := string(payload)
	for _, expected := range []string{
		`"depot_code":"AU-VIC-MEL-DC-01"`,
		`"channel":"online"`,
		`"timezone":"Australia/Melbourne"`,
		`"availability_at_run"`,
		`"sellable_available_base_units":0`,
		`"revision":17`,
		`"as_of":"2026-08-04T03:02:01Z"`,
	} {
		if !strings.Contains(text, expected) {
			t.Fatalf("forecast JSON missing %s: %s", expected, payload)
		}
	}
	if strings.Contains(text, `"current_stock"`) {
		t.Fatalf("forecast JSON contains removed current_stock: %s", payload)
	}
}
