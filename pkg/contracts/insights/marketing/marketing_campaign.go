package marketing

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v24/pkg/contracts/common/audit"
	geography "github.com/Potato-Mart/Backend-Shared-Contract/v24/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v24/pkg/contracts/common/metadata"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v24/pkg/contracts/common/security"

	"github.com/Potato-Mart/Backend-Shared-Contract/v24/pkg/contracts/common/party"
	"github.com/Potato-Mart/Backend-Shared-Contract/v24/pkg/contracts/insights/marketing/marketing_enums"
)

// MarketingCampaign records a single EDM / SMS / LINE push broadcast.
// Actual delivery is handled by edge functions (Resend, Twilio, etc.);
// this contract is the shared data shape for admin UI and reporting.
type MarketingCampaign struct {
	ID               string                                  `json:"id"`
	Name             string                                  `json:"name"`
	Channel          marketing_enums.MarketingChannel        `json:"channel"`
	SegmentKey       string                                  `json:"segment_key,omitempty"` // e.g. "churn_high" / "vip" / "all"
	Subject          string                                  `json:"subject,omitempty"`
	Body             string                                  `json:"body,omitempty"`
	RecipientCount   int                                     `json:"recipient_count"`
	Status           marketing_enums.MarketingCampaignStatus `json:"status"`
	GeographicScope  geography.GeographicScope               `json:"geographic_scope"`
	ScheduleTimezone string                                  `json:"schedule_timezone"`
	SentAt           *time.Time                              `json:"sent_at,omitempty"`
	Metadata         metadata.Metadata                       `json:"metadata,omitempty"`
	History          []security.HistoryEntry                 `json:"history,omitempty"`

	audit.AuditFields
}

// MarketingCampaignRecipient is the delivery record for a single recipient
// within a campaign. Status progresses as the message is sent and tracked.
type MarketingCampaignRecipient struct {
	ID             string                                   `json:"id"`
	CampaignID     string                                   `json:"campaign_id"`
	CustomerNumber string                                   `json:"customer_number,omitempty"`
	Contacts       party.ContactChannels                    `json:"contacts,omitempty"`
	CustomerName   string                                   `json:"customer_name,omitempty"`
	Status         marketing_enums.MarketingRecipientStatus `json:"status"`
	SentAt         *time.Time                               `json:"sent_at,omitempty"`
	Error          string                                   `json:"error,omitempty"`
	Metadata       metadata.Metadata                        `json:"metadata,omitempty"`
	History        []security.HistoryEntry                  `json:"history,omitempty"`
	CreatedAt      time.Time                                `json:"created_at"`
}
