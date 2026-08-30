package benefit

import "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/localization"

// BenefitRef is a customer-safe open reference to a Pricing-owned benefit.
// Kind and Code are deliberately open so Pricing can introduce benefit
// families without a shared-contract enum release.
type BenefitRef struct {
	Kind string                       `json:"kind"`
	Code string                       `json:"code"`
	Name []localization.LocalizedName `json:"name,omitempty"`
}
