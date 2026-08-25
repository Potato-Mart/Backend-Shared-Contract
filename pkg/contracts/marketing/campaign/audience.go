package campaign

import "github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/marketing/campaign/campaign_enums"

// Audience narrows a campaign by customer type and client platform.
type Audience struct {
	CustomerType campaign_enums.CampaignCustomerType `json:"customer_type,omitempty"`
	Platform     campaign_enums.CampaignPlatform     `json:"platform,omitempty"`
}
