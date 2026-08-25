package geography

import "github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/geography/geography_enums"

// GeographicContext is the immutable geographic resolution snapshot carried
// by pricing, eligibility, and order projections.
type GeographicContext struct {
	Source             geography_enums.GeographicContextSource `json:"source"`
	MarketCode         string                                  `json:"market_code,omitempty"`
	CountryCode        CountryCode                             `json:"country_code,omitempty"`
	SubdivisionCode    SubdivisionCode                         `json:"subdivision_code,omitempty"`
	DepotRegionCode    string                                  `json:"depot_region_code,omitempty"`
	DepotCode          string                                  `json:"depot_code,omitempty"`
	MatchedTargetKind  geography_enums.GeographicTargetKind    `json:"matched_target_kind,omitempty"`
	MatchedTargetCode  string                                  `json:"matched_target_code,omitempty"`
	ScopeRevision      int64                                   `json:"scope_revision"`
	RuleRevision       int64                                   `json:"rule_revision"`
	EvaluationTimezone string                                  `json:"evaluation_timezone"`
}
