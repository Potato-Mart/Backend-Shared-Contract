package membership

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/localization"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/pricing/membership/membership_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/catalogue/classification"
)

// TierProgressTier is the tier summary embedded in customer tier progress.
// DisplayColour is the optional colour token clients use for the tier chip;
// it is omitted when Pricing has no colour for the tier.
type TierProgressTier struct {
	TierKey             string                             `json:"tier_key"`
	TierCard            *classification.ObjectMediaRef     `json:"tier_card,omitempty"`
	Label               []localization.LocalizedText       `json:"label,omitempty"`
	QualifyingThreshold *money.Money                       `json:"qualifying_threshold,omitempty"`
	DisplayColour       membership_enums.TierDisplayColour `json:"display_colour,omitempty"`
}
