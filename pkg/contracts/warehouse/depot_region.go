package warehouse

import "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/common"

// DepotRegion identifies an operational region between an administrative area
// and its depots. Code is the canonical business reference.
type DepotRegion struct {
	ID                     string                 `json:"id"`
	Code                   string                 `json:"code"`
	Name                   string                 `json:"name"`
	CountryCode            common.CountryCode     `json:"country_code"`
	AdministrativeAreaCode common.SubdivisionCode `json:"administrative_area_code"`
	IsActive               bool                   `json:"is_active"`

	common.AuditFields
}
