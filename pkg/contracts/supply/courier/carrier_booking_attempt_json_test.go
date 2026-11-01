package courier

import (
	"encoding/json"
	"reflect"
	"sort"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/audit"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/temporal"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/courier/courier_enums"
)

func TestCarrierBookingAttemptJSON(t *testing.T) {
	const fixture = `{"id":"cba_2","order_number":"MAMA261101ABC123","shipment_ids":["shp_1","shp_2"],"delivery_company":{"code":"melbourne-fleet","name":"Local delivery","revision":7},"outcome":"rejected","request_summary":[{"key":"job_type","value":"delivery"},{"key":"postal_code","value":"3000"}],"error_code":"invalid_address","error_fields":["address.line1"],"error_message":"Address could not be geocoded","override_reason":"Selected carrier cannot serve this area today","edits":{"intake_date":"2026-11-02","delivery_date":"2026-11-03","time_window":{"from_time":"09:00","to_time":"12:00"},"box_count":3,"instructions":"","delivery_type":"Home delivery"},"sandbox":true,"attempted_by":"staff_1","attempted_at":"2026-11-01T09:30:00Z","resolved_at":"2026-11-01T09:30:02Z","classification":"confidential","contains_pii":true,"retention_policy_key":"carrier-booking","created_at":"2026-11-01T09:30:00Z","created_by":"staff_1","updated_at":"2026-11-01T09:30:02Z"}`
	var attempt CarrierBookingAttempt
	if err := json.Unmarshal([]byte(fixture), &attempt); err != nil {
		t.Fatal(err)
	}
	if attempt.Outcome != courier_enums.CarrierBookingOutcomeRejected || !attempt.Outcome.IsValid() {
		t.Fatalf("outcome = %q", attempt.Outcome)
	}
	if attempt.ErrorCode != courier_enums.CarrierBookingErrorCodeInvalidAddress || !attempt.ErrorCode.IsValid() {
		t.Fatalf("error code = %q", attempt.ErrorCode)
	}
	if len(attempt.ShipmentIDs) != 2 || attempt.DeliveryCompany.Code != "melbourne-fleet" || attempt.DeliveryCompany.Revision != 7 {
		t.Fatalf("order-level booking identity lost: %+v", attempt)
	}
	edits := attempt.Edits
	if edits == nil || edits.IntakeDate == nil || *edits.IntakeDate != "2026-11-02" || edits.DeliveryDate == nil || *edits.DeliveryDate != "2026-11-03" || edits.TimeWindow == nil || edits.TimeWindow.FromTime == nil || *edits.TimeWindow.FromTime != "09:00" || edits.TimeWindow.ToTime == nil || *edits.TimeWindow.ToTime != "12:00" || edits.BoxCount == nil || *edits.BoxCount != 3 {
		t.Fatalf("retry edits lost: %+v", edits)
	}
	// A present empty string clears the value; nil leaves it unedited.
	if edits.Instructions == nil || *edits.Instructions != "" || edits.DeliveryType == nil || *edits.DeliveryType != "Home delivery" {
		t.Fatalf("clearable text edits lost: %+v", edits)
	}
	if attempt.CreatedBy != "staff_1" || attempt.UpdatedAt.IsZero() {
		t.Fatalf("audit fields lost: %+v", attempt.AuditFields)
	}
	if !attempt.Sandbox || attempt.ResolvedAt == nil || !attempt.ContainsPII || attempt.Classification != "confidential" {
		t.Fatalf("sandbox, resolution or protection labels lost: %+v", attempt)
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
		t.Fatalf("carrier booking attempt JSON changed:\n got: %s\nwant: %s", encoded, fixture)
	}
}

