package coupon_enums

// Distribution describes how the service issues coupon entitlements.
type Distribution string

const (
	DistributionSelfClaimByCode Distribution = "SELF_CLAIM_BY_CODE"
	// DistributionAssigned retains cohort-based assignment.
	DistributionAssigned Distribution = "ASSIGNED"
	// DistributionReferralAssigned reserves issuance for qualified referral awards.
	DistributionReferralAssigned Distribution = "REFERRAL_ASSIGNED"
	DistributionAutoMemberClaim  Distribution = "AUTO_MEMBER_CLAIM"
)

func (v Distribution) IsValid() bool {
	return v == DistributionSelfClaimByCode || v == DistributionAssigned || v == DistributionReferralAssigned || v == DistributionAutoMemberClaim
}
func (v Distribution) String() string { return string(v) }
