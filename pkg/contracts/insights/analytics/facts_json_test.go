package analytics

import (
	"encoding/json"

	"strings"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/packaging"
)

// TestAnalyticsFactsCarryCountryAttribution locks the country a scoped read
// filters on. Before v28.0.0 these models carried market_code only, so a
// countryAdmin could not be filtered at all and Insights had to refuse the
// read rather than widen it — the rank-2 principal could see no analytics.
// The field is omitempty, so a fact recorded without it decodes unchanged and
// the consumer treats the absence as "unattributed", fail-closed.
func TestAnalyticsFactsCarryCountryAttribution(t *testing.T) {
	composition := packaging.PackageCompositionSnapshot{TotalBaseUnits: 0, Components: []packaging.PackageComponentSnapshot{}}
	populated := map[string]any{
		"order":    OrderItemFact{SKUCode: "A0001", MarketCode: "mkt_au_vic", CountryCode: "AU", PackageComposition: composition},
		"refund":   RefundItemFact{SKUCode: "A0001", MarketCode: "mkt_au_vic", CountryCode: "AU", PackageComposition: composition},
		"rollup":   MetricRollup{Metric: "gross_sales", Granularity: "daily", MarketCode: "mkt_au_vic", CountryCode: "AU"},
		"forecast": SKUDemandForecast{SKUCode: "A0001", MarketCode: "mkt_au_vic", CountryCode: "AU"},
	}
	for name, value := range populated {
		t.Run(name, func(t *testing.T) {
			body, err := json.Marshal(value)
			if err != nil {
				t.Fatal(err)
			}
			if !strings.Contains(string(body), `"country_code":"AU"`) {
				t.Fatalf("%s fact carries no country attribution: %s", name, body)
			}
		})
	}

	bare := map[string]any{
		"order":    OrderItemFact{SKUCode: "A0001", PackageComposition: composition},
		"refund":   RefundItemFact{SKUCode: "A0001", PackageComposition: composition},
		"rollup":   MetricRollup{Metric: "gross_sales"},
		"forecast": SKUDemandForecast{SKUCode: "A0001"},
	}
	for name, value := range bare {
		t.Run(name+"_absent", func(t *testing.T) {
			body, err := json.Marshal(value)
			if err != nil {
				t.Fatal(err)
			}
			if strings.Contains(string(body), "country_code") {
				t.Fatalf("%s fact must omit an absent country rather than emit an empty one: %s", name, body)
			}
		})
	}
}

func TestItemFactsUseBrandCode(t *testing.T) {
	for name, value := range map[string]any{
		"order":  OrderItemFact{SKUCode: "A0001", BrandCode: "64c13ab08edf48a008793ca1", PackageComposition: packaging.PackageCompositionSnapshot{TotalBaseUnits: 0, Components: []packaging.PackageComponentSnapshot{}}},
		"refund": RefundItemFact{SKUCode: "A0001", BrandCode: "64c13ab08edf48a008793ca1", PackageComposition: packaging.PackageCompositionSnapshot{TotalBaseUnits: 0, Components: []packaging.PackageComponentSnapshot{}}},
	} {
		t.Run(name, func(t *testing.T) {
			body, err := json.Marshal(value)
			if err != nil {
				t.Fatal(err)
			}
			if !strings.Contains(string(body), `"brand_code":"64c13ab08edf48a008793ca1"`) || strings.Contains(string(body), `"brand_key"`) {
				t.Fatalf("item fact JSON = %s", body)
			}
			if !strings.Contains(string(body), `"package_composition":{"total_base_units":0`) || strings.Contains(string(body), `"quantity"`) {
				t.Fatalf("item fact did not use package composition: %s", body)
			}
		})
	}
}
