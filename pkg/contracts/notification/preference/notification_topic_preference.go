package preference

// NotificationTopicPreference records a choice for an open backend topic within
// its containing channel. Missing choices defer to backend policy. Social media
// requires an explicit destination allow-list and separate destination consent;
// an absent or empty allow-list never means all destinations. Mandatory topics
// remain governed by NotificationTopic policy, regardless of Enabled.
type NotificationTopicPreference struct {
	TopicCode        string   `json:"topic_code"`
	Enabled          bool     `json:"enabled"`
	DestinationCodes []string `json:"destination_codes,omitempty"`
}
