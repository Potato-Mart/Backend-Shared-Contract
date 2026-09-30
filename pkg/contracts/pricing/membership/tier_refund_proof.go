package membership

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v39/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v39/pkg/contracts/pricing/membership/membership_enums"
)

// TierRefundProof is an immutable projection of a Pricing-owned net tier-refund
// reservation, not authority by itself. Pricing verifies every field and the
// enclosing refund/order/customer/market/currency against its protected native
// record before consuming it. Opaque identities/fingerprints compare exactly;
// never trim, convert, derive aliases or hash again. Selection, unit lineage,
// preparation, replay/cumulative caps and cancellation are service-owned.
//
// QuotedCredit requires the original credited quote pair, including a verified
// zero grant. ExcludedGiftPurchase requires an immutable original gift-purchase
// witness, zero net reversal and no quote pair. Missing proof, quote or credit
// never implies exclusion or verified zero. No historical reconstruction applies.
type TierRefundProof struct {
	Basis            membership_enums.TierRefundProofBasis `json:"basis"`
	ProofID          string                                `json:"proof_id"`
	ProofFingerprint string                                `json:"proof_fingerprint"`
	AllocationID     string                                `json:"allocation_id"`
	// AllocationVersion is the positive immutable Pricing allocation format/
	// algorithm version (initially 1), never a numeric quote revision.
	AllocationVersion     int64  `json:"allocation_version"`
	AllocationFingerprint string `json:"allocation_fingerprint"`
	QuoteKey              string `json:"quote_key,omitempty"`
	QuoteFingerprint      string `json:"quote_fingerprint,omitempty"`
	// NetQualifyingSpendReversal is the verified nonnegative incremental NET goods
	// tier amount. Required Money preserves currency-bearing verified zero. It is
	// distinct from gross earned-point clawback and from refunded settlement money.
	NetQualifyingSpendReversal money.Money `json:"net_qualifying_spend_reversal"`
}
