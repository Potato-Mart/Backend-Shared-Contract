package pos

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/pricing/quote/quote_enums"
)

// CashRoundingSnapshot is the separate settlement evidence recorded when a
// sale is tendered entirely in cash.
//
// The order's consideration and tax stay at exact minor units: rounding is a
// tender adjustment, never a change to the priced amounts. Electronic tenders
// remain exact and record no adjustment, and a tender that mixes cash with
// another method is rejected.
type CashRoundingSnapshot struct {
	// ExactAmountDue is the unrounded amount payable, equal to the order
	// total at exact minor units.
	ExactAmountDue money.Money `json:"exact_amount_due"`
	// RoundingAdjustment is CashAmountDue minus ExactAmountDue and may be
	// negative.
	RoundingAdjustment money.Money `json:"rounding_adjustment"`
	// CashAmountDue is the rounded amount actually tendered in cash.
	CashAmountDue money.Money `json:"cash_amount_due"`
	// IncrementMinor is the market's cash increment in minor units, for
	// example 5 for Australian five-cent rounding.
	IncrementMinor int64                    `json:"increment_minor"`
	Mode           quote_enums.RoundingMode `json:"mode"`
	AppliedAt      time.Time                `json:"applied_at"`
}
