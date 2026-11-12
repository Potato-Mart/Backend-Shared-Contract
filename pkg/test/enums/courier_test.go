package enums_test

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/courier/courier_enums"
	"testing"
)

func TestCourierEnumsValidateKnownValues(t *testing.T) {
	if courier_enums.PostalCodeMatchMode("").IsValid() || courier_enums.DeliverySlotSource("").IsValid() {
		t.Fatal("empty configuration values must not acquire implicit defaults")
	}
	assertStringEnums(t, []enumCase{
		{name: "PostalCodeMatchMode", valid: []stringEnum{courier_enums.PostalCodeMatchModeIncludeOnly, courier_enums.PostalCodeMatchModeAllExcept}, invalid: courier_enums.PostalCodeMatchMode("__invalid__")},
		{name: "DeliverySlotSource", valid: []stringEnum{courier_enums.DeliverySlotSourceNone, courier_enums.DeliverySlotSourceConfiguredSchedule, courier_enums.DeliverySlotSourceProvider, courier_enums.DeliverySlotSourceStaffOverride}, invalid: courier_enums.DeliverySlotSource("__invalid__")},
		{name: "CarrierBookingOutcome", valid: []stringEnum{courier_enums.CarrierBookingOutcomePending, courier_enums.CarrierBookingOutcomeBooked, courier_enums.CarrierBookingOutcomeAlreadyBooked, courier_enums.CarrierBookingOutcomeRejected, courier_enums.CarrierBookingOutcomeUnreachable}, invalid: courier_enums.CarrierBookingOutcome("__invalid__")},
		{name: "CarrierBookingErrorCode", valid: []stringEnum{courier_enums.CarrierBookingErrorCodeInvalidAddress, courier_enums.CarrierBookingErrorCodeInvalidContact, courier_enums.CarrierBookingErrorCodeDuplicateJob, courier_enums.CarrierBookingErrorCodeInvalidDate, courier_enums.CarrierBookingErrorCodeInvalidParcel, courier_enums.CarrierBookingErrorCodeAuthFailed, courier_enums.CarrierBookingErrorCodeRateLimited, courier_enums.CarrierBookingErrorCodeTimeout, courier_enums.CarrierBookingErrorCodeUnsupportedArea, courier_enums.CarrierBookingErrorCodeCarrierError}, invalid: courier_enums.CarrierBookingErrorCode("__invalid__")},
		{name: "CarrierManualActionKind", valid: []stringEnum{courier_enums.CarrierManualActionKindCancelBooking}, invalid: courier_enums.CarrierManualActionKind("__invalid__")},
		{name: "CarrierManualActionStatus", valid: []stringEnum{courier_enums.CarrierManualActionStatusOpen, courier_enums.CarrierManualActionStatusResolved}, invalid: courier_enums.CarrierManualActionStatus("__invalid__")},
		{name: "CarrierCancellationOutcome", valid: []stringEnum{courier_enums.CarrierCancellationOutcomeCancelled, courier_enums.CarrierCancellationOutcomeAlreadyCancelled, courier_enums.CarrierCancellationOutcomeNotConfigured, courier_enums.CarrierCancellationOutcomeFailed, courier_enums.CarrierCancellationOutcomeUnreachable}, invalid: courier_enums.CarrierCancellationOutcome("__invalid__")},
	})
}

func TestCarrierBookingEnumWireValuesAreStable(t *testing.T) {
	if courier_enums.CarrierBookingOutcome("").IsValid() || courier_enums.CarrierBookingErrorCode("").IsValid() {
		t.Fatal("empty carrier booking values must not acquire implicit defaults")
	}
	for value, want := range map[stringEnum]string{
		courier_enums.CarrierBookingOutcomePending:          "pending",
		courier_enums.CarrierBookingOutcomeBooked:           "booked",
		courier_enums.CarrierBookingOutcomeAlreadyBooked:    "already_booked",
		courier_enums.CarrierBookingOutcomeRejected:         "rejected",
		courier_enums.CarrierBookingOutcomeUnreachable:      "unreachable",
		courier_enums.CarrierBookingErrorCodeInvalidAddress: "invalid_address",
		courier_enums.CarrierBookingErrorCodeInvalidContact: "invalid_contact",
		courier_enums.CarrierBookingErrorCodeDuplicateJob:   "duplicate_job",
		courier_enums.CarrierBookingErrorCodeInvalidDate:    "invalid_date",
		courier_enums.CarrierBookingErrorCodeInvalidParcel:  "invalid_parcel",
		courier_enums.CarrierBookingErrorCodeAuthFailed:     "auth_failed",
		courier_enums.CarrierBookingErrorCodeRateLimited:    "rate_limited",
		courier_enums.CarrierBookingErrorCodeCarrierError:   "carrier_error",
	} {
		if got := value.String(); got != want {
			t.Errorf("%T wire value = %q, want %q", value, got, want)
		}
	}
}

func TestCarrierSwitchEnumWireValuesAreStable(t *testing.T) {
	if courier_enums.CarrierManualActionKind("").IsValid() || courier_enums.CarrierManualActionStatus("").IsValid() || courier_enums.CarrierCancellationOutcome("").IsValid() {
		t.Fatal("empty carrier switch values must not acquire implicit defaults")
	}
	for value, want := range map[stringEnum]string{
		courier_enums.DeliverySlotSourceStaffOverride:            "staff_override",
		courier_enums.CarrierManualActionKindCancelBooking:       "cancel_booking",
		courier_enums.CarrierManualActionStatusOpen:              "open",
		courier_enums.CarrierManualActionStatusResolved:          "resolved",
		courier_enums.CarrierCancellationOutcomeCancelled:        "cancelled",
		courier_enums.CarrierCancellationOutcomeAlreadyCancelled: "already_cancelled",
		courier_enums.CarrierCancellationOutcomeNotConfigured:    "not_configured",
		courier_enums.CarrierCancellationOutcomeFailed:           "failed",
		courier_enums.CarrierCancellationOutcomeUnreachable:      "unreachable",
	} {
		if got := value.String(); got != want {
			t.Errorf("%T wire value = %q, want %q", value, got, want)
		}
	}
}
