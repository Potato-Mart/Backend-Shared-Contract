package orders_test

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/customers/retail/retail_enums"
	notification "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/pubsub/orders"
)

func TestPickingFactJSONRoundTripsLinesAndPickers(t *testing.T) {
	printed := time.Date(2026, 11, 6, 0, 30, 0, 0, time.UTC)
	firstPick := printed.Add(5 * time.Minute)
	done := printed.Add(20 * time.Minute)
	fact := notification.PickingFact{
		FactID: "picking:SO-1:2", OrderNumber: "SO-1", FulfillmentGeneration: 2,
		MarketCode: "MEL", DepotCode: "MAMA", BuyerType: retail_enums.BuyerTypeWholesaleOrganisation,
		PickingStartedAt: &printed, PickingCompletedAt: done.Add(time.Minute),
		Lines: []notification.PickingFactLine{
			{OrderItemID: "item_1", SKUCode: "SKU-1", OrderedBaseUnits: 48, PickedBaseUnits: 46, ShortageBaseUnits: 2},
		},
		Pickers: []notification.PickingFactPicker{{
			Role: "primary", PickerSubject: "staff_7", StartedAt: &firstPick, CompletedAt: &done,
			Lines: []notification.PickingFactPickerLine{{OrderItemID: "item_1", PickedBaseUnits: 46, ShortageBaseUnits: 2}},
		}},
	}
	payload, err := json.Marshal(fact)
	if err != nil {
		t.Fatalf("marshal picking fact: %v", err)
	}
	for _, want := range []string{
		`"fact_id":"picking:SO-1:2"`, `"buyer_type":"wholesale_organisation"`,
		`"picking_started_at":"2026-11-06T00:30:00Z"`, `"ordered_base_units":48`, `"picker_subject":"staff_7"`,
	} {
		if !strings.Contains(string(payload), want) {
			t.Fatalf("picking fact JSON missing %s: %s", want, payload)
		}
	}
	var decoded notification.PickingFact
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal picking fact: %v", err)
	}
	if !reflect.DeepEqual(decoded, fact) {
		t.Fatalf("picking fact round trip changed: %+v", decoded)
	}
}

func TestPickingFactWithoutAssignmentOmitsPickers(t *testing.T) {
	payload, err := json.Marshal(notification.PickingFact{FactID: "picking:SO-1:1", OrderNumber: "SO-1", Lines: []notification.PickingFactLine{}})
	if err != nil {
		t.Fatalf("marshal picking fact: %v", err)
	}
	if strings.Contains(string(payload), `"pickers"`) {
		t.Fatalf("an order picked without an assignment must omit pickers: %s", payload)
	}
}
