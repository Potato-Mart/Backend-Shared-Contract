package enums_test

import (
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v38/pkg/contracts/supply/courier/courier_enums"
)

func TestCourierEnumsValidateKnownValues(t *testing.T) {
	if courier_enums.PostalCodeMatchMode("").IsValid() ||
		courier_enums.DeliverySlotSource("").IsValid() ||
		courier_enums.DeliveryConnectionHealth("").IsValid() ||
		courier_enums.DeliveryAPIHTTPMethod("").IsValid() ||
		courier_enums.DeliveryAPIRequestPurpose("").IsValid() {
		t.Fatal("empty configuration values must not acquire implicit defaults")
	}
	assertStringEnums(t, []enumCase{
		{name: "DeliveryAPIHTTPMethod", valid: []stringEnum{courier_enums.DeliveryAPIHTTPMethodGet, courier_enums.DeliveryAPIHTTPMethodPost}, invalid: courier_enums.DeliveryAPIHTTPMethod("PATCH")},
		{name: "DeliveryAPIRequestPurpose", valid: []stringEnum{courier_enums.DeliveryAPIRequestPurposeConnection, courier_enums.DeliveryAPIRequestPurposeShippingAreas, courier_enums.DeliveryAPIRequestPurposeTimeSlots}, invalid: courier_enums.DeliveryAPIRequestPurpose("booking")},
		{name: "DeliveryConnectionHealth", valid: []stringEnum{courier_enums.DeliveryConnectionHealthUnknown, courier_enums.DeliveryConnectionHealthHealthy, courier_enums.DeliveryConnectionHealthUnhealthy}, invalid: courier_enums.DeliveryConnectionHealth("__invalid__")},
		{name: "PostalCodeMatchMode", valid: []stringEnum{courier_enums.PostalCodeMatchModeIncludeOnly, courier_enums.PostalCodeMatchModeAllExcept}, invalid: courier_enums.PostalCodeMatchMode("__invalid__")},
		{name: "DeliverySlotSource", valid: []stringEnum{courier_enums.DeliverySlotSourceNone, courier_enums.DeliverySlotSourceConfiguredSchedule, courier_enums.DeliverySlotSourceProvider}, invalid: courier_enums.DeliverySlotSource("__invalid__")},
	})
}
