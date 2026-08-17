package marketing

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/audit"
	geography "github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/localization"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/marketing/marketing_enums"
)

// MarketingMessage is the public aggregate for an authored marketing send.
// It carries campaign-name snapshot content only and never recipient, contact,
// delivery-provider, or delivery-outcome data.
type MarketingMessage struct {
	Code             string                                 `json:"code"`
	CampaignCode     string                                 `json:"campaign_code"`
	CampaignName     []localization.LocalizedName           `json:"campaign_name"`
	Channel          marketing_enums.MarketingChannel       `json:"channel"`
	Subject          []localization.LocalizedText           `json:"subject,omitempty"`
	Message          []localization.LocalizedText           `json:"message,omitempty"`
	Images           []security.ObjectMedia                 `json:"images,omitempty"`
	Status           marketing_enums.MarketingMessageStatus `json:"status"`
	GeographicScope  geography.GeographicScope              `json:"geographic_scope"`
	ScheduleTimezone string                                 `json:"schedule_timezone"`
	ScheduledSendAt  *time.Time                             `json:"scheduled_send_at,omitempty"`

	audit.AuditFields
}
