package review_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/localization"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/catalogue/review"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/catalogue/review/review_enums"
)

func TestReviewProjectionsKeepPrivateAndModerationFieldsOutOfCustomerPayloads(t *testing.T) {
	now := time.Date(2026, 8, 24, 1, 2, 3, 0, time.UTC)
	published := review.PublishedReview{
		ID: "review_1",
		Subject: review.PublishedReviewSubject{
			Type:        review_enums.ReviewTypeOrder,
			DisplayName: []localization.LocalizedName{{Language: "en-AU", Name: "Recent delivery"}},
		},
		ReviewerDisplayName: "Pat",
		SubmissionChannel:   "online",
		Content: review.ReviewContent{
			Body:   []localization.LocalizedText{{Language: "en-AU", Text: "Fresh potatoes."}},
			Score:  5,
			Images: []security.ObjectMedia{{Code: "review-image-1", URL: "https://cdn.example/image"}},
		},
		Revision:    3,
		CreatedAt:   now,
		UpdatedAt:   now,
		PublishedAt: now,
	}

	internal := review.Review{
		ID:                  published.ID,
		Subject:             review.ReviewSubject{Type: review_enums.ReviewTypeOrder, Reference: "ORD-123"},
		CustomerNumber:      "CUS-123",
		ReviewerDisplayName: published.ReviewerDisplayName,
		SubmissionChannel:   published.SubmissionChannel,
		Qualification: review.ReviewQualification{
			Kind:       review_enums.ReviewQualificationKindVerifiedOrder,
			Reference:  "ORD-123",
			VerifiedAt: &now,
		},
		Content:        published.Content,
		Status:         review_enums.ReviewStatusPublished,
		Revision:       published.Revision,
		LastPublished:  &published,
		ModerationNote: "Internal review workflow note",
		History:        []security.HistoryEntry{{Type: "review.published", OccurredAt: now}},
		AuditFields: audit.AuditFields{
			CreatedAt: now,
			UpdatedAt: now,
		},
		DataProtectionFields: security.DataProtectionFields{
			ContainsPII: true,
		},
	}

	internalJSON := marshalReviewJSON(t, internal)
	assertReviewJSONContains(t, internalJSON,
		`"customer_number":"CUS-123"`,
		`"reference":"ORD-123"`,
		`"moderation_note":"Internal review workflow note"`,
		`"contains_pii":true`)

	publicJSON := marshalReviewJSON(t, published)
	assertReviewJSONContains(t, publicJSON,
		`"reviewer_display_name":"Pat"`,
		`"score":5`,
		`"published_at":"2026-08-24T01:02:03Z"`)
	assertReviewJSONOmits(t, publicJSON,
		"customer_number", "qualification", "moderation_note", "history", "contains_pii", "reference")

	mine := review.MyReview{
		ID:                  internal.ID,
		Subject:             internal.Subject,
		ReviewerDisplayName: internal.ReviewerDisplayName,
		SubmissionChannel:   internal.SubmissionChannel,
		Qualification:       internal.Qualification,
		Content:             internal.Content,
		Status:              review_enums.ReviewStatusPendingApproval,
		Revision:            4,
		LastPublished:       internal.LastPublished,
		CreatedAt:           now,
		UpdatedAt:           now,
	}
	mineJSON := marshalReviewJSON(t, mine)
	assertReviewJSONContains(t, mineJSON, `"status":"pending_approval"`, `"last_published"`)
	assertReviewJSONOmits(t, mineJSON,
		"customer_number", "moderation_note", "history", "contains_pii")
}

func TestRatingSummaryJSONHasFiveScoreCounts(t *testing.T) {
	summary := review.RatingSummary{
		Subject:      review.PublishedReviewSubject{Type: review_enums.ReviewTypeProduct, Reference: "SKU-1"},
		AverageScore: 4.25,
		RatingCount:  12,
		ScoreCounts:  [5]int64{1, 0, 2, 2, 7},
	}

	payload, err := json.Marshal(summary)
	if err != nil {
		t.Fatalf("marshal rating summary: %v", err)
	}

	var decoded struct {
		ScoreCounts []int64 `json:"score_counts"`
	}
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal rating summary: %v", err)
	}
	if len(decoded.ScoreCounts) != 5 {
		t.Fatalf("score count length = %d, want 5", len(decoded.ScoreCounts))
	}
	if got, want := decoded.ScoreCounts[4], int64(7); got != want {
		t.Fatalf("five-star count = %d, want %d", got, want)
	}
}

func TestReviewEnumsUseCanonicalValues(t *testing.T) {
	for _, value := range []review_enums.ReviewType{
		review_enums.ReviewTypeOrder,
		review_enums.ReviewTypeProduct,
		review_enums.ReviewTypeCampaign,
		review_enums.ReviewTypeAppFeedback,
	} {
		if !value.IsValid() {
			t.Fatalf("review type %q is invalid", value)
		}
	}
	if review_enums.ReviewStatus("modified").IsValid() {
		t.Fatal("modified must be represented by revision/history, not a status")
	}
	if !review_enums.ReviewErrorCodeQualificationRequired.IsValid() {
		t.Fatal("generic qualification-required error must be valid")
	}
}

func marshalReviewJSON(t *testing.T, value any) string {
	t.Helper()
	payload, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("marshal review value: %v", err)
	}
	return string(payload)
}

func assertReviewJSONContains(t *testing.T, payload string, values ...string) {
	t.Helper()
	for _, value := range values {
		if !strings.Contains(payload, value) {
			t.Fatalf("review JSON = %s, want %s", payload, value)
		}
	}
}

func assertReviewJSONOmits(t *testing.T, payload string, keys ...string) {
	t.Helper()
	for _, key := range keys {
		if strings.Contains(payload, `"`+key+`"`) {
			t.Fatalf("review JSON unexpectedly exposes %q: %s", key, payload)
		}
	}
}
