package quote

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/pricing/pricebook/pricebook_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/pricing/quote/quote_enums"
)

// RoundingEvidence records the exact value a rounded minor amount came from
// and the rule that reduced it, so every rounded cent is reproducible.
type RoundingEvidence struct {
	Mode        quote_enums.RoundingMode          `json:"mode"`
	PriceEnding pricebook_enums.PriceEndingPolicy `json:"price_ending"`
	Exponent    int32                             `json:"exponent"`
	// ExactNumerator over ExactDenominator is the unrounded value in minor
	// units before the mode was applied.
	ExactNumerator   int64       `json:"exact_numerator"`
	ExactDenominator int64       `json:"exact_denominator"`
	RoundedAmount    money.Money `json:"rounded_amount"`
	// RemainderRank and RemainderMinorApplied record this line's position
	// and share when document minor units were allocated by largest
	// remainder. TieBreakKey is the stable line identity used to break
	// equal remainders.
	RemainderRank         int32  `json:"remainder_rank,omitempty"`
	RemainderMinorApplied int64  `json:"remainder_minor_applied,omitempty"`
	TieBreakKey           string `json:"tie_break_key,omitempty"`
}
