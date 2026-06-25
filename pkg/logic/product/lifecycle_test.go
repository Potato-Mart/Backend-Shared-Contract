package productlogic

import (
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v9/pkg/contracts/product"
)

func ts(t time.Time) *time.Time { return &t }

// Tests 8–12: NEW is driven solely by FirstListedAt, never by creation
// date, and lapses exactly 14 days later.
func TestIsNew(t *testing.T) {
	listed := time.Date(2026, 1, 10, 0, 0, 0, 0, time.UTC)
	created := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)

	p := product.Product{FirstListedAt: ts(listed)}
	p.CreatedAt = created // creation date must not matter

	cases := []struct {
		name string
		now  time.Time
		want bool
	}{
		{"before first listing (created but not yet listed window)", listed.Add(-time.Hour), false},
		{"at first listing", listed, true},
		{"day 13", listed.Add(13 * 24 * time.Hour), true},
		{"one second before expiry", listed.Add(LifecycleTagTTL - time.Second), true},
		{"exactly 14 days — expired", listed.Add(LifecycleTagTTL), false},
		{"well after expiry", listed.Add(30 * 24 * time.Hour), false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := IsNew(p, tc.now); got != tc.want {
				t.Errorf("IsNew(%s) = %v, want %v", tc.now, got, tc.want)
			}
		})
	}

	// Product created on Jan 1, listed Jan 10 → NEW expires Jan 24,
	// regardless of the Jan 1 creation date.
	wantExpiry := time.Date(2026, 1, 24, 0, 0, 0, 0, time.UTC)
	if got := NewExpiresAt(p); got == nil || !got.Equal(wantExpiry) {
		t.Errorf("NewExpiresAt = %v, want %v", got, wantExpiry)
	}

	// Never listed → never NEW, no expiry.
	never := product.Product{}
	never.CreatedAt = created
	if IsNew(never, created.Add(time.Hour)) {
		t.Error("unlisted product reported NEW")
	}
	if NewExpiresAt(never) != nil {
		t.Error("unlisted product has a NEW expiry")
	}
}

// Tests 13–18 (tag side): RESTOCKED is driven solely by RestockedAt and
// refreshes when the timestamp moves forward.
func TestIsRestocked(t *testing.T) {
	restocked := time.Date(2026, 3, 1, 9, 0, 0, 0, time.UTC)
	p := product.Product{RestockedAt: ts(restocked)}

	if IsRestocked(p, restocked.Add(-time.Minute)) {
		t.Error("RESTOCKED before the restock instant")
	}
	if !IsRestocked(p, restocked) {
		t.Error("not RESTOCKED at the restock instant")
	}
	if !IsRestocked(p, restocked.Add(LifecycleTagTTL-time.Second)) {
		t.Error("not RESTOCKED just before the 14-day expiry")
	}
	if IsRestocked(p, restocked.Add(LifecycleTagTTL)) {
		t.Error("still RESTOCKED at exactly 14 days")
	}

	// A later 0→>=1 transition refreshes the window (test 17).
	again := restocked.Add(20 * 24 * time.Hour)
	p.RestockedAt = ts(again)
	if !IsRestocked(p, again.Add(13*24*time.Hour)) {
		t.Error("refreshed restock window not honoured")
	}

	// Never restocked → never RESTOCKED.
	if IsRestocked(product.Product{}, restocked) {
		t.Error("product without RestockedAt reported RESTOCKED")
	}
}

func TestLifecycleTags(t *testing.T) {
	now := time.Date(2026, 5, 1, 0, 0, 0, 0, time.UTC)
	p := product.Product{
		FirstListedAt: ts(now.Add(-24 * time.Hour)),
		RestockedAt:   ts(now.Add(-time.Hour)),
	}
	got := LifecycleTags(p, now)
	if len(got) != 2 || got[0] != LifecycleTagNew || got[1] != LifecycleTagRestocked {
		t.Errorf("LifecycleTags = %v, want [NEW RESTOCKED]", got)
	}
	if tags := LifecycleTags(product.Product{}, now); len(tags) != 0 {
		t.Errorf("LifecycleTags on bare product = %v, want empty", tags)
	}
}
