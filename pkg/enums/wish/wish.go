package wishenum

// WishProposalState is the lifecycle state of a customer product proposal.
type WishProposalState string

const (
	WishProposalStatePending   WishProposalState = "pending"
	WishProposalStateConverted WishProposalState = "converted"
	WishProposalStateRejected  WishProposalState = "rejected"
)

func (s WishProposalState) String() string { return string(s) }

func (s WishProposalState) IsValid() bool {
	switch s {
	case WishProposalStatePending, WishProposalStateConverted, WishProposalStateRejected:
		return true
	default:
		return false
	}
}

// WishCandidateState is the lifecycle state of an admin-authored ballot choice.
type WishCandidateState string

const (
	WishCandidateStateDraft     WishCandidateState = "draft"
	WishCandidateStatePublished WishCandidateState = "published"
	WishCandidateStateRetired   WishCandidateState = "retired"
	WishCandidateStateFulfilled WishCandidateState = "fulfilled"
)

func (s WishCandidateState) String() string { return string(s) }

func (s WishCandidateState) IsValid() bool {
	switch s {
	case WishCandidateStateDraft, WishCandidateStatePublished,
		WishCandidateStateRetired, WishCandidateStateFulfilled:
		return true
	default:
		return false
	}
}

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
