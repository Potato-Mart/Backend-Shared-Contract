package geography

// CountryCode is an ISO 3166-1 alpha-2 country code such as AU, TW, JP, or US.
type CountryCode string

// SubdivisionCode is an ISO 3166-2 subdivision code such as AU-VIC or US-CA.
type SubdivisionCode string

// CountryRef carries an authoritative country code and optional display name.
type CountryRef struct {
	Code CountryCode `json:"code"`
	Name string      `json:"name,omitempty"`
}

// AdministrativeAreaRef carries an authoritative subdivision code and
// optional display metadata.
type AdministrativeAreaRef struct {
	Code SubdivisionCode        `json:"code"`
	Name string                 `json:"name,omitempty"`
	Type AdministrativeAreaType `json:"type,omitempty"`
}
