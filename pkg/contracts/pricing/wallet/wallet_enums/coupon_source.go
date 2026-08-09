package wallet_enums

// CouponSource records how a CouponAssignment was created.
type CouponSource string

const (
	CouponSourceManual      CouponSource = "MANUAL"
	CouponSourceRFMComeback CouponSource = "RFM_COMEBACK"
	CouponSourceBirthday    CouponSource = "BIRTHDAY"
	CouponSourceReferral    CouponSource = "REFERRAL"
	CouponSourceSignupBonus CouponSource = "SIGNUP_BONUS"
	CouponSourceCampaign    CouponSource = "CAMPAIGN"
)

// IsValid reports whether c is a known CouponSource.
func (c CouponSource) IsValid() bool {
	switch c {
	case CouponSourceManual, CouponSourceRFMComeback, CouponSourceBirthday,
		CouponSourceReferral, CouponSourceSignupBonus, CouponSourceCampaign:
		return true
	}
	return false
}

// String returns the wire value for c.
func (c CouponSource) String() string { return string(c) }
