package promotion

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v12/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v12/pkg/contracts/shared"
	promotionenum "github.com/Potato-Mart/Backend-Shared-Contract/v12/pkg/enums/promotion"
	salesenum "github.com/Potato-Mart/Backend-Shared-Contract/v12/pkg/enums/sales"
)

// Promotion is the rule-based, auto-applied discount engine entity.
// It covers all promotion types: auto_discount, spend_gift, addon_purchase,
// bogo, bundle, and tiered_pricing.
type Promotion struct {
	ID          string                      `json:"id"`
	Name        string                      `json:"name"`
	Description string                      `json:"description,omitempty"`
	Type        promotionenum.PromotionType `json:"type"`

	// ── Classification & targeting (additive, v5.2.0) ─────────────────
	// Class separates standing promotions (常態特價, normal_promotion)
	// from time-boxed special campaigns (特殊活動, special_campaign).
	// Empty is read as normal_promotion so existing documents keep their
	// behaviour; see EffectiveClass.
	Class promotionenum.PromotionClass `json:"class,omitempty"`
	// TargetScope narrows the promotion to one product or one category tag.
	// Empty / "all" applies cart-wide.
	// Product targets are keyed by the product SKU code.
	TargetScope promotionenum.DiscountScope `json:"target_scope,omitempty"`
	// TargetProductSKUCode is required when TargetScope is "product".
	TargetProductSKUCode string `json:"target_product_sku_code,omitempty"`
	// TargetCategoryTagID and TargetCategoryTagName are required when
	// TargetScope is "category_tag". Matching uses the ID; the name keeps
	// the target displayable in rule-management UIs and audit history.
	TargetCategoryTagID   string                 `json:"target_category_tag_id,omitempty"`
	TargetCategoryTagName []common.LocalizedName `json:"target_category_tag_name,omitempty"`

	// ── Trigger conditions ────────────────────────────────────────────
	MinCartAmount           *common.Money `json:"min_cart_amount,omitempty"`
	MinCartQty              int           `json:"min_cart_qty,omitempty"`
	RequiredProductSKUCodes []string      `json:"required_product_sku_codes,omitempty"`
	RequiredQtyEach         int           `json:"required_qty_each"`

	// RequiredQtyMode controls whether RequiredQtyEach applies per product
	// or as a combined total across all required products.
	RequiredQtyMode promotionenum.PromotionQtyMode `json:"required_qty_mode,omitempty"`
	// RequiredQtyPerProduct overrides per-SKU qty when mode is PER_PRODUCT
	// and quantities differ across products. Map key is SKU, value is qty.
	RequiredQtyPerProduct common.Metadata `json:"required_qty_per_product,omitempty"`
	// RequiredQtyCombined is the total quantity threshold for COMBINED mode.
	RequiredQtyCombined int `json:"required_qty_combined,omitempty"`

	// ── auto_discount ─────────────────────────────────────────────────
	DiscountSpec
	MaxDiscount    *common.Money                         `json:"max_discount,omitempty"`
	DiscountTarget promotionenum.PromotionDiscountTarget `json:"discount_target,omitempty"`

	// ── spend_gift / bogo ─────────────────────────────────────────────
	GiftProductSKUCode string `json:"gift_product_sku_code,omitempty"`
	GiftQty            int    `json:"gift_qty,omitempty"`
	// GiftOptional allows the customer to decline the free gift at cart.
	GiftOptional bool `json:"gift_optional,omitempty"`
	// GiftTiers enables multiple spend thresholds for spend_gift promotions.
	GiftTiers []GiftTier `json:"gift_tiers,omitempty"`

	// ── addon_purchase ────────────────────────────────────────────────
	AddonProductSKUCode string                              `json:"addon_product_sku_code,omitempty"`
	AddonPrice          *common.Money                       `json:"addon_price,omitempty"`
	AddonMaxQty         int                                 `json:"addon_max_qty,omitempty"`
	AddonTriggerMode    promotionenum.PromotionAddonTrigger `json:"addon_trigger_mode,omitempty"`

	// ── bundle ────────────────────────────────────────────────────────
	BundlePrice *common.Money `json:"bundle_price,omitempty"`

	// ── tiered_pricing ────────────────────────────────────────────────
	PricingTiers      []PricingTier `json:"pricing_tiers,omitempty"`
	PricingMixAllowed bool          `json:"pricing_mix_allowed,omitempty"`

	// ── Control ───────────────────────────────────────────────────────
	Priority    int  `json:"priority"`
	IsStackable bool `json:"is_stackable"`
	UsageLimits
	ActiveWindow
	Channels []salesenum.OrderType `json:"channels,omitempty"` // e.g. ["online","pos"]

	// Source tracking for promotions synced from external systems.
	Source    string                `json:"source,omitempty"`
	SourceRef string                `json:"source_ref,omitempty"`
	History   []shared.HistoryEntry `json:"history,omitempty"`

	common.AuditFields
}

// GiftTier defines a single spend threshold and its associated free gift
// for spend_gift tier promotions.
type GiftTier struct {
	MinAmount          common.Money `json:"min_amount"`
	GiftProductSKUCode string       `json:"gift_product_sku_code"`
	GiftQty            int          `json:"gift_qty"`
}

// PricingTier defines a quantity break-point for tiered_pricing promotions.
// Either UnitPrice or TotalPrice should be set, not both.
type PricingTier struct {
	Qty        int           `json:"qty"`
	UnitPrice  *common.Money `json:"unit_price,omitempty"`
	TotalPrice *common.Money `json:"total_price,omitempty"`
}
