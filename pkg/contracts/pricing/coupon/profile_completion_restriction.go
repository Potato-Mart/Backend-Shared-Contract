package coupon

// ProfileCompletionRestriction requires authoritative 100% profile completion
// when active. The service enforces one entitlement per immutable coupon ID and
// eligible customer, including retries and subsequent coupon-code edits.
type ProfileCompletionRestriction struct {
	Active bool `json:"active"`
}
