package inventory

// UnknownReturnPhysicalEvidence represents physically present returned stock
// without a known lot, unit identity or verified cost. Sources must be unique
// within this allocation or movement, positive, stably ordered by the owner,
// and total exactly its base units. The same source may support several distinct owner-verified claims.
type UnknownReturnPhysicalEvidence struct {
	ReturnSources []UnknownReturnSource `json:"return_sources"`
}
