package notification_enums

// NotificationStatus is the aggregate lifecycle of a protected Notification.
type NotificationStatus string

const (
	NotificationStatusPending            NotificationStatus = "pending"
	NotificationStatusScheduled          NotificationStatus = "scheduled"
	NotificationStatusDispatching        NotificationStatus = "dispatching"
	NotificationStatusPartiallyDelivered NotificationStatus = "partially_delivered"
	NotificationStatusDelivered          NotificationStatus = "delivered"
	NotificationStatusFailed             NotificationStatus = "failed"
	NotificationStatusCancelled          NotificationStatus = "cancelled"
	NotificationStatusExpired            NotificationStatus = "expired"
)

// IsValid reports whether s is a supported notification lifecycle state.
func (s NotificationStatus) IsValid() bool {
	switch s {
	case NotificationStatusPending, NotificationStatusScheduled, NotificationStatusDispatching,
		NotificationStatusPartiallyDelivered, NotificationStatusDelivered, NotificationStatusFailed,
		NotificationStatusCancelled, NotificationStatusExpired:
		return true
	default:
		return false
	}
}

// String returns the wire value for s.
func (s NotificationStatus) String() string { return string(s) }
