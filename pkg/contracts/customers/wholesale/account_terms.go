package wholesale

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/metadata"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/money"
)

// WholesaleTerms groups B2B checkout policy, credit,
// and freight terms for a wholesale organisation/business account.
type WholesaleTerms struct {
	RebateRate             *float64          `json:"rebate_rate,omitempty"`
	ShippingFee            *money.Money      `json:"shipping_fee,omitempty"`
	FreightRule            metadata.Metadata `json:"freight_rule,omitempty"`
	AccountCheckoutEnabled bool              `json:"account_checkout_enabled,omitempty"`
	DueDays                int               `json:"due_days,omitempty"`
	PurchaseOrderRequired  bool              `json:"purchase_order_required,omitempty"`
	PaymentInstructions    string            `json:"payment_instructions,omitempty"`
	CreditLimit            *money.Money      `json:"credit_limit,omitempty"`
}
