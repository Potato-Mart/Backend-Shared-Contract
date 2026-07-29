package enums_test

import (
	"testing"

	wishenum "github.com/Potato-Mart/Backend-Shared-Contract/v20/pkg/enums/wish"
)

func TestWishEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{
			name: "wishenum.WishProposalState",
			valid: []stringEnum{
				wishenum.WishProposalStatePending,
				wishenum.WishProposalStateConverted,
				wishenum.WishProposalStateRejected,
			},
			invalid: wishenum.WishProposalState("__invalid__"),
		},
		{
			name: "wishenum.WishCandidateState",
			valid: []stringEnum{
				wishenum.WishCandidateStateDraft,
				wishenum.WishCandidateStatePublished,
				wishenum.WishCandidateStateRetired,
				wishenum.WishCandidateStateFulfilled,
			},
			invalid: wishenum.WishCandidateState("__invalid__"),
		},
		{
			name: "wishenum.WishBallotState",
			valid: []stringEnum{
				wishenum.WishBallotStateScheduled,
				wishenum.WishBallotStateOpen,
				wishenum.WishBallotStateClosed,
			},
			invalid: wishenum.WishBallotState("__invalid__"),
		},
		{
			name: "wishenum.WishErrorCode",
			valid: []stringEnum{
				wishenum.WishErrorCodeNoActiveBallot,
				wishenum.WishErrorCodeBallotClosed,
				wishenum.WishErrorCodeCandidateUnavailable,
			},
			invalid: wishenum.WishErrorCode("__invalid__"),
		},
	})
}

func TestWishEnumWireValuesAreStable(t *testing.T) {
	want := map[stringEnum]string{
		wishenum.WishProposalStatePending:          "pending",
		wishenum.WishProposalStateConverted:        "converted",
		wishenum.WishProposalStateRejected:         "rejected",
		wishenum.WishCandidateStateDraft:           "draft",
		wishenum.WishCandidateStatePublished:       "published",
		wishenum.WishCandidateStateRetired:         "retired",
		wishenum.WishCandidateStateFulfilled:       "fulfilled",
		wishenum.WishBallotStateScheduled:          "scheduled",
		wishenum.WishBallotStateOpen:               "open",
		wishenum.WishBallotStateClosed:             "closed",
		wishenum.WishErrorCodeNoActiveBallot:       "WISH_NO_ACTIVE_BALLOT",
		wishenum.WishErrorCodeBallotClosed:         "WISH_BALLOT_CLOSED",
		wishenum.WishErrorCodeCandidateUnavailable: "WISH_CANDIDATE_UNAVAILABLE",
	}
	for value, wire := range want {
		if got := value.String(); got != wire {
			t.Fatalf("%T.String() = %q, want stable wire value %q", value, got, wire)
		}
	}
}
