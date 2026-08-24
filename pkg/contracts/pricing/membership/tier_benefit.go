package membership

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/localization"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/pricing/membership/membership_enums"
)

// TierBenefit is one typed, localized membership tier benefit. Exactly one
// TierBenefitValue field is set for a given Kind; that invariant is enforced
// by the owning service, not the model.
type TierBenefit struct {
	BenefitKey  string                           `json:"benefit_key"`
	Kind        membership_enums.TierBenefitKind `json:"kind"`
	Title       []localization.LocalizedText     `json:"title,omitempty"`
	Description []localization.LocalizedText     `json:"description,omitempty"`
	Value       TierBenefitValue                 `json:"value"`
}
