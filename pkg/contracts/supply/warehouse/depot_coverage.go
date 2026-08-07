package warehouse

import geography "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/common/geography"

import common "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/common/shared"

// DepotCoverageRule is the geographic eligibility scope for one depot.
// CountryCode is required; the administrative area and postal code are
// optional so a rule can represent country-level coverage.
type DepotCoverageRule struct {
	ID                     string                    `json:"id"`
	DepotCode              string                    `json:"depot_code"`
	CountryCode            geography.CountryCode     `json:"country_code"`
	AdministrativeAreaCode geography.SubdivisionCode `json:"administrative_area_code,omitempty"`
	PostalCode             string                    `json:"postal_code,omitempty"`
	Priority               int                       `json:"priority"`
	IsActive               bool                      `json:"is_active"`

	common.AuditFields
}
