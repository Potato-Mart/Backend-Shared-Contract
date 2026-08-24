package order

import (
	"time"

	geography "github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/geography"

	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/commerce/commerce_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/orders/shipping"

	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/money"
)

// Channel is the order channel this cart is being built for
// (online/pos/b2b/...). Optional and additive.
type Cart struct {
	ID             string `json:"id"`
	SessionID      string `json:"session_id"`
	CustomerNumber string `json:"customer_number,omitempty"`
	// MarketCode is the immutable commercial market the cart is built for.
	// It is mandatory at cart creation and cannot change afterwards.
	MarketCode string `json:"market_code"`
	// CountryCode is the denormalized country of MarketCode, carried so a
	// country-scoped staff query is a plain indexed match.
	CountryCode geography.CountryCode `json:"country_code,omitempty"`

	Channel commerce_enums.OrderType `json:"channel,omitempty"`
	// Buyer describes who is buying, independently of Channel. POS is a
	// channel, not a buyer type — see sales.BuyerContext. Optional pointer
	// so it is omitted entirely when unset.
	Buyer              *BuyerContext                       `json:"buyer,omitempty"`
	FulfilmentLocation shipping.FulfilmentLocationSnapshot `json:"fulfilment_location"`
	Items              []CartItem                          `json:"items"`
	CouponCode         string                              `json:"coupon_code,omitempty"`
	Subtotal           money.Money                         `json:"subtotal"`
	ExpiresAt          time.Time                           `json:"expires_at"`

	audit.AuditFields
}
