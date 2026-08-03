package review_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/contracts/review"
	reviewenum "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/enums/review"
)

func TestRatingSummaryJSONHasFiveOrderedBuckets(t *testing.T) {
	summary := review.RatingSummary{
		AverageScore:             4.25,
		RatingCount:              12,
		PublishedTextReviewCount: 7,
		Distribution: [5]review.RatingDistributionBucket{
			{Score: 1, Count: 1},
			{Score: 2, Count: 0},
			{Score: 3, Count: 2},
			{Score: 4, Count: 2},
			{Score: 5, Count: 7},
		},
	}

	payload, err := json.Marshal(summary)
	if err != nil {
		t.Fatalf("marshal rating summary: %v", err)
	}

	var decoded struct {
		AverageScore             float64                           `json:"average_score"`
		RatingCount              int64                             `json:"rating_count"`
		PublishedTextReviewCount int64                             `json:"published_text_review_count"`
		Distribution             []review.RatingDistributionBucket `json:"distribution"`
	}
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal rating summary: %v", err)
	}
	if decoded.AverageScore != 4.25 || decoded.RatingCount != 12 || decoded.PublishedTextReviewCount != 7 {
		t.Fatalf("rating summary did not round-trip: %+v", decoded)
	}
	if len(decoded.Distribution) != 5 {
		t.Fatalf("distribution length = %d, want 5", len(decoded.Distribution))
	}
	for index, bucket := range decoded.Distribution {
		if wantScore := index + 1; bucket.Score != wantScore {
			t.Fatalf("distribution[%d].score = %d, want %d", index, bucket.Score, wantScore)
		}
	}
}

func TestProductReviewProjectionsKeepIdentityAndModerationSeparated(t *testing.T) {
	now := time.Date(2026, 7, 20, 1, 2, 3, 0, time.UTC)
	public := review.ProductReview{
		ID:               "review_1",
		ProductSKUCode:   "SKU-1",
		Score:            5,
		Title:            "Great",
		Body:             "Approved public copy",
		Locale:           "en-AU",
		VerifiedPurchase: true,
		CreatedAt:        now,
		UpdatedAt:        now,
	}

	publicJSON := marshalReviewJSON(t, public)
	for _, want := range []string{
		`"product_sku_code":"SKU-1"`,
		`"verified_purchase":true`,
		`"created_at":"2026-07-20T01:02:03Z"`,
	} {
		if !strings.Contains(publicJSON, want) {
			t.Fatalf("public review JSON = %s, want %s", publicJSON, want)
		}
	}
	assertReviewJSONOmits(t, publicJSON,
		"customer_id", "user_id", "account_id", "original_title", "original_body",
		"moderation_status", "rejection_reason", "moderation_note", "moderated_at")

	mine := review.MyProductReview{
		ID:               public.ID,
		ProductSKUCode:   public.ProductSKUCode,
		Score:            public.Score,
		Title:            public.Title,
		Body:             public.Body,
		OriginalTitle:    "Original title",
		OriginalBody:     "Original body",
		Locale:           public.Locale,
		VerifiedPurchase: public.VerifiedPurchase,
		ModerationStatus: reviewenum.ReviewModerationStatusRejected,
		RejectionReason:  reviewenum.ReviewRejectionReasonOffTopic,
		CreatedAt:        public.CreatedAt,
		UpdatedAt:        public.UpdatedAt,
	}
	mineJSON := marshalReviewJSON(t, mine)
	for _, want := range []string{
		`"original_title":"Original title"`,
		`"original_body":"Original body"`,
		`"moderation_status":"rejected"`,
		`"rejection_reason":"off_topic"`,
	} {
		if !strings.Contains(mineJSON, want) {
			t.Fatalf("my review JSON = %s, want %s", mineJSON, want)
		}
	}
	assertReviewJSONOmits(t, mineJSON,
		"customer_id", "user_id", "account_id", "moderation_note", "moderated_at")

	moderation := review.ProductReviewModeration{
		ID:               mine.ID,
		ProductSKUCode:   mine.ProductSKUCode,
		Score:            mine.Score,
		Title:            mine.Title,
		Body:             mine.Body,
		OriginalTitle:    mine.OriginalTitle,
		OriginalBody:     mine.OriginalBody,
		Locale:           mine.Locale,
		VerifiedPurchase: mine.VerifiedPurchase,
		ModerationStatus: mine.ModerationStatus,
		RejectionReason:  mine.RejectionReason,
		ModerationNote:   "Policy guidance only; no customer identity.",
		CreatedAt:        mine.CreatedAt,
		UpdatedAt:        mine.UpdatedAt,
		ModeratedAt:      &now,
	}
	moderationJSON := marshalReviewJSON(t, moderation)
	for _, want := range []string{
		`"moderation_note":"Policy guidance only; no customer identity."`,
		`"moderated_at":"2026-07-20T01:02:03Z"`,
	} {
		if !strings.Contains(moderationJSON, want) {
			t.Fatalf("moderation review JSON = %s, want %s", moderationJSON, want)
		}
	}
	assertReviewJSONOmits(t, moderationJSON, "customer_id", "user_id", "account_id", "moderated_by")
}

func TestOptionalReviewFieldsAreOmitted(t *testing.T) {
	payload := marshalReviewJSON(t, review.MyProductReview{
		ModerationStatus: reviewenum.ReviewModerationStatusNotRequired,
	})
	assertReviewJSONOmits(t, payload,
		"title", "body", "locale", "original_title", "original_body", "rejection_reason")
}

func marshalReviewJSON(t *testing.T, value any) string {
	t.Helper()
	payload, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("marshal review projection: %v", err)
	}
	return string(payload)
}

func assertReviewJSONOmits(t *testing.T, payload string, keys ...string) {
	t.Helper()
	for _, key := range keys {
		if strings.Contains(payload, `"`+key+`"`) {
			t.Fatalf("review JSON unexpectedly exposes %q: %s", key, payload)
		}
	}
}
