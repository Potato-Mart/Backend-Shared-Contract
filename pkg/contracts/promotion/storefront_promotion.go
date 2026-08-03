package promotion

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/common"
	promotionenum "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/enums/promotion"
)

// StorefrontPromotion is the customer-safe catalogue projection of one
// promotion. It intentionally omits pricing rules, discount configuration,
// usage counters, source metadata, internal authoring history, and priority.
type StorefrontPromotion struct {
	ID          string                       `json:"id"`
	SeriesKey   string                       `json:"series_key"`
	Name        string                       `json:"name"`
	Description string                       `json:"description,omitempty"`
	Type        promotionenum.PromotionType  `json:"type"`
	Class       promotionenum.PromotionClass `json:"class"`
	TargetScope promotionenum.DiscountScope  `json:"target_scope"`

	ProductSKUCodes     []string               `json:"product_sku_codes,omitempty"`
	CategoryTagIDs      []string               `json:"category_tag_ids,omitempty"`
	CategoryTagNames    []common.LocalizedName `json:"category_tag_names,omitempty"`
	RequiredProductSKUs []string               `json:"required_product_sku_codes,omitempty"`
	GiftProductSKUCode  string                 `json:"gift_product_sku_code,omitempty"`
	AddonProductSKUCode string                 `json:"addon_product_sku_code,omitempty"`
	StartsAt            *time.Time             `json:"starts_at,omitempty"`
	ExpiresAt           *time.Time             `json:"expires_at,omitempty"`
	ScheduleTimezone    string                 `json:"schedule_timezone"`
	GeographicScope     common.GeographicScope `json:"geographic_scope"`
	IsActive            bool                   `json:"is_active"`
}
