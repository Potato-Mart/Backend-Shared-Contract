package template

import "time"

// TranslationReview is server-stamped evidence that a human confirmed the final
// target text against a particular source. Fingerprints are independent of
// draft counters and include the applicable document/token interpretation.
// Notification binds this evidence to the enclosing country/resource, verifies
// source and target freshness, and never trusts client-supplied reviewer fields.
// Translated text edits invalidate that target; source/structure/token changes
// invalidate both targets. Visual-only edits require a new render preview,
// without invalidating otherwise-current translation review.
type TranslationReview struct {
	SourceFingerprint string    `json:"source_fingerprint"`
	TargetFingerprint string    `json:"target_fingerprint"`
	ReviewedAt        time.Time `json:"reviewed_at"`
	ReviewedBy        string    `json:"reviewed_by"`
}
