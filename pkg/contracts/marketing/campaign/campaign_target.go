package campaign

// CampaignTarget identifies product and category content presented by a
// campaign. Product identity is always SKUCode.
type CampaignTarget struct {
	SKUCodes   []string                 `json:"sku_codes,omitempty"`
	Categories []CampaignCategoryTarget `json:"categories,omitempty"`
}
