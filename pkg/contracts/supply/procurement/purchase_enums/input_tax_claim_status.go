package purchase_enums

// InputTaxClaimStatus records whether recorded input tax may be claimed. It is
// never set to claimable without qualifying tax-invoice evidence.
type InputTaxClaimStatus string

const (
	InputTaxClaimStatusClaimable    InputTaxClaimStatus = "claimable"
	InputTaxClaimStatusNotClaimable InputTaxClaimStatus = "not_claimable"
	// InputTaxClaimStatusInsufficientEvidence is a claim blocked because
	// the qualifying tax-invoice evidence is missing or incomplete.
	InputTaxClaimStatusInsufficientEvidence InputTaxClaimStatus = "insufficient_evidence"
	InputTaxClaimStatusPendingReview        InputTaxClaimStatus = "pending_review"
)

// IsValid reports whether s is a known InputTaxClaimStatus.
func (s InputTaxClaimStatus) IsValid() bool {
	switch s {
	case InputTaxClaimStatusClaimable, InputTaxClaimStatusNotClaimable,
		InputTaxClaimStatusInsufficientEvidence, InputTaxClaimStatusPendingReview:
		return true
	}
	return false
}

func (s InputTaxClaimStatus) String() string { return string(s) }
