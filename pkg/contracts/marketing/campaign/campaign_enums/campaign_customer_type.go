package campaign_enums

// CampaignCustomerType identifies the customer population for a campaign.
type CampaignCustomerType string

const (
	CampaignCustomerTypeRetail    CampaignCustomerType = "retail"
	CampaignCustomerTypeWholesale CampaignCustomerType = "wholesale"
)

func (c CampaignCustomerType) IsValid() bool {
	switch c {
	case CampaignCustomerTypeRetail, CampaignCustomerTypeWholesale:
		return true
	}
	return false
}
func (c CampaignCustomerType) String() string { return string(c) }
