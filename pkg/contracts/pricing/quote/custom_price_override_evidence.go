package quote

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/pricing/quote/quote_enums"
	"time"
)

// CustomPriceOverrideEvidence freezes a tax-inclusive custom unit amount,
// entered at POS or resolved from an approved scoped customization, together
// with its actor, mandatory reason, replaced approved price and cost evidence.
type CustomPriceOverrideEvidence struct {
	ActorUserID string `json:"actor_user_id"`
	Reason      string `json:"reason"`
	// ReasonCode categorizes the customization without replacing the
	// mandatory human explanation in Reason. The optional code and revision
	// identify the approved custom-price definition frozen at checkout.
	ReasonCode          quote_enums.CustomPriceReason `json:"reason_code,omitempty"`
	CustomPriceCode     string                        `json:"custom_price_code,omitempty"`
	CustomPriceRevision *int64                        `json:"custom_price_revision,omitempty"`
	// SourceApprovedPrice is the approved price the override replaced.
	SourceApprovedPrice money.Money `json:"source_approved_price"`
	// OverrideGrossAmount is the tax-inclusive custom unit amount. Tax is
	// extracted from it rather than added to it, regardless of its source.
	OverrideGrossAmount money.Money                `json:"override_gross_amount"`
	CostComparison      quote_enums.CostComparison `json:"cost_comparison"`
	ComparedCost        *money.Money               `json:"compared_cost,omitempty"`
	BelowCostWarning    bool                       `json:"below_cost_warning"`
	OverriddenAt        time.Time                  `json:"overridden_at"`
}
