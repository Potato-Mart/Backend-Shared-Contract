package shipping

import (
	"time"

	geography "github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/geography"
)

type Zone struct {
	ID                      string                      `json:"id"`
	Name                    string                      `json:"name"`
	CountryCode             geography.CountryCode       `json:"country_code"`
	AdministrativeAreaCodes []geography.SubdivisionCode `json:"administrative_area_codes,omitempty"`
	PostalCodes             []string                    `json:"postal_codes,omitempty"`
	IsActive                bool                        `json:"is_active"`
	CreatedAt               time.Time                   `json:"created_at"`
}
