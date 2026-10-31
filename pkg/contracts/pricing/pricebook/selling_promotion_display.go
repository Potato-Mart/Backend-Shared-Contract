package pricebook

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/localization"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/pricing/pricebook/pricebook_enums"
)

// SellingPromotionDisplay describes a customer-visible promotion without exposing
// internal rule, actor, approval, or cost records. Kind is open so configured
// presentations such as everyday_special, weekly_promotion, bogo, group_order,
// and discount can evolve without changing shared commercial rule semantics.
// Conditions describe customer qualification requirements, not executable
// rules. A displayed promotion need not have affected the effective unit price.
type SellingPromotionDisplay struct {
	// PromotionID is the corresponding customer-visible storefront promotion.
	// It is omitted when the display does not originate from that catalogue.
	PromotionID             string                                        `json:"promotion_id,omitempty"`
	DiscountKind            *pricebook_enums.SellingPromotionDiscountType `json:"discount_type,omitempty"`
	DiscountValue           *SellingPromotionDiscountValue                `json:"discount_value,omitempty"`
	Kind                    string                                        `json:"kind"`
	Messages                []localization.LocalizedText                  `json:"messages,omitempty"`
	Conditions              []localization.LocalizedText                  `json:"conditions,omitempty"`
	Conditional             bool                                          `json:"conditional"`
	AppliedToEffectivePrice bool                                          `json:"applied_to_effective_price"`
	ValidFrom               *time.Time                                    `json:"valid_from,omitempty"`
	ValidUntil              *time.Time                                    `json:"valid_until,omitempty"`
	// RequiredQuantity, MemberSKUCodes and MemberSeriesKeys describe what
	// qualifies a conditional promotion. They are display facts, not rules.
	RequiredQuantity *int64   `json:"required_quantity,omitempty"`
	MemberSKUCodes   []string `json:"member_sku_codes,omitempty"`
	MemberSeriesKeys []string `json:"member_series_keys,omitempty"`
}
