package wish

// WishRankingEntry records one candidate's ordered, backend-computed position.
type WishRankingEntry struct {
	CandidateID string `json:"candidate_id"`
	Rank        int    `json:"rank"`
	VoteCount   int64  `json:"vote_count"`
}
