package membership

import "github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/money"

// TierBenefitValue carries exactly one value arm for a membership benefit.
type TierBenefitValue struct {
	Decimal string       `json:"decimal,omitempty"`
	Money   *money.Money `json:"money,omitempty"`
	Integer *int64       `json:"integer,omitempty"`
	Boolean *bool        `json:"boolean,omitempty"`
}
