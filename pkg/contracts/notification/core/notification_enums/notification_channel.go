package notification_enums

// NotificationChannel identifies a provider-neutral notification channel.
type NotificationChannel string

const (
	NotificationChannelEmail       NotificationChannel = "email"
	NotificationChannelPush        NotificationChannel = "push"
	NotificationChannelSMS         NotificationChannel = "sms"
	NotificationChannelInApp       NotificationChannel = "in_app"
	NotificationChannelSocialMedia NotificationChannel = "social_media"
)

// IsValid reports whether c is a supported notification channel.
func (c NotificationChannel) IsValid() bool {
	switch c {
	case NotificationChannelEmail, NotificationChannelPush, NotificationChannelSMS,
		NotificationChannelInApp, NotificationChannelSocialMedia:
		return true
	default:
		return false
	}
}

// String returns the wire value for c.
func (c NotificationChannel) String() string { return string(c) }
