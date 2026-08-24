package campaign_enums

// CampaignSeverity drives the visual tone of a campaign.
type CampaignSeverity string

const (
	CampaignSeverityInfo     CampaignSeverity = "info"
	CampaignSeveritySuccess  CampaignSeverity = "success"
	CampaignSeverityWarning  CampaignSeverity = "warning"
	CampaignSeverityCritical CampaignSeverity = "critical"
)

func (s CampaignSeverity) IsValid() bool {
	switch s {
	case CampaignSeverityInfo, CampaignSeveritySuccess, CampaignSeverityWarning, CampaignSeverityCritical:
		return true
	}
	return false
}
func (s CampaignSeverity) String() string { return string(s) }
