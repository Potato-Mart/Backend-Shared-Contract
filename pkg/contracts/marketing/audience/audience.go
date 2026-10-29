package audience

import "github.com/Potato-Mart/Backend-Shared-Contract/v50/pkg/contracts/marketing/audience/audience_enums"

// Audience narrows a Pricing-owned offer. Omitted dimensions are unrestricted.
type Audience struct {
	CustomerType audience_enums.CustomerType `json:"customer_type,omitempty"`
	Platform     audience_enums.Platform     `json:"platform,omitempty"`
}
