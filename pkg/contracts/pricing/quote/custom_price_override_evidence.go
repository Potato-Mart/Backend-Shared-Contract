package quote

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/pricing/quote/quote_enums"
	"time"
)

// CustomPriceOverrideEvidence freezes a cashier-entered tax-inclusive unit
// amount together with the actor, the mandatory reason, the approved price it
// replaced, and the cost comparison recorded at the time.
type CustomPriceOverrideEvidence struct {
	ActorUserID string `json:"actor_user_id"`
	Reason      string `json:"reason"`
	// SourceApprovedPrice is the approved price the override replaced.
	SourceApprovedPrice money.Money `json:"source_approved_price"`
	// OverrideGrossAmount is the tax-inclusive unit amount the cashier
	// entered. Tax is extracted from it rather than added to it.
	OverrideGrossAmount money.Money                `json:"override_gross_amount"`
	CostComparison      quote_enums.CostComparison `json:"cost_comparison"`
	ComparedCost        *money.Money               `json:"compared_cost,omitempty"`
	BelowCostWarning    bool                       `json:"below_cost_warning"`
	OverriddenAt        time.Time                  `json:"overridden_at"`
}
