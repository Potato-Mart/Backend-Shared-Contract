package campaign

import (
	"time"

	geography "github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/customers/campaign/campaign_enums"
)

type CampaignComparableEvent struct {
	CampaignKey      string                                  `json:"campaign_key"`
	SeriesKey        string                                  `json:"series_key,omitempty"`
	MatchSource      campaign_enums.CampaignPredictionSource `json:"match_source"`
	ResolvedSKUIDs   []string                                `json:"resolved_sku_ids"`
	StartsAt         time.Time                               `json:"starts_at"`
	EndsAt           time.Time                               `json:"ends_at"`
	ScheduleTimezone string                                  `json:"schedule_timezone"`
	Audience         *Audience                               `json:"audience,omitempty"`
	GeographicScope  geography.GeographicScope               `json:"geographic_scope"`
	Placement        campaign_enums.CampaignPlacement        `json:"placement"`
}

type CampaignPredictionEvidence struct {
	CampaignKey         string                                  `json:"campaign_key,omitempty"`
	Source              campaign_enums.CampaignPredictionSource `json:"source"`
	WindowStart         time.Time                               `json:"window_start"`
	WindowEnd           time.Time                               `json:"window_end"`
	RawNetBaseUnits     int64                                   `json:"raw_net_base_units"`
	NormalizedBaseUnits int64                                   `json:"normalized_base_units"`
	Weight              float64                                 `json:"weight"`
}

type CampaignProductPrediction struct {
	SKUID                      string                                  `json:"sku_id"`
	SupplierCode               string                                  `json:"supplier_code,omitempty"`
	Source                     campaign_enums.CampaignPredictionSource `json:"source"`
	Evidence                   []CampaignPredictionEvidence            `json:"evidence,omitempty"`
	PredictedDemandBaseUnits   int64                                   `json:"predicted_demand_base_units"`
	SellableAvailableBaseUnits int64                                   `json:"sellable_available_base_units"`
	ConfirmedInboundBaseUnits  int64                                   `json:"confirmed_inbound_base_units"`
	NetRequiredBaseUnits       int64                                   `json:"net_required_base_units"`
	SuggestedOrderBaseUnits    int64                                   `json:"suggested_order_base_units"`
	SuggestedComposition       packaging.PackageCompositionSnapshot    `json:"suggested_composition"`
	MinimumOrderBaseUnits      int64                                   `json:"minimum_order_base_units,omitempty"`
	Orderable                  bool                                    `json:"orderable"`
	Warnings                   []string                                `json:"warnings,omitempty"`
}

type CampaignSupplierPrediction struct {
	SupplierCode   string                      `json:"supplier_code"`
	Products       []CampaignProductPrediction `json:"products"`
	TotalBaseUnits int64                       `json:"total_base_units"`
	Warnings       []string                    `json:"warnings,omitempty"`
}

type CampaignPrediction struct {
	PredictionKey    string                                  `json:"prediction_key"`
	CampaignKey      string                                  `json:"campaign_key,omitempty"`
	Revision         int                                     `json:"revision"`
	Status           campaign_enums.CampaignPredictionStatus `json:"status"`
	AlgorithmVersion string                                  `json:"algorithm_version"`
	ResolvedSKUIDs   []string                                `json:"resolved_sku_ids,omitempty"`
	Products         []CampaignProductPrediction             `json:"products,omitempty"`
	Suppliers        []CampaignSupplierPrediction            `json:"suppliers,omitempty"`
	Warnings         []string                                `json:"warnings,omitempty"`
	PredictedAt      time.Time                               `json:"predicted_at"`
}
