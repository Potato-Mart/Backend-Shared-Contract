package campaign

import "github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/marketing/campaign/campaign_enums"

// CTADestination is a validated, allowlisted campaign call-to-action route.
type CTADestination struct {
	Type            campaign_enums.CampaignCTADestinationType `json:"type"`
	SKUCode         string                                    `json:"sku_code,omitempty"`
	CollectionCode  string                                    `json:"collection_code,omitempty"`
	CategoryTagCode string                                    `json:"category_tag_code,omitempty"`
	PromotionCode   string                                    `json:"promotion_code,omitempty"`
}
