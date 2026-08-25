package wish

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/supply/wish/wish_enums"
)

// WishBallot is a revisioned, time-bounded set of candidate identifiers.
type WishBallot struct {
	ID           string                     `json:"id"`
	State        wish_enums.WishBallotState `json:"state"`
	CandidateIDs []string                   `json:"candidate_ids"`
	OpensAt      time.Time                  `json:"opens_at"`
	ClosesAt     time.Time                  `json:"closes_at"`
	Revision     int64                      `json:"revision"`
	AsOf         time.Time                  `json:"as_of"`
}
