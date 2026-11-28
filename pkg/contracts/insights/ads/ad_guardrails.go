package ads

import "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"

// AdGuardrails are the spending limits a delivery change is checked against
// before it reaches a platform.
//
// MaxDailyBudgetPerAccount is a hard ceiling, one amount per account currency:
// a change that would exceed it is refused outright and cannot be waved
// through with a confirmation. ConfirmIncreaseOverPercent is the softer limit:
// an increase above that percentage is allowed but only once the person
// submitting it has confirmed the increase explicitly.
type AdGuardrails struct {
	MaxDailyBudgetPerAccount   []money.Money `json:"max_daily_budget_per_account,omitempty"`
	ConfirmIncreaseOverPercent int32         `json:"confirm_increase_over_percent,omitempty"`
}
