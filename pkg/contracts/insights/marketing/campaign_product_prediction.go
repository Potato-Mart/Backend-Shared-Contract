package marketing

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/insights/marketing/marketing_enums"
)

// CampaignProductPrediction is a per-product forecast in canonical base units.
type CampaignProductPrediction struct {
	SKUCode                    string                                   `json:"sku_code"`
	SupplierCode               string                                   `json:"supplier_code,omitempty"`
	Source                     marketing_enums.CampaignPredictionSource `json:"source"`
	Evidence                   []CampaignPredictionEvidence             `json:"evidence,omitempty"`
	PredictedDemandBaseUnits   int64                                    `json:"predicted_demand_base_units"`
	SellableAvailableBaseUnits int64                                    `json:"sellable_available_base_units"`
	ConfirmedInboundBaseUnits  int64                                    `json:"confirmed_inbound_base_units"`
	NetRequiredBaseUnits       int64                                    `json:"net_required_base_units"`
	SuggestedOrderBaseUnits    int64                                    `json:"suggested_order_base_units"`
	SuggestedComposition       packaging.PackageCompositionSnapshot     `json:"suggested_composition"`
	MinimumOrderBaseUnits      int64                                    `json:"minimum_order_base_units,omitempty"`
	Orderable                  bool                                     `json:"orderable"`
	Warnings                   []string                                 `json:"warnings,omitempty"`
}
