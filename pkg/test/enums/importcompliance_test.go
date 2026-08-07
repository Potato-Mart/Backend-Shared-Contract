package enums_test

import (
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v24/pkg/contracts/supply/importcompliance/importcompliance_enums"
)

func TestImportComplianceEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "importcomplianceenum.ReviewState", valid: []stringEnum{importcompliance_enums.ReviewStateDraft, importcompliance_enums.ReviewStateInReview, importcompliance_enums.ReviewStateApproved, importcompliance_enums.ReviewStateRejected, importcompliance_enums.ReviewStateArchived}, invalid: importcompliance_enums.ReviewState("__invalid__")},
		{name: "importcomplianceenum.Jurisdiction", valid: []stringEnum{importcompliance_enums.JurisdictionAustralia, importcompliance_enums.JurisdictionTaiwan}, invalid: importcompliance_enums.Jurisdiction("__invalid__")},
		{name: "importcomplianceenum.ImportMode", valid: []stringEnum{importcompliance_enums.ImportModeAirCargo, importcompliance_enums.ImportModeSeaAmbient, importcompliance_enums.ImportModeSeaFrozen}, invalid: importcompliance_enums.ImportMode("__invalid__")},
		{name: "importcomplianceenum.EvidenceKind", valid: []stringEnum{importcompliance_enums.EvidenceKindMedia, importcompliance_enums.EvidenceKindOfficialSource, importcompliance_enums.EvidenceKindInternalRecord, importcompliance_enums.EvidenceKindManualCitation, importcompliance_enums.EvidenceKindAutomationProposal}, invalid: importcompliance_enums.EvidenceKind("__invalid__")},
		{name: "importcomplianceenum.RFIChannel", valid: []stringEnum{importcompliance_enums.RFIChannelBiosecurityPortal, importcompliance_enums.RFIChannelEmailException}, invalid: importcompliance_enums.RFIChannel("__invalid__")},
		{name: "importcomplianceenum.RFISubmissionState", valid: []stringEnum{importcompliance_enums.RFISubmissionStateNotSubmitted, importcompliance_enums.RFISubmissionStateSubmitted, importcompliance_enums.RFISubmissionStateConfirmed, importcompliance_enums.RFISubmissionStateRescheduled, importcompliance_enums.RFISubmissionStateCancelled, importcompliance_enums.RFISubmissionStateClosed}, invalid: importcompliance_enums.RFISubmissionState("__invalid__")},
		{name: "importcomplianceenum.RFIRequestedTime", valid: []stringEnum{importcompliance_enums.RFIRequestedTimeAnytime, importcompliance_enums.RFIRequestedTimeAM, importcompliance_enums.RFIRequestedTimePM}, invalid: importcompliance_enums.RFIRequestedTime("__invalid__")},
		{name: "importcomplianceenum.LabelSize", valid: []stringEnum{importcompliance_enums.LabelSize100x60, importcompliance_enums.LabelSize100x80, importcompliance_enums.LabelSize80x50, importcompliance_enums.LabelSize65x45, importcompliance_enums.LabelSize60x40, importcompliance_enums.LabelSize50x30}, invalid: importcompliance_enums.LabelSize("__invalid__")},
		{name: "importcomplianceenum.LabelOrientation", valid: []stringEnum{importcompliance_enums.LabelOrientationPortrait, importcompliance_enums.LabelOrientationLandscape}, invalid: importcompliance_enums.LabelOrientation("__invalid__")},
		{name: "importcomplianceenum.ArtifactKind", valid: []stringEnum{importcompliance_enums.ArtifactKindPDF, importcompliance_enums.ArtifactKindDOCX, importcompliance_enums.ArtifactKindXLSX, importcompliance_enums.ArtifactKindZIP, importcompliance_enums.ArtifactKindEML}, invalid: importcompliance_enums.ArtifactKind("__invalid__")},
	})
}
