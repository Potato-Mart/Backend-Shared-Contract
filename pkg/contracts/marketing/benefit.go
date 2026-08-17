package marketing

import "github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/localization"

// BenefitRef is the customer-safe, code-only reference to a coupon or
// promotion. It deliberately carries no pricing rule, eligibility, or usage
// state. Path is server-generated as exactly /promotions/{campaignCode}: it
// always targets the Campaign landing page and never a coupon or Pricing
// Promotion lookup.
type BenefitRef struct {
	Code string                       `json:"code"`
	Name []localization.LocalizedName `json:"name"`
	// Path is the server-generated Campaign landing sub-path exactly
	// /promotions/{campaignCode}. It is not a full URL or a benefit lookup.
	Path string `json:"path"`
}
