package membership

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/localization"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/money"

	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/pricing/membership/membership_enums"
)

// CustomerTierProgress is the customer-safe projection of progress toward the
// next membership tier. When Available is false, Reason explains why and the
// progress fields are omitted.
type CustomerTierProgress struct {
	Available                bool                                   `json:"available"`
	Reason                   membership_enums.TierProgressReason    `json:"reason,omitempty"`
	QualificationMetric      *membership_enums.MembershipTierMetric `json:"qualification_metric,omitempty"`
	QualificationWindow      *QualificationWindow                   `json:"qualification_window,omitempty"`
	QualifyingSpend          *money.Money                           `json:"qualifying_spend,omitempty"`
	CurrentTier              *TierProgressTier                      `json:"current_tier,omitempty"`
	NextTier                 *TierProgressTier                      `json:"next_tier,omitempty"`
	RemainingQualifyingSpend *money.Money                           `json:"remaining_qualifying_spend,omitempty"`
	ProgressBasisPoints      *int                                   `json:"progress_basis_points,omitempty"`
	IsMaxTier                *bool                                  `json:"is_max_tier,omitempty"`
	CalculatedAt             time.Time                              `json:"calculated_at"`
}

// TierProgressTier is the tier summary embedded in customer tier progress.
type TierProgressTier struct {
	TierKey             string                       `json:"tier_key"`
	Label               []localization.LocalizedText `json:"label,omitempty"`
	QualifyingThreshold *money.Money                 `json:"qualifying_threshold,omitempty"`
}

// QualificationWindow is the period over which qualifying spend is measured.
type QualificationWindow struct {
	StartsAt *time.Time `json:"starts_at,omitempty"`
	EndsAt   *time.Time `json:"ends_at,omitempty"`
}
