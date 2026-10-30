package quote

// PromotionCandidateRequirement preserves an authored role-vector index even
// when two requirements use the same line and role. WholePackage states the
// actual requirement, not a guess from the chosen unit count. Every peer is
// necessary for every indicated original group; zero-reduction qualifiers are
// explicit and physically disjoint from other competing participants.
type PromotionCandidateRequirement struct {
	Index         int32                    `json:"index"`
	Role          string                   `json:"role"`
	WholePackage  bool                     `json:"whole_package"`
	UnitsPerGroup int64                    `json:"units_per_group"`
	Peers         []PromotionCandidatePeer `json:"peers"`
}
