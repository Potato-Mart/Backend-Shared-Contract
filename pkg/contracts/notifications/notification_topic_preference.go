package notifications

// NotificationTopicPreference records the chosen channels for one open,
// backend-created topic code. A missing topic or channel preference defers to
// the topic's backend policy default.
type NotificationTopicPreference struct {
	TopicCode string                          `json:"topic_code"`
	Channels  []NotificationChannelPreference `json:"channels"`
}
