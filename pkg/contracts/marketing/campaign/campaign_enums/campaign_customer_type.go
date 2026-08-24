package campaign_enums

// CampaignCustomerType identifies the customer population for a campaign.
type CampaignCustomerType string

const (
	CampaignCustomerTypeGuest     CampaignCustomerType = "guest"
	CampaignCustomerTypeRetail    CampaignCustomerType = "retail"
	CampaignCustomerTypeWholesale CampaignCustomerType = "wholesale"
)

func (c CampaignCustomerType) IsValid() bool {
	switch c {
	case CampaignCustomerTypeGuest, CampaignCustomerTypeRetail, CampaignCustomerTypeWholesale:
		return true
	}
	return false
}
func (c CampaignCustomerType) String() string { return string(c) }
