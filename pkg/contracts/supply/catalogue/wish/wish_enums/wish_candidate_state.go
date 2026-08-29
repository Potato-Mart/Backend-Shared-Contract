package wish_enums

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
