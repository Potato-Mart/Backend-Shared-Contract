package pricebook

import "github.com/Potato-Mart/Backend-Shared-Contract/v34/pkg/contracts/common/money"

// SellingPriceDisplay is customer-safe presentation evidence resolved by
// Pricing. All amounts use the same sellable base unit, currency, and tax basis
// as the enclosing SellingPrice. It neither selects a price nor replaces a
// checkout PriceSnapshot.
type SellingPriceDisplay struct {
	// RegularUnitPrice is the approved regular selling price in this context,
	// never supplier acquisition cost. EffectiveUnitPrice is the applicable
	// selling price at SellingPrice.AsOf, consistent with its UnitPrice.
	RegularUnitPrice   money.Money `json:"regular_unit_price"`
	EffectiveUnitPrice money.Money `json:"effective_unit_price"`
	// CompareAtUnitPrice is present only when the approved regular selling
	// price is valid as a comparison for this effective price. Absence is
	// distinct from a zero amount.
	CompareAtUnitPrice *money.Money `json:"compare_at_unit_price,omitempty"`
	// UnitPricePerMeasure is derived from EffectiveUnitPrice and its declared
	// measurement. Conditional promotion savings must not be presented as an
	// unconditional lower effective or comparison-unit price.
	UnitPricePerMeasure *SellingUnitPriceDisplay  `json:"unit_price_per_measure,omitempty"`
	PromotionDisplays   []SellingPromotionDisplay `json:"promotions,omitempty"`
	Offers              []SellingPriceOffer       `json:"offers,omitempty"`
}
