package account

// NotificationTopicGroupChannels is one preference-centre topic group's
// channel switches. Transactional email (the payment, order, and receipt
// groups) is normalized true server-side and rendered as a locked toggle.
type NotificationTopicGroupChannels struct {
	Email bool `json:"email"`
	Push  bool `json:"push"`
	SMS   bool `json:"sms"`
}

// UserNotificationTopicGroupPreferences is the customer preference-centre
// matrix: five fixed topic groups x three channels.
type UserNotificationTopicGroupPreferences struct {
	Payment        NotificationTopicGroupChannels `json:"payment"`
	Order          NotificationTopicGroupChannels `json:"order"`
	Receipt        NotificationTopicGroupChannels `json:"receipt"`
	ProductUpdates NotificationTopicGroupChannels `json:"product_updates"`
	Promotions     NotificationTopicGroupChannels `json:"promotions"`
}
