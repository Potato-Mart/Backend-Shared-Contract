package order

import (
	"time"

	geography "github.com/Potato-Mart/Backend-Shared-Contract/v25/pkg/contracts/common/geography"

	"github.com/Potato-Mart/Backend-Shared-Contract/v25/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v25/pkg/contracts/common/commerce/commerce_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v25/pkg/contracts/supply/product"

	"github.com/Potato-Mart/Backend-Shared-Contract/v25/pkg/contracts/common/money"
)

// Channel is the order channel this cart is being built for
// (online/pos/b2b/...). Optional and additive.
type Cart struct {
	ID             string `json:"id"`
	SessionID      string `json:"session_id"`
	CustomerNumber string `json:"customer_number,omitempty"`

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
	Product            product.Snapshot                `json:"product"`
	Components         []PricedPackageComponent        `json:"components"`
	TotalBaseUnits     int64                           `json:"total_base_units"`
	Pricing            *PricingContext                 `json:"pricing,omitempty"`
	SubstitutionPolicy LooseSubstitutionPolicySnapshot `json:"substitution_policy"`
	Preorder           *PreorderItemSnapshot           `json:"preorder,omitempty"`
}
