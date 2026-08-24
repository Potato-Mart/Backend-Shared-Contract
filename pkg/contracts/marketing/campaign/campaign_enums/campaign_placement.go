package campaign_enums

// CampaignPlacement identifies where a campaign renders in a storefront.
type CampaignPlacement string

const (
	CampaignPlacementTopBanner      CampaignPlacement = "top_banner"
	CampaignPlacementHomeHero       CampaignPlacement = "home_hero"
	CampaignPlacementModal          CampaignPlacement = "modal"
	CampaignPlacementCheckoutNotice CampaignPlacement = "checkout_notice"
	CampaignPlacementProductNotice  CampaignPlacement = "product_notice"
)

func (p CampaignPlacement) IsValid() bool {
	switch p {
	case CampaignPlacementTopBanner, CampaignPlacementHomeHero, CampaignPlacementModal, CampaignPlacementCheckoutNotice, CampaignPlacementProductNotice:
		return true
	}
	return false
}
func (p CampaignPlacement) String() string { return string(p) }
