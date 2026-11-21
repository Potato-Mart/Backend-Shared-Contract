package temporal_enums

// DayOfWeek names a calendar weekday for recurring weekly schedules.
type DayOfWeek string

const (
	DayOfWeekMonday    DayOfWeek = "MONDAY"
	DayOfWeekTuesday   DayOfWeek = "TUESDAY"
	DayOfWeekWednesday DayOfWeek = "WEDNESDAY"
	DayOfWeekThursday  DayOfWeek = "THURSDAY"
	DayOfWeekFriday    DayOfWeek = "FRIDAY"
	DayOfWeekSaturday  DayOfWeek = "SATURDAY"
	DayOfWeekSunday    DayOfWeek = "SUNDAY"
)

func (d DayOfWeek) IsValid() bool {
	switch d {
	case DayOfWeekMonday, DayOfWeekTuesday, DayOfWeekWednesday, DayOfWeekThursday,
		DayOfWeekFriday, DayOfWeekSaturday, DayOfWeekSunday:
		return true
	}
	return false
}

func (d DayOfWeek) String() string { return string(d) }
