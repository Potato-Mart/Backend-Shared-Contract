package ads

// AdDeliveryEstimate is the platform's own estimate of how far a targeting
// definition would reach, shown while someone is editing an audience.
//
// It is the platform's number, not ours, and it is advisory: the range is wide,
// it changes between reads, and it is never evidence of what an ad will
// actually achieve. The bounds are a range because platforms report one;
// LowerBound and UpperBound are both absent when the platform declined to
// estimate, for example because the audience is too small to report on.
type AdDeliveryEstimate struct {
	LowerBound int64 `json:"lower_bound,omitempty"`
	UpperBound int64 `json:"upper_bound,omitempty"`
}
