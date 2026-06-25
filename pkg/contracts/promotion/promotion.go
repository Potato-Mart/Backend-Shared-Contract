package promotion

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v9/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v9/pkg/contracts/shared"
	"github.com/Potato-Mart/Backend-Shared-Contract/v9/pkg/enums"
)

// Promotion is the rule-based, auto-applied discount engine entity.
// It covers all promotion types: auto_discount, spend_gift, addon_purchase,
// bogo, bundle, and tiered_pricing.
type Promotion struct {
	ID          string              `json:"id"`
	Name        string              `json:"name"`
	Description string              `json:"description,omitempty"`
	Type        enums.PromotionType `json:"type"`

	// ── Classification & targeting (additive, v5.2.0) ─────────────────
	// Class separates standing promotions (常態特價, normal_promotion)
	// from time-boxed special campaigns (特殊活動, special_campaign).
	// Empty is read as normal_promotion so existing documents keep their
	// behaviour; see EffectiveClass.
	Class enums.PromotionClass `json:"class,omitempty"`
	// TargetScope narrows the promotion to one product or one category
	// subtree. Empty / "all" keeps the legacy cart-wide behaviour.
	// SKU-level targeting is deliberately not supported.
	TargetScope enums.DiscountScope `json:"target_scope,omitempty"`
	// TargetProductID is required when TargetScope is "product".
	TargetProductID string `json:"target_product_id,omitempty"`
	// TargetCategoryKey is required when TargetScope is "category". It
	// references a node of the ops product-category tree by key; the
	// promotion applies to every product whose category path contains
	// the key (see TargetIncludesDescendants).
	TargetCategoryKey string `json:"target_category_key,omitempty"`
	// TargetIncludesDescendants controls whether a category target also
	// covers descendant categories. nil means true (the default rule:
	// category promotions apply to the whole subtree).
	TargetIncludesDescendants *bool `json:"target_includes_descendants,omitempty"`

	// ── Trigger conditions ────────────────────────────────────────────
	MinCartAmount      *common.Money `json:"min_cart_amount,omitempty"`
	MinCartQty         int           `json:"min_cart_qty,omitempty"`
	RequiredProductIDs []string      `json:"required_product_ids,omitempty"`
	RequiredQtyEach    int           `json:"required_qty_each"`

	// RequiredQtyMode controls whether RequiredQtyEach applies per product
	// or as a combined total across all required products.
	RequiredQtyMode enums.PromotionQtyMode `json:"required_qty_mode,omitempty"`
	// RequiredQtyPerProduct overrides per-SKU qty when mode is PER_PRODUCT
	// and quantities differ across products. Map key is SKU, value is qty.
	RequiredQtyPerProduct common.Metadata `json:"required_qty_per_product,omitempty"`
	// RequiredQtyCombined is the total quantity threshold for COMBINED mode.
	RequiredQtyCombined int `json:"required_qty_combined,omitempty"`

	// ── auto_discount ─────────────────────────────────────────────────
	DiscountSpec   `bson:",inline"`
	MaxDiscount    *common.Money                 `json:"max_discount,omitempty"`
	DiscountTarget enums.PromotionDiscountTarget `json:"discount_target,omitempty"`

	// ── spend_gift / bogo ─────────────────────────────────────────────
	GiftProductID string `json:"gift_product_id,omitempty"`
	GiftQty       int    `json:"gift_qty,omitempty"`
	// GiftOptional allows the customer to decline the free gift at cart.
	GiftOptional bool `json:"gift_optional,omitempty"`
	// GiftTiers enables multiple spend thresholds for spend_gift promotions.
	GiftTiers []GiftTier `json:"gift_tiers,omitempty"`

	// ── addon_purchase ────────────────────────────────────────────────
	AddonProductID   string                      `json:"addon_product_id,omitempty"`
	AddonPrice       *common.Money               `json:"addon_price,omitempty"`
	AddonMaxQty      int                         `json:"addon_max_qty,omitempty"`
	AddonTriggerMode enums.PromotionAddonTrigger `json:"addon_trigger_mode,omitempty"`

	// ── bundle ────────────────────────────────────────────────────────
	BundlePrice *common.Money `json:"bundle_price,omitempty"`

	// ── tiered_pricing ────────────────────────────────────────────────
	PricingTiers      []PricingTier `json:"pricing_tiers,omitempty"`
	PricingMixAllowed bool          `json:"pricing_mix_allowed,omitempty"`

	// ── Control ───────────────────────────────────────────────────────
	Priority     int  `json:"priority"`
	IsStackable  bool `json:"is_stackable"`
	UsageLimits  `bson:",inline"`
	ActiveWindow `bson:",inline"`
	Channels     []enums.OrderType `json:"channels,omitempty"` // e.g. ["online","pos"]

	// Source tracking for promotions synced from external systems.
	Source    string                `json:"source,omitempty"`
	SourceRef string                `json:"source_ref,omitempty"`
	History   []shared.HistoryEntry `json:"history,omitempty"`

	common.AuditFields `bson:",inline"`
}

// GiftTier defines a single spend threshold and its associated free gift
// for spend_gift tier promotions.
type GiftTier struct {
	MinAmount     common.Money `json:"min_amount"`
	GiftProductID string       `json:"gift_product_id"`
	GiftQty       int          `json:"gift_qty"`
}

// PricingTier defines a quantity break-point for tiered_pricing promotions.
// Either UnitPrice or TotalPrice should be set, not both.
type PricingTier struct {
	Qty        int           `json:"qty"`
	UnitPrice  *common.Money `json:"unit_price,omitempty"`
	TotalPrice *common.Money `json:"total_price,omitempty"`
}
