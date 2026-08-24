package wallet_enums

// PointLedgerReason records why a wallet points transaction occurred.
type PointLedgerReason string

const (
	PointLedgerReasonOrder        PointLedgerReason = "ORDER"
	PointLedgerReasonBirthday     PointLedgerReason = "BIRTHDAY"
	PointLedgerReasonRedeem       PointLedgerReason = "REDEEM"
	PointLedgerReasonRefund       PointLedgerReason = "REFUND"
	PointLedgerReasonRewardRedeem PointLedgerReason = "REWARD_REDEEM"
	PointLedgerReasonAdminAdjust  PointLedgerReason = "ADMIN_ADJUST"
	PointLedgerReasonExpired      PointLedgerReason = "EXPIRED"
	PointLedgerReasonReferral     PointLedgerReason = "REFERRAL"
	PointLedgerReasonSignupBonus  PointLedgerReason = "SIGNUP_BONUS"
	PointLedgerReasonTierUpgrade  PointLedgerReason = "TIER_UPGRADE"
	PointLedgerReasonDebtIncurred PointLedgerReason = "DEBT_INCURRED"
	PointLedgerReasonDebtRepaid   PointLedgerReason = "DEBT_REPAID"
	PointLedgerReasonManual       PointLedgerReason = "MANUAL"
)

func (r PointLedgerReason) IsValid() bool {
	switch r {
	case PointLedgerReasonOrder, PointLedgerReasonBirthday, PointLedgerReasonRedeem, PointLedgerReasonRefund, PointLedgerReasonRewardRedeem, PointLedgerReasonAdminAdjust, PointLedgerReasonExpired, PointLedgerReasonReferral, PointLedgerReasonSignupBonus, PointLedgerReasonTierUpgrade, PointLedgerReasonDebtIncurred, PointLedgerReasonDebtRepaid, PointLedgerReasonManual:
		return true
	}
	return false
}
func (r PointLedgerReason) String() string { return string(r) }
