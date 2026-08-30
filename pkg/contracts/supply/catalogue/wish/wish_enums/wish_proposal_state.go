package wish_enums

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
