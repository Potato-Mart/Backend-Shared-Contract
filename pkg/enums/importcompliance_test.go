package enums_test

import (
	"testing"

	importcomplianceenum "github.com/Potato-Mart/Backend-Shared-Contract/v21/pkg/enums/importcompliance"
)

func TestImportComplianceEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "importcomplianceenum.ReviewState", valid: []stringEnum{importcomplianceenum.ReviewStateDraft, importcomplianceenum.ReviewStateInReview, importcomplianceenum.ReviewStateApproved, importcomplianceenum.ReviewStateRejected, importcomplianceenum.ReviewStateArchived}, invalid: importcomplianceenum.ReviewState("__invalid__")},
		{name: "importcomplianceenum.Jurisdiction", valid: []stringEnum{importcomplianceenum.JurisdictionAustralia, importcomplianceenum.JurisdictionTaiwan}, invalid: importcomplianceenum.Jurisdiction("__invalid__")},
		{name: "importcomplianceenum.ImportMode", valid: []stringEnum{importcomplianceenum.ImportModeAirCargo, importcomplianceenum.ImportModeSeaAmbient, importcomplianceenum.ImportModeSeaFrozen}, invalid: importcomplianceenum.ImportMode("__invalid__")},
		{name: "importcomplianceenum.EvidenceKind", valid: []stringEnum{importcomplianceenum.EvidenceKindMedia, importcomplianceenum.EvidenceKindOfficialSource, importcomplianceenum.EvidenceKindInternalRecord, importcomplianceenum.EvidenceKindManualCitation, importcomplianceenum.EvidenceKindAutomationProposal}, invalid: importcomplianceenum.EvidenceKind("__invalid__")},
		{name: "importcomplianceenum.RFIChannel", valid: []stringEnum{importcomplianceenum.RFIChannelBiosecurityPortal, importcomplianceenum.RFIChannelEmailException}, invalid: importcomplianceenum.RFIChannel("__invalid__")},
		{name: "importcomplianceenum.RFISubmissionState", valid: []stringEnum{importcomplianceenum.RFISubmissionStateNotSubmitted, importcomplianceenum.RFISubmissionStateSubmitted, importcomplianceenum.RFISubmissionStateConfirmed, importcomplianceenum.RFISubmissionStateRescheduled, importcomplianceenum.RFISubmissionStateCancelled, importcomplianceenum.RFISubmissionStateClosed}, invalid: importcomplianceenum.RFISubmissionState("__invalid__")},
		{name: "importcomplianceenum.RFIRequestedTime", valid: []stringEnum{importcomplianceenum.RFIRequestedTimeAnytime, importcomplianceenum.RFIRequestedTimeAM, importcomplianceenum.RFIRequestedTimePM}, invalid: importcomplianceenum.RFIRequestedTime("__invalid__")},
		{name: "importcomplianceenum.LabelSize", valid: []stringEnum{importcomplianceenum.LabelSize100x60, importcomplianceenum.LabelSize100x80, importcomplianceenum.LabelSize80x50, importcomplianceenum.LabelSize65x45, importcomplianceenum.LabelSize60x40, importcomplianceenum.LabelSize50x30}, invalid: importcomplianceenum.LabelSize("__invalid__")},
		{name: "importcomplianceenum.LabelOrientation", valid: []stringEnum{importcomplianceenum.LabelOrientationPortrait, importcomplianceenum.LabelOrientationLandscape}, invalid: importcomplianceenum.LabelOrientation("__invalid__")},
		{name: "importcomplianceenum.ArtifactKind", valid: []stringEnum{importcomplianceenum.ArtifactKindPDF, importcomplianceenum.ArtifactKindDOCX, importcomplianceenum.ArtifactKindXLSX, importcomplianceenum.ArtifactKindZIP, importcomplianceenum.ArtifactKindEML}, invalid: importcomplianceenum.ArtifactKind("__invalid__")},
	})
}
