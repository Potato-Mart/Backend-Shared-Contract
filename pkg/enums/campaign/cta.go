package campaignenum

// CampaignCTADestinationType selects one allowlisted native/storefront route.
// The accompanying campaign.CTADestination carries the type-specific key;
// arbitrary URLs are not represented by this enum.
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
	case CampaignCTADestinationProduct, CampaignCTADestinationCollection,
		CampaignCTADestinationCategory, CampaignCTADestinationCart,
		CampaignCTADestinationPromotions:
		return true
	default:
		return false
	}
}

func (t CampaignCTADestinationType) String() string { return string(t) }
