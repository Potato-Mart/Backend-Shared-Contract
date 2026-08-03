package sales

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/contracts/product"
	salesenum "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/enums/sales"
)

type Cart struct {
	ID             string `json:"id"`
	SessionID      string `json:"session_id"`
	CustomerNumber string `json:"customer_number,omitempty"`
	// Channel is the order channel this cart is being built for
	// (online/pos/b2b/...). Optional and additive.
	Channel salesenum.OrderType `json:"channel,omitempty"`
	// Buyer describes who is buying, independently of Channel. POS is a
	// channel, not a buyer type — see sales.BuyerContext. Optional pointer
	// so it is omitted entirely when unset.
	Buyer             *BuyerContext            `json:"buyer,omitempty"`
	GeographicContext common.GeographicContext `json:"geographic_context"`
	Items             []CartItem               `json:"items"`
	CouponCode        string                   `json:"coupon_code,omitempty"`
	Subtotal          common.Money             `json:"subtotal"`
	ExpiresAt         time.Time                `json:"expires_at"`

	common.AuditFields
}

type CartItem struct {
	Product            product.Snapshot                `json:"product"`
	Components         []PricedPackageComponent        `json:"components"`
	TotalBaseUnits     int64                           `json:"total_base_units"`
	Pricing            *PricingContext                 `json:"pricing,omitempty"`
	SubstitutionPolicy LooseSubstitutionPolicySnapshot `json:"substitution_policy"`
	Preorder           *PreorderItemSnapshot           `json:"preorder,omitempty"`
}
