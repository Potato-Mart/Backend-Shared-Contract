package courier

import (
	"encoding/json"
	"reflect"
	"sort"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/courier/courier_enums"
)

func TestCarrierManualActionJSON(t *testing.T) {
	const fixture = `{"id":"cma_1","order_number":"MAMA261112ABC123","shipment_ids":["shp_1","shp_2"],"depot_code":"AU-VIC-MEL-DC-01","market_code":"AU-VIC","country_code":"AU","kind":"cancel_booking","status":"resolved","delivery_company":{"code":"melbourne-fleet","name":"Local delivery","revision":7},"carrier_reference":"JOB-123","source_attempt_id":"cba_3","cause":"failed","error_code":"carrier_error","error_message":"Cancel endpoint returned 500","switch_reason":"Carrier cannot deliver tomorrow","replacement_company":{"code":"aupost","name":"Australia Post","revision":2},"resolved_at":"2026-11-12T05:00:00Z","resolved_by":{"actor_id":"staff_2","actor_email":"ops@example.test","actor_domain":"workforce","actor_role":"marketing"},"resolution_note":"Cancelled by phone","created_at":"2026-11-12T04:00:00Z","created_by":"staff_1","updated_at":"2026-11-12T05:00:00Z","updated_by":"staff_2"}`
	var action CarrierManualAction
	if err := json.Unmarshal([]byte(fixture), &action); err != nil {
		t.Fatal(err)
	}
	if action.Kind != courier_enums.CarrierManualActionKindCancelBooking || !action.Kind.IsValid() || action.Status != courier_enums.CarrierManualActionStatusResolved || !action.Status.IsValid() {
		t.Fatalf("kind or status lost: %+v", action)
	}
	if action.Cause != courier_enums.CarrierCancellationOutcomeFailed || !action.Cause.IsValid() || action.ErrorCode != courier_enums.CarrierBookingErrorCodeCarrierError {
		t.Fatalf("cancellation cause lost: %+v", action)
	}
	if action.DeliveryCompany.Code != "melbourne-fleet" || action.CarrierReference != "JOB-123" || action.SourceAttemptID != "cba_3" || len(action.ShipmentIDs) != 2 {
		t.Fatalf("booking identity lost: %+v", action)
	}
	if action.ReplacementCompany == nil || action.ReplacementCompany.Code != "aupost" || action.ResolvedAt == nil || action.ResolvedBy == nil || action.ResolvedBy.ActorID != "staff_2" {
		t.Fatalf("switch or resolution lost: %+v", action)
	}
	encoded, err := json.Marshal(action)
	if err != nil {
		t.Fatal(err)
	}
	var got, want map[string]any
	if err := json.Unmarshal(encoded, &got); err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal([]byte(fixture), &want); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("carrier manual action JSON changed:\n got: %s\nwant: %s", encoded, fixture)
	}
}

func TestOpenCarrierManualActionOmitsOptionalFields(t *testing.T) {
	createdAt := time.Date(2026, 11, 12, 4, 0, 0, 0, time.UTC)
	encoded, err := json.Marshal(CarrierManualAction{
		ID:              "cma_2",
		OrderNumber:     "MAMA261112ABC123",
		DepotCode:       "AU-VIC-MEL-DC-01",
		Kind:            courier_enums.CarrierManualActionKindCancelBooking,
		Status:          courier_enums.CarrierManualActionStatusOpen,
		DeliveryCompany: DeliveryCompanyRef{Code: "melbourne-fleet", Name: "Local delivery", Revision: 7},
		Cause:           courier_enums.CarrierCancellationOutcomeNotConfigured,
		AuditFields:     audit.AuditFields{CreatedAt: createdAt, UpdatedAt: createdAt},
	})
	if err != nil {
		t.Fatal(err)
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(encoded, &fields); err != nil {
		t.Fatal(err)
	}
	keys := make([]string, 0, len(fields))
	for key := range fields {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	want := []string{"cause", "created_at", "delivery_company", "depot_code", "id", "kind", "order_number", "status", "updated_at"}
	if !reflect.DeepEqual(keys, want) {
		t.Fatalf("open manual action keys = %v, want %v", keys, want)
	}
}

func TestCarrierBookingAttemptSupersessionJSON(t *testing.T) {
	const fixture = `{"id":"cba_3","order_number":"MAMA261112ABC123","shipment_ids":["shp_1"],"delivery_company":{"code":"melbourne-fleet","name":"Local delivery","revision":7},"outcome":"booked","carrier_reference":"JOB-123","sandbox":false,"attempted_by":"staff_1","attempted_at":"2026-11-12T01:00:00Z","resolved_at":"2026-11-12T01:00:02Z","supersession":{"at":"2026-11-12T04:00:00Z","by":{"actor_id":"staff_1","actor_domain":"workforce"},"reason":"Carrier cannot deliver tomorrow","replacement_company":{"code":"aupost","name":"Australia Post","revision":2},"cancellation_outcome":"failed","manual_action_id":"cma_1"},"created_at":"2026-11-12T01:00:00Z","updated_at":"2026-11-12T04:00:00Z"}`
	var attempt CarrierBookingAttempt
	if err := json.Unmarshal([]byte(fixture), &attempt); err != nil {
		t.Fatal(err)
	}
	// Supersession never rewrites the attempt's own outcome.
	if attempt.Outcome != courier_enums.CarrierBookingOutcomeBooked {
		t.Fatalf("outcome = %q, want booked", attempt.Outcome)
	}
	supersession := attempt.Supersession
	if supersession == nil || supersession.At.IsZero() || supersession.By.ActorID != "staff_1" || supersession.Reason == "" {
		t.Fatalf("supersession identity lost: %+v", supersession)
	}
	if supersession.ReplacementCompany == nil || supersession.ReplacementCompany.Code != "aupost" || supersession.CancellationOutcome != courier_enums.CarrierCancellationOutcomeFailed || supersession.ManualActionID != "cma_1" {
		t.Fatalf("supersession result lost: %+v", supersession)
	}
	encoded, err := json.Marshal(attempt)
	if err != nil {
		t.Fatal(err)
	}
	var got, want map[string]any
	if err := json.Unmarshal(encoded, &got); err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal([]byte(fixture), &want); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("superseded attempt JSON changed:\n got: %s\nwant: %s", encoded, fixture)
	}
}
