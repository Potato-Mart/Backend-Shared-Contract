package quote

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/money"
	"time"
)

// AppliedPriceRule is one frozen rule outcome. Kind is an open string so new
// commercial mechanics do not require a contract major bump; current values
// include group_15, group_20, expiry_20, expiry_bogo, damage_tier_30,
// damage_tier_50, damage_tier_80, and custom_pos. Exclusive rules never stack
// with each other or with ordinary promotions.
type AppliedPriceRule struct {
	Kind         string `json:"kind"`
	RuleID       string `json:"rule_id,omitempty"`
	RuleRevision int64  `json:"rule_revision,omitempty"`
	Exclusive    bool   `json:"exclusive"`
	// FactorNumerator over FactorDenominator is the exact multiplier the
	// rule applied, for example 85/100 for group_15 or 8/10 for
	// damage_tier_30.
	FactorNumerator   int64       `json:"factor_numerator,omitempty"`
	FactorDenominator int64       `json:"factor_denominator,omitempty"`
	AmountBefore      money.Money `json:"amount_before"`
	AmountAfter       money.Money `json:"amount_after"`
	// ChargeableBaseUnits records how many base units the rule charged for,
	// which differs from the reserved quantity under a buy-one-get-one
	// expiry mechanic.
	ChargeableBaseUnits int64     `json:"chargeable_base_units,omitempty"`
	Reason              string    `json:"reason,omitempty"`
	AppliedAt           time.Time `json:"applied_at"`
}
