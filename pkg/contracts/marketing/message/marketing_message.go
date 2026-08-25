package message

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/audit"
	geography "github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/localization"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/marketing/message/message_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/notifications/notification_enums"
)

// MarketingMessage is an authored, aggregated outbound marketing send. It
// deliberately contains no recipients, contact details, provider data, or
// delivery outcomes; Notification owns those operational records.
type MarketingMessage struct {
	Code                  string                                 `json:"code"`
	CampaignCode          string                                 `json:"campaign_code,omitempty"`
	CampaignName          []localization.LocalizedName           `json:"campaign_name,omitempty"`
	Channel               notification_enums.NotificationChannel `json:"channel"`
	SegmentCode           string                                 `json:"segment_code,omitempty"`
	NotificationTopicCode string                                 `json:"notification_topic_code,omitempty"`
	Subject               []localization.LocalizedText           `json:"subject,omitempty"`
	Body                  []localization.LocalizedText           `json:"body,omitempty"`
	Images                []security.ObjectMedia                 `json:"images,omitempty"`
	Status                message_enums.MarketingMessageStatus   `json:"status"`
	SendSummary           *MarketingMessageSendSummary           `json:"send_summary,omitempty"`
	GeographicScope       geography.GeographicScope              `json:"geographic_scope"`
	ScheduleTimezone      string                                 `json:"schedule_timezone"`
	ScheduledSendAt       *time.Time                             `json:"scheduled_send_at,omitempty"`
	MarketCode            string                                 `json:"market_code,omitempty"`
	CountryCode           geography.CountryCode                  `json:"country_code,omitempty"`

	audit.AuditFields
}
