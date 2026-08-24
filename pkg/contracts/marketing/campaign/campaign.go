package campaign

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/audit"
	geography "github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/localization"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/marketing/campaign/campaign_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/pricing/benefit"
)

// Campaign is scheduled, targeted storefront content. BenefitRefs are
// presentation references only; Pricing owns eligibility and value.
type Campaign struct {
	ID                    string                           `json:"id,omitempty"`
	CampaignCode          string                           `json:"campaign_code"`
	MarketCode            string                           `json:"market_code"`
	CountryCode           geography.CountryCode            `json:"country_code,omitempty"`
	SeriesCode            string                           `json:"series_code,omitempty"`
	Title                 []localization.LocalizedName     `json:"title"`
	Message               []localization.LocalizedText     `json:"message,omitempty"`
	CTAText               []localization.LocalizedText     `json:"cta_text,omitempty"`
	CTAHref               string                           `json:"cta_href,omitempty"`
	CTA                   *CTADestination                  `json:"cta,omitempty"`
	Media                 *security.ObjectMedia            `json:"media,omitempty"`
	BackgroundToken       string                           `json:"background_token,omitempty"`
	Placement             campaign_enums.CampaignPlacement `json:"placement"`
	Severity              campaign_enums.CampaignSeverity  `json:"severity"`
	NotificationTopicCode string                           `json:"notification_topic_code,omitempty"`
	BenefitRefs           []benefit.BenefitRef             `json:"benefit_refs,omitempty"`
	Priority              int                              `json:"priority"`
	IsActive              bool                             `json:"is_active"`
	Dismissible           bool                             `json:"dismissible"`
	DismissKey            string                           `json:"dismiss_key,omitempty"`
	StartsAt              *time.Time                       `json:"starts_at,omitempty"`
	EndsAt                *time.Time                       `json:"ends_at,omitempty"`
	ScheduleTimezone      string                           `json:"schedule_timezone"`
	Audience              *Audience                        `json:"audience,omitempty"`
	GeographicScope       geography.GeographicScope        `json:"geographic_scope"`
	Targets               CampaignTarget                   `json:"targets,omitempty"`
	Status                campaign_enums.CampaignStatus    `json:"status"`
	Revision              int64                            `json:"revision"`
	ActivationRevision    int64                            `json:"activation_revision"`
	ActivatedAt           *time.Time                       `json:"activated_at,omitempty"`
	DeactivatedAt         *time.Time                       `json:"deactivated_at,omitempty"`
	ArchivedAt            *time.Time                       `json:"archived_at,omitempty"`

	audit.AuditFields
}
