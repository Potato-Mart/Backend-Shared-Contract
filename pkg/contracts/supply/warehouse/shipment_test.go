package warehouse_test

import (
	"encoding/json"
	"testing"
	"time"

	warehousecontract "github.com/Potato-Mart/Backend-Shared-Contract/v24/pkg/contracts/supply/warehouse"
	"github.com/Potato-Mart/Backend-Shared-Contract/v24/pkg/contracts/supply/warehouse/warehouse_enums"
)

func TestOutboundShipmentDeliveredContractRoundTrips(t *testing.T) {
	deliveredAt := time.Date(2026, time.July, 11, 9, 30, 0, 0, time.UTC)
	want := warehousecontract.OutboundShipment{
		ID:          "shipment-1",
		Status:      warehouse_enums.OutboundShipmentStatusDelivered,
		DeliveredAt: &deliveredAt,
		CreatedAt:   deliveredAt.Add(-time.Hour),
	}

	raw, err := json.Marshal(want)
	if err != nil {
		t.Fatalf("marshal delivered shipment: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(raw, &payload); err != nil {
		t.Fatalf("unmarshal delivered shipment JSON shape: %v", err)
	}
	if _, ok := payload["delivered_at"]; !ok {
		t.Fatalf("delivered shipment JSON = %s, want delivered_at", raw)
	}
	var got warehousecontract.OutboundShipment
	if err := json.Unmarshal(raw, &got); err != nil {
		t.Fatalf("unmarshal delivered shipment: %v", err)
	}
	if got.Status != warehouse_enums.OutboundShipmentStatusDelivered || got.DeliveredAt == nil || !got.DeliveredAt.Equal(deliveredAt) {
		t.Fatalf("delivered shipment round trip = %+v", got)
	}

	raw, err = json.Marshal(warehousecontract.OutboundShipment{
		ID:        "shipment-2",
		Status:    warehouse_enums.OutboundShipmentStatusDispatched,
		CreatedAt: deliveredAt.Add(-2 * time.Hour),
	})
	if err != nil {
		t.Fatalf("marshal shipment without delivery time: %v", err)
	}
	payload = nil
	if err := json.Unmarshal(raw, &payload); err != nil {
		t.Fatalf("unmarshal shipment JSON without delivery time: %v", err)
	}
	if _, ok := payload["delivered_at"]; ok {
		t.Fatalf("shipment JSON = %s, want delivered_at omitted", raw)
	}
}
