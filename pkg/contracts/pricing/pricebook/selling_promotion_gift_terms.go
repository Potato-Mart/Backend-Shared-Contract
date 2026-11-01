package pricebook

import "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"

// SellingPromotionGiftTerms is one tier of a multi-tier gift deal, as shown to
// a customer. Every field is optional and omitted when the tier has no such
// term. These are display facts, never executable rules.
type SellingPromotionGiftTerms struct {
	// QualifierBaseUnits counts paid units only, as it does on
	// SellingPromotionTerms. The free units are extra.
	QualifierBaseUnits *int64 `json:"qualifier_base_units,omitempty"`
	// GiftBaseUnits is how many free units this tier gives.
	GiftBaseUnits      *int64       `json:"gift_base_units,omitempty"`
	MinimumOrderAmount *money.Money `json:"minimum_order_amount,omitempty"`
}
