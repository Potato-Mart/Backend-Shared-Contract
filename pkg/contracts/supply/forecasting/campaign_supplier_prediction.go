package forecasting

// CampaignSupplierPrediction groups product forecasts for one supplier.
type CampaignSupplierPrediction struct {
	SupplierCode   string                      `json:"supplier_code"`
	Products       []CampaignProductPrediction `json:"products"`
	TotalBaseUnits int64                       `json:"total_base_units"`
	Warnings       []string                    `json:"warnings,omitempty"`
}
