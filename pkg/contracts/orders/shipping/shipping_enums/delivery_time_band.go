package shipping_enums

// DeliveryTimeBand classifies a delivery window by local time of day,
// evaluated in the slot's own IANA timezone. A window is late_night when it
// overlaps local 21:00–24:00, otherwise early_morning when it overlaps local
// 00:00–07:00, otherwise standard. Overlap is strict: a window that ends
// exactly at 21:00 or starts exactly at 07:00 is standard. For example
// 19:00–22:00 and 23:00–01:00 are late_night, 06:00–09:00 is early_morning,
// and 18:00–21:00 and 07:00–10:00 are standard.
type DeliveryTimeBand string

const (
	// DeliveryTimeBandStandard is a window wholly within local 07:00–21:00.
	DeliveryTimeBandStandard DeliveryTimeBand = "standard"
	// DeliveryTimeBandLateNight is a window that overlaps local 21:00–24:00.
	DeliveryTimeBandLateNight DeliveryTimeBand = "late_night"
	// DeliveryTimeBandEarlyMorning is a window that overlaps local
	// 00:00–07:00 and not 21:00–24:00.
	DeliveryTimeBandEarlyMorning DeliveryTimeBand = "early_morning"
)

// IsValid reports whether b is a known DeliveryTimeBand.
func (b DeliveryTimeBand) IsValid() bool {
	switch b {
	case DeliveryTimeBandStandard, DeliveryTimeBandLateNight, DeliveryTimeBandEarlyMorning:
		return true
	}
	return false
}

func (b DeliveryTimeBand) String() string { return string(b) }
