package membershipenum

// MembershipOwnerType identifies the business entity that owns a membership
// wallet and tier.
type MembershipOwnerType string

const (
	MembershipOwnerTypeRetailCustomer        MembershipOwnerType = "retail_customer"
	MembershipOwnerTypeWholesaleOrganisation MembershipOwnerType = "wholesale_organisation"
)

func (m MembershipOwnerType) IsValid() bool {
	switch m {
	case MembershipOwnerTypeRetailCustomer, MembershipOwnerTypeWholesaleOrganisation:
		return true
	}
	return false
}

func (m MembershipOwnerType) String() string { return string(m) }

// MembershipAccountStatus describes the lifecycle of a global membership
// account.
type MembershipAccountStatus string

const (
	MembershipAccountStatusActive    MembershipAccountStatus = "active"
	MembershipAccountStatusSuspended MembershipAccountStatus = "suspended"
	MembershipAccountStatusClosed    MembershipAccountStatus = "closed"
)

func (m MembershipAccountStatus) IsValid() bool {
	switch m {
	case MembershipAccountStatusActive, MembershipAccountStatusSuspended, MembershipAccountStatusClosed:
		return true
	}
	return false
}

func (m MembershipAccountStatus) String() string { return string(m) }

// MembershipTierMetric controls which value qualifies a member for a tier.
type MembershipTierMetric string

const (
	MembershipTierMetricAnnualSpend   MembershipTierMetric = "annual_spend"
	MembershipTierMetricLifetimeSpend MembershipTierMetric = "lifetime_spend"
	MembershipTierMetricManual        MembershipTierMetric = "manual"
)

func (m MembershipTierMetric) IsValid() bool {
	switch m {
	case MembershipTierMetricAnnualSpend, MembershipTierMetricLifetimeSpend, MembershipTierMetricManual:
		return true
	}
	return false
}

func (m MembershipTierMetric) String() string { return string(m) }

// MembershipPointReason records why a membership points transaction occurred.
type MembershipPointReason string

const (
	MembershipPointReasonOrder        MembershipPointReason = "ORDER"
	MembershipPointReasonBirthday     MembershipPointReason = "BIRTHDAY"
	MembershipPointReasonRedeem       MembershipPointReason = "REDEEM"
	MembershipPointReasonRewardRedeem MembershipPointReason = "REWARD_REDEEM"
	MembershipPointReasonAdminAdjust  MembershipPointReason = "ADMIN_ADJUST"
	MembershipPointReasonExpired      MembershipPointReason = "EXPIRED"
	MembershipPointReasonReferral     MembershipPointReason = "REFERRAL"
	MembershipPointReasonSignupBonus  MembershipPointReason = "SIGNUP_BONUS"
	MembershipPointReasonTierUpgrade  MembershipPointReason = "TIER_UPGRADE"
	MembershipPointReasonManual       MembershipPointReason = "MANUAL"
)

func (m MembershipPointReason) IsValid() bool {
	switch m {
	case MembershipPointReasonOrder, MembershipPointReasonBirthday, MembershipPointReasonRedeem,
		MembershipPointReasonRewardRedeem, MembershipPointReasonAdminAdjust, MembershipPointReasonExpired,
		MembershipPointReasonReferral, MembershipPointReasonSignupBonus, MembershipPointReasonTierUpgrade,
		MembershipPointReasonManual:
		return true
	}
	return false
}

func (m MembershipPointReason) String() string { return string(m) }

// MembershipPromotionTarget restricts which membership segment a points
// multiplier applies to.
type MembershipPromotionTarget string

const (
	MembershipPromotionTargetAll          MembershipPromotionTarget = "ALL"
	MembershipPromotionTargetWholesale    MembershipPromotionTarget = "WHOLESALE"
	MembershipPromotionTargetRetail       MembershipPromotionTarget = "RETAIL"
	MembershipPromotionTargetTierSpecific MembershipPromotionTarget = "TIER_SPECIFIC"
)

func (m MembershipPromotionTarget) IsValid() bool {
	switch m {
	case MembershipPromotionTargetAll, MembershipPromotionTargetWholesale,
		MembershipPromotionTargetRetail, MembershipPromotionTargetTierSpecific:
		return true
	}
	return false
}

func (m MembershipPromotionTarget) String() string { return string(m) }

// PointReservationStatus describes a points reservation before it becomes a
// committed ledger spend.
type PointReservationStatus string

const (
	PointReservationStatusReserved  PointReservationStatus = "RESERVED"
	PointReservationStatusCommitted PointReservationStatus = "COMMITTED"
	PointReservationStatusCancelled PointReservationStatus = "CANCELLED"
	PointReservationStatusExpired   PointReservationStatus = "EXPIRED"
)

func (p PointReservationStatus) IsValid() bool {
	switch p {
	case PointReservationStatusReserved, PointReservationStatusCommitted,
		PointReservationStatusCancelled, PointReservationStatusExpired:
		return true
	}
	return false
}

func (p PointReservationStatus) String() string { return string(p) }

// MembershipRedemptionType identifies what a points spend is reserving.
type MembershipRedemptionType string

const (
	MembershipRedemptionTypeCheckoutDiscount MembershipRedemptionType = "CHECKOUT_DISCOUNT"
	MembershipRedemptionTypeRewardCatalog    MembershipRedemptionType = "REWARD_CATALOG"
)

func (m MembershipRedemptionType) IsValid() bool {
	switch m {
	case MembershipRedemptionTypeCheckoutDiscount, MembershipRedemptionTypeRewardCatalog:
		return true
	}
	return false
}

func (m MembershipRedemptionType) String() string { return string(m) }

// MembershipRewardType describes what a catalog reward gives the member.
type MembershipRewardType string

const (
	MembershipRewardTypeOrderDiscount MembershipRewardType = "ORDER_DISCOUNT"
	MembershipRewardTypeProduct       MembershipRewardType = "PRODUCT"
	MembershipRewardTypeFreeShipping  MembershipRewardType = "FREE_SHIPPING"
	MembershipRewardTypeVoucher       MembershipRewardType = "VOUCHER"
)

func (m MembershipRewardType) IsValid() bool {
	switch m {
	case MembershipRewardTypeOrderDiscount, MembershipRewardTypeProduct,
		MembershipRewardTypeFreeShipping, MembershipRewardTypeVoucher:
		return true
	}
	return false
}

func (m MembershipRewardType) String() string { return string(m) }

// MembershipRewardRedemptionStatus describes the lifecycle of a reward
// redemption.
type MembershipRewardRedemptionStatus string

const (
	MembershipRewardRedemptionStatusReserved  MembershipRewardRedemptionStatus = "RESERVED"
	MembershipRewardRedemptionStatusRedeemed  MembershipRewardRedemptionStatus = "REDEEMED"
	MembershipRewardRedemptionStatusCancelled MembershipRewardRedemptionStatus = "CANCELLED"
	MembershipRewardRedemptionStatusExpired   MembershipRewardRedemptionStatus = "EXPIRED"
)

func (m MembershipRewardRedemptionStatus) IsValid() bool {
	switch m {
	case MembershipRewardRedemptionStatusReserved, MembershipRewardRedemptionStatusRedeemed,
		MembershipRewardRedemptionStatusCancelled, MembershipRewardRedemptionStatusExpired:
		return true
	}
	return false
}

func (m MembershipRewardRedemptionStatus) String() string { return string(m) }
