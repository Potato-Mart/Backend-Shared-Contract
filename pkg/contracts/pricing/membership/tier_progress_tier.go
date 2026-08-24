package membership

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/localization"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/money"
)

// TierProgressTier is the tier summary embedded in customer tier progress.
type TierProgressTier struct {
	TierKey             string                       `json:"tier_key"`
	Label               []localization.LocalizedText `json:"label,omitempty"`
	QualifyingThreshold *money.Money                 `json:"qualifying_threshold,omitempty"`
}
