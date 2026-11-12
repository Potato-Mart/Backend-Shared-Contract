package fulfilment_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	shippingcontract "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/fulfilment"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/warehouse/warehouse_enums"
)

func TestOutboundShipmentWarningsRoundTrip(t *testing.T) {
	createdAt := time.Date(2026, time.November, 12, 4, 0, 0, 0, time.UTC)
	raw, err := json.Marshal(shippingcontract.OutboundShipment{
		ID:     "shipment-5",
		Status: warehouse_enums.OutboundShipmentStatusPacked,
		Warnings: []shippingcontract.ShipmentWarning{
			{
				Code:             warehouse_enums.ShipmentWarningCodeCarrierCancellationRequired,
				Message:          "Cancel the earlier booking with the carrier",
				ManualActionID:   "cma_1",
				CarrierCode:      "melbourne-fleet",
				CarrierReference: "JOB-123",
				CreatedAt:        &createdAt,
			},
			{Code: warehouse_enums.ShipmentWarningCodeOrderUnpaid},
		},
	})
	if err != nil {
		t.Fatalf("marshal shipment: %v", err)
	}
	var payload map[string]json.RawMessage
	if err := json.Unmarshal(raw, &payload); err != nil {
		t.Fatalf("unmarshal shipment JSON shape: %v", err)
	}
	const wantWarnings = `[{"code":"carrier_cancellation_required","message":"Cancel the earlier booking with the carrier","manual_action_id":"cma_1","carrier_code":"melbourne-fleet","carrier_reference":"JOB-123","created_at":"2026-11-12T04:00:00Z"},{"code":"order_unpaid"}]`
	if string(payload["warnings"]) != wantWarnings {
		t.Fatalf("warnings JSON = %s, want %s", payload["warnings"], wantWarnings)
	}
	var got shippingcontract.OutboundShipment
	if err := json.Unmarshal(raw, &got); err != nil {
		t.Fatalf("unmarshal shipment: %v", err)
	}
	if len(got.Warnings) != 2 || !got.Warnings[0].Code.IsValid() || got.Warnings[0].CreatedAt == nil || !got.Warnings[0].CreatedAt.Equal(createdAt) || got.Warnings[1].Code != warehouse_enums.ShipmentWarningCodeOrderUnpaid {
		t.Fatalf("warnings round trip = %+v", got.Warnings)
	}

	raw, err = json.Marshal(shippingcontract.OutboundShipment{ID: "shipment-6", Status: warehouse_enums.OutboundShipmentStatusPacked})
	if err != nil {
		t.Fatalf("marshal shipment without warnings: %v", err)
	}
	if strings.Contains(string(raw), `"warnings"`) {
		t.Fatalf("shipment JSON = %s, want warnings omitted", raw)
	}
}
