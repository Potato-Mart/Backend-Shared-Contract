package warehouse_test

import (
	"encoding/json"
	"testing"
	"time"

	warehousecontract "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/contracts/warehouse"
	warehouseenum "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/enums/warehouse"
)

func TestOutboundShipmentDeliveredContractRoundTrips(t *testing.T) {
	deliveredAt := time.Date(2026, time.July, 11, 9, 30, 0, 0, time.UTC)
	want := warehousecontract.OutboundShipment{
		ID:          "shipment-1",
		Status:      warehouseenum.OutboundShipmentStatusDelivered,
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
	if got.Status != warehouseenum.OutboundShipmentStatusDelivered || got.DeliveredAt == nil || !got.DeliveredAt.Equal(deliveredAt) {
		t.Fatalf("delivered shipment round trip = %+v", got)
	}

	raw, err = json.Marshal(warehousecontract.OutboundShipment{
		ID:        "shipment-2",
		Status:    warehouseenum.OutboundShipmentStatusDispatched,
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
