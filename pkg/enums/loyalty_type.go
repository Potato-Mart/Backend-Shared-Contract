package enums

// LoyaltyLedgerReason records why a loyalty points transaction occurred.
type LoyaltyLedgerReason string

const (
	LoyaltyLedgerReasonOrder       LoyaltyLedgerReason = "ORDER"
	LoyaltyLedgerReasonBirthday    LoyaltyLedgerReason = "BIRTHDAY"
	LoyaltyLedgerReasonRedeem      LoyaltyLedgerReason = "REDEEM"
	LoyaltyLedgerReasonAdminAdjust LoyaltyLedgerReason = "ADMIN_ADJUST"
	LoyaltyLedgerReasonExpired     LoyaltyLedgerReason = "EXPIRED"
	LoyaltyLedgerReasonReferral    LoyaltyLedgerReason = "REFERRAL"
	LoyaltyLedgerReasonSignupBonus LoyaltyLedgerReason = "SIGNUP_BONUS"
	LoyaltyLedgerReasonTierUpgrade LoyaltyLedgerReason = "TIER_UPGRADE"
	LoyaltyLedgerReasonManual      LoyaltyLedgerReason = "MANUAL"
)

func (l LoyaltyLedgerReason) IsValid() bool {
	switch l {
	case LoyaltyLedgerReasonOrder, LoyaltyLedgerReasonBirthday, LoyaltyLedgerReasonRedeem,
		LoyaltyLedgerReasonAdminAdjust, LoyaltyLedgerReasonExpired, LoyaltyLedgerReasonReferral,
		LoyaltyLedgerReasonSignupBonus, LoyaltyLedgerReasonTierUpgrade, LoyaltyLedgerReasonManual:
		return true
	}
	return false
}

func (l LoyaltyLedgerReason) String() string { return string(l) }

// LoyaltyPromotionTarget restricts which customer segment a loyalty
// points-multiplier promotion applies to.
type LoyaltyPromotionTarget string

const (
	LoyaltyPromotionTargetAll          LoyaltyPromotionTarget = "ALL"
	LoyaltyPromotionTargetWholesale    LoyaltyPromotionTarget = "WHOLESALE"
	LoyaltyPromotionTargetRetail       LoyaltyPromotionTarget = "RETAIL"
	LoyaltyPromotionTargetTierSpecific LoyaltyPromotionTarget = "TIER_SPECIFIC"
)

func (l LoyaltyPromotionTarget) IsValid() bool {
	switch l {
	case LoyaltyPromotionTargetAll, LoyaltyPromotionTargetWholesale,
		LoyaltyPromotionTargetRetail, LoyaltyPromotionTargetTierSpecific:
		return true
	}
	return false
}

func (l LoyaltyPromotionTarget) String() string { return string(l) }
