package listing

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v44/pkg/contracts/common/commerce/commerce_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v44/pkg/contracts/orders/shipping/shipping_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v44/pkg/contracts/pricing/listing/listing_enums"
)

// SaleRestriction is one market-specific restriction recorded against a
// listing. Channels narrows applicable non-age restrictions to specific order
// channels. Age rules apply to retail buyers independently of order channel.
// Kind-specific validation and enforcement remain service-owned.
type SaleRestriction struct {
	Kind     listing_enums.SaleRestrictionKind `json:"kind"`
	Channels []commerce_enums.OrderType        `json:"channels,omitempty"`
	Value    int64                             `json:"value,omitempty"`
	Note     string                            `json:"note,omitempty"`
	// AgeYears is the completed-calendar-years threshold for kind age. A
	// pointer distinguishes an explicit zero from absent evidence.
	AgeYears *int32 `json:"age_years,omitempty"`
	// AgeComparison identifies the allowed group: below permits age < AgeYears;
	// above permits age >= AgeYears. This rule applies only to retail buyers,
	// independent of channel. Services calculate completed years on the market
	// calendar; a 29 February DOB reaches its birthday on 1 March in non-leap years.
	// Services use the trusted saved Customers DOB. Guests or missing/unusable
	// DOB may view listings but cannot purchase an age-restricted item.
	// Wholesale buyers bypass the age rule. No proof verification is implied.
	AgeComparison listing_enums.AgeComparison `json:"age_comparison,omitempty"`
	// ExcludedDeliveryMethods enumerates the blocked physical methods for kind
	// delivery_excluded. It does not redefine order channels or buyer types.
	ExcludedDeliveryMethods []shipping_enums.DeliveryMethod `json:"excluded_delivery_methods,omitempty"`
}
