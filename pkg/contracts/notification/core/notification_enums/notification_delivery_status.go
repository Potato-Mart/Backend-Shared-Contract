package notification_enums

// NotificationDeliveryStatus is the lifecycle of one notification endpoint.
type NotificationDeliveryStatus string

const (
	NotificationDeliveryStatusPending     NotificationDeliveryStatus = "pending"
	NotificationDeliveryStatusDispatching NotificationDeliveryStatus = "dispatching"
	NotificationDeliveryStatusDelivered   NotificationDeliveryStatus = "delivered"
	NotificationDeliveryStatusFailed      NotificationDeliveryStatus = "failed"
	NotificationDeliveryStatusCancelled   NotificationDeliveryStatus = "cancelled"
)

// IsValid reports whether s is a supported notification delivery state.
func (s NotificationDeliveryStatus) IsValid() bool {
	switch s {
	case NotificationDeliveryStatusPending, NotificationDeliveryStatusDispatching,
		NotificationDeliveryStatusDelivered, NotificationDeliveryStatusFailed,
		NotificationDeliveryStatusCancelled:
		return true
	default:
		return false
	}
}

// String returns the wire value for s.
func (s NotificationDeliveryStatus) String() string { return string(s) }
