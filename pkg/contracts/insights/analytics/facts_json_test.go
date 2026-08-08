package analytics

import (
	"encoding/json"

	"strings"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v25/pkg/contracts/common/packaging"
)

func TestItemFactsUseBrandID(t *testing.T) {
	for name, value := range map[string]any{
		"order":  OrderItemFact{ProductSKUCode: "A0001", BrandID: "64c13ab08edf48a008793ca1", PackageComposition: packaging.PackageCompositionSnapshot{TotalBaseUnits: 0, Components: []packaging.PackageComponentSnapshot{}}},
		"refund": RefundItemFact{ProductSKUCode: "A0001", BrandID: "64c13ab08edf48a008793ca1", PackageComposition: packaging.PackageCompositionSnapshot{TotalBaseUnits: 0, Components: []packaging.PackageComponentSnapshot{}}},
	} {
		t.Run(name, func(t *testing.T) {
			body, err := json.Marshal(value)
			if err != nil {
				t.Fatal(err)
			}
			if !strings.Contains(string(body), `"brand_id":"64c13ab08edf48a008793ca1"`) || strings.Contains(string(body), `"brand_key"`) {
				t.Fatalf("item fact JSON = %s", body)
			}
			if !strings.Contains(string(body), `"package_composition":{"total_base_units":0`) || strings.Contains(string(body), `"quantity"`) {
				t.Fatalf("item fact did not use package composition: %s", body)
			}
		})
	}
}
