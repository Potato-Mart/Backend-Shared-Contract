package pricebook

import "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"

// SellingPromotionTerms holds the customer-safe deal terms of a promotion on a
// selling-price display. Every field is optional and omitted, not zeroed, when
// the deal has no such term. A pointer keeps a real zero, such as a free
// add-on price, distinct from an absent term.
//
// These are display facts, never executable rules. They carry no rule ids,
// priorities, stacking or usage counts, and the publishing service alone
// decides whether a deal applies.
type SellingPromotionTerms struct {
	// QualifierBaseUnits counts paid units only. The free units are extra, so
	// on a same-SKU BOGO, Buy 1 get 1 is qualifier 1, gift 1.
	QualifierBaseUnits *int64 `json:"qualifier_base_units,omitempty"`
	// GiftBaseUnits is how many free units one award gives.
	GiftBaseUnits *int64 `json:"gift_base_units,omitempty"`
	// BaseUnitsPerAward is how many units the cart needs to earn one award. On
	// a same-SKU BOGO it is qualifier plus gift (Buy 1 get 1 = 2, Buy 2 get 1
	// = 3). On a cross-SKU BOGO it is the qualifier alone. Clients use this
	// value rather than deriving it.
	BaseUnitsPerAward *int64 `json:"base_units_per_award,omitempty"`
	// MinimumOrderAmount is the spend the deal requires, where it has one.
	MinimumOrderAmount *money.Money `json:"minimum_order_amount,omitempty"`
	AddonPackagePrice  *money.Money `json:"addon_package_price,omitempty"`
	BundlePackagePrice *money.Money `json:"bundle_package_price,omitempty"`
	// DiscountBasisPoints, DiscountAmount and FixedPackagePrice state the
	// offer value of a discount deal. A deal carries the one that matches its
	// discount type.
	DiscountBasisPoints *int64       `json:"discount_basis_points,omitempty"`
	DiscountAmount      *money.Money `json:"discount_amount,omitempty"`
	FixedPackagePrice   *money.Money `json:"fixed_package_price,omitempty"`
	// MaximumApplicationsPerCustomer is the per-customer limit, when one is
	// set. It is the limit only, never how much of it has been used.
	MaximumApplicationsPerCustomer *int64 `json:"maximum_applications_per_customer,omitempty"`
	GroupPackageCount              *int64 `json:"group_package_count,omitempty"`
	// GiftTiers lists each tier of a multi-tier gift deal.
	GiftTiers []SellingPromotionGiftTerms `json:"gift_tiers,omitempty"`
}
