package campaign_enums

// CampaignStatus is the lifecycle state of a campaign.
type CampaignStatus string

const (
	CampaignStatusDraft     CampaignStatus = "draft"
	CampaignStatusPublished CampaignStatus = "published"
	CampaignStatusPaused    CampaignStatus = "paused"
	CampaignStatusArchived  CampaignStatus = "archived"
)

func (s CampaignStatus) IsValid() bool {
	switch s {
	case CampaignStatusDraft, CampaignStatusPublished, CampaignStatusPaused, CampaignStatusArchived:
		return true
	}
	return false
}
func (s CampaignStatus) String() string { return string(s) }
