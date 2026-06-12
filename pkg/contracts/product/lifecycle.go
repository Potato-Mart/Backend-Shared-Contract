package product

import "time"

// Lifecycle tags are computed, never stored: a product is NEW for 14
// days from FirstListedAt and RESTOCKED for 14 days from RestockedAt.
// Deriving the tags from the two timestamps keeps expiry deterministic
// and testable, needs no background job, and cannot drift — the same
// instant always produces the same answer everywhere.
//
// All timestamps are stored and compared in UTC.
const (
	// LifecycleTagNew marks a 新品: listed for the first time within
	// the last LifecycleTagTTL.
	LifecycleTagNew = "NEW"
	// LifecycleTagRestocked marks a 補貨 product: total sellable stock
	// went 0 → >=1 within the last LifecycleTagTTL.
	LifecycleTagRestocked = "RESTOCKED"
	// LifecycleTagTTL is how long either tag stays valid.
	LifecycleTagTTL = 14 * 24 * time.Hour
)

// IsNew reports whether the product carries the NEW tag at the given
// instant. A product that has never been listed is never NEW: the
// window starts at FirstListedAt, not at creation/import time.
func (p Product) IsNew(now time.Time) bool {
	if p.FirstListedAt == nil || p.FirstListedAt.IsZero() {
		return false
	}
	return !now.Before(*p.FirstListedAt) && now.Before(p.FirstListedAt.Add(LifecycleTagTTL))
}

// NewExpiresAt returns the instant the NEW tag lapses, or nil if the
// product has never been listed.
func (p Product) NewExpiresAt() *time.Time {
	if p.FirstListedAt == nil || p.FirstListedAt.IsZero() {
		return nil
	}
	t := p.FirstListedAt.Add(LifecycleTagTTL)
	return &t
}

// IsRestocked reports whether the product carries the RESTOCKED tag at
// the given instant.
func (p Product) IsRestocked(now time.Time) bool {
	if p.RestockedAt == nil || p.RestockedAt.IsZero() {
		return false
	}
	return !now.Before(*p.RestockedAt) && now.Before(p.RestockedAt.Add(LifecycleTagTTL))
}

// RestockedExpiresAt returns the instant the RESTOCKED tag lapses, or
// nil if the product has never restocked.
func (p Product) RestockedExpiresAt() *time.Time {
	if p.RestockedAt == nil || p.RestockedAt.IsZero() {
		return nil
	}
	t := p.RestockedAt.Add(LifecycleTagTTL)
	return &t
}

// LifecycleTags returns the currently valid lifecycle tags, in a stable
// order (NEW before RESTOCKED). Empty when neither applies.
func (p Product) LifecycleTags(now time.Time) []string {
	var tags []string
	if p.IsNew(now) {
		tags = append(tags, LifecycleTagNew)
	}
	if p.IsRestocked(now) {
		tags = append(tags, LifecycleTagRestocked)
	}
	return tags
}
