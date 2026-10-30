package quote

// PromotionCandidatePeerSpan binds a complete root group interval to its
// original contiguous one-based line ordinals. A span represents exactly
// GroupCount*UnitsPerGroup units, ordered by group then within-role ordinal.
// Retained clipped slices reference this root span and the authenticated
// predecessor chain; they do not claim omitted peers are still present.
type PromotionCandidatePeerSpan struct {
	FirstGroup int64 `json:"first_group"`
	GroupCount int64 `json:"group_count"`
	FirstUnit  int64 `json:"first_unit"`
}
