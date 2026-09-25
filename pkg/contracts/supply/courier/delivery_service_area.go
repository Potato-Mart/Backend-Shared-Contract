package courier

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v36/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v36/pkg/contracts/supply/courier/courier_enums"
)

// DeliveryServiceArea is one market- and zone-bound eligibility rule within a
// company. Postal codes are exact country-normalized strings, never
// numeric ranges or inferred metropolitan boundaries. Exclusions win.
// Include-only with no inclusions matches nothing; all-except explicitly
// considers the live zone outside its exclusions, subject to coverage checks.
// All-except requires ProviderCoverageRequired=true; it is not a country-wide
// coverage assertion. CountryCode must also appear in the parent CountryCodes.
// Supply validates MarketCode against the country and resolves the live zone.
// Selecting all country zones materializes explicit rows; no missing field or
// empty list is a country wildcard. Overlap validation remains service-owned.
type DeliveryServiceArea struct {
	Code        string                `json:"code"`
	CountryCode geography.CountryCode `json:"country_code"`
	// MarketCode identifies the Pricing-owned commercial market, not its ID.
	// An absent legacy market requires explicit migration, never country inference.
	MarketCode string `json:"market_code"`
	// ShippingZone identifies an Orders-owned shipping zone by ID only. Supply
	// resolves the live zone record when checking coverage; it is not a snapshot.
	// Nil permits reading legacy records but grants no wildcard coverage.
	ShippingZone *ShippingZoneRef `json:"shipping_zone,omitempty"`
	// StateCodes is derived from the referenced zone, not independent coverage.
	StateCodes         []geography.SubdivisionCode       `json:"state_codes,omitempty"`
	PostalCodeMode     courier_enums.PostalCodeMatchMode `json:"postal_code_mode"`
	IncludePostalCodes []string                          `json:"include_postal_codes,omitempty"`
	ExcludePostalCodes []string                          `json:"exclude_postal_codes,omitempty"`
	Enabled            bool                              `json:"enabled"`
	// ProviderCoverageRequired requires an authoritative provider coverage
	// confirmation in addition to this configured filter. Unknown is not covered.
	ProviderCoverageRequired bool `json:"provider_coverage_required"`
}
