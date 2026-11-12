package courier_enums

// DeliverySlotSource distinguishes configured backend windows from windows
// returned by a provider. None explicitly declares no selectable slot source.
// StaffOverride marks a delivery selection that staff replaced on an existing
// order; its window is the staff-entered value, not a published slot.
type DeliverySlotSource string

const (
	DeliverySlotSourceNone               DeliverySlotSource = "none"
	DeliverySlotSourceConfiguredSchedule DeliverySlotSource = "configured_schedule"
	DeliverySlotSourceProvider           DeliverySlotSource = "provider"
	DeliverySlotSourceStaffOverride      DeliverySlotSource = "staff_override"
)

func (s DeliverySlotSource) IsValid() bool {
	switch s {
	case DeliverySlotSourceNone, DeliverySlotSourceConfiguredSchedule, DeliverySlotSourceProvider, DeliverySlotSourceStaffOverride:
		return true
	default:
		return false
	}
}

func (s DeliverySlotSource) String() string { return string(s) }
