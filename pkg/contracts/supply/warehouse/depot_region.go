package warehouse

import geography "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/common/geography"

import common "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/common/shared"

// DepotRegion identifies an operational region between an administrative area
// and its depots. Code is the canonical business reference.
type DepotRegion struct {
	ID                     string                    `json:"id"`
	Code                   string                    `json:"code"`
	Name                   string                    `json:"name"`
	CountryCode            geography.CountryCode     `json:"country_code"`
	AdministrativeAreaCode geography.SubdivisionCode `json:"administrative_area_code"`
	IsActive               bool                      `json:"is_active"`

	common.AuditFields
}
