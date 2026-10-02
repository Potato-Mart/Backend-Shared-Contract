package payments

import (
	"encoding/json"
	"github.com/Potato-Mart/Backend-Shared-Contract/v42/pkg/contracts/payments/payment"
	"github.com/Potato-Mart/Backend-Shared-Contract/v42/pkg/contracts/payments/payment/payment_enums"
	"reflect"
	"testing"
	"time"
)

func TestCaptureV2PreservesConfirmationWithoutInventingExactTime(t *testing.T) {
	bound := time.Date(2026, 9, 29, 3, 0, 1, 0, time.UTC)
	evidence := payment.CaptureTimingEvidence{Confirmation: &payment.CaptureConfirmationEvidence{
		UpperBoundAt: bound, Source: payment_enums.CaptureConfirmationSourceProviderEvent, Provider: "stripe", EvidenceReference: "evt_capture_1",
	}}
	data, err := json.Marshal(evidence)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"confirmation":{"upper_bound_at":"2026-09-29T03:00:01Z","source":"provider_event","provider":"stripe","evidence_reference":"evt_capture_1"}}`
	if string(data) != want {
		t.Fatalf("capture timing = %s, want %s", data, want)
	}
	event := PaymentCapturedEventV2{PaymentID: "payment-1", OrderNumber: "GC000000000001", CaptureTiming: evidence}
	data, err = json.Marshal(event)
	if err != nil {
		t.Fatal(err)
	}
	var decoded PaymentCapturedEventV2
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(event, decoded) || decoded.CaptureTiming.CapturedAt != nil {
		t.Fatal("capture evidence changed on replay")
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		t.Fatal(err)
	}
	if _, found := fields["captured_at"]; found {
		t.Fatal("v2 must not expose fabricated top-level exact time")
	}
	if _, found := fields["capture_timing"]; !found {
		t.Fatal("v2 timing evidence is required")
	}
}

func TestCaptureV2KeepsExactTimeDistinctFromLaterConfirmation(t *testing.T) {
	exact := time.Date(2026, 9, 29, 3, 0, 0, 0, time.UTC)
	bound := exact.Add(20 * time.Minute)
	evidence := payment.CaptureTimingEvidence{CapturedAt: &exact, Confirmation: &payment.CaptureConfirmationEvidence{
		UpperBoundAt: bound, Source: payment_enums.CaptureConfirmationSourceProviderStatusObservation, Provider: "terminal", EvidenceReference: "observation-1",
	}}
	data, err := json.Marshal(evidence)
	if err != nil {
		t.Fatal(err)
	}
	var decoded payment.CaptureTimingEvidence
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(evidence, decoded) {
		t.Fatal("exact capture time was replaced by confirmation or receipt time")
	}
	evidence.CapturedAt = nil
	data, err = json.Marshal(evidence)
	if err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatal(err)
	}
	// Decode into a fresh value: absent exact time remains unknown even when the
	// confirmation is later than an order's deadline. Decisions stay service-owned.
	var unknown payment.CaptureTimingEvidence
	if err := json.Unmarshal(data, &unknown); err != nil {
		t.Fatal(err)
	}
	if unknown.CapturedAt != nil || !unknown.Confirmation.UpperBoundAt.Equal(bound) {
		t.Fatal("unknown exact time was lost")
	}
}

func TestCaptureTimingAbsenceAndLegacyV1RemainDistinct(t *testing.T) {
	data, err := json.Marshal(payment.Payment{})
	if err != nil {
		t.Fatal(err)
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		t.Fatal(err)
	}
	if _, present := fields["capture_timing"]; present {
		t.Fatal("optional historical timing evidence must be absent")
	}
	field, _ := reflect.TypeOf(PaymentCapturedEvent{}).FieldByName("CapturedAt")
	if field.Type != reflect.TypeOf(time.Time{}) || field.Tag.Get("json") != "captured_at" {
		t.Fatal("additive release changed legacy v1 shape")
	}
}

func TestPaymentRetainsImmutableCaptureEvidence(t *testing.T) {
	bound := time.Date(2026, 9, 29, 3, 0, 1, 0, time.UTC)
	record := payment.Payment{CaptureTiming: &payment.CaptureTimingEvidence{Confirmation: &payment.CaptureConfirmationEvidence{
		UpperBoundAt: bound, Source: payment_enums.CaptureConfirmationSourceProviderEvent, Provider: "stripe", EvidenceReference: "evt_capture_1",
	}}}
	data, err := json.Marshal(record)
	if err != nil {
		t.Fatal(err)
	}
	var decoded payment.Payment
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(record.CaptureTiming, decoded.CaptureTiming) {
		t.Fatal("persisted immutable timing evidence changed")
	}
}
