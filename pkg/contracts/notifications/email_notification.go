package notifications

// EmailNotification is the provider-neutral authored content for one email
// notification delivery.
type EmailNotification struct {
	Subject     string `json:"subject"`
	Body        string `json:"body"`
	PreviewText string `json:"preview_text,omitempty"`
}
