package quote

// PromotionGroupSpan identifies complete repeated candidate groups. CandidateID
// is stable within the immutable Pricing decision and shared by all qualifier
// and target peers. Group ordinals and FirstGroupUnitOffset are one-based;
// GroupCount is the number of root groups touched and UnitsPerGroup is this
// role's number of units in each group. A fresh allocation contains complete
// groups. A predecessor-bound retained slice may start or end within a group;
// its omitted peers are authenticated from the immutable predecessor, not
// asserted as retained. For zero-based run offset j, the original group is
// FirstGroup + (FirstGroupUnitOffset - 1 + j) / UnitsPerGroup; the within-role
// position is 1 + (FirstGroupUnitOffset - 1 + j) % UnitsPerGroup. GroupCount
// equals the number of groups touched by BaseUnits. For a complete fresh run,
// the offset is 1 and BaseUnits = GroupCount * UnitsPerGroup. Consumers
// authenticate the entire original peer set.
type PromotionGroupSpan struct {
	CandidateID          string `json:"candidate_id"`
	FirstGroup           int64  `json:"first_group"`
	GroupCount           int64  `json:"group_count"`
	UnitsPerGroup        int64  `json:"units_per_group"`
	FirstGroupUnitOffset int64  `json:"first_group_unit_offset"`
}
