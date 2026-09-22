package courier_enums

// DeliverySlotSource distinguishes configured backend windows from windows
// returned by a provider. None explicitly declares no selectable slot source.
type DeliverySlotSource string

const (
	DeliverySlotSourceNone               DeliverySlotSource = "none"
	DeliverySlotSourceConfiguredSchedule DeliverySlotSource = "configured_schedule"
	DeliverySlotSourceProvider           DeliverySlotSource = "provider"
)

func (s DeliverySlotSource) IsValid() bool {
	switch s {
	case DeliverySlotSourceNone, DeliverySlotSourceConfiguredSchedule, DeliverySlotSourceProvider:
		return true
	default:
		return false
	}
}

func (s DeliverySlotSource) String() string { return string(s) }
