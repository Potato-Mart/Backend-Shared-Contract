package enums_test

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v44/pkg/contracts/supply/courier/courier_enums"
	"testing"
)

func TestCourierEnumsValidateKnownValues(t *testing.T) {
	if courier_enums.PostalCodeMatchMode("").IsValid() || courier_enums.DeliverySlotSource("").IsValid() {
		t.Fatal("empty configuration values must not acquire implicit defaults")
	}
	assertStringEnums(t, []enumCase{
		{name: "PostalCodeMatchMode", valid: []stringEnum{courier_enums.PostalCodeMatchModeIncludeOnly, courier_enums.PostalCodeMatchModeAllExcept}, invalid: courier_enums.PostalCodeMatchMode("__invalid__")},
		{name: "DeliverySlotSource", valid: []stringEnum{courier_enums.DeliverySlotSourceNone, courier_enums.DeliverySlotSourceConfiguredSchedule, courier_enums.DeliverySlotSourceProvider}, invalid: courier_enums.DeliverySlotSource("__invalid__")},
	})
}
