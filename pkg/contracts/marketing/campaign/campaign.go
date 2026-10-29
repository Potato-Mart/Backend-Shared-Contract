package campaign

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v50/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v50/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v50/pkg/contracts/common/localization"
	"github.com/Potato-Mart/Backend-Shared-Contract/v50/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v50/pkg/contracts/marketing/campaign/campaign_enums"
	"time"
)

// Campaign permits incomplete authored drafts. Customers owns publication and
// effective-state evaluation; Pricing owns commercial decisions.
type Campaign struct {
	ID                 string                                `json:"id,omitempty"`
	CampaignCode       string                                `json:"campaign_code"`
	MarketCode         string                                `json:"market_code,omitempty"`
	CountryCode        geography.CountryCode                 `json:"country_code,omitempty"`
	CurrencyCode       string                                `json:"currency_code,omitempty"`
	Title              []localization.LocalizedName          `json:"title,omitempty"`
	Description        []localization.LocalizedText          `json:"description,omitempty"`
	Media              *security.ObjectMedia                 `json:"media,omitempty"`
	Placements         []Placement                           `json:"placements"`
	Offers             []OfferLink                           `json:"offers"`
	Published          bool                                  `json:"published"`
	StartsAt           *time.Time                            `json:"starts_at,omitempty"`
	EndsAt             *time.Time                            `json:"ends_at,omitempty"`
	ScheduleTimezone   string                                `json:"schedule_timezone,omitempty"`
	Audience           []campaign_enums.CampaignCustomerType `json:"audience"`
	GeographicScope    geography.GeographicScope             `json:"geographic_scope,omitempty"`
	Targets            CampaignTarget                        `json:"targets,omitempty"`
	Status             campaign_enums.CampaignStatus         `json:"publication_status"`
	EffectiveStatus    string                                `json:"effective_status"`
	Revision           int64                                 `json:"revision"`
	ActivationRevision int64                                 `json:"publication_revision"`
	ActivatedAt        *time.Time                            `json:"published_at,omitempty"`
	DeactivatedAt      *time.Time                            `json:"paused_at,omitempty"`
	ArchivedAt         *time.Time                            `json:"archived_at,omitempty"`
	audit.AuditFields
}
