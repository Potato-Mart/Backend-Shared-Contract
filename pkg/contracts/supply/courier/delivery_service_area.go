package courier

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v35/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v35/pkg/contracts/supply/courier/courier_enums"
)

// DeliveryServiceArea is one configured country/postal-code eligibility rule
// within a company. Postal codes are exact country-normalized strings, never
// numeric ranges or inferred metropolitan boundaries. Exclusions win.
// Include-only with no inclusions matches nothing; all-except explicitly
// considers the country outside its exclusions, subject to coverage checks.
// All-except requires ProviderCoverageRequired=true; it is not a country-wide
// coverage assertion. CountryCode must also appear in the parent CountryCodes.
// Lower RoutingPriority values rank first; the service rejects ambiguous ties.
type DeliveryServiceArea struct {
	Code        string                `json:"code"`
	CountryCode geography.CountryCode `json:"country_code"`
	// ShippingZone identifies an Orders-owned shipping zone by ID only. Supply
	// resolves the live zone record when checking coverage; it is not a snapshot.
	ShippingZone       *ShippingZoneRef                  `json:"shipping_zone,omitempty"`
	StateCodes         []geography.SubdivisionCode       `json:"state_codes,omitempty"`
	PostalCodeMode     courier_enums.PostalCodeMatchMode `json:"postal_code_mode"`
	IncludePostalCodes []string                          `json:"include_postal_codes,omitempty"`
	ExcludePostalCodes []string                          `json:"exclude_postal_codes,omitempty"`
	RoutingPriority    int                               `json:"routing_priority"`
	Enabled            bool                              `json:"enabled"`
	// ProviderCoverageRequired requires an authoritative provider coverage
	// confirmation in addition to this configured filter. Unknown is not covered.
	ProviderCoverageRequired bool `json:"provider_coverage_required"`
}
