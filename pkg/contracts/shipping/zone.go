package shipping

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/common"
)

type Zone struct {
	ID                      string                   `json:"id"`
	Name                    string                   `json:"name"`
	CountryCode             common.CountryCode       `json:"country_code"`
	AdministrativeAreaCodes []common.SubdivisionCode `json:"administrative_area_codes,omitempty"`
	PostalCodes             []string                 `json:"postal_codes,omitempty"`
	IsActive                bool                     `json:"is_active"`
	CreatedAt               time.Time                `json:"created_at"`
}
