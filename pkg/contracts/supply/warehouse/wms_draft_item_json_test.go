package warehouse_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/metadata"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/warehouse"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/warehouse/warehouse_enums"
)

func TestWMSDraftItemExpiryMonthAndMoveRoundTrip(t *testing.T) {
	bare, err := json.Marshal(warehouse.WMSDraftItem{ID: "item_1", SKUCode: "A00001"})
	if err != nil {
		t.Fatalf("marshal bare item: %v", err)
	}
	var bareShape map[string]any
	if err := json.Unmarshal(bare, &bareShape); err != nil {
		t.Fatalf("decode bare item: %v", err)
	}
	for _, key := range []string{"expiry_month", "move"} {
		if _, ok := bareShape[key]; ok {
			t.Fatalf("empty %s must be omitted: %s", key, bare)
		}
	}

	item := warehouse.WMSDraftItem{
		ID: "item_1", SKUCode: "A00001", ExpiryMonth: "2027-03",
		Move: &warehouse.WMSDraftItemMove{
			Key:  "move_key_1",
			Step: warehouse_enums.WMSDraftItemMoveStepMoved,
			Request: metadata.Metadata{
				"sku_code": "A00001",
				"from":     map[string]any{"depot_code": "A7", "location_code": "SYS-GOODS-POOL"},
				"lines":    []any{map[string]any{"lot_id": "lot_1", "note": "top shelf"}},
			},
		},
	}
	raw, err := json.Marshal(item)
	if err != nil {
		t.Fatalf("marshal item: %v", err)
	}
	var back warehouse.WMSDraftItem
	if err := json.Unmarshal(raw, &back); err != nil {
		t.Fatalf("unmarshal item: %v", err)
	}
	if back.ExpiryMonth != "2027-03" || back.Move == nil || back.Move.Key != "move_key_1" || back.Move.Step != warehouse_enums.WMSDraftItemMoveStepMoved {
		t.Fatalf("expiry month or move did not round-trip: %s", raw)
	}
	if !reflect.DeepEqual(back.Move.Request, item.Move.Request) {
		t.Fatalf("move request did not round-trip: got %+v, want %+v", back.Move.Request, item.Move.Request)
	}

	withoutRequest, err := json.Marshal(warehouse.WMSDraftItemMove{Key: "move_key_1", Step: warehouse_enums.WMSDraftItemMoveStepPending})
	if err != nil {
		t.Fatalf("marshal move: %v", err)
	}
	if string(withoutRequest) != `{"key":"move_key_1","step":"PENDING"}` {
		t.Fatalf("move without request = %s", withoutRequest)
	}
}

func TestInventoryLotOriginLotIDIsOptional(t *testing.T) {
	bare, err := json.Marshal(warehouse.InventoryLot{ID: "lot_1", SKUCode: "A00001"})
	if err != nil {
		t.Fatalf("marshal lot: %v", err)
	}
	var bareShape map[string]any
	if err := json.Unmarshal(bare, &bareShape); err != nil {
		t.Fatalf("decode lot: %v", err)
	}
	if _, ok := bareShape["origin_lot_id"]; ok {
		t.Fatalf("empty origin_lot_id must be omitted: %s", bare)
	}

	raw, err := json.Marshal(warehouse.InventoryLot{ID: "lot_2", SKUCode: "A00001", OriginLotID: "lot_1"})
	if err != nil {
		t.Fatalf("marshal split lot: %v", err)
	}
	var back warehouse.InventoryLot
	if err := json.Unmarshal(raw, &back); err != nil {
		t.Fatalf("unmarshal split lot: %v", err)
	}
	if back.OriginLotID != "lot_1" {
		t.Fatalf("origin_lot_id did not round-trip: %s", raw)
	}
}
