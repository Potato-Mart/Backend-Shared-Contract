package geography

import "github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/geography/geography_enums"

// CountryRef carries an authoritative country code and optional display name.
type CountryRef struct {
	Code CountryCode `json:"code"`
	Name string      `json:"name,omitempty"`
}

// AdministrativeAreaRef carries an authoritative subdivision code and
// optional display metadata.
type AdministrativeAreaRef struct {
	Code SubdivisionCode                        `json:"code"`
	Name string                                 `json:"name,omitempty"`
	Type geography_enums.AdministrativeAreaType `json:"type,omitempty"`
}