func TestPendingCarrierBookingAttemptOmitsOptionalFields(t *testing.T) {
	attemptedAt := time.Date(2026, 11, 1, 9, 30, 0, 0, time.UTC)
	encoded, err := json.Marshal(CarrierBookingAttempt{
		ID:                   "cba_1",
		OrderNumber:          "MAMA261101ABC123",
		ShipmentIDs:          []string{"shp_1"},
		DeliveryCompany:      DeliveryCompanyRef{Code: "melbourne-fleet", Name: "Local delivery", Revision: 7},
		AuditFields:          audit.AuditFields{CreatedAt: attemptedAt, UpdatedAt: attemptedAt},
		Outcome:              courier_enums.CarrierBookingOutcomePending,
		AttemptedBy:          "staff_1",
		AttemptedAt:          attemptedAt,
		DataProtectionFields: security.DataProtectionFields{},
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
	want := []string{"attempted_at", "attempted_by", "created_at", "delivery_company", "id", "order_number", "outcome", "sandbox", "shipment_ids", "updated_at"}
	if !reflect.DeepEqual(keys, want) {
		t.Fatalf("pending attempt keys = %v, want %v", keys, want)
	}
}

func TestCarrierBookingSummaryJSON(t *testing.T) {
	const fixture = `{"attempt_id":"cba_3","outcome":"booked","carrier_reference":"JOB-123","outcome_at":"2026-11-01T09:31:00Z"}`
	var summary CarrierBookingSummary
	if err := json.Unmarshal([]byte(fixture), &summary); err != nil {
		t.Fatal(err)
	}
	if summary.AttemptID != "cba_3" || summary.Outcome != courier_enums.CarrierBookingOutcomeBooked || summary.CarrierReference != "JOB-123" || summary.OutcomeAt.IsZero() {
		t.Fatalf("summary lost: %+v", summary)
	}
	encoded, err := json.Marshal(summary)
	if err != nil {
		t.Fatal(err)
	}
	if string(encoded) != fixture {
		t.Fatalf("summary JSON = %s, want %s", encoded, fixture)
	}
	encoded, err = json.Marshal(CarrierBookingSummary{AttemptID: "cba_4", Outcome: courier_enums.CarrierBookingOutcomeUnreachable, OutcomeAt: summary.OutcomeAt})
	if err != nil {
		t.Fatal(err)
	}
	if want := `{"attempt_id":"cba_4","outcome":"unreachable","outcome_at":"2026-11-01T09:31:00Z"}`; string(encoded) != want {
		t.Fatalf("summary without reference = %s, want %s", encoded, want)
	}
}

func TestCarrierBookingEditsDistinguishUneditedFromCleared(t *testing.T) {
	encoded, err := json.Marshal(CarrierBookingEdits{})
	if err != nil {
		t.Fatal(err)
	}
	if string(encoded) != `{}` {
		t.Fatalf("unedited edits JSON = %s, want {}", encoded)
	}
	cleared, zero := "", 0
	encoded, err = json.Marshal(CarrierBookingEdits{Instructions: &cleared, DeliveryType: &cleared, BoxCount: &zero})
	if err != nil {
		t.Fatal(err)
	}
	if want := `{"box_count":0,"instructions":"","delivery_type":""}`; string(encoded) != want {
		t.Fatalf("cleared edits JSON = %s, want %s", encoded, want)
	}
	var decoded CarrierBookingEdits
	if err := json.Unmarshal([]byte(`{"instructions":""}`), &decoded); err != nil {
		t.Fatal(err)
	}
	if decoded.Instructions == nil || *decoded.Instructions != "" || decoded.DeliveryType != nil || decoded.BoxCount != nil || decoded.IntakeDate != nil || decoded.DeliveryDate != nil || decoded.TimeWindow != nil {
		t.Fatalf("cleared instructions must decode as a present empty value only: %+v", decoded)
	}
}

// A window edit may change one side only; the missing side keeps the value
// originally sent to the carrier.
func TestCarrierBookingTimeWindowMayBeOneSided(t *testing.T) {
	var edits CarrierBookingEdits
	if err := json.Unmarshal([]byte(`{"time_window":{"to_time":"14:00"}}`), &edits); err != nil {
		t.Fatal(err)
	}
	window := edits.TimeWindow
	if window == nil || window.FromTime != nil || window.ToTime == nil || *window.ToTime != "14:00" {
		t.Fatalf("one-sided window = %+v", window)
	}
	encoded, err := json.Marshal(edits)
	if err != nil {
		t.Fatal(err)
	}
	if want := `{"time_window":{"to_time":"14:00"}}`; string(encoded) != want {
		t.Fatalf("one-sided window JSON = %s, want %s", encoded, want)
	}
	from := temporal.TimeOfDay("08:30")
	encoded, err = json.Marshal(CarrierBookingTimeWindow{FromTime: &from})
	if err != nil {
		t.Fatal(err)
	}
	if want := `{"from_time":"08:30"}`; string(encoded) != want {
		t.Fatalf("start-only window JSON = %s, want %s", encoded, want)
	}
}

// The summary carries the values actually sent in the latest booking.
func TestCarrierBookingSummaryCarriesBookedValues(t *testing.T) {
	const fixture = `{"attempt_id":"cba_5","outcome":"booked","carrier_reference":"JOB-456","outcome_at":"2026-11-01T10:00:00Z","delivery_date":"2026-11-03","intake_date":"2026-11-02","time_window":{"from_time":"09:00","to_time":"12:00"},"box_count":2,"delivery_type":"Home delivery","instructions":""}`
	var summary CarrierBookingSummary
	if err := json.Unmarshal([]byte(fixture), &summary); err != nil {
		t.Fatal(err)
	}
	if summary.DeliveryDate == nil || *summary.DeliveryDate != "2026-11-03" || summary.IntakeDate == nil || *summary.IntakeDate != "2026-11-02" {
		t.Fatalf("booked dates lost: %+v", summary)
	}
	if summary.TimeWindow == nil || summary.TimeWindow.FromTime == nil || *summary.TimeWindow.FromTime != "09:00" || summary.TimeWindow.ToTime == nil || *summary.TimeWindow.ToTime != "12:00" {
		t.Fatalf("booked window lost: %+v", summary.TimeWindow)
	}
	if summary.BoxCount == nil || *summary.BoxCount != 2 || summary.DeliveryType == nil || *summary.DeliveryType != "Home delivery" || summary.Instructions == nil || *summary.Instructions != "" {
		t.Fatalf("booked values lost: %+v", summary)
	}
	encoded, err := json.Marshal(summary)
	if err != nil {
		t.Fatal(err)
	}
	if string(encoded) != fixture {
		t.Fatalf("summary JSON = %s, want %s", encoded, fixture)
	}
}
