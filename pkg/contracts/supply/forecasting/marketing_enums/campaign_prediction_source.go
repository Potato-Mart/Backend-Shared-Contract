package marketing_enums

// CampaignPredictionSource identifies a comparable-data source.
type CampaignPredictionSource string

const (
	CampaignPredictionSourceSameSeries        CampaignPredictionSource = "same_series"
	CampaignPredictionSourceSimilarEvent      CampaignPredictionSource = "similar_event"
	CampaignPredictionSourceLast14DaysDoubled CampaignPredictionSource = "last_14_days_doubled"
)

func (s CampaignPredictionSource) IsValid() bool {
	switch s {
	case CampaignPredictionSourceSameSeries, CampaignPredictionSourceSimilarEvent, CampaignPredictionSourceLast14DaysDoubled:
		return true
	}
	return false
}
func (s CampaignPredictionSource) String() string { return string(s) }
