package sales

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v9/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v9/pkg/contracts/product"
	"github.com/Potato-Mart/Backend-Shared-Contract/v9/pkg/enums"
)

type Cart struct {
	ID         string `json:"id"`
	SessionID  string `json:"session_id"`
	CustomerID string `json:"customer_id,omitempty"`
	// Channel is the order channel this cart is being built for
	// (online/pos/b2b/...). Optional and additive; absent on legacy carts.
	Channel enums.OrderType `json:"channel,omitempty"`
	// Buyer describes who is buying, independently of Channel. POS is a
	// channel, not a buyer type — see sales.BuyerContext. Optional pointer
	// so it is omitted entirely on legacy carts.
	Buyer      *BuyerContext `json:"buyer,omitempty"`
	Items      []CartItem    `json:"items"`
	CouponCode string        `json:"coupon_code,omitempty"`
	Subtotal   common.Money  `json:"subtotal"`
	ExpiresAt  time.Time     `json:"expires_at"`

	common.AuditFields `bson:",inline"`
}

type CartItem struct {
	Product product.Snapshot `json:"product"`
	Price   common.Money     `json:"price"`
	// Pricing is the commercial pricing context under which Price was set
	// (retail vs wholesale audience, visibility). Optional pointer so it is
	// omitted entirely on legacy cart items.
	Pricing  *PricingContext `json:"pricing,omitempty"`
	Quantity int             `json:"quantity"`
}
