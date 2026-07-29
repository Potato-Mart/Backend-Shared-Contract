package wish

import (
	"time"

	wishenum "github.com/Potato-Mart/Backend-Shared-Contract/v21/pkg/enums/wish"
)

// WishBallot is a revisioned, time-bounded set of candidate identifiers.
type WishBallot struct {
	ID           string                   `json:"id"`
	State        wishenum.WishBallotState `json:"state"`
	CandidateIDs []string                 `json:"candidate_ids"`
	OpensAt      time.Time                `json:"opens_at"`
	ClosesAt     time.Time                `json:"closes_at"`
	Revision     int64                    `json:"revision"`
	AsOf         time.Time                `json:"as_of"`
}

// WishRankingEntry records one candidate's ordered, backend-computed position.
type WishRankingEntry struct {
	CandidateID string `json:"candidate_id"`
	Rank        int    `json:"rank"`
	VoteCount   int64  `json:"vote_count"`
}

// WishSelection is a customer's identity-free selection for one ballot.
// CandidateIDs preserves the customer's selected order when multiple choices
// are allowed by the backend-owned ballot rules.
type WishSelection struct {
	BallotID     string    `json:"ballot_id"`
	CandidateIDs []string  `json:"candidate_ids"`
	UpdatedAt    time.Time `json:"updated_at"`
}
