package marketing

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v9/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v9/pkg/contracts/shared"
	"github.com/Potato-Mart/Backend-Shared-Contract/v9/pkg/enums"
)

// MarketingCampaign records a single EDM / SMS / LINE push broadcast.
// Actual delivery is handled by edge functions (Resend, Twilio, etc.);
// this contract is the shared data shape for admin UI and reporting.
type MarketingCampaign struct {
	ID             string                        `json:"id"`
	Name           string                        `json:"name"`
	Channel        enums.MarketingChannel        `json:"channel"`
	SegmentKey     string                        `json:"segment_key,omitempty"` // e.g. "churn_high" / "vip" / "all"
	Subject        string                        `json:"subject,omitempty"`
	Body           string                        `json:"body,omitempty"`
	RecipientCount int                           `json:"recipient_count"`
	Status         enums.MarketingCampaignStatus `json:"status"`
	SentAt         *time.Time                    `json:"sent_at,omitempty"`
	Metadata       common.Metadata               `json:"metadata,omitempty"`
	History        []shared.HistoryEntry         `json:"history,omitempty"`

	common.AuditFields `bson:",inline"`
}

// MarketingCampaignRecipient is the delivery record for a single recipient
// within a campaign. Status progresses as the message is sent and tracked.
type MarketingCampaignRecipient struct {
	ID                string                         `json:"id"`
	CampaignID        string                         `json:"campaign_id"`
	CustomerProfileID string                         `json:"customer_profile_id,omitempty"`
	Contacts          common.ContactChannels         `json:"contacts,omitempty"`
	CustomerName      string                         `json:"customer_name,omitempty"`
	Status            enums.MarketingRecipientStatus `json:"status"`
	SentAt            *time.Time                     `json:"sent_at,omitempty"`
	Error             string                         `json:"error,omitempty"`
	Metadata          common.Metadata                `json:"metadata,omitempty"`
	History           []shared.HistoryEntry          `json:"history,omitempty"`
	CreatedAt         time.Time                      `json:"created_at"`
}
