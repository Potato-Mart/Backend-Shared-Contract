package campaign_enums

// CampaignCTADestinationType selects a supported storefront destination.
type CampaignCTADestinationType string

const (
	CampaignCTADestinationProduct    CampaignCTADestinationType = "product"
	CampaignCTADestinationCollection CampaignCTADestinationType = "collection"
	CampaignCTADestinationCategory   CampaignCTADestinationType = "category"
	CampaignCTADestinationCart       CampaignCTADestinationType = "cart"
	CampaignCTADestinationPromotions CampaignCTADestinationType = "promotions"
)

func (t CampaignCTADestinationType) IsValid() bool {
	switch t {
	case CampaignCTADestinationProduct, CampaignCTADestinationCollection, CampaignCTADestinationCategory, CampaignCTADestinationCart, CampaignCTADestinationPromotions:
		return true
	}
	return false
}
func (t CampaignCTADestinationType) String() string { return string(t) }
