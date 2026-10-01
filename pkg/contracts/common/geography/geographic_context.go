package geography

import "github.com/Potato-Mart/Backend-Shared-Contract/v41/pkg/contracts/common/geography/geography_enums"

// GeographicContext is the immutable geographic resolution snapshot carried
// by pricing, eligibility, and order projections.
// MarketCode remains the authoritative single commercial market. Path records
// resolved geography and MatchedPath records a matched scope selection.
type GeographicContext struct {
	Path               *GeographicPath                         `json:"path,omitempty"`
	MatchedPath        *GeographicPath                         `json:"matched_path,omitempty"`
	Source             geography_enums.GeographicContextSource `json:"source"`
	MarketCode         string                                  `json:"market_code,omitempty"`
	ScopeRevision      int64                                   `json:"scope_revision"`
	RuleRevision       int64                                   `json:"rule_revision"`
	EvaluationTimezone string                                  `json:"evaluation_timezone"`
}
