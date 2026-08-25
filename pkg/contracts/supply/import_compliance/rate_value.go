package import_compliance

// RateValue preserves official source text because tariff schedules may use
// compound or non-percentage rates. BasisPoints is populated only when the
// source can be faithfully represented as a percentage.
type RateValue struct {
	Raw         string `json:"raw"`
	BasisPoints *int64 `json:"basis_points,omitempty"`
}
