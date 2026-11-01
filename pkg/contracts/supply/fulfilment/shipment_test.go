package fulfilment_test

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/party"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/courier"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/courier/courier_enums"
	shippingcontract "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/fulfilment"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/warehouse/warehouse_enums"
)

func TestOutboundShipmentDeliveredContractRoundTrips(t *testing.T) {
	deliveredAt := time.Date(2026, time.July, 11, 9, 30, 0, 0, time.UTC)
	want := shippingcontract.OutboundShipment{
		ID:          "shipment-1",
		Status:      warehouse_enums.OutboundShipmentStatusDelivered,
		DeliveredAt: &deliveredAt,
		AuditFields: audit.AuditFields{CreatedAt: deliveredAt.Add(-time.Hour), UpdatedAt: deliveredAt.Add(-time.Hour)},
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
	var got shippingcontract.OutboundShipment
	if err := json.Unmarshal(raw, &got); err != nil {
		t.Fatalf("unmarshal delivered shipment: %v", err)
	}
	if got.Status != warehouse_enums.OutboundShipmentStatusDelivered || got.DeliveredAt == nil || !got.DeliveredAt.Equal(deliveredAt) {
		t.Fatalf("delivered shipment round trip = %+v", got)
	}

	raw, err = json.Marshal(shippingcontract.OutboundShipment{
		ID:          "shipment-2",
		Status:      warehouse_enums.OutboundShipmentStatusDispatched,
		AuditFields: audit.AuditFields{CreatedAt: deliveredAt.Add(-2 * time.Hour), UpdatedAt: deliveredAt.Add(-2 * time.Hour)},
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

func TestOutboundShipmentRecipientAndLatestCarrierBookingRoundTrip(t *testing.T) {
	outcomeAt := time.Date(2026, time.November, 1, 9, 31, 0, 0, time.UTC)
	want := shippingcontract.OutboundShipment{
		ID:          "shipment-3",
		OrderNumber: "MAMA261101ABC123",
		Status:      warehouse_enums.OutboundShipmentStatusPacked,
		Recipient:   &party.Recipient{Name: "Alex Doe", Phone: "+61400000000", Email: "alex@example.test"},
		LatestCarrierBooking: &courier.CarrierBookingSummary{
			AttemptID:        "cba_3",
			Outcome:          courier_enums.CarrierBookingOutcomeBooked,
			CarrierReference: "JOB-123",
			OutcomeAt:        outcomeAt,
		},
		AuditFields:          audit.AuditFields{CreatedAt: outcomeAt, UpdatedAt: outcomeAt},
		DataProtectionFields: security.DataProtectionFields{ContainsPII: true},
	}
	raw, err := json.Marshal(want)
	if err != nil {
		t.Fatalf("marshal shipment: %v", err)
	}
	var payload map[string]json.RawMessage
	if err := json.Unmarshal(raw, &payload); err != nil {
		t.Fatalf("unmarshal shipment JSON shape: %v", err)
	}
	if string(payload["recipient"]) != `{"name":"Alex Doe","phone":"+61400000000","email":"alex@example.test"}` {
		t.Fatalf("recipient JSON = %s", payload["recipient"])
	}
	if string(payload["latest_carrier_booking"]) != `{"attempt_id":"cba_3","outcome":"booked","carrier_reference":"JOB-123","outcome_at":"2026-11-01T09:31:00Z"}` {
		t.Fatalf("latest_carrier_booking JSON = %s", payload["latest_carrier_booking"])
	}
	if string(payload["contains_pii"]) != "true" {
		t.Fatalf("shipment with a recipient must carry its PII label: %s", raw)
	}
	var got shippingcontract.OutboundShipment
	if err := json.Unmarshal(raw, &got); err != nil {
		t.Fatalf("unmarshal shipment: %v", err)
	}
	if !reflect.DeepEqual(got.Recipient, want.Recipient) || got.LatestCarrierBooking == nil || got.LatestCarrierBooking.AttemptID != "cba_3" || !got.LatestCarrierBooking.OutcomeAt.Equal(outcomeAt) {
		t.Fatalf("shipment round trip = %+v", got)
	}

	raw, err = json.Marshal(shippingcontract.OutboundShipment{ID: "shipment-4", Status: warehouse_enums.OutboundShipmentStatusPacked})
	if err != nil {
		t.Fatalf("marshal legacy shipment: %v", err)
	}
	payload = nil
	if err := json.Unmarshal(raw, &payload); err != nil {
		t.Fatalf("unmarshal legacy shipment: %v", err)
	}
	for _, key := range []string{"recipient", "latest_carrier_booking"} {
		if _, ok := payload[key]; ok {
			t.Fatalf("shipment JSON = %s, want %s omitted", raw, key)
		}
	}
}
