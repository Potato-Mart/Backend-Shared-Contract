package notification_enums

// InAppNotificationStatus is the customer-visible state of an in-app item.
type InAppNotificationStatus string

const (
	InAppNotificationStatusUnread    InAppNotificationStatus = "unread"
	InAppNotificationStatusRead      InAppNotificationStatus = "read"
	InAppNotificationStatusDismissed InAppNotificationStatus = "dismissed"
)

// IsValid reports whether s is a supported in-app notification state.
func (s InAppNotificationStatus) IsValid() bool {
	switch s {
	case InAppNotificationStatusUnread, InAppNotificationStatusRead, InAppNotificationStatusDismissed:
		return true
	default:
		return false
	}
}

// String returns the wire value for s.
func (s InAppNotificationStatus) String() string { return string(s) }
