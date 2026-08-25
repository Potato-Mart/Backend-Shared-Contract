package wish

import (
	"time"
)

// WishSelection is a customer's identity-free selection for one ballot.
// CandidateIDs preserves the customer's selected order when multiple choices
// are allowed by the backend-owned ballot rules.
type WishSelection struct {
	BallotID     string    `json:"ballot_id"`
	CandidateIDs []string  `json:"candidate_ids"`
	UpdatedAt    time.Time `json:"updated_at"`
}
