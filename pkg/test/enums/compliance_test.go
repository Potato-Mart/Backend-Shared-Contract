package enums_test

import (
	"testing"

	compliance_enums "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/compliance/compliance_enums"
)

func TestComplianceEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "compliance_enums.ReviewState", valid: []stringEnum{compliance_enums.ReviewStateDraft, compliance_enums.ReviewStateInReview, compliance_enums.ReviewStateApproved, compliance_enums.ReviewStateRejected, compliance_enums.ReviewStateArchived}, invalid: compliance_enums.ReviewState("__invalid__")},
		{name: "compliance_enums.Jurisdiction", valid: []stringEnum{compliance_enums.JurisdictionAustralia, compliance_enums.JurisdictionTaiwan}, invalid: compliance_enums.Jurisdiction("__invalid__")},
		{name: "compliance_enums.ImportMode", valid: []stringEnum{compliance_enums.ImportModeAirCargo, compliance_enums.ImportModeSeaAmbient, compliance_enums.ImportModeSeaFrozen}, invalid: compliance_enums.ImportMode("__invalid__")},
		{name: "compliance_enums.EvidenceKind", valid: []stringEnum{compliance_enums.EvidenceKindMedia, compliance_enums.EvidenceKindOfficialSource, compliance_enums.EvidenceKindInternalRecord, compliance_enums.EvidenceKindManualCitation, compliance_enums.EvidenceKindAutomationProposal}, invalid: compliance_enums.EvidenceKind("__invalid__")},
		{name: "compliance_enums.RFIChannel", valid: []stringEnum{compliance_enums.RFIChannelBiosecurityPortal, compliance_enums.RFIChannelEmailException}, invalid: compliance_enums.RFIChannel("__invalid__")},
		{name: "compliance_enums.RFISubmissionState", valid: []stringEnum{compliance_enums.RFISubmissionStateNotSubmitted, compliance_enums.RFISubmissionStateSubmitted, compliance_enums.RFISubmissionStateConfirmed, compliance_enums.RFISubmissionStateRescheduled, compliance_enums.RFISubmissionStateCancelled, compliance_enums.RFISubmissionStateClosed}, invalid: compliance_enums.RFISubmissionState("__invalid__")},
		{name: "compliance_enums.RFIRequestedTime", valid: []stringEnum{compliance_enums.RFIRequestedTimeAnytime, compliance_enums.RFIRequestedTimeAM, compliance_enums.RFIRequestedTimePM}, invalid: compliance_enums.RFIRequestedTime("__invalid__")},
		{name: "compliance_enums.LabelSize", valid: []stringEnum{compliance_enums.LabelSize100x60, compliance_enums.LabelSize100x80, compliance_enums.LabelSize80x50, compliance_enums.LabelSize65x45, compliance_enums.LabelSize60x40, compliance_enums.LabelSize50x30}, invalid: compliance_enums.LabelSize("__invalid__")},
		{name: "compliance_enums.LabelOrientation", valid: []stringEnum{compliance_enums.LabelOrientationPortrait, compliance_enums.LabelOrientationLandscape}, invalid: compliance_enums.LabelOrientation("__invalid__")},
		{name: "compliance_enums.ArtifactKind", valid: []stringEnum{compliance_enums.ArtifactKindPDF, compliance_enums.ArtifactKindDOCX, compliance_enums.ArtifactKindXLSX, compliance_enums.ArtifactKindZIP, compliance_enums.ArtifactKindEML}, invalid: compliance_enums.ArtifactKind("__invalid__")},
	})
}
