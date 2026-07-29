package campaign

import (
	"time"

	campaignenum "github.com/Potato-Mart/Backend-Shared-Contract/v21/pkg/enums/campaign"
)

type CampaignComparableEvent struct {
	CampaignKey             string                                `json:"campaign_key"`
	SeriesKey               string                                `json:"series_key,omitempty"`
	MatchSource             campaignenum.CampaignPredictionSource `json:"match_source"`
	ResolvedProductSKUCodes []string                              `json:"resolved_product_sku_codes"`
	StartsAt                time.Time                             `json:"starts_at"`
	EndsAt                  time.Time                             `json:"ends_at"`
	Audience                *Audience                             `json:"audience,omitempty"`
	Placement               campaignenum.CampaignPlacement        `json:"placement"`
}

type CampaignPredictionEvidence struct {
	CampaignKey     string                                `json:"campaign_key,omitempty"`
	Source          campaignenum.CampaignPredictionSource `json:"source"`
	WindowStart     time.Time                             `json:"window_start"`
	WindowEnd       time.Time                             `json:"window_end"`
	RawNetUnits     int                                   `json:"raw_net_units"`
	NormalizedUnits int                                   `json:"normalized_units"`
	Weight          float64                               `json:"weight"`
}

type CampaignProductPrediction struct {
	ProductSKUCode         string                                `json:"product_sku_code"`
	SupplierCode           string                                `json:"supplier_code,omitempty"`
	Source                 campaignenum.CampaignPredictionSource `json:"source"`
	Evidence               []CampaignPredictionEvidence          `json:"evidence,omitempty"`
	PredictedDemandUnits   int                                   `json:"predicted_demand_units"`
	SellableAvailableUnits int                                   `json:"sellable_available_units"`
	ConfirmedInboundUnits  int                                   `json:"confirmed_inbound_units"`
	NetRequiredUnits       int                                   `json:"net_required_units"`
	SuggestedOrderUnits    int                                   `json:"suggested_order_units"`
	SuggestedCartons       int                                   `json:"suggested_cartons,omitempty"`
	CartonSize             int                                   `json:"carton_size,omitempty"`
	MinimumOrderQuantity   int                                   `json:"minimum_order_quantity,omitempty"`
	Orderable              bool                                  `json:"orderable"`
	Warnings               []string                              `json:"warnings,omitempty"`
}

type CampaignSupplierPrediction struct {
	SupplierCode string                      `json:"supplier_code"`
	Products     []CampaignProductPrediction `json:"products"`
	TotalUnits   int                         `json:"total_units"`
	Warnings     []string                    `json:"warnings,omitempty"`
}

type CampaignPrediction struct {
	PredictionKey           string                                `json:"prediction_key"`
	CampaignKey             string                                `json:"campaign_key,omitempty"`
	Revision                int                                   `json:"revision"`
	Status                  campaignenum.CampaignPredictionStatus `json:"status"`
	AlgorithmVersion        string                                `json:"algorithm_version"`
	ResolvedProductSKUCodes []string                              `json:"resolved_product_sku_codes,omitempty"`
	Products                []CampaignProductPrediction           `json:"products,omitempty"`
	Suppliers               []CampaignSupplierPrediction          `json:"suppliers,omitempty"`
	Warnings                []string                              `json:"warnings,omitempty"`
	PredictedAt             time.Time                             `json:"predicted_at"`
}
