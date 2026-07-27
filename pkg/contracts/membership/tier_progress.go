package membership

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v19/pkg/common"
	membershipenum "github.com/Potato-Mart/Backend-Shared-Contract/v19/pkg/enums/membership"
)

// CustomerTierProgress is the customer-safe projection of progress toward the
// next membership tier. When Available is false, Reason explains why and the
// progress fields are omitted.
type CustomerTierProgress struct {
	Available                bool                                 `json:"available"`
	Reason                   membershipenum.TierProgressReason    `json:"reason,omitempty"`
	QualificationMetric      *membershipenum.MembershipTierMetric `json:"qualification_metric,omitempty"`
	QualificationWindow      *QualificationWindow                 `json:"qualification_window,omitempty"`
	QualifyingSpend          *common.Money                        `json:"qualifying_spend,omitempty"`
	CurrentTier              *TierProgressTier                    `json:"current_tier,omitempty"`
	NextTier                 *TierProgressTier                    `json:"next_tier,omitempty"`
	RemainingQualifyingSpend *common.Money                        `json:"remaining_qualifying_spend,omitempty"`
	ProgressBasisPoints      *int                                 `json:"progress_basis_points,omitempty"`
	IsMaxTier                *bool                                `json:"is_max_tier,omitempty"`
	CalculatedAt             time.Time                            `json:"calculated_at"`
}

// TierProgressTier is the tier summary embedded in customer tier progress.
type TierProgressTier struct {
	TierKey             string                 `json:"tier_key"`
	Label               []common.LocalizedText `json:"label,omitempty"`
	QualifyingThreshold *common.Money          `json:"qualifying_threshold,omitempty"`
}

// QualificationWindow is the period over which qualifying spend is measured.
type QualificationWindow struct {
	StartsAt *time.Time `json:"starts_at,omitempty"`
	EndsAt   *time.Time `json:"ends_at,omitempty"`
}
