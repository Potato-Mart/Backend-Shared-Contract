package campaign

import (
	"encoding/json"

	"strings"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/packaging/packaging_enums"
)

func TestCampaignPredictionUsesPackageComposition(t *testing.T) {
	payload, err := json.Marshal(CampaignSupplierPrediction{
		SupplierCode: "SUP-1",
		Products: []CampaignProductPrediction{{
			SKUID: "SKU-1",
			Evidence: []CampaignPredictionEvidence{{
				RawNetBaseUnits: 35, NormalizedBaseUnits: 30,
			}},
			PredictedDemandBaseUnits:   30,
			SellableAvailableBaseUnits: 3,
			ConfirmedInboundBaseUnits:  0,
			NetRequiredBaseUnits:       27,
			SuggestedOrderBaseUnits:    27,
			SuggestedComposition: packaging.PackageCompositionSnapshot{
				TotalBaseUnits: 27,
				Components: []packaging.PackageComponentSnapshot{
					{PackageOptionID: "case_12", HandlingUnit: packaging_enums.PackageHandlingUnitCase, PackageCount: 2, UnitsPerPackage: 12, BaseUnits: 24},
					{PackageOptionID: "each_1", HandlingUnit: packaging_enums.PackageHandlingUnitEach, PackageCount: 3, UnitsPerPackage: 1, BaseUnits: 3},
				},
			},
			MinimumOrderBaseUnits: 12,
		}},
		TotalBaseUnits: 27,
	})
	if err != nil {
		t.Fatalf("marshal campaign product prediction: %v", err)
	}
	if !strings.Contains(string(payload), `"suggested_composition":{"total_base_units":27,"components"`) {
		t.Fatalf("campaign prediction missing package composition: %s", payload)
	}
	for _, expected := range []string{`"raw_net_base_units":35`, `"normalized_base_units":30`, `"predicted_demand_base_units":30`, `"sellable_available_base_units":3`, `"confirmed_inbound_base_units":0`, `"net_required_base_units":27`, `"suggested_order_base_units":27`, `"minimum_order_base_units":12`, `"total_base_units":27`} {
		if !strings.Contains(string(payload), expected) {
			t.Fatalf("campaign prediction missing explicit base-unit field %s: %s", expected, payload)
		}
	}
	for _, removed := range []string{`"suggested_cartons"`, `"carton_size"`, `"raw_net_units"`, `"normalized_units"`, `"predicted_demand_units"`, `"sellable_available_units"`, `"confirmed_inbound_units"`, `"net_required_units"`, `"suggested_order_units"`, `"minimum_order_quantity"`, `"total_units"`} {
		if strings.Contains(string(payload), removed) {
			t.Fatalf("campaign prediction retained removed field %s: %s", removed, payload)
		}
	}
}
