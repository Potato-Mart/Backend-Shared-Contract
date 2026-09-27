package orders_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v37/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v37/pkg/contracts/common/packaging"
	orders "github.com/Potato-Mart/Backend-Shared-Contract/v37/pkg/contracts/pubsub/orders"
)

func TestOrderEditedEventCarriesWholeRequestedCompositionAndStableEditIdentity(t *testing.T) {
	before := orders.OrderEditedItemSnapshot{
		OrderItemID: "line_1", SKUCode: "POTATO-A", ProductName: "Potatoes",
		PackageOptionCode: "CASE-12",
		RequestedComposition: packaging.PackageCompositionSnapshot{TotalBaseUnits: 36, Components: []packaging.PackageComponentSnapshot{{
			PackageOptionCode: "CASE-12", PackageCount: 3, UnitsPerPackage: 12, BaseUnits: 36,
		}}},
	}
	after := before
	after.RequestedComposition = packaging.PackageCompositionSnapshot{TotalBaseUnits: 24, Components: []packaging.PackageComponentSnapshot{{
		PackageOptionCode: "CASE-12", PackageCount: 2, UnitsPerPackage: 12, BaseUnits: 24,
	}}}
	event := orders.OrderEditedEvent{
		OrderID: "order_1", OrderNumber: "SO-1", EditID: "edit_1",
		PreviousFulfillmentGeneration: 4, FulfillmentGeneration: 5,
		RetailCustomerNumber: "customer_1", MarketCode: "AU", CountryCode: "AU", DepotCode: "depot_1",
		PreviousItems: []orders.OrderEditedItemSnapshot{before}, RevisedItems: []orders.OrderEditedItemSnapshot{after},
		PreviousTotal: money.Money{AmountMinor: 3600, Currency: "AUD"},
		RevisedTotal:  money.Money{AmountMinor: 2400, Currency: "AUD"},
		CommittedAt:   time.Date(2026, 9, 27, 1, 2, 3, 0, time.UTC),
	}
	encoded, err := json.Marshal(event)
	if err != nil {
		t.Fatal(err)
	}
	for _, field := range []string{`"edit_id":"edit_1"`, `"fulfillment_generation":5`, `"previous_items"`, `"revised_items"`, `"package_count":2`, `"revised_total"`} {
		if !strings.Contains(string(encoded), field) {
			t.Errorf("committed edit event is missing %s: %s", field, encoded)
		}
	}
	for _, forbidden := range []string{`"packed_composition"`, `"payment_method"`, `"password"`} {
		if strings.Contains(string(encoded), forbidden) {
			t.Errorf("committed edit event exposed %s: %s", forbidden, encoded)
		}
	}
	var decoded orders.OrderEditedEvent
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		t.Fatal(err)
	}
	if decoded.EditID != "edit_1" || decoded.FulfillmentGeneration != 5 || len(decoded.RevisedItems) != 1 || decoded.RevisedItems[0].RequestedComposition.Components[0].PackageCount != 2 {
		t.Fatalf("committed edit evidence changed: %+v", decoded)
	}
}
