package membership

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/pricing/membership/membership_enums"
)

// CustomerTierProgress is the customer-safe projection of progress toward the
// next membership tier. When Available is false, Reason explains why and the
// progress fields are omitted.
//
// NextTier, RemainingQualifyingSpend, ProgressBasisPoints and IsMaxTier are
// relative to the customer's effective tier: the earned or the assigned tier,
// whichever ranks higher. An assigned member therefore sees the tier above
// the assigned one, and IsMaxTier is true when the effective tier is the top
// tier.
type CustomerTierProgress struct {
	Available                bool                                   `json:"available"`
	Reason                   membership_enums.TierProgressReason    `json:"reason,omitempty"`
	QualificationMetric      *membership_enums.MembershipTierMetric `json:"qualification_metric,omitempty"`
	QualificationWindow      *QualificationWindow                   `json:"qualification_window,omitempty"`
	QualifyingSpend          *money.Money                           `json:"qualifying_spend,omitempty"`
	CurrentTier              *TierProgressTier                      `json:"current_tier,omitempty"`
	NextTier                 *TierProgressTier                      `json:"next_tier,omitempty"`
	RemainingQualifyingSpend *money.Money                           `json:"remaining_qualifying_spend,omitempty"`
	// ProgressBasisPoints measures qualifying spend from $0 toward the threshold
	// of the next tier above the effective tier: floor(spend / next threshold
	// * 10000), clamped to 0..10000; max tier is 10000.
	// TierProgressBasisPoints is optional within-tier progress, using the configured current threshold as its floor.
	TierProgressBasisPoints *int      `json:"tier_progress_basis_points,omitempty"`
	ProgressBasisPoints     *int      `json:"progress_basis_points,omitempty"`
	IsMaxTier               *bool     `json:"is_max_tier,omitempty"`
	CalculatedAt            time.Time `json:"calculated_at"`
}
