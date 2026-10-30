package quote

import "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"

// PromotionApplicationPeerProof uses aggregate amounts for this root ordinal
// interval, not a per-unit amount. BookAfterRounding is the actual scorer
// event and may reference a larger complete population via its pool evidence.
type PromotionApplicationPeerProof struct {
	RootLineID              string              `json:"root_line_id"`
	RequirementIndex        int32               `json:"requirement_index"`
	FirstUnit               int64               `json:"first_unit"`
	BaseUnits               int64               `json:"base_units"`
	ExactBookReduction      PromotionExactMinor `json:"exact_book_reduction"`
	ExactBookAfter          PromotionExactMinor `json:"exact_book_after"`
	BookAfterRounding       RoundingEvidence    `json:"book_after_rounding"`
	ExactNetCandidateEffect PromotionExactMinor `json:"exact_net_candidate_effect"`
	RecordedNetReduction    money.Money         `json:"recorded_net_reduction"`
	NetRounding             *RoundingEvidence   `json:"net_rounding,omitempty"`
}
