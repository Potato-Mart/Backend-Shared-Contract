package promotion

import (
	"time"

	geography "github.com/Potato-Mart/Backend-Shared-Contract/v25/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v25/pkg/contracts/common/localization"
	"github.com/Potato-Mart/Backend-Shared-Contract/v25/pkg/contracts/pricing/promotion/promotion_enums"
)

// StorefrontPromotion is the customer-safe catalogue projection of one
// promotion. It intentionally omits pricing rules, discount configuration,
// usage counters, source metadata, internal authoring history, and priority.
type StorefrontPromotion struct {
	ID          string                         `json:"id"`
	SeriesKey   string                         `json:"series_key"`
	Name        string                         `json:"name"`
	Description string                         `json:"description,omitempty"`
	Type        promotion_enums.PromotionType  `json:"type"`
	Class       promotion_enums.PromotionClass `json:"class"`
	TargetScope promotion_enums.DiscountScope  `json:"target_scope"`

	ProductSKUCodes     []string                     `json:"product_sku_codes,omitempty"`
	CategoryTagIDs      []string                     `json:"category_tag_ids,omitempty"`
	CategoryTagNames    []localization.LocalizedName `json:"category_tag_names,omitempty"`
	RequiredProductSKUs []string                     `json:"required_product_sku_codes,omitempty"`
	GiftProductSKUCode  string                       `json:"gift_product_sku_code,omitempty"`
	AddonProductSKUCode string                       `json:"addon_product_sku_code,omitempty"`
	StartsAt            *time.Time                   `json:"starts_at,omitempty"`
	ExpiresAt           *time.Time                   `json:"expires_at,omitempty"`
	ScheduleTimezone    string                       `json:"schedule_timezone"`
	GeographicScope     geography.GeographicScope    `json:"geographic_scope"`
	IsActive            bool                         `json:"is_active"`
}
