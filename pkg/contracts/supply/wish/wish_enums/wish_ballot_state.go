package wish_enums

// WishBallotState is the publication lifecycle of a wish ballot.
type WishBallotState string

const (
	WishBallotStateScheduled WishBallotState = "scheduled"
	WishBallotStateOpen      WishBallotState = "open"
	WishBallotStateClosed    WishBallotState = "closed"
)

func (s WishBallotState) String() string { return string(s) }

func (s WishBallotState) IsValid() bool {
	switch s {
	case WishBallotStateScheduled, WishBallotStateOpen, WishBallotStateClosed:
		return true
	default:
		return false
	}
}
