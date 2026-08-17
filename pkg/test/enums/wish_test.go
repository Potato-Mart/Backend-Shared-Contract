package enums_test

import (
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/supply/wish/wish_enums"
)

func TestWishEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{
			name: "wishenum.WishProposalState",
			valid: []stringEnum{
				wish_enums.WishProposalStatePending,
				wish_enums.WishProposalStateConverted,
				wish_enums.WishProposalStateRejected,
			},
			invalid: wish_enums.WishProposalState("__invalid__"),
		},
		{
			name: "wishenum.WishCandidateState",
			valid: []stringEnum{
				wish_enums.WishCandidateStateDraft,
				wish_enums.WishCandidateStatePublished,
				wish_enums.WishCandidateStateRetired,
				wish_enums.WishCandidateStateFulfilled,
			},
			invalid: wish_enums.WishCandidateState("__invalid__"),
		},
		{
			name: "wishenum.WishBallotState",
			valid: []stringEnum{
				wish_enums.WishBallotStateScheduled,
				wish_enums.WishBallotStateOpen,
				wish_enums.WishBallotStateClosed,
			},
			invalid: wish_enums.WishBallotState("__invalid__"),
		},
		{
			name: "wishenum.WishErrorCode",
			valid: []stringEnum{
				wish_enums.WishErrorCodeNoActiveBallot,
				wish_enums.WishErrorCodeBallotClosed,
				wish_enums.WishErrorCodeCandidateUnavailable,
			},
			invalid: wish_enums.WishErrorCode("__invalid__"),
		},
	})
}

func TestWishEnumWireValuesAreStable(t *testing.T) {
	want := map[stringEnum]string{
		wish_enums.WishProposalStatePending:          "pending",
		wish_enums.WishProposalStateConverted:        "converted",
		wish_enums.WishProposalStateRejected:         "rejected",
		wish_enums.WishCandidateStateDraft:           "draft",
		wish_enums.WishCandidateStatePublished:       "published",
		wish_enums.WishCandidateStateRetired:         "retired",
		wish_enums.WishCandidateStateFulfilled:       "fulfilled",
		wish_enums.WishBallotStateScheduled:          "scheduled",
		wish_enums.WishBallotStateOpen:               "open",
		wish_enums.WishBallotStateClosed:             "closed",
		wish_enums.WishErrorCodeNoActiveBallot:       "WISH_NO_ACTIVE_BALLOT",
		wish_enums.WishErrorCodeBallotClosed:         "WISH_BALLOT_CLOSED",
		wish_enums.WishErrorCodeCandidateUnavailable: "WISH_CANDIDATE_UNAVAILABLE",
	}
	for value, wire := range want {
		if got := value.String(); got != wire {
			t.Fatalf("%T.String() = %q, want stable wire value %q", value, got, wire)
		}
	}
}
