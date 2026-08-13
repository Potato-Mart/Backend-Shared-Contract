package campaign_enums

// CampaignMessagingCategory classifies what a campaign announces so
// notification fan-out can map it onto the customer preference-centre topic
// groups. It carries no pricing or delivery logic.
type CampaignMessagingCategory string

const (
	CampaignMessagingCategoryPromotion    CampaignMessagingCategory = "promotion"
	CampaignMessagingCategoryAnnouncement CampaignMessagingCategory = "announcement"
	CampaignMessagingCategoryNewProduct   CampaignMessagingCategory = "new_product"
	CampaignMessagingCategoryPreorder     CampaignMessagingCategory = "preorder"
)

func (c CampaignMessagingCategory) IsValid() bool {
	switch c {
	case CampaignMessagingCategoryPromotion, CampaignMessagingCategoryAnnouncement,
		CampaignMessagingCategoryNewProduct, CampaignMessagingCategoryPreorder:
		return true
	default:
		return false
	}
}

func (c CampaignMessagingCategory) String() string { return string(c) }
