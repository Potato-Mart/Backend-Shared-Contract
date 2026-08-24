package message_enums

// MarketingMessageStatus is the aggregate lifecycle of an authored send.
type MarketingMessageStatus string

const (
	MarketingMessageStatusDraft     MarketingMessageStatus = "draft"
	MarketingMessageStatusScheduled MarketingMessageStatus = "scheduled"
	MarketingMessageStatusSending   MarketingMessageStatus = "sending"
	MarketingMessageStatusSent      MarketingMessageStatus = "sent"
	MarketingMessageStatusPartial   MarketingMessageStatus = "partial"
	MarketingMessageStatusFailed    MarketingMessageStatus = "failed"
	MarketingMessageStatusCancelled MarketingMessageStatus = "cancelled"
)

func (s MarketingMessageStatus) IsValid() bool {
	switch s {
	case MarketingMessageStatusDraft, MarketingMessageStatusScheduled, MarketingMessageStatusSending, MarketingMessageStatusSent, MarketingMessageStatusPartial, MarketingMessageStatusFailed, MarketingMessageStatusCancelled:
		return true
	}
	return false
}
func (s MarketingMessageStatus) String() string { return string(s) }
