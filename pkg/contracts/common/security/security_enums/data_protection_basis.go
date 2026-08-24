package security_enums

// DataProtectionBasis records the business/legal basis for processing a
// record that may contain personal information.
type DataProtectionBasis string

const (
	DataProtectionBasisNotApplicable      DataProtectionBasis = "not_applicable"
	DataProtectionBasisConsent            DataProtectionBasis = "consent"
	DataProtectionBasisContract           DataProtectionBasis = "contract"
	DataProtectionBasisLegalObligation    DataProtectionBasis = "legal_obligation"
	DataProtectionBasisLegitimateInterest DataProtectionBasis = "legitimate_interest"
)

func (b DataProtectionBasis) IsValid() bool {
	switch b {
	case DataProtectionBasisNotApplicable, DataProtectionBasisConsent,
		DataProtectionBasisContract, DataProtectionBasisLegalObligation,
		DataProtectionBasisLegitimateInterest:
		return true
	}
	return false
}

func (b DataProtectionBasis) String() string { return string(b) }
