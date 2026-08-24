package wish_enums

// WishErrorCode is a stable domain error code shared by wish clients.
type WishErrorCode string

const (
	WishErrorCodeNoActiveBallot       WishErrorCode = "WISH_NO_ACTIVE_BALLOT"
	WishErrorCodeBallotClosed         WishErrorCode = "WISH_BALLOT_CLOSED"
	WishErrorCodeCandidateUnavailable WishErrorCode = "WISH_CANDIDATE_UNAVAILABLE"
)

func (c WishErrorCode) String() string { return string(c) }

func (c WishErrorCode) IsValid() bool {
	switch c {
	case WishErrorCodeNoActiveBallot, WishErrorCodeBallotClosed,
		WishErrorCodeCandidateUnavailable:
		return true
	default:
		return false
	}
}
