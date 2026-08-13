package marketing

import "github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/localization"

// BenefitRef is the customer-safe, code-only reference to a coupon or
// promotion. It deliberately carries no pricing rule, eligibility, or usage
// state; consumers follow Path to obtain the authoritative public aggregate.
type BenefitRef struct {
	Code string                       `json:"code"`
	Name []localization.LocalizedName `json:"name"`
	Path string                       `json:"path"`
}
