package promotion

//import (
//	"time"
//
//	"github.com/Potato-Mart/Backend-Shared-Contract/pkg/enums"
//)
//
//type Promotion struct {
//	ID                 string                        `json:"id"`
//	Name               string                        `json:"name"`
//	Description        string                        `json:"description,omitempty"`
//	Type               enums.PromotionType           `json:"type"`
//	MinCartAmount      float64                       `json:"min_cart_amount,omitempty"`
//	MinCartQty         int                           `json:"min_cart_qty,omitempty"`
//	RequiredProductIDs []string                      `json:"required_product_ids,omitempty"`
//	RequiredQtyEach    int                           `json:"required_qty_each"`
//	DiscountType       enums.PromotionDiscountType   `json:"discount_type,omitempty"`
//	DiscountValue      float64                       `json:"discount_value,omitempty"`
//	MaxDiscount        float64                       `json:"max_discount,omitempty"`
//	DiscountTarget     enums.PromotionDiscountTarget `json:"discount_target,omitempty"`
//	GiftProductID      string                        `json:"gift_product_id,omitempty"`
//	GiftQty            int                           `json:"gift_qty,omitempty"`
//	AddonProductID     string                        `json:"addon_product_id,omitempty"`
//	AddonPrice         float64                       `json:"addon_price,omitempty"`
//	AddonMaxQty        int                           `json:"addon_max_qty,omitempty"`
//	BundlePrice        float64                       `json:"bundle_price,omitempty"`
//	Priority           int                           `json:"priority"`
//	IsStackable        bool                          `json:"is_stackable"`
//	UsageLimit         int                           `json:"usage_limit,omitempty"`
//	UsedCount          int                           `json:"used_count"`
//	PerCustomerLimit   int                           `json:"per_customer_limit"`
//	StartsAt           *time.Time                    `json:"starts_at,omitempty"`
//	ExpiresAt          *time.Time                    `json:"expires_at,omitempty"`
//	IsActive           bool                          `json:"is_active"`
//	OrderTypes         []enums.OrderType             `json:"order_types,omitempty"`
//	CreatedAt          time.Time                     `json:"created_at"`
//	UpdatedAt          time.Time                     `json:"updated_at"`
//}
