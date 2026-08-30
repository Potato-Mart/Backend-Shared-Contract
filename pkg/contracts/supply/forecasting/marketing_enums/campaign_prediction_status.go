package marketing_enums

// CampaignPredictionStatus reports whether prediction data is ready for use.
type CampaignPredictionStatus string

const (
	CampaignPredictionStatusNotApplicable CampaignPredictionStatus = "not_applicable"
	CampaignPredictionStatusReady         CampaignPredictionStatus = "ready"
	CampaignPredictionStatusWarning       CampaignPredictionStatus = "warning"
)

func (s CampaignPredictionStatus) IsValid() bool {
	return s == CampaignPredictionStatusNotApplicable || s == CampaignPredictionStatusReady || s == CampaignPredictionStatusWarning
}
func (s CampaignPredictionStatus) String() string { return string(s) }
