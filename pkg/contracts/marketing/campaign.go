package marketing

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/audit"
	geography "github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/localization"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/marketing/marketing_enums"
)

// Campaign is the public, campaign-owned composition of storefront copy and
// references to the benefits it presents. Pricing remains the authority for
// coupon and promotion mechanics.
type Campaign struct {
	CampaignCode     string                       `json:"campaign_code"`
	Title            []localization.LocalizedName `json:"title"`
	Cover            *security.ObjectMedia        `json:"cover,omitempty"`
	CampaignDetail   CampaignDetail               `json:"campaign_detail"`
	CampaignPosition CampaignPosition             `json:"campaign_position"`
	CampaignStatus   CampaignStatus               `json:"campaign_status"`
	Audience         Audience                     `json:"audience"`

	audit.AuditFields
}

// CampaignDetail contains public campaign content and code-only benefit
// references. It does not duplicate coupon or promotion pricing rules.
type CampaignDetail struct {
	Message          []localization.LocalizedText `json:"message,omitempty"`
	CampaignType     marketing_enums.CampaignType `json:"campaign_type"`
	CouponDetails    []BenefitRef                 `json:"coupon_details,omitempty"`
	PromotionDetails []BenefitRef                 `json:"promotion_details,omitempty"`
}

// CampaignPosition defines where and in which commercial geography a campaign
// can be presented.
type CampaignPosition struct {
	Placement        marketing_enums.CampaignPlacement `json:"placement"`
	GeographicScope  geography.GeographicScope         `json:"geographic_scope"`
	ScheduleTimezone string                            `json:"schedule_timezone"`
}

// CampaignStatus combines lifecycle state with the campaign's display window.
type CampaignStatus struct {
	Status      marketing_enums.CampaignStatus `json:"status"`
	Dismissible bool                           `json:"dismissible"`
	StartsAt    *time.Time                     `json:"starts_at,omitempty"`
	EndsAt      *time.Time                     `json:"ends_at,omitempty"`
}

// Audience selects the customer and platform audience for a campaign.
type Audience struct {
	CustomerType marketing_enums.CampaignCustomerType `json:"customer_type"`
	Platform     marketing_enums.CampaignPlatform     `json:"platform"`
}
