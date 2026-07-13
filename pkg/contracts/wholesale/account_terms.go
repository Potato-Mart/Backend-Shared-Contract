package wholesale

import "github.com/Potato-Mart/Backend-Shared-Contract/v16/pkg/common"

// WholesaleTerms groups B2B price-tier configuration, checkout policy, credit,
// and freight terms for a wholesale organisation/business account.
type WholesaleTerms struct {
	TierKey                string          `json:"tier_key,omitempty"`
	PriceTier              int             `json:"price_tier,omitempty"`
	PriceTierSea           *int            `json:"price_tier_sea,omitempty"`
	PriceTierAir           *int            `json:"price_tier_air,omitempty"`
	RebateRate             *float64        `json:"rebate_rate,omitempty"`
	ShippingFee            *common.Money   `json:"shipping_fee,omitempty"`
	FreightRule            common.Metadata `json:"freight_rule,omitempty"`
	AccountCheckoutEnabled bool            `json:"account_checkout_enabled,omitempty"`
	DueDays                int             `json:"due_days,omitempty"`
	PurchaseOrderRequired  bool            `json:"purchase_order_required,omitempty"`
	PaymentInstructions    string          `json:"payment_instructions,omitempty"`
	CreditLimit            *common.Money   `json:"credit_limit,omitempty"`
}
