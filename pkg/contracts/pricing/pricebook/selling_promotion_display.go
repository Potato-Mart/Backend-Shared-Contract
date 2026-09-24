package pricebook

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v35/pkg/contracts/common/localization"
)

// SellingPromotionDisplay describes a customer-visible promotion without exposing
// internal rule, actor, approval, or cost records. Kind is open so configured
// presentations such as everyday_special, weekly_promotion, bogo, group_order,
// and discount can evolve without changing shared commercial rule semantics.
// Conditions describe customer qualification requirements, not executable
// rules. A displayed promotion need not have affected the effective unit price.
type SellingPromotionDisplay struct {
	Kind                    string                       `json:"kind"`
	Messages                []localization.LocalizedText `json:"messages,omitempty"`
	Conditions              []localization.LocalizedText `json:"conditions,omitempty"`
	Conditional             bool                         `json:"conditional"`
	AppliedToEffectivePrice bool                         `json:"applied_to_effective_price"`
	ValidFrom               *time.Time                   `json:"valid_from,omitempty"`
	ValidUntil              *time.Time                   `json:"valid_until,omitempty"`
}
