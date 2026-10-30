package quote

import "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"

// MixedAdminDiscountWitness is versioned evidence for an administrative
// discount applied to an amendment containing mixed-source price allocations.
// It covers every quoted line, including ordinary and added lines. Version 1
// operations are new, replace, clear and carry. New/replace allocate the
// current desired amount over authenticated pre-admin line gross amounts by
// largest remainder, with line ID as tie key. Clear applies zero after
// reversing any retained carry; carry has no new allocation and preserves the
// previously allocated reductions on the selected native packages.
//
// The complete witness and authoritative priced-line and administrative
// discount wrappers must be included in Pricing's immutable pre-revision
// decision and persisted by transaction services for predecessor-bound replay.
// A line's native selection and pre-admin money must be independently
// authenticated; sums or a caller-supplied target alone are not authority.
// PricedLinesHash and AdministrativeDiscountHash bind the complete frozen
// priced-line mapping and administrative-discount wrappers. Each is lowercase
// SHA-256 of structural canonical JSON (sorted object keys, exact integers,
// preserved array order, UTC millisecond timestamps); a nil wrapper is JSON
// null and is distinct from present zero or empty evidence. These wrapper
// hashes are computed before the global witness hash and quote revision.
type MixedAdminDiscountWitness struct {
	Version                    int32                    `json:"version"`
	Operation                  string                   `json:"operation"`
	SourceVersionID            string                   `json:"source_version_id"`
	SourceQuoteKey             string                   `json:"source_quote_key"`
	SourceQuoteRevision        string                   `json:"source_quote_revision"`
	EditID                     string                   `json:"edit_id"`
	PricedLinesHash            string                   `json:"priced_lines_hash"`
	AdministrativeDiscountHash string                   `json:"administrative_discount_hash"`
	ActorID                    string                   `json:"actor_id,omitempty"`
	Reason                     string                   `json:"reason,omitempty"`
	PreviousDesired            money.Money              `json:"previous_desired"`
	CurrentDesired             money.Money              `json:"current_desired"`
	PreAdminTotal              money.Money              `json:"pre_admin_total"`
	Lines                      []MixedAdminDiscountLine `json:"lines"`
}
