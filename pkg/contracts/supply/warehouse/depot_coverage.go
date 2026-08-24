package warehouse

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/audit"
	geography "github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/geography"
)

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

	audit.AuditFields
}
