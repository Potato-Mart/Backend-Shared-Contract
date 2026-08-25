package geography

// CountryRef carries an authoritative country code and optional display name.
type CountryRef struct {
	Code CountryCode `json:"code"`
	Name string      `json:"name,omitempty"`
}
