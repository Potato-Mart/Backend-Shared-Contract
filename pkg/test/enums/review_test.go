package enums_test

import (
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v25/pkg/contracts/supply/review/review_enums"
)

func TestReviewEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{
			name: "reviewenum.ReviewModerationStatus",
			valid: []stringEnum{
				review_enums.ReviewModerationStatusNotRequired,
				review_enums.ReviewModerationStatusPending,
				review_enums.ReviewModerationStatusApproved,
				review_enums.ReviewModerationStatusRejected,
				review_enums.ReviewModerationStatusSuppressed,
			},
			invalid: review_enums.ReviewModerationStatus("__invalid__"),
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
				review_enums.ReviewErrorCodePurchaseRequired,
			},
			invalid: review_enums.ReviewErrorCode("__invalid__"),
		},
	})
}

func TestReviewEnumWireValuesAreStable(t *testing.T) {
	want := map[stringEnum]string{
		review_enums.ReviewModerationStatusNotRequired:        "not_required",
		review_enums.ReviewModerationStatusPending:            "pending",
		review_enums.ReviewModerationStatusApproved:           "approved",
		review_enums.ReviewModerationStatusRejected:           "rejected",
		review_enums.ReviewModerationStatusSuppressed:         "suppressed",
		review_enums.ReviewRejectionReasonSpam:                "spam",
		review_enums.ReviewRejectionReasonOffTopic:            "off_topic",
		review_enums.ReviewRejectionReasonInappropriate:       "inappropriate",
		review_enums.ReviewRejectionReasonPersonalInformation: "personal_information",
		review_enums.ReviewRejectionReasonUnsupportedLanguage: "unsupported_language",
		review_enums.ReviewRejectionReasonOther:               "other",
		review_enums.ReviewErrorCodeNotFound:                  "REVIEW_NOT_FOUND",
		review_enums.ReviewErrorCodePurchaseRequired:          "REVIEW_PURCHASE_REQUIRED",
	}
	for value, wire := range want {
		if got := value.String(); got != wire {
			t.Fatalf("%T.String() = %q, want stable wire value %q", value, got, wire)
		}
	}
}
