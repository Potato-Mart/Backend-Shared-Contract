package review_enums

// ReviewQualificationKind records the authorization evidence persisted by the
// service when a customer submits a review.
type ReviewQualificationKind string

const (
	ReviewQualificationKindNotRequired      ReviewQualificationKind = "not_required"
	ReviewQualificationKindVerifiedOrder    ReviewQualificationKind = "verified_order"
	ReviewQualificationKindVerifiedPurchase ReviewQualificationKind = "verified_purchase"
)

func (k ReviewQualificationKind) String() string { return string(k) }

func (k ReviewQualificationKind) IsValid() bool {
	switch k {
	case ReviewQualificationKindNotRequired, ReviewQualificationKindVerifiedOrder,
		ReviewQualificationKindVerifiedPurchase:
		return true
	default:
		return false
	}
}
