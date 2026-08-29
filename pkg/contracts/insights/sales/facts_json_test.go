package sales

import (
	"encoding/json"

	"strings"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/insights/analytics"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/forecasting"
)

// TestAnalyticsFactsCarryCountryAttribution locks the country a scoped read
// filters on. The field is optional so facts without country attribution stay
// distinguishable from facts attributed to a specific country.
func TestAnalyticsFactsCarryCountryAttribution(t *testing.T) {
	composition := packaging.PackageCompositionSnapshot{TotalBaseUnits: 0, Components: []packaging.PackageComponentSnapshot{}}
	populated := map[string]any{
		"order":    OrderItemFact{ProductFactDimensions: ProductFactDimensions{SKUCode: "A0001", MarketCode: "mkt_au_vic", CountryCode: "AU", PackageComposition: composition}},
		"refund":   RefundItemFact{ProductFactDimensions: ProductFactDimensions{SKUCode: "A0001", MarketCode: "mkt_au_vic", CountryCode: "AU", PackageComposition: composition}},
		"rollup":   analytics.MetricRollup{Metric: "gross_sales", Granularity: "daily", MarketCode: "mkt_au_vic", CountryCode: "AU"},
		"forecast": forecasting.SKUDemandForecast{SKUCode: "A0001", MarketCode: "mkt_au_vic", CountryCode: "AU"},
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
		"order":    OrderItemFact{ProductFactDimensions: ProductFactDimensions{SKUCode: "A0001", PackageComposition: composition}},
		"refund":   RefundItemFact{ProductFactDimensions: ProductFactDimensions{SKUCode: "A0001", PackageComposition: composition}},
		"rollup":   analytics.MetricRollup{Metric: "gross_sales"},
		"forecast": forecasting.SKUDemandForecast{SKUCode: "A0001"},
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
		"order":  OrderItemFact{ProductFactDimensions: ProductFactDimensions{SKUCode: "A0001", BrandCode: "64c13ab08edf48a008793ca1", PackageComposition: packaging.PackageCompositionSnapshot{TotalBaseUnits: 0, Components: []packaging.PackageComponentSnapshot{}}}},
		"refund": RefundItemFact{ProductFactDimensions: ProductFactDimensions{SKUCode: "A0001", BrandCode: "64c13ab08edf48a008793ca1", PackageComposition: packaging.PackageCompositionSnapshot{TotalBaseUnits: 0, Components: []packaging.PackageComponentSnapshot{}}}},
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
