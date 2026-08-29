package sms

// SMSNotification is the provider-neutral authored content for one SMS
// notification delivery.
type SMSNotification struct {
	Body string `json:"body"`
}
