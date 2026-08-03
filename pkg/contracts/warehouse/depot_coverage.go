package warehouse

import "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/common"

// DepotCoverageRule is the geographic eligibility scope for one depot.
// CountryCode is required; the administrative area and postal code are
// optional so a rule can represent country-level coverage.
type DepotCoverageRule struct {
	ID                     string                 `json:"id"`
	DepotCode              string                 `json:"depot_code"`
	CountryCode            common.CountryCode     `json:"country_code"`
	AdministrativeAreaCode common.SubdivisionCode `json:"administrative_area_code,omitempty"`
	PostalCode             string                 `json:"postal_code,omitempty"`
	Priority               int                    `json:"priority"`
	IsActive               bool                   `json:"is_active"`

	common.AuditFields
}
