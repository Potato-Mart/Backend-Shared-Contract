package benefit

import "github.com/Potato-Mart/Backend-Shared-Contract/v34/pkg/contracts/common/localization"

// BenefitRef is a customer-safe open reference to a Pricing-owned benefit.
// Kind and Code are deliberately open so Pricing can introduce benefit
// families without a shared-contract enum release.
// Code is the opaque canonical identity selected by Kind: promotion references
// use Promotion.ID, and coupon references use Coupon.ID, never the redeemable
// coupon code. The reference does not confer eligibility or ownership.
type BenefitRef struct {
	Kind string                       `json:"kind"`
	Code string                       `json:"code"`
	Name []localization.LocalizedName `json:"name,omitempty"`
}
