package campaign_enums

// CampaignPlatform identifies a client family targeted by a campaign.
type CampaignPlatform string

const (
	CampaignPlatformWeb    CampaignPlatform = "web"
	CampaignPlatformMobile CampaignPlatform = "mobile"
)

func (p CampaignPlatform) IsValid() bool {
	switch p {
	case CampaignPlatformWeb, CampaignPlatformMobile:
		return true
	}
	return false
}
func (p CampaignPlatform) String() string { return string(p) }
