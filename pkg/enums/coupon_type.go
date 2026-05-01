package enums

// CouponAppliesTo restricts which products or categories a coupon covers.
type CouponAppliesTo string

const (
	CouponAppliesToAll                CouponAppliesTo = "all"
	CouponAppliesToSpecificProducts   CouponAppliesTo = "specific_products"
	CouponAppliesToSpecificCategories CouponAppliesTo = "specific_categories"
)

func (c CouponAppliesTo) IsValid() bool {
	switch c {
	case CouponAppliesToAll, CouponAppliesToSpecificProducts, CouponAppliesToSpecificCategories:
		return true
	}
	return false
}

func (c CouponAppliesTo) String() string { return string(c) }

// CouponSource records how a customer_coupon assignment was created.
type CouponSource string

const (
	CouponSourceManual      CouponSource = "MANUAL"
	CouponSourceRFMComeback CouponSource = "RFM_COMEBACK"
	CouponSourceBirthday    CouponSource = "BIRTHDAY"
	CouponSourceReferral    CouponSource = "REFERRAL"
	CouponSourceSignupBonus CouponSource = "SIGNUP_BONUS"
	CouponSourceCampaign    CouponSource = "CAMPAIGN"
)

func (c CouponSource) IsValid() bool {
	switch c {
	case CouponSourceManual, CouponSourceRFMComeback, CouponSourceBirthday,
		CouponSourceReferral, CouponSourceSignupBonus, CouponSourceCampaign:
		return true
	}
	return false
}

func (c CouponSource) String() string { return string(c) }
