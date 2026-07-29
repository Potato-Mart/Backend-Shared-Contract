package enums_test

import (
	"testing"

	reviewenum "github.com/Potato-Mart/Backend-Shared-Contract/v20/pkg/enums/review"
)

func TestReviewEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{
			name: "reviewenum.ReviewModerationStatus",
			valid: []stringEnum{
				reviewenum.ReviewModerationStatusNotRequired,
				reviewenum.ReviewModerationStatusPending,
				reviewenum.ReviewModerationStatusApproved,
				reviewenum.ReviewModerationStatusRejected,
				reviewenum.ReviewModerationStatusSuppressed,
			},
			invalid: reviewenum.ReviewModerationStatus("__invalid__"),
		},
		{
			name: "reviewenum.ReviewRejectionReason",
			valid: []stringEnum{
				reviewenum.ReviewRejectionReasonSpam,
				reviewenum.ReviewRejectionReasonOffTopic,
				reviewenum.ReviewRejectionReasonInappropriate,
				reviewenum.ReviewRejectionReasonPersonalInformation,
				reviewenum.ReviewRejectionReasonUnsupportedLanguage,
				reviewenum.ReviewRejectionReasonOther,
			},
			invalid: reviewenum.ReviewRejectionReason("__invalid__"),
		},
		{
			name: "reviewenum.ReviewErrorCode",
			valid: []stringEnum{
				reviewenum.ReviewErrorCodeNotFound,
				reviewenum.ReviewErrorCodePurchaseRequired,
			},
			invalid: reviewenum.ReviewErrorCode("__invalid__"),
		},
	})
}

func TestReviewEnumWireValuesAreStable(t *testing.T) {
	want := map[stringEnum]string{
		reviewenum.ReviewModerationStatusNotRequired:        "not_required",
		reviewenum.ReviewModerationStatusPending:            "pending",
		reviewenum.ReviewModerationStatusApproved:           "approved",
		reviewenum.ReviewModerationStatusRejected:           "rejected",
		reviewenum.ReviewModerationStatusSuppressed:         "suppressed",
		reviewenum.ReviewRejectionReasonSpam:                "spam",
		reviewenum.ReviewRejectionReasonOffTopic:            "off_topic",
		reviewenum.ReviewRejectionReasonInappropriate:       "inappropriate",
		reviewenum.ReviewRejectionReasonPersonalInformation: "personal_information",
		reviewenum.ReviewRejectionReasonUnsupportedLanguage: "unsupported_language",
		reviewenum.ReviewRejectionReasonOther:               "other",
		reviewenum.ReviewErrorCodeNotFound:                  "REVIEW_NOT_FOUND",
		reviewenum.ReviewErrorCodePurchaseRequired:          "REVIEW_PURCHASE_REQUIRED",
	}
	for value, wire := range want {
		if got := value.String(); got != wire {
			t.Fatalf("%T.String() = %q, want stable wire value %q", value, got, wire)
		}
	}
}
