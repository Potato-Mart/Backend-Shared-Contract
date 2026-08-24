package message

import "time"

// MarketingMessageSendSummary is aggregate reporting for an authored send.
// Recipient-level records remain part of Notification.
type MarketingMessageSendSummary struct {
	RecipientCount int        `json:"recipient_count"`
	SentCount      int        `json:"sent_count,omitempty"`
	DeliveredCount int        `json:"delivered_count,omitempty"`
	FailedCount    int        `json:"failed_count,omitempty"`
	SentAt         *time.Time `json:"sent_at,omitempty"`
}
