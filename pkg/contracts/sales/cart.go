package sales

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v15/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v15/pkg/contracts/product"
	salesenum "github.com/Potato-Mart/Backend-Shared-Contract/v15/pkg/enums/sales"
)

type Cart struct {
	ID             string `json:"id"`
	SessionID      string `json:"session_id"`
	CustomerNumber string `json:"customer_number,omitempty"`
	// Channel is the order channel this cart is being built for
	// (online/pos/b2b/...). Optional and additive.
	Channel salesenum.OrderType `json:"channel,omitempty"`
	// Buyer describes who is buying, independently of Channel. POS is a
	// channel, not a buyer type â€” see sales.BuyerContext. Optional pointer
	// so it is omitted entirely when unset.
	Buyer      *BuyerContext `json:"buyer,omitempty"`
	Items      []CartItem    `json:"items"`
	CouponCode string        `json:"coupon_code,omitempty"`
	Subtotal   common.Money  `json:"subtotal"`
	ExpiresAt  time.Time     `json:"expires_at"`

	common.AuditFields
}

type CartItem struct {
	Product product.Snapshot `json:"product"`
	Price   common.Money     `json:"price"`
	// Pricing is the commercial pricing context under which Price was set
	// (retail vs wholesale audience, visibility). Optional pointer so it is
	// omitted entirely when unset.
	Pricing    *PricingContext `json:"pricing,omitempty"`
	Quantity   int             `json:"quantity"`
	Properties common.Metadata `json:"properties,omitempty"`
}
