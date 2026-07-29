package membership

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v20/pkg/common"
	membershipenum "github.com/Potato-Mart/Backend-Shared-Contract/v20/pkg/enums/membership"
)

// TierBenefit is one typed, localized membership tier benefit. Exactly one
// TierBenefitValue field is set for a given Kind; that invariant is enforced
// by the owning service, not the model.
type TierBenefit struct {
	BenefitKey  string                         `json:"benefit_key"`
	Kind        membershipenum.TierBenefitKind `json:"kind"`
	Title       []common.LocalizedText         `json:"title,omitempty"`
	Description []common.LocalizedText         `json:"description,omitempty"`
	Value       TierBenefitValue               `json:"value"`
}

// TierBenefitValue carries the benefit's value in exactly one representation.
// Decimal is a fixed-point decimal string (e.g. "1.5" points multiplier).
type TierBenefitValue struct {
	Decimal string        `json:"decimal,omitempty"`
	Money   *common.Money `json:"money,omitempty"`
	Integer *int64        `json:"integer,omitempty"`
	Boolean *bool         `json:"boolean,omitempty"`
}
