package campaign_enums

// CampaignPlacement identifies where a campaign renders in a storefront or on
// a POS customer-facing display.
type CampaignPlacement string

const (
	// CampaignPlacementAnnouncementBar renders as a thin top header strip.
	CampaignPlacementAnnouncementBar CampaignPlacement = "announcement_bar"
	// CampaignPlacementHomeBanner renders as the web home hero or the mobile
	// home banner above the collection rail.
	CampaignPlacementHomeBanner CampaignPlacement = "home_banner"
	// CampaignPlacementAccountBanner renders on mobile Account below the
	// order-status strip.
	CampaignPlacementAccountBanner CampaignPlacement = "account_banner"
	// CampaignPlacementHomeModal renders as the initial home-page popup.
	CampaignPlacementHomeModal CampaignPlacement = "home_modal"
	// CampaignPlacementPOSCustomerDisplay renders on the customer-facing screen
	// of a POS till.
	CampaignPlacementPOSCustomerDisplay CampaignPlacement = "pos_customer_display"
)

func (p CampaignPlacement) IsValid() bool {
	switch p {
	case CampaignPlacementAnnouncementBar, CampaignPlacementHomeBanner, CampaignPlacementAccountBanner, CampaignPlacementHomeModal, CampaignPlacementPOSCustomerDisplay:
		return true
	}
	return false
}
func (p CampaignPlacement) String() string { return string(p) }
