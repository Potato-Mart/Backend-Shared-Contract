package order

import (
	"time"

	geography "github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/geography"

	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/commerce/commerce_enums"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/supply/product"

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
	Buyer             *BuyerContext               `json:"buyer,omitempty"`
	GeographicContext geography.GeographicContext `json:"geographic_context"`
	Items             []CartItem                  `json:"items"`
	CouponCode        string                      `json:"coupon_code,omitempty"`
	Subtotal          money.Money                 `json:"subtotal"`
	ExpiresAt         time.Time                   `json:"expires_at"`

	audit.AuditFields
}

type CartItem struct {
	// SKUCode is the frozen SKU code captured when the line was priced.
	SKUCode              string                          `json:"sku_code"`
	ProductName          string                          `json:"product_name"`
	ProductImage         *security.ObjectMedia           `json:"product_image,omitempty"`
	ProductPackageOption product.ProductPackageOption    `json:"product_package_option"`
	CapturedAt           time.Time                       `json:"captured_at"`
	Components           []PricedPackageComponent        `json:"components"`
	TotalBaseUnits       int64                           `json:"total_base_units"`
	Pricing              *PricingContext                 `json:"pricing,omitempty"`
	SubstitutionPolicy   LooseSubstitutionPolicySnapshot `json:"substitution_policy"`
	Preorder             *PreorderItemSnapshot           `json:"preorder,omitempty"`
}
