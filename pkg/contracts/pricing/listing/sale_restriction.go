package listing

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v41/pkg/contracts/common/commerce/commerce_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v41/pkg/contracts/orders/shipping/shipping_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v41/pkg/contracts/pricing/listing/listing_enums"
)

// SaleRestriction is one market-specific restriction recorded against a
// listing. Channels narrows the restriction to specific order channels when it
// does not apply everywhere. Channels never broadens age restrictions beyond
// retail buyers. Kind-specific validation and enforcement remain service-owned.
type SaleRestriction struct {
	Kind     listing_enums.SaleRestrictionKind `json:"kind"`
	Channels []commerce_enums.OrderType        `json:"channels,omitempty"`
	Value    int64                             `json:"value,omitempty"`
	Note     string                            `json:"note,omitempty"`
	// AgeYears is the completed-calendar-years threshold for kind age. A
	// pointer distinguishes an explicit zero from absent evidence.
	AgeYears *int32 `json:"age_years,omitempty"`
	// AgeComparison describes which side of AgeYears is blocked; equality is
	// allowed. This rule applies only to retail buyers, independent of channel.
	// Services use the trusted saved Customers DOB. Guests or missing/unusable
	// DOB may view listings but cannot purchase an age-restricted item.
	// Wholesale buyers bypass the age rule. No proof verification is implied.
	AgeComparison listing_enums.AgeComparison `json:"age_comparison,omitempty"`
	// ExcludedDeliveryMethods enumerates the blocked physical methods for kind
	// delivery_excluded. It does not redefine order channels or buyer types.
	ExcludedDeliveryMethods []shipping_enums.DeliveryMethod `json:"excluded_delivery_methods,omitempty"`
}
