package enums_test

import (
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/review/review_enums"
)

func TestReviewEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{
			name: "review.ReviewType",
			valid: []stringEnum{
				review_enums.ReviewTypeOrder,
				review_enums.ReviewTypeProduct,
				review_enums.ReviewTypeCampaign,
				review_enums.ReviewTypeAppFeedback,
			},
			invalid: review_enums.ReviewType("__invalid__"),
		},
		{
			name: "review.ReviewStatus",
			valid: []stringEnum{
				review_enums.ReviewStatusDraft,
				review_enums.ReviewStatusPendingApproval,
				review_enums.ReviewStatusPublished,
				review_enums.ReviewStatusRejected,
				review_enums.ReviewStatusSuppressed,
			},
			invalid: review_enums.ReviewStatus("__invalid__"),
		},
		{
			name: "review.ReviewQualificationKind",
			valid: []stringEnum{
				review_enums.ReviewQualificationKindNotRequired,
				review_enums.ReviewQualificationKindVerifiedOrder,
				review_enums.ReviewQualificationKindVerifiedPurchase,
			},
			invalid: review_enums.ReviewQualificationKind("__invalid__"),
		},
		{
			name: "reviewenum.ReviewRejectionReason",
			valid: []stringEnum{
				review_enums.ReviewRejectionReasonSpam,
				review_enums.ReviewRejectionReasonOffTopic,
				review_enums.ReviewRejectionReasonInappropriate,
				review_enums.ReviewRejectionReasonPersonalInformation,
				review_enums.ReviewRejectionReasonUnsupportedLanguage,
				review_enums.ReviewRejectionReasonOther,
			},
			invalid: review_enums.ReviewRejectionReason("__invalid__"),
		},
		{
			name: "reviewenum.ReviewErrorCode",
			valid: []stringEnum{
				review_enums.ReviewErrorCodeNotFound,
				review_enums.ReviewErrorCodeQualificationRequired,
			},
			invalid: review_enums.ReviewErrorCode("__invalid__"),
		},
	})
}
