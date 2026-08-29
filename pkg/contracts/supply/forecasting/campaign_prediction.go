package forecasting

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/forecasting/marketing_enums"
)

// CampaignPrediction is a revisioned forecast and its evidence summary.
type CampaignPrediction struct {
	PredictionKey    string                                   `json:"prediction_key"`
	CampaignCode     string                                   `json:"campaign_code,omitempty"`
	Revision         int                                      `json:"revision"`
	Status           marketing_enums.CampaignPredictionStatus `json:"status"`
	AlgorithmVersion string                                   `json:"algorithm_version"`
	ResolvedSKUCodes []string                                 `json:"resolved_sku_codes,omitempty"`
	Products         []CampaignProductPrediction              `json:"products,omitempty"`
	Suppliers        []CampaignSupplierPrediction             `json:"suppliers,omitempty"`
	Warnings         []string                                 `json:"warnings,omitempty"`
	PredictedAt      time.Time                                `json:"predicted_at"`
}
