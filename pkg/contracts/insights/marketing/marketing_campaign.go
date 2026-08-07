package marketing

import (
	geography "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/common/geography"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/common/security"
	"time"

	common "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/common/shared"
)

// MarketingCampaign records a single EDM / SMS / LINE push broadcast.
// Actual delivery is handled by edge functions (Resend, Twilio, etc.);
// this contract is the shared data shape for admin UI and reporting.
type MarketingCampaign struct {
	ID               string                    `json:"id"`
	Name             string                    `json:"name"`
	Channel          MarketingChannel          `json:"channel"`
	SegmentKey       string                    `json:"segment_key,omitempty"` // e.g. "churn_high" / "vip" / "all"
	Subject          string                    `json:"subject,omitempty"`
	Body             string                    `json:"body,omitempty"`
	RecipientCount   int                       `json:"recipient_count"`
	Status           MarketingCampaignStatus   `json:"status"`
	GeographicScope  geography.GeographicScope `json:"geographic_scope"`
	ScheduleTimezone string                    `json:"schedule_timezone"`
	SentAt           *time.Time                `json:"sent_at,omitempty"`
	Metadata         common.Metadata           `json:"metadata,omitempty"`
	History          []security.HistoryEntry   `json:"history,omitempty"`

	common.AuditFields
}

// MarketingCampaignRecipient is the delivery record for a single recipient
// within a campaign. Status progresses as the message is sent and tracked.
type MarketingCampaignRecipient struct {
	ID             string                   `json:"id"`
	CampaignID     string                   `json:"campaign_id"`
	CustomerNumber string                   `json:"customer_number,omitempty"`
	Contacts       common.ContactChannels   `json:"contacts,omitempty"`
	CustomerName   string                   `json:"customer_name,omitempty"`
	Status         MarketingRecipientStatus `json:"status"`
	SentAt         *time.Time               `json:"sent_at,omitempty"`
	Error          string                   `json:"error,omitempty"`
	Metadata       common.Metadata          `json:"metadata,omitempty"`
	History        []security.HistoryEntry  `json:"history,omitempty"`
	CreatedAt      time.Time                `json:"created_at"`
}
