package membership_enums

// TierProgressReason explains why customer tier progress is unavailable.
type TierProgressReason string

const (
	TierProgressReasonNoActiveTiers         TierProgressReason = "no_active_tiers"
	TierProgressReasonManualQualification   TierProgressReason = "manual_qualification"
	TierProgressReasonUnsupportedMetric     TierProgressReason = "unsupported_metric"
	TierProgressReasonCurrencyMismatch      TierProgressReason = "currency_mismatch"
	TierProgressReasonMembershipNotAssigned TierProgressReason = "membership_not_assigned"
)

func (r TierProgressReason) IsValid() bool {
	switch r {
	case TierProgressReasonNoActiveTiers, TierProgressReasonManualQualification,
		TierProgressReasonUnsupportedMetric, TierProgressReasonCurrencyMismatch,
		TierProgressReasonMembershipNotAssigned:
		return true
	default:
		return false
	}
}

func (r TierProgressReason) String() string { return string(r) }
