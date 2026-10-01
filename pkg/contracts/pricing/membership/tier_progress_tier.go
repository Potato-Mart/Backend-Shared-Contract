package membership

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v40/pkg/contracts/common/localization"
	"github.com/Potato-Mart/Backend-Shared-Contract/v40/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v40/pkg/contracts/supply/catalogue/classification"
)

// TierProgressTier is the tier summary embedded in customer tier progress.
type TierProgressTier struct {
	TierKey             string                         `json:"tier_key"`
	TierCard            *classification.ObjectMediaRef `json:"tier_card,omitempty"`
	Label               []localization.LocalizedText   `json:"label,omitempty"`
	QualifyingThreshold *money.Money                   `json:"qualifying_threshold,omitempty"`
}
