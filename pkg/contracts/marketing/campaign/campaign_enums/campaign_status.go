package campaign_enums

// CampaignStatus is the lifecycle state of a campaign.
type CampaignStatus string

const (
	CampaignStatusDraft     CampaignStatus = "draft"
	CampaignStatusScheduled CampaignStatus = "scheduled"
	CampaignStatusActive    CampaignStatus = "active"
	CampaignStatusCompleted CampaignStatus = "completed"
	CampaignStatusArchived  CampaignStatus = "archived"
)

func (s CampaignStatus) IsValid() bool {
	switch s {
	case CampaignStatusDraft, CampaignStatusScheduled, CampaignStatusActive, CampaignStatusCompleted, CampaignStatusArchived:
		return true
	}
	return false
}
func (s CampaignStatus) String() string { return string(s) }
