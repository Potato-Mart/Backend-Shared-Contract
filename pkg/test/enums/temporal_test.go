package enums_test

import (
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/temporal/temporal_enums"
)

func TestTemporalEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "temporal.DayOfWeek", valid: []stringEnum{temporal_enums.DayOfWeekMonday, temporal_enums.DayOfWeekTuesday, temporal_enums.DayOfWeekWednesday, temporal_enums.DayOfWeekThursday, temporal_enums.DayOfWeekFriday, temporal_enums.DayOfWeekSaturday, temporal_enums.DayOfWeekSunday}, invalid: temporal_enums.DayOfWeek("__invalid__")},
	})
	for value, want := range map[temporal_enums.DayOfWeek]string{
		temporal_enums.DayOfWeekMonday:    "MONDAY",
		temporal_enums.DayOfWeekTuesday:   "TUESDAY",
		temporal_enums.DayOfWeekWednesday: "WEDNESDAY",
		temporal_enums.DayOfWeekThursday:  "THURSDAY",
		temporal_enums.DayOfWeekFriday:    "FRIDAY",
		temporal_enums.DayOfWeekSaturday:  "SATURDAY",
		temporal_enums.DayOfWeekSunday:    "SUNDAY",
	} {
		if got := value.String(); got != want {
			t.Errorf("DayOfWeek wire value = %q, want %q", got, want)
		}
	}
}
