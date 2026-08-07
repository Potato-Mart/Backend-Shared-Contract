package geography

// GeographicTarget names one inclusive target in the geographic hierarchy.
type GeographicTarget struct {
	Kind GeographicTargetKind `json:"kind"`
	Code string               `json:"code"`
}

// GeographicScope is either explicitly global or an inclusive list of
// country, subdivision, depot-region, or depot targets.
type GeographicScope struct {
	Mode    GeographicScopeMode `json:"mode"`
	Targets []GeographicTarget  `json:"targets,omitempty"`
}

// GeographicContext is the immutable geographic resolution snapshot carried
// by pricing, eligibility, and order projections.
type GeographicContext struct {
	Source             GeographicContextSource `json:"source"`
	CountryCode        CountryCode             `json:"country_code,omitempty"`
	SubdivisionCode    SubdivisionCode         `json:"subdivision_code,omitempty"`
	DepotRegionCode    string                  `json:"depot_region_code,omitempty"`
	DepotCode          string                  `json:"depot_code,omitempty"`
	MatchedTargetKind  GeographicTargetKind    `json:"matched_target_kind,omitempty"`
	MatchedTargetCode  string                  `json:"matched_target_code,omitempty"`
	ScopeRevision      int64                   `json:"scope_revision"`
	RuleRevision       int64                   `json:"rule_revision"`
	EvaluationTimezone string                  `json:"evaluation_timezone"`
}
