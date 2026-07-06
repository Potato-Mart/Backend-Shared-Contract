// Package productlogic holds the computed lifecycle/display helpers for
// product.Product. These are read-time derivations (never stored); they live
// outside pkg/contracts so the contract files stay pure data shapes.
package productlogic

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v12/pkg/contracts/product"
	productenum "github.com/Potato-Mart/Backend-Shared-Contract/v12/pkg/enums/product"
)

// Display states are the read-time, computed presentation states layered on top
// of an active product. They are never stored (storing them would require a
// background job and could drift); compute them with DisplayStatus wherever a
// live product is shown.
const (
	// DisplayStatusNew marks an active product still inside its NEW (新品)
	// window. Takes precedence over restocked and out_of_stock.
	DisplayStatusNew = "new"
	// DisplayStatusRestocked marks an active product still inside its RESTOCKED
	// (補貨) window.
	DisplayStatusRestocked = "restocked"
	// DisplayStatusOutOfStock marks an active product with no sellable stock (缺貨).
	DisplayStatusOutOfStock = "out_of_stock"
	// DisplayStatusActive marks an active product with nothing special to flag.
	DisplayStatusActive = "active"
)

// Lifecycle tags are computed, never stored: a product is NEW for 14 days from
// FirstListedAt and RESTOCKED for 14 days from RestockedAt. Deriving the tags
// from the two timestamps keeps expiry deterministic and testable, needs no
// background job, and cannot drift — the same instant always produces the same
// answer everywhere. All timestamps are stored and compared in UTC.
const (
	// LifecycleTagNew marks a 新品: listed for the first time within the last
	// LifecycleTagTTL.
	LifecycleTagNew = "NEW"
	// LifecycleTagRestocked marks a 補貨 product: total sellable stock went
	// 0 → >=1 within the last LifecycleTagTTL.
	LifecycleTagRestocked = "RESTOCKED"
	// LifecycleTagTTL is how long either tag stays valid.
	LifecycleTagTTL = 14 * 24 * time.Hour
)

// IsNew reports whether the product carries the NEW tag at the given instant. A
// product that has never been listed is never NEW: the window starts at
// FirstListedAt, not at creation/import time.
func IsNew(p product.Product, now time.Time) bool {
	if p.FirstListedAt == nil || p.FirstListedAt.IsZero() {
		return false
	}
	return !now.Before(*p.FirstListedAt) && now.Before(p.FirstListedAt.Add(LifecycleTagTTL))
}

// NewExpiresAt returns the instant the NEW tag lapses, or nil if the product has
// never been listed.
func NewExpiresAt(p product.Product) *time.Time {
	if p.FirstListedAt == nil || p.FirstListedAt.IsZero() {
		return nil
	}
	t := p.FirstListedAt.Add(LifecycleTagTTL)
	return &t
}

// IsRestocked reports whether the product carries the RESTOCKED tag at the given
// instant.
func IsRestocked(p product.Product, now time.Time) bool {
	if p.RestockedAt == nil || p.RestockedAt.IsZero() {
		return false
	}
	return !now.Before(*p.RestockedAt) && now.Before(p.RestockedAt.Add(LifecycleTagTTL))
}

// RestockedExpiresAt returns the instant the RESTOCKED tag lapses, or nil if the
// product has never restocked.
func RestockedExpiresAt(p product.Product) *time.Time {
	if p.RestockedAt == nil || p.RestockedAt.IsZero() {
		return nil
	}
	t := p.RestockedAt.Add(LifecycleTagTTL)
	return &t
}

// DisplayStatus is the read-time merge of admin status + recency + stock, for
// badges/labels on live product views. It is never persisted: an
// archived/draft/discontinued product passes its stored status straight
// through, while an active product resolves to new → restocked → out_of_stock →
// active by precedence. Because the NEW/RESTOCKED windows lapse on their own and
// out_of_stock clears once stock returns, an active product reverts to "active"
// automatically — no background job.
func DisplayStatus(p product.Product, now time.Time) string {
	if p.Status != productenum.ProductStatusActive {
		return string(p.Status)
	}
	switch {
	case IsNew(p, now):
		return DisplayStatusNew
	case IsRestocked(p, now):
		return DisplayStatusRestocked
	case p.CurrentStock <= 0:
		return DisplayStatusOutOfStock
	default:
		return DisplayStatusActive
	}
}

// LifecycleTags returns the currently valid lifecycle tags, in a stable order
// (NEW before RESTOCKED). Empty when neither applies.
func LifecycleTags(p product.Product, now time.Time) []string {
	var tags []string
	if IsNew(p, now) {
		tags = append(tags, LifecycleTagNew)
	}
	if IsRestocked(p, now) {
		tags = append(tags, LifecycleTagRestocked)
	}
	return tags
}
