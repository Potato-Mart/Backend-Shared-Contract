package campaign

// CampaignCategoryTarget selects a collection or category for campaign copy.
type CampaignCategoryTarget struct {
	CollectionCode  string `json:"collection_code"`
	CategoryTagCode string `json:"category_tag_code"`
}
