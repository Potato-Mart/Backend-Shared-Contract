package order

// AdClick is the paid-advertising click evidence a client hands to checkout:
// the visitor's most recent ad click, their first one, and the advertising
// consent that governs both.
//
// Last and First are kept apart because the attribution rule chooses between
// them: the agreed default credits the last ad click inside the window, while
// a first-click rule credits the earliest. A later direct or organic visit
// never overwrites Last, so an order can still be credited to the ad that
// originally brought the visitor. Either may be absent.
//
// Consent governs the whole record. Without granted advertising storage there
// is no click to hand over at all, and without granted user-data consent no
// hashed contact digest may accompany the purchase to a platform.
//
// It holds no personal data. Which platform the click belongs to is decided by
// the receiving service from the tracking parameters, never taken from the
// client, and the stored result is the read-only AdAttribution on the order.
type AdClick struct {
	Last    *AdClickSnapshot `json:"last,omitempty"`
	First   *AdClickSnapshot `json:"first,omitempty"`
	Consent *AdConsent       `json:"consent,omitempty"`
}
