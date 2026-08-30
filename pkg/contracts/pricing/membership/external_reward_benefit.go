package membership

import "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/localization"

// ExternalRewardBenefit configures a reward fulfilled by an external partner
// system, such as another company's subscription or service. ProviderCode is a
// backend-configured open code so a new partner ships without a shared-contract
// release, and ExternalProductCode and ExternalPlanCode are that partner's own
// identifiers for what a redemption provisions.
type ExternalRewardBenefit struct {
	ProviderCode        string                       `json:"provider_code"`
	ExternalProductCode string                       `json:"external_product_code,omitempty"`
	ExternalPlanCode    string                       `json:"external_plan_code,omitempty"`
	DisplayNames        []localization.LocalizedName `json:"display_names,omitempty"`
}
