package audience

import "github.com/Potato-Mart/Backend-Shared-Contract/v36/pkg/contracts/marketing/campaign/campaign_enums"

// Audience narrows a campaign or Pricing-owned offer by customer type and
// client platform. Omitted dimensions are unrestricted. Platform is distinct
// from an order channel; owning services enforce these dimensions consistently.
type Audience struct {
	CustomerType campaign_enums.CampaignCustomerType `json:"customer_type,omitempty"`
	Platform     campaign_enums.CampaignPlatform     `json:"platform,omitempty"`
}
