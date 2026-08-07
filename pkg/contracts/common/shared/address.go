package shared

import geography "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/common/geography"

type Address struct {
	Label string `json:"label,omitempty"`
	Line1 string `json:"line1"`
	Line2 string `json:"line2,omitempty"`

	// Locality covers a city, suburb, town, municipality, or equivalent
	// postal locality.
	Locality string `json:"locality"`

	AdministrativeArea *geography.AdministrativeAreaRef `json:"administrative_area,omitempty"`
	PostalCode         string                           `json:"postal_code"`
	Country            geography.CountryRef             `json:"country"`

	FormattedAddress string   `json:"formatted_address,omitempty"`
	PlaceID          string   `json:"place_id,omitempty"`
	Latitude         *float64 `json:"latitude,omitempty"`
	Longitude        *float64 `json:"longitude,omitempty"`
}
