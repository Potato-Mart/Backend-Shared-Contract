package retail

import "time"

// RetailCustomerReferralProfile groups the referral-programme state of a
// retail customer.
type RetailCustomerReferralProfile struct {
	Code                      string     `json:"code,omitempty"`
	ReferrerCustomerNumber    string     `json:"referrer_customer_number,omitempty"`
	UsedReferralCodeConfirmed bool       `json:"used_referral_code_confirmed,omitempty"`
	UsedByCount               int        `json:"used_by_count,omitempty"`
	RewardCouponsIssued       int        `json:"reward_coupons_issued,omitempty"`
	RewardCouponCodes         []string   `json:"reward_coupon_codes,omitempty"`
	LastRewardIssuedAt        *time.Time `json:"last_reward_issued_at,omitempty"`
}
