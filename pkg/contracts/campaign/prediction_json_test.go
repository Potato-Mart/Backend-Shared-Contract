package campaign

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/common"
)

func TestCampaignPredictionUsesPackageComposition(t *testing.T) {
	payload, err := json.Marshal(CampaignProductPrediction{
		ProductSKUCode:      "SKU-1",
		SuggestedOrderUnits: 27,
		SuggestedComposition: common.PackageCompositionSnapshot{
			TotalBaseUnits: 27,
			Components: []common.PackageComponentSnapshot{
				{PackageOptionID: "case_12", HandlingUnit: common.PackageHandlingUnitCase, PackageCount: 2, UnitsPerPackage: 12, BaseUnits: 24},
				{PackageOptionID: "each_1", HandlingUnit: common.PackageHandlingUnitEach, PackageCount: 3, UnitsPerPackage: 1, BaseUnits: 3},
			},
		},
	})
	if err != nil {
		t.Fatalf("marshal campaign product prediction: %v", err)
	}
	if !strings.Contains(string(payload), `"suggested_composition":{"total_base_units":27,"components"`) {
		t.Fatalf("campaign prediction missing package composition: %s", payload)
	}
	for _, removed := range []string{`"suggested_cartons"`, `"carton_size"`} {
		if strings.Contains(string(payload), removed) {
			t.Fatalf("campaign prediction retained removed field %s: %s", removed, payload)
		}
	}
}
