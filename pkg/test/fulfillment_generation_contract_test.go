package pkg_test

import (
	"encoding/json"
	"strings"
	"testing"

	orderfulfilment "github.com/Potato-Mart/Backend-Shared-Contract/v42/pkg/contracts/orders/fulfilment"
	"github.com/Potato-Mart/Backend-Shared-Contract/v42/pkg/contracts/orders/order"
	pubsuborders "github.com/Potato-Mart/Backend-Shared-Contract/v42/pkg/contracts/pubsub/orders"
	supplyfulfilment "github.com/Potato-Mart/Backend-Shared-Contract/v42/pkg/contracts/supply/fulfilment"
)

func TestPhysicalWorkCarriesOrderFulfillmentGeneration(t *testing.T) {
	models := []any{
		order.Order{FulfillmentGeneration: 5, Packing: &orderfulfilment.OrderPackingProgress{FulfillmentGeneration: 5}},
		orderfulfilment.OrderPackingProgress{FulfillmentGeneration: 5},
		supplyfulfilment.PickingList{FulfillmentGeneration: 5, AllocationFingerprint: "allocation-set-5"},
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
	legacy, err := json.Marshal(supplyfulfilment.PickingList{})
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(legacy), "fulfillment_generation") {
		t.Fatalf("legacy picking list gained an asserted generation: %s", legacy)
	}
	legacyProjection, err := json.Marshal(pubsuborders.OrderPackingProjection{})
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(legacyProjection), "fulfillment_generation") {
		t.Fatalf("legacy packing projection gained an asserted generation: %s", legacyProjection)
	}
}
