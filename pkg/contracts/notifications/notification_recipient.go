package notifications

// NotificationRecipient is the protected ownership identity for a
// Notification. At least one identifier is required and services must ensure
// supplied identifiers resolve to the same person.
type NotificationRecipient struct {
	UserID         string `json:"user_id,omitempty"`
	AccountID      string `json:"account_id,omitempty"`
	CustomerNumber string `json:"customer_number,omitempty"`
}
