package pkg_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v50/pkg/contracts/common/packaging"
	orderfulfilment "github.com/Potato-Mart/Backend-Shared-Contract/v50/pkg/contracts/orders/fulfilment"
	"github.com/Potato-Mart/Backend-Shared-Contract/v50/pkg/contracts/orders/order"
	pubsuborders "github.com/Potato-Mart/Backend-Shared-Contract/v50/pkg/contracts/pubsub/orders"
	supplyfulfilment "github.com/Potato-Mart/Backend-Shared-Contract/v50/pkg/contracts/supply/fulfilment"
)

func TestPhysicalWorkCarriesOrderFulfillmentGeneration(t *testing.T) {
	models := []any{
		order.Order{FulfillmentGeneration: 5, Packing: &orderfulfilment.OrderPackingProgress{FulfillmentGeneration: 5}},
		orderfulfilment.OrderPackingProgress{FulfillmentGeneration: 5},
		supplyfulfilment.PackingRecord{FulfillmentGeneration: 5},
		pubsuborders.OrderPackingProjection{Packing: orderfulfilment.OrderPackingProgress{FulfillmentGeneration: 5}},
	}
	for _, model := range models {
		encoded, err := json.Marshal(model)
		if err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(string(encoded), `"fulfillment_generation":5`) {
			t.Errorf("%T omitted composition generation: %s", model, encoded)
		}
	}
	legacyProjection, err := json.Marshal(pubsuborders.OrderPackingProjection{})
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(legacyProjection), "fulfillment_generation") {
		t.Fatalf("legacy packing projection gained an asserted generation: %s", legacyProjection)
	}
}

func TestPackingProjectionCarriesQuantitiesWithoutPickingListEvidence(t *testing.T) {
	projection := pubsuborders.OrderPackingProjection{
		OrderNumber: "ORDER-1", Revision: 4,
		Packing: orderfulfilment.OrderPackingProgress{FulfillmentGeneration: 2,
			Lines: []supplyfulfilment.PackingLine{{OrderItemID: "item-1",
				RequestedComposition: packaging.PackageCompositionSnapshot{TotalBaseUnits: 4},
				PackedComposition:    packaging.PackageCompositionSnapshot{TotalBaseUnits: 3}}}},
	}
	raw, err := json.Marshal(projection)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(raw), "picked_composition") || strings.Contains(string(raw), "picking_list") {
		t.Fatalf("retired picking state in packing projection: %s", raw)
	}
	var got pubsuborders.OrderPackingProjection
	if err := json.Unmarshal(raw, &got); err != nil {
		t.Fatal(err)
	}
	if got.OrderNumber != "ORDER-1" || got.Revision != 4 || got.Packing.FulfillmentGeneration != 2 || len(got.Packing.Lines) != 1 || got.Packing.Lines[0].RequestedComposition.TotalBaseUnits != 4 || got.Packing.Lines[0].PackedComposition.TotalBaseUnits != 3 {
		t.Fatalf("packing projection lost ordered or packed quantities: %+v", got)
	}
}
