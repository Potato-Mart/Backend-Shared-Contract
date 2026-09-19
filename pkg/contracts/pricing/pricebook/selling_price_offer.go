package pricebook

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/localization"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/money"
)

// SellingPriceOffer is a customer-safe conditional or alternative price
// presentation resolved by Pricing. Optional package and membership targets
// describe qualification; consumers must not use an offer to replace the
// enclosing guest effective price.
type SellingPriceOffer struct {
	PackageOptionCode string                       `json:"package_option_code,omitempty"`
	MembershipTierKey string                       `json:"membership_tier_key,omitempty"`
	BaseUnits         int64                        `json:"base_units"`
	RegularAmount     money.Money                  `json:"regular_amount"`
	EffectiveAmount   money.Money                  `json:"effective_amount"`
	CompareAtAmount   *money.Money                 `json:"compare_at_amount,omitempty"`
	Messages          []localization.LocalizedText `json:"messages,omitempty"`
	Conditions        []localization.LocalizedText `json:"conditions,omitempty"`
	Conditional       bool                         `json:"conditional"`
	PromotionDisplays []SellingPromotionDisplay    `json:"promotions,omitempty"`
	ValidFrom         time.Time                    `json:"valid_from"`
	ValidUntil        *time.Time                   `json:"valid_until,omitempty"`
}
