package compliance_enums

type EvidenceKind string

const (
	EvidenceKindMedia              EvidenceKind = "media"
	EvidenceKindOfficialSource     EvidenceKind = "official_source"
	EvidenceKindInternalRecord     EvidenceKind = "internal_record"
	EvidenceKindManualCitation     EvidenceKind = "manual_citation"
	EvidenceKindAutomationProposal EvidenceKind = "automation_proposal"
)

func (k EvidenceKind) IsValid() bool {
	switch k {
	case EvidenceKindMedia, EvidenceKindOfficialSource, EvidenceKindInternalRecord,
		EvidenceKindManualCitation, EvidenceKindAutomationProposal:
		return true
	}
	return false
}

func (k EvidenceKind) String() string { return string(k) }
